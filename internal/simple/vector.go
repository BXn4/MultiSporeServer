package simple

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Vector3 [3]float32

func NewVector3() Vector3 {
	return Vector3{}
}

func (v Vector3) X() float32 {
	return v[0]
}

func (v Vector3) Y() float32 {
	return v[1]
}

func (v Vector3) Z() float32 {
	return v[2]
}

func (v *Vector3) Scan(value interface{}) error {
	var str string

	switch val := value.(type) {
	case []byte:
		str = string(val)
	case string:
		str = val
	case nil:
		*v = Vector3{}
		return nil
	default:
		return fmt.Errorf("Failed to convert database value to bytes")
	}

	if str == "" {
		*v = Vector3{}
		return nil
	}

	elements := strings.Split(str, "+")
	if len(elements) != 3 {
		return fmt.Errorf("Invalid vector3 size: expected 3 elements, got %d", len(elements))
	}

	var vec Vector3
	for i := range 3 {
		num, err := strconv.ParseFloat(elements[i], 32)
		if err != nil {
			return fmt.Errorf("Failed to convert element to float: %v", err)
		}
		vec[i] = float32(num)
	}

	*v = vec
	return nil
}

func (v Vector3) Value() (driver.Value, error) {
	elements := make([]string, 3)
	for i := range 3 {
		elements[i] = strconv.FormatFloat(float64(v[i]), 'f', 6, 32)
	}

	return strings.Join(elements, "+"), nil
}

func (v Vector3) String() string {
	elements := make([]string, 3)
	for i := range 3 {
		elements[i] = strconv.FormatFloat(float64(v[i]), 'f', 6, 32)
	}

	return strings.Join(elements, "+")
}

func (v Vector3) Get(index int) (float32, error) {
	if index < 0 || index >= 3 {
		return 0, fmt.Errorf("Index out of bounds")
	}
	return v[index], nil
}

func (v *Vector3) Set(index int, value float32) error {
	if index < 0 || index >= 3 {
		return fmt.Errorf("Index out of bounds")
	}
	v[index] = value
	return nil
}
