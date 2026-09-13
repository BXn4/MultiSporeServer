package types

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	X float32
	Y float32
	Z float32
}

func NewPosition(x, y, z float32) Position {
	return Position{X: x, Y: y, Z: z}
}

func (p *Position) Scan(value interface{}) error {
	var str string

	switch val := value.(type) {
	case []byte:
		str = string(val)
	case string:
		str = val
	case nil:
		*p = Position{}
		return nil
	default:
		return fmt.Errorf("Failed to convert database value to bytes")
	}

	if str == "" {
		*p = Position{}
		return nil
	}

	elements := strings.Split(str, "+")
	if len(elements) != 3 {
		return fmt.Errorf("Invalid position size: expected 3 elements, got %d", len(elements))
	}

	var pos Position
	fields := [3]*float32{&pos.X, &pos.Y, &pos.Z}
	for i := range 3 {
		num, err := strconv.ParseFloat(elements[i], 32)
		if err != nil {
			return fmt.Errorf("Failed to convert element to float: %v", err)
		}
		*fields[i] = float32(num)
	}

	*p = pos
	return nil
}

func (p Position) String() string {
	elements := []string{
		strconv.FormatFloat(float64(p.X), 'f', 6, 32),
		strconv.FormatFloat(float64(p.Y), 'f', 6, 32),
		strconv.FormatFloat(float64(p.Z), 'f', 6, 32),
	}

	return strings.Join(elements, "+")
}
