package creature

import (
	"fmt"
	"multispore/internal/types"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

/* CELL STAGE
2026/09/12 09:53:04 INFO Received 11 rigblocks from client 127.0.0.1
2026/09/12 09:53:04 INFO Rigblock 0="[0 -1 -1 11 5 0.4000 -0.0000 -0.2851 0.0000 0.0000 0.0000 1.0000 0 1080123392 3244096132]"
2026/09/12 09:53:04 INFO Rigblock 1="[1 0 -1 9 5 0.6000 -0.0000 -0.1734 0.0000 0.2500 0.0000 1.0000 0 1080123392 3244096132]"
2026/09/12 09:53:04 INFO Rigblock 2="[2 1 -1 9 5 1.0750 -0.0000 -0.0618 0.0000 0.8438 0.0000 1.0000 0 1080123392 3244096132]"
2026/09/12 09:53:04 INFO Rigblock 3="[3 2 -1 9 5 0.9750 -0.0000 0.0499 0.0000 0.7188 0.0000 1.0000 0 1080123392 3244096132]"
2026/09/12 09:53:04 INFO Rigblock 4="[4 3 -1 9 5 0.6250 -0.0000 0.1616 0.0000 0.2812 0.0000 1.0000 0 1080123392 3244096132]"
2026/09/12 09:53:04 INFO Rigblock 5="[5 4 -1 9 5 0.5625 -0.0000 0.2721 -0.0112 0.2031 0.0000 1.0000 0 1080123392 3244096132]"
2026/09/12 09:53:04 INFO Rigblock 6="[6 0 -1 1088 2 0.9277 -0.0000 -0.3648 -0.0963 0.4277 0.0000 1.0000 1014894446 1080123392 2714007141]"
2026/09/12 09:53:04 INFO Rigblock 7="[7 2 -1 512 2 0.8607 -0.0000 -0.0390 0.0634 0.3607 0.0000 1.0000 0 1080123392 1729351260]"
2026/09/12 09:53:04 INFO Rigblock 8="[8 3 -1 512 2 1.0779 -0.0000 0.1459 0.0480 0.5779 0.0000 1.0000 0 1080123392 1729351260]"
2026/09/12 09:53:04 INFO Rigblock 9="[9 3 10 2304 2 1.0779 0.2018 0.0860 -0.1000 0.5779 0.0000 1.0000 0 1080123392 898723801]"
2026/09/12 09:53:04 INFO Rigblock 10="[10 3 9 2048 2 1.0779 -0.2019 0.0860 -0.1000 0.5779 0.0000 1.0000 0 1080123392 2053914431]"
*/

/* CREATURE STAGE
2026/09/12 12:06:27 INFO Received 18 rigblocks from client 127.0.0.1
2026/09/12 12:06:27 INFO Rigblock 0="[0 -1 -1 11 5 0.8170 0.0000 0.0968 0.8599 0.2797 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 1="[1 0 -1 9 5 0.7119 0.0000 0.1537 0.7616 0.2249 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 2="[2 1 -1 9 5 0.7325 0.0000 0.2004 0.6666 0.2357 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 3="[3 2 -1 9 5 0.9559 0.0000 0.2111 0.5618 0.3520 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 4="[4 3 -1 9 5 0.9936 0.0000 0.1863 0.4580 0.3717 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 5="[5 4 -1 9 5 1.1280 0.0000 0.1343 0.3634 0.4417 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 6="[6 5 -1 9 5 0.9696 0.0000 0.0617 0.2833 0.3592 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 7="[7 6 -1 9 5 0.7822 0.0000 -0.0282 0.2255 0.2615 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 8="[8 7 -1 9 5 0.4779 0.0000 -0.1294 0.2059 0.1031 0.0000 1.0000 0 1080188928 3495711651]"
2026/09/12 12:06:27 INFO Rigblock 9="[9 5 11 15 3 1.0000 0.2451 -0.0384 0.4762 0.2857 0.3250 0.6000 0 1080188928 3737371178]"
2026/09/12 12:06:27 INFO Rigblock 10="[10 9 12 13 3 1.0000 0.3460 -0.1719 0.3066 0.2857 0.1250 0.8000 0 1080188928 872205457]"
2026/09/12 12:06:27 INFO Rigblock 11="[11 5 9 15 3 1.0000 -0.2451 -0.0384 0.4762 0.2857 0.3250 0.6000 0 1080188928 3737371178]"
2026/09/12 12:06:27 INFO Rigblock 12="[12 11 10 13 3 1.0000 -0.3460 -0.1719 0.3066 0.2857 0.1250 0.8000 0 1080188928 872205457]"
2026/09/12 12:06:27 INFO Rigblock 13="[13 0 -1 1088 2 0.8607 0.0000 -0.1936 0.8534 0.2360 0.0000 1.0000 1014894446 1080188928 850171424]"
2026/09/12 12:06:27 INFO Rigblock 14="[14 0 15 256 2 0.6376 0.0894 -0.1616 0.8687 0.1563 0.0000 1.0000 0 1080188928 3963308926]"
2026/09/12 12:06:27 INFO Rigblock 15="[15 0 14 0 2 0.6376 -0.0896 -0.1619 0.8689 0.1563 0.0000 1.0000 0 1080188928 2116431088]"
2026/09/12 12:06:27 INFO Rigblock 16="[16 10 17 292 2 1.2523 0.2704 -0.0024 0.1490 0.5015 0.0000 1.0483 1249339683 1080188928 2139184273]"
2026/09/12 12:06:27 INFO Rigblock 17="[17 12 16 36 2 1.2523 -0.2704 -0.0024 0.1490 0.5015 0.0000 1.0483 1249339683 1080188928 1257409767]"
*/

// https://modapi-docs.sporecommunity.com/struct_editors_1_1c_creature_data_resource_1_1_rigblock_data.html
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

// RigblockDataType
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
			return fmt.Errorf("Failed to unmarshal rigblock value: %v", value)
		}
		str = string(bytes)
	}

	creature, err := NewRigblockFromString(str)
	if err != nil {
		return fmt.Errorf("Failed to parse creature string: %s", str)
	}

	*r = creature
	return nil
}

