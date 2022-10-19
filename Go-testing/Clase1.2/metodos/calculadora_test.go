package metodos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	n1, n2, exp := 5, 12, -7
	current := Subtract(n1, n2)
	assert.Equal(t, exp, current)
}
func TestAdd(t *testing.T) {
	n1, n2, exp := 13, 52, 65
	current := Add(n1, n2)
	assert.Equal(t, exp, current)
}

func TestDivision(t *testing.T) {
	n1, n2, exp := 10, 2, 5
	current, err := division(n1, n2)
	assert.Equal(t, exp, current)
	assert.Nil(t, err)
}
func TestDivisionByZero(t *testing.T) {
	n1, n2, exp := 10, 0, 0
	current, err := division(n1, n2)
	assert.Equal(t, exp, current)
	assert.ErrorContains(t, err, NotZeroDivision.Error())
}
