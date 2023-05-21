package tensor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTensor(t *testing.T) {
	shape := []int{2, 3, 4}
	t1 := NewTensor(shape...)
	t2 := NewTensor(shape...)

	assert.Equal(t, shape, t1.GetShape())
	assert.Equal(t, shape, t2.GetShape())

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				t1.SetValue(float64(i*j*k), i, j, k)
				t2.SetValue(float64(10+i*j*k), i, j, k)
			}
		}
	}

	assert.Equal(t, float64(2*3), t1.Slice(1).GetValue(2, 3))
	assert.Equal(t, float64(1*2*3), t1.Slice(1, 2).GetValue(3))

	t3 := t1.Flatten()
	assert.Equal(t, t1.Value(), t3.Value())
	assert.Equal(t, 3*4*2, t3.GetShape()[0])

	t.Log(t1.Add(t2))
	t.Log(t1.Sub(t2))
	t.Log(t1.Mul(t2))
	t.Log(t1.Mul(t2))
}
