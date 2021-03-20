package xcompare

//a==b
func equal(a, b interface{}) bool {
	switch a.(type) {
	case string:
		return a.(string) == b.(string)
	case int:
		return a.(int) == b.(int)
	case uint:
		return a.(uint) == b.(uint)
	case int8:
		return a.(int8) == b.(int8)
	case uint8:
		return a.(uint8) == b.(uint8)
	case int16:
		return a.(int16) == b.(int16)
	case uint16:
		return a.(uint16) == b.(uint16)
	case int32:
		return a.(int32) == b.(int32)
	case uint32:
		return a.(uint32) == b.(uint32)
	case int64:
		return a.(int64) == b.(int64)
	case uint64:
		return a.(uint64) == b.(uint64)
	case float32:
		return a.(float32) == b.(float32)
	case float64:
		return a.(float64) == b.(float64)
	}
	return false
}

//a!=b
func notEqual(a, b interface{}) bool {
	switch a.(type) {
	case string:
		return a.(string) != b.(string)
	case int:
		return a.(int) != b.(int)
	case uint:
		return a.(uint) != b.(uint)
	case int8:
		return a.(int8) != b.(int8)
	case uint8:
		return a.(uint8) != b.(uint8)
	case int16:
		return a.(int16) != b.(int16)
	case uint16:
		return a.(uint16) != b.(uint16)
	case int32:
		return a.(int32) != b.(int32)
	case uint32:
		return a.(uint32) != b.(uint32)
	case int64:
		return a.(int64) != b.(int64)
	case uint64:
		return a.(uint64) != b.(uint64)
	case float32:
		return a.(float32) != b.(float32)
	case float64:
		return a.(float64) != b.(float64)
	}
	return false
}

//a>b
func great(a, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) > b.(int)
	case uint:
		return a.(uint) > b.(uint)
	case int8:
		return a.(int8) > b.(int8)
	case uint8:
		return a.(uint8) > b.(uint8)
	case int16:
		return a.(int16) > b.(int16)
	case uint16:
		return a.(uint16) > b.(uint16)
	case int32:
		return a.(int32) > b.(int32)
	case uint32:
		return a.(uint32) > b.(uint32)
	case int64:
		return a.(int64) > b.(int64)
	case uint64:
		return a.(uint64) > b.(uint64)
	case float32:
		return a.(float32) > b.(float32)
	case float64:
		return a.(float64) > b.(float64)
	}
	return false
}

//a>=b
func greatEqual(a, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) >= b.(int)
	case uint:
		return a.(uint) >= b.(uint)
	case int8:
		return a.(int8) >= b.(int8)
	case uint8:
		return a.(uint8) >= b.(uint8)
	case int16:
		return a.(int16) >= b.(int16)
	case uint16:
		return a.(uint16) >= b.(uint16)
	case int32:
		return a.(int32) >= b.(int32)
	case uint32:
		return a.(uint32) >= b.(uint32)
	case int64:
		return a.(int64) >= b.(int64)
	case uint64:
		return a.(uint64) >= b.(uint64)
	case float32:
		return a.(float32) >= b.(float32)
	case float64:
		return a.(float64) >= b.(float64)
	}
	return false
}

//a<b
func little(a, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case uint:
		return a.(uint) < b.(uint)
	case int8:
		return a.(int8) < b.(int8)
	case uint8:
		return a.(uint8) < b.(uint8)
	case int16:
		return a.(int16) < b.(int16)
	case uint16:
		return a.(uint16) < b.(uint16)
	case int32:
		return a.(int32) < b.(int32)
	case uint32:
		return a.(uint32) < b.(uint32)
	case int64:
		return a.(int64) < b.(int64)
	case uint64:
		return a.(uint64) < b.(uint64)
	case float32:
		return a.(float32) < b.(float32)
	case float64:
		return a.(float64) < b.(float64)
	}
	return false
}

//a<=b
func littleEqual(a, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) <= b.(int)
	case uint:
		return a.(uint) <= b.(uint)
	case int8:
		return a.(int8) <= b.(int8)
	case uint8:
		return a.(uint8) <= b.(uint8)
	case int16:
		return a.(int16) <= b.(int16)
	case uint16:
		return a.(uint16) <= b.(uint16)
	case int32:
		return a.(int32) <= b.(int32)
	case uint32:
		return a.(uint32) <= b.(uint32)
	case int64:
		return a.(int64) <= b.(int64)
	case uint64:
		return a.(uint64) <= b.(uint64)
	case float32:
		return a.(float32) <= b.(float32)
	case float64:
		return a.(float64) <= b.(float64)
	}
	return false
}

//a in b(x,x,x,x)
func in(a, b interface{}) bool {
	switch b.(type) {
	case []string:
		av := a.(string)
		for _, bv := range b.([]string) {
			if av == bv {
				return true
			}
		}
	case []int:
		av := a.(int)
		for _, bv := range b.([]int) {
			if av == bv {
				return true
			}
		}
	case []uint:
		av := a.(uint)
		for _, bv := range b.([]uint) {
			if av == bv {
				return true
			}
		}
	case []int8:
		av := a.(int8)
		for _, bv := range b.([]int8) {
			if av == bv {
				return true
			}
		}
	case []uint8:
		av := a.(uint8)
		for _, bv := range b.([]uint8) {
			if av == bv {
				return true
			}
		}
	case []int16:
		av := a.(int16)
		for _, bv := range b.([]int16) {
			if av == bv {
				return true
			}
		}
	case []uint16:
		av := a.(uint16)
		for _, bv := range b.([]uint16) {
			if av == bv {
				return true
			}
		}
	case []int32:
		av := a.(int32)
		for _, bv := range b.([]int32) {
			if av == bv {
				return true
			}
		}
	case []uint32:
		av := a.(uint32)
		for _, bv := range b.([]uint32) {
			if av == bv {
				return true
			}
		}
	case []int64:
		av := a.(int64)
		for _, bv := range b.([]int64) {
			if av == bv {
				return true
			}
		}
	case []uint64:
		av := a.(uint64)
		for _, bv := range b.([]uint64) {
			if av == bv {
				return true
			}
		}
	case []float32:
		av := a.(float32)
		for _, bv := range b.([]float32) {
			if av == bv {
				return true
			}
		}
	case []float64:
		av := a.(float64)
		for _, bv := range b.([]float64) {
			if av == bv {
				return true
			}
		}
	}
	return false
}
