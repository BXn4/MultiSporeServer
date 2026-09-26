package creature

import (
	"database/sql/driver"
	"fmt"
	"multispore/internal/types"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Rigblock struct {
	ID                int16          // int16_t mIndex
	Field_2           int16          // int16_t field_2
	Parent            int16          // int16_t mParentIndex
	Symmetric         int16          // int16_t mSymmetricIndex
	Flags             int16          // int16_t mFlags
	Field_A           int8           // int8_t field_A
	Type              int            // RigblockDataType mType
	Capability        int16          // int16_t mCapabilityIndex
	Morphs            int16          // int16_t mMorphsIndex
	Capabilites       int8           // int8_t mNumCapabilities
	Bounding          int            // Math::BoundingBox mBoundingBox
	Scale             float32        // float mScale
	Orientation       types.Matrix3  // Math::Matrix3 mTotalOrientation
	Position          types.Position // Math::Vector3 mPosition
	Offset            types.Vector3  // Math::Vector3 mEffectOffset
	ScaleRelative     float32        // float mScaleRelative
	Distance          float32        // float mSocketConnectorDistance
	MuscleScale       float32        // float mMuscleScale
	MuscleScaleBase   float32        // float mBaseMuscleScale
	Field_7           float32        // float field_7C
	FootWeaponOrMouth uint32         // uint32_t mFootWeaponOrMouthType
	GroupID           uint32         // uint32_t mGroupID
	InstanceID        uint32         // uint32_t mInstanceID
}

const (
	NullBlock = 0
	Unk1      = 1
	Standard  = 2
	Unk3      = 3
	PlantRoot = 4
	Vertebra  = 5
)

// Flags
const (
	notUseSkin   = 1
	isBaked      = 8
	isFoot       = 32
	isWeapon     = 128
	isJiggable   = 512
	extraJiggly  = 2048
	isAsymmetric = 4096
)

type Rigblocks []Rigblock

func (r *Rigblocks) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("Failed to unmarshal Rigblocks value: %v", value)
		}
		str = string(bytes)
	}

	rigblocks := NewRigblockFromString(str)
	if rigblocks == nil {
		return fmt.Errorf("Failed to parse rigblocks string: %s", str)
	}

	*r = *rigblocks
	return nil
}
func (r Rigblocks) Value() (driver.Value, error) {
	if len(r) == 0 {
		return "", nil
	}
	return r.String(), nil
}

func NewRigblockFromString(s string) *Rigblocks {
	ids := strings.Split(s, "#")
	r := make(Rigblocks, 0, len(ids))
	for _, part := range ids {
		if part == "" {
			continue
		}

		parts := strings.Split(part, "+")

		partsLen := len(parts)

		switch partsLen {
		// other not known yet
		case 15:
			{
				rId, _ := strconv.ParseInt(parts[0], 10, 16)
				rParent, _ := strconv.ParseInt(parts[1], 10, 16)
				rSymmetric, _ := strconv.ParseInt(parts[2], 10, 16)
				rFlags, _ := strconv.ParseInt(parts[3], 10, 16)
				rType, _ := strconv.Atoi(parts[4])
				rScale, _ := strconv.ParseFloat(parts[5], 32)
				var rPos types.Position
				rPos.Scan(strings.Join(parts[6:9], "+"))
				rScaleRelative, _ := strconv.ParseFloat(parts[9], 32)
				rMuscleScale, _ := strconv.ParseFloat(parts[10], 32)
				rMuscleScaleBase, _ := strconv.ParseFloat(parts[11], 32)
				rFootWeaponOrMouth, _ := strconv.ParseUint(parts[12], 10, 32)
				rGroupID, _ := strconv.ParseUint(parts[13], 10, 32)
				rInstanceID, _ := strconv.ParseUint(parts[14], 10, 32)

				r = append(r, Rigblock{
					ID:                int16(rId),
					Parent:            int16(rParent),
					Symmetric:         int16(rSymmetric),
					Flags:             int16(rFlags),
					Type:              rType,
					Scale:             float32(rScale),
					Position:          rPos,
					ScaleRelative:     float32(rScaleRelative),
					MuscleScale:       float32(rMuscleScale),
					MuscleScaleBase:   float32(rMuscleScaleBase),
					FootWeaponOrMouth: uint32(rFootWeaponOrMouth),
					GroupID:           uint32(rGroupID),
					InstanceID:        uint32(rInstanceID),
				})
			}
		default:
			log.Errorf("Failed to create a new rigblock from string, because the length of the parts: (%d) not supported!", partsLen)
			return nil
		}
	}

	return &r
}

func (rb Rigblock) GetPartType() string {
	switch {
	case rb.Type == Vertebra:
		return "Spine"

	case int(rb.Flags)&isFoot != 0:
		return "Foot"

	case int(rb.Flags)&isWeapon != 0:
		return "Weapon"

	case rb.Type == Unk3:
		return "Unk3"

	case rb.FootWeaponOrMouth != 0:
		return "Foot or Weapon or Mouth"

	case rb.Parent == 0 && rb.Symmetric >= 0:
		return "Head Detail"

	case rb.Type == Standard:
		return "Attached Part"

	default:
		return "Unknown"
	}
}

func (rb Rigblock) string() string {
	return fmt.Sprintf("%d+%d+%d+%d+%d+%f+%s+%f+%f+%f+%d+%d+%d",
		rb.ID,
		rb.Parent,
		rb.Symmetric,
		rb.Flags,
		rb.Type,
		rb.Scale,
		rb.Position.String(),
		rb.ScaleRelative,
		rb.MuscleScale,
		rb.MuscleScaleBase,
		rb.FootWeaponOrMouth,
		rb.GroupID,
		rb.InstanceID)
}

func (r Rigblocks) String() string {
	entries := make([]string, len(r))
	for i, rb := range r {
		entries[i] = rb.string()
	}
	return strings.Join(entries, "#")
}
