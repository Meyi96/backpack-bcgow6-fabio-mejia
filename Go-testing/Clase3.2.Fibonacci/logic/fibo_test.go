package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiboCaseOne(t *testing.T) {
	amount := 1
	expList := []int{0}

	result := Fibonacci(amount)

	assert.Equal(t, expList, result)
}
func TestFiboCaseTwo(t *testing.T) {
	amount := 2
	expList := []int{0, 1}

	result := Fibonacci(amount)

	assert.Equal(t, expList, result)
}
func TestFiboMoreThanTwo(t *testing.T) {
	amount := 15

	result := Fibonacci(amount)

	for i := range result {
		if i == 0 {
			assert.Equal(t, 0, result[i])
		} else if i == 1 {
			assert.Equal(t, 1, result[i])
		} else if i > 1 {
			temp := result[i-2] + result[i-1]
			assert.Equal(t, temp, result[i])
		}
	}
}
func TestFiboZeroNegativeSize(t *testing.T) {
	expOutput := []int{}

	resultZero := Fibonacci(0)
	resultNegative := Fibonacci(-3)

	assert.Equal(t, expOutput, resultZero)
	assert.Equal(t, expOutput, resultNegative)

}
