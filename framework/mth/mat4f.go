package mth

import "github.com/sirupsen/logrus"

const (
	mat4frowlen int = 4
	mat4flen        = mat4frowlen * mat4frowlen
)

type (
	// Mat4f ...
	Mat4f struct {
		val [mat4flen]float32
	}
)

// NewUnitMat4f ...
func NewUnitMat4f() Mat4f {
	var val [mat4flen]float32
	for i := 0; i < mat4frowlen; i++ {
		val[i*mat4frowlen+i] = 1
	}
	return Mat4f{
		val: val,
	}
}

// Set ...
func (m *Mat4f) Set(i, j int, val float64) *Mat4f {
	if i < 0 ||
		j < 0 ||
		i >= mat4frowlen ||
		j >= mat4frowlen {
		logrus.Panic(mathInfoTag + ": invalid idx")
		return m
	}
	m.val[i*mat4frowlen+j] = float32(val)
	return m
}

func isCompatible(m, n Mat4f) bool {
	return len(m.val) == len(n.val) && len(m.val) == mat4flen
}

// Mult ...
func (m Mat4f) Mult(n Mat4f) Mat4f {
	var result [mat4flen]float32
	if !isCompatible(m, n) {
		logrus.Panic(matInfoTag + ": mult args aren't compatible")
	}

	for i := 0; i < mat4frowlen; i++ {
		for j := 0; j < mat4frowlen; j++ {
			for k := 0; k < mat4frowlen; k++ {
				result[i*mat4frowlen+j] += m.val[i*mat4frowlen+k] * n.val[k*mat4frowlen+j]
			}
		}
	}

	return Mat4f{val: result}
}

// Rotate ...
func (m *Mat4f) Rotate(rm [3][3]float64) *Mat4f {
	n := NewUnitMat4f()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			n.val[i*mat4frowlen+j] = float32(rm[i][j])
		}
	}
	*m = m.Mult(n)
	return m
}

// Values ...
func (m *Mat4f) Values() [mat4flen]float32 {
	return m.val
}
