package metodos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	array := []int{2, 4, 64, 7, 4, 6, 3, 7}
	insercionSort(&array)
	for i := range array {
		if i < len(array)-1 {
			assert.True(t, array[i] <= array[i+1])
		}
	}
}
