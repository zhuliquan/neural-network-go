package tensor

type Tensor interface {
	GetShape() []int          // 获取tensor的形状
	Slice(...int) Tensor      // 获取tensor的切片
	GetValue(...int) float64  // 获取坐标为(i,j,k,...)的值
	SetValue(float64, ...int) // 给坐标(i,j,k,...)的数据赋值
	// 以下操作必须是同维度运算
	Add(Tensor) Tensor // tensor 相加
	Sub(Tensor) Tensor // tensor 相减
	Mul(Tensor) Tensor // tensor 点乘
	Div(Tensor) Tensor // tensor 点除

	Flatten() Tensor // 将tensor 转化为一维向量

	Value() []float64 // 返回内部的value
}

type matrix struct {
	shape []int
	value []float64
	size  []int
}

func NewTensor(shape ...int) Tensor {
	t := new(matrix)
	t.shape = shape
	n := 1
	for _, sh := range shape {
		n *= sh
	}
	t.value = make([]float64, n)

	m := len(shape)
	t.size = make([]int, m)

	for i := m - 1; i >= 0; i-- {
		if i == m-1 {
			t.size[i] = 1
		} else {
			t.size[i] = t.shape[i+1] * t.size[i+1]
		}
	}
	return t
}

func (t *matrix) GetShape() []int {
	return t.shape
}

func (t *matrix) GetValue(index ...int) float64 {
	return t.value[docMul(t.size, index)]
}

func (t *matrix) SetValue(val float64, index ...int) {
	t.value[docMul(t.size, index)] = val
}

func (t *matrix) Slice(slice ...int) Tensor {
	if len(slice) == len(t.shape) {
		return &matrix{
			value: []float64{t.GetValue(slice...)},
			shape: []int{1},
			size:  []int{1},
		}
	}
	i := docMul(t.size[:len(slice)], slice)
	j := i + prodInt(t.shape[len(slice):])
	return &matrix{
		value: t.value[i:j],
		size:  t.size[len(slice):],
		shape: t.shape[len(slice):],
	}
}

func (t *matrix) Add(v Tensor) Tensor {
	v1 := t.Value()
	v2 := v.Value()
	v3 := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		v3[i] = v1[i] + v2[i]
	}
	return &matrix{
		shape: t.GetShape(),
		size:  t.size,
		value: v3,
	}
}

func (t *matrix) Sub(v Tensor) Tensor {
	v1 := t.Value()
	v2 := v.Value()
	v3 := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		v3[i] = v1[i] - v2[i]
	}
	return &matrix{
		shape: t.GetShape(),
		size:  t.size,
		value: v3,
	}
}

func (t *matrix) Mul(v Tensor) Tensor {
	v1 := t.Value()
	v2 := v.Value()
	v3 := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		v3[i] = v1[i] * v2[i]
	}
	return &matrix{
		shape: t.GetShape(),
		size:  t.size,
		value: v3,
	}
}

func (t *matrix) Div(v Tensor) Tensor {
	v1 := t.Value()
	v2 := v.Value()
	v3 := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		v3[i] = v1[i] / v2[i]
	}
	return &matrix{
		shape: t.GetShape(),
		size:  t.size,
		value: v3,
	}
}

func (t *matrix) Flatten() Tensor {
	return &matrix{
		shape: []int{len(t.value)},
		size:  []int{1},
		value: t.value,
	}

}

func (t *matrix) Value() []float64 {
	return t.value
}
