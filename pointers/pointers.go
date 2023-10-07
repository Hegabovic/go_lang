package pointers

var Age int = 27

// PointerAge : pointers syntax
var PointerAge *int = &Age

var DoubleAge = double(PointerAge)

func double(number *int) (result int) {
	result = *number * 2
	return
}
