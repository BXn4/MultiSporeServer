package simple

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Matrix3 [3][3]float32

func NewMatrix3() Matrix3 {
	return Matrix3{}
}

func (m *Matrix3) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case nil:
		*m = Matrix3{}
		return nil
	default:
		return fmt.Errorf("Failed to convert database value to bytes")
	}

	if str == "" {
		*m = Matrix3{}
		return nil
	}

	elements := strings.Split(str, "+")
	if len(elements) != 9 {
		return fmt.Errorf("Invalid matrix3 size: expected 9 elements, got %d", len(elements))
	}

	var matrix Matrix3
	for i := range 9 {
		num, err := strconv.ParseFloat(elements[i], 32)
		if err != nil {
			return fmt.Errorf("Failed to convert element to float: %v", err)
		}
		row := i / 3
		col := i % 3
		matrix[row][col] = float32(num)
	}

	*m = matrix
	return nil
}

func (m Matrix3) Value() (driver.Value, error) {
	elements := make([]string, 9)
	for i := range 3 {
		for j := range 3 {
			elements[i*3+j] = strconv.FormatFloat(float64(m[i][j]), 'f', 6, 32)
		}
	}

	return strings.Join(elements, "+"), nil
}

func (m Matrix3) String() string {
	elements := make([]string, 9)
	for i := range 3 {
		for j := range 3 {
			elements[i*3+j] = strconv.FormatFloat(float64(m[i][j]), 'f', 6, 32)
		}
	}

	return strings.Join(elements, "+")
}

func (m Matrix3) Get(row, col int) (float32, error) {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return 0, fmt.Errorf("Index out of bounds")
	}
	return m[row][col], nil
}

func (m *Matrix3) Set(row, col int, value float32) error {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return fmt.Errorf("Index out of bounds")
	}
	m[row][col] = value
	return nil
}