func NewRigblockFromString(s string) ([]Rigblock, error) {
	/* CELL
	|| ID: 0  PARENT: -1  || SYMMETRIC: -1 || FLAGS: 11   TYPE: 5 || SCALE: 0.4000 || POSITION: (-0.0000) (-0.2851) (0.0000) || SCALERELATIVE: 0.0000 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 3244096132 ||
	|| ID: 1  PARENT:  0  || SYMMETRIC: -1 || FLAGS: 9    TYPE: 5 || SCALE: 0.6000 || POSITION: (-0.0000) (-0.1734) (0.0000) || SCALERELATIVE: 0.2500 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 3244096132 ||
	|| ID: 2  PARENT:  1  || SYMMETRIC: -1 || FLAGS: 9    TYPE: 5 || SCALE: 1.0750 || POSITION: (-0.0000) (-0.0618) (0.0000) || SCALERELATIVE: 0.8438 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 3244096132 ||
	|| ID: 3  PARENT:  2  || SYMMETRIC: -1 || FLAGS: 9    TYPE: 5 || SCALE: 0.9750 || POSITION: (-0.0000) (0.0499)  (0.0000) || SCALERELATIVE: 0.7188 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0	      || GROUPID: 1080123392 INSTANCEID: 3244096132 ||
	|| ID: 4  PARENT:  3  || SYMMETRIC: -1 || FLAGS: 9    TYPE: 5 || SCALE: 0.6250 || POSITION: (-0.0000) (0.1616)  (0.0000) || SCALERELATIVE: 0.2812 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 3244096132 ||
	|| ID: 5  PARENT:  4  || SYMMETRIC: -1 || FLAGS: 9    TYPE: 5 || SCALE: 0.5625 || POSITION: (-0.0000) (0.2721)  (-0.0112)|| SCALERELATIVE: 0.2031 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 3244096132 ||
	|| ID: 6  PARENT:  0  || SYMMETRIC: -1 || FLAGS: 1088 TYPE: 2 || SCALE: 0.9277 || POSITION: (-0.0000) (-0.3648) (-0.0963)|| SCALERELATIVE: 0.4277 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 1014894446 || GROUPID: 1080123392 INSTANCEID: 2714007141 ||
	|| ID: 7  PARENT:  2  || SYMMETRIC: -1 || FLAGS: 512  TYPE: 2 || SCALE: 0.8607 || POSITION: (-0.0000) (-0.0390) (0.0634) || SCALERELATIVE: 0.3607 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 1729351260 ||
	|| ID: 8  PARENT:  3  || SYMMETRIC: -1 || FLAGS: 512  TYPE: 2 || SCALE: 1.0779 || POSITION: (-0.0000) (0.1459)  (0.0480) || SCALERELATIVE: 0.5779 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 1729351260 ||
	|| ID: 9  PARENT:  3  || SYMMETRIC: 10 || FLAGS: 2304 TYPE: 2 || SCALE: 1.0779 || POSITION: ( 0.2018) (0.0860)  (-0.1000)|| SCALERELATIVE: 0.5779 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 898723801  ||
	|| ID: 10 PARENT:  3  || SYMMETRIC: 9  || FLAGS: 2048 TYPE: 2 || SCALE: 1.0779 || POSITION: (-0.2019) (0.0860)  (-0.1000)|| SCALERELATIVE: 0.5779 || MUSCLESCALE: 0.0000 MUSCLESCALERELATIVE: 1.0000 || FOOTWEAPONORMOUTH: 0 	      || GROUPID: 1080123392 INSTANCEID: 2053914431 ||
	/*
		InstanceID        uint32        // uint32_t mInstanceID */
	// InstanceID need to store in the db, because the game needs to create the instance id?
	ids := strings.Split(s, "#")
	r := make([]Rigblock, 0, len(ids))
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
			return nil, fmt.Errorf("Failed to create a new rigblock from string, because the length of the parts: (%d) not supported!", partsLen)
		}
	}

	return r, nil
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

/*
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
*/

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

func String(rigblocks []Rigblock) string {
	entries := make([]string, len(rigblocks))
	for i, rb := range rigblocks {
		log.Info(rb.GetPartType())
		entries[i] = rb.string()
	}
	return strings.Join(entries, "#")
}
