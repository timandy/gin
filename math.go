package gin

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
//
// Note that floating-point types may contain NaN ("not-a-number") values.
// An operator such as == or < will always report false when
// comparing a NaN value with any other value, NaN or not.
// See the [Compare] function for a consistent way to compare NaN values.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// The max built-in function returns the largest value of a fixed number of
// arguments of [cmp.Ordered] types. There must be at least one argument.
// If T is a floating-point type and any of the arguments are NaNs,
// max will return NaN.
func max[T Ordered](x T, y ...T) T {
	var value = x
	for _, v := range y {
		if v > value {
			value = v
		}
	}
	return value
}

// The min built-in function returns the smallest value of a fixed number of
// arguments of [cmp.Ordered] types. There must be at least one argument.
// If T is a floating-point type and any of the arguments are NaNs,
// min will return NaN.
func min[T Ordered](x T, y ...T) T {
	var value = x
	for _, v := range y {
		if v < value {
			value = v
		}
	}
	return value
}
