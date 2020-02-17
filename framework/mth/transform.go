package mth

import "math"

// PerspectiveRH ...
func PerspectiveRH(fovY, aspect, zNear, zFar float64) Mat4f {
	var p Mat4f
	tanHalfFovY := math.Tan(fovY / 2.)
	p.Set(0, 0, 1./(aspect*tanHalfFovY)).
		Set(1, 1, 1./tanHalfFovY).
		Set(2, 2, -(zFar+zNear)/(zFar-zNear)).
		Set(2, 3, -1.).
		Set(3, 2, (2.*zFar*zNear)/(zNear-zFar))
	return p
}

// PerspectiveLH ...
func PerspectiveLH(fovY, aspect, zNear, zFar float64) Mat4f {
	var p Mat4f
	tanHalfFovY := math.Tan(fovY / 2.)
	p.Set(0, 0, 1./(aspect*tanHalfFovY)).
		Set(1, 1, 1./tanHalfFovY).
		Set(2, 3, 1.).
		Set(2, 2, (zFar+zNear)/(zFar-zNear)).
		Set(3, 2, -(2.*zFar*zNear)/(zNear-zFar))
	return p
}

// OrthoRH ...
func OrthoRH(left, right, bottom, top, zNear, zFar float64) Mat4f {
	var o = NewUnitMat4f()
	o.Set(0, 0, 2./(right-left))
	o.Set(1, 1, 2./(top-bottom))
	o.Set(3, 0, -(right+left)/(right-left))
	o.Set(3, 1, -(top+bottom)/(top-bottom))
	o.Set(2, 2, 2./(zFar-zNear))
	o.Set(3, 2, -(zFar+zNear)/(zFar-zNear))
	return o
}

// OrthoLH ...
func OrthoLH(left, right, bottom, top, zNear, zFar float64) Mat4f {
	var o = NewUnitMat4f()
	o.Set(0, 0, 2./(right-left))
	o.Set(1, 1, 2./(top-bottom))
	o.Set(2, 2, -2./(zFar-zNear))
	o.Set(3, 0, -(right+left)/(right-left))
	o.Set(3, 1, -(top+bottom)/(top-bottom))
	o.Set(3, 2, -(zFar+zNear)/(zFar-zNear))
	return o
}
