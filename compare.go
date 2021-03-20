package xcompare

type inOperator struct{}

func (inOperator) String(a string, b []string) bool {
	return in(a, b)
}
func (inOperator) Int32(a int32, b []int32) bool {
	return in(a, b)
}
func (inOperator) Int64(a int64, b []int64) bool {
	return in(a, b)
}
func (inOperator) Float32(a float32, b []float32) bool {
	return in(a, b)
}
func (inOperator) Float64(a float64, b []float64) bool {
	return in(a, b)
}

type equalOperator struct{}

func (equalOperator) String(a, b string) bool {
	return equal(a, b)
}
func (equalOperator) Int32(a, b int32) bool {
	return equal(a, b)
}
func (equalOperator) Int64(a, b int64) bool {
	return equal(a, b)
}
func (equalOperator) Float32(a, b float32) bool {
	return equal(a, b)
}
func (equalOperator) Float64(a, b float64) bool {
	return equal(a, b)
}

type notEqualOperator struct{}

func (notEqualOperator) String(a, b string) bool {
	return notEqual(a, b)
}
func (notEqualOperator) Int32(a, b int32) bool {
	return notEqual(a, b)
}
func (notEqualOperator) Int64(a, b int64) bool {
	return notEqual(a, b)
}
func (notEqualOperator) Float32(a, b float32) bool {
	return notEqual(a, b)
}
func (notEqualOperator) Float64(a, b float64) bool {
	return notEqual(a, b)
}

type litterOperator struct{}

func (litterOperator) Int32(a, b int32) bool {
	return little(a, b)
}
func (litterOperator) Int64(a, b int64) bool {
	return little(a, b)
}
func (litterOperator) Float32(a, b float32) bool {
	return little(a, b)
}
func (litterOperator) Float64(a, b float64) bool {
	return little(a, b)
}

type litterEqualOperator struct{}

func (litterEqualOperator) Int32(a, b int32) bool {
	return littleEqual(a, b)
}
func (litterEqualOperator) Int64(a, b int64) bool {
	return littleEqual(a, b)
}
func (litterEqualOperator) Float32(a, b float32) bool {
	return littleEqual(a, b)
}
func (litterEqualOperator) Float64(a, b float64) bool {
	return littleEqual(a, b)
}

type greatOperator struct{}

func (greatOperator) Int32(a, b int32) bool {
	return great(a, b)
}
func (greatOperator) Int64(a, b int64) bool {
	return great(a, b)
}
func (greatOperator) Float32(a, b float32) bool {
	return great(a, b)
}
func (greatOperator) Float64(a, b float64) bool {
	return great(a, b)
}

type greatEqualOperator struct{}

func (greatEqualOperator) Int32(a, b int32) bool {
	return greatEqual(a, b)
}
func (greatEqualOperator) Int64(a, b int64) bool {
	return greatEqual(a, b)
}
func (greatEqualOperator) Float32(a, b float32) bool {
	return greatEqual(a, b)
}
func (greatEqualOperator) Float64(a, b float64) bool {
	return greatEqual(a, b)
}
