package climax

import (
	"time"
)

// ParseAddress accepts a string and returns the named type as defined in in this source file if possible, and indicates success with a return value of true
func ParseAddress(s string, v interface{}) (parsed bool) {
	return
}

// ParseBool accepts a string and returns the named type as defined in in this source file if possible, and indicates success with a return value of true
func ParseBool(s string, v interface{}) (parsed bool) {
	return
}

// ParseDuration accepts a string and returns the named type as defined in in this source file if possible, and indicates success with a return value of true
func ParseDuration(s string, v interface{}) (parsed bool) {
	return
}

// ParseFloat accepts a string and returns the named type as defined in in this source file if possible, and indicates success with a return value of true
func ParseFloat(s string, v interface{}) (parsed bool) {
	return
}

// ParseInt accepts a string and returns the named type as defined in in this source file if possible, and indicates success with a return value of true
func ParseInt(s string, v interface{}) (parsed bool) {
	return
}

// ParseString accepts a string and returns the named type as defined in in this source file if possible, and indicates success with a return value of true
func ParseString(s string, v interface{}) (parsed bool) {
	return
}

// ParseURL accepts a string and returns the named type as defined in in this source file if possible, and indicates success with a return value of true
func ParseURL(s string, v interface{}) (parsed bool) {
	return
}

// ToAddress takes a variable of the builtin type the alias defined in this file either pointer or by value, and returns the variable type-converted to the alias
func ToAddress(in interface{}) Address {
	switch at := in.(type) {
	case string:
		return &at
	case *string:
		return at
	}
	return nil
}

// ToBool takes a variable of the builtin type the alias defined in this file either pointer or by value, and returns the variable type-converted to the alias
func ToBool(in interface{}) Bool {
	switch at := in.(type) {
	case *bool:
		return at
	case bool:
		return &at
	}
	return nil
}

// ToDuration takes a variable of the builtin type the alias defined in this file either pointer or by value, and returns the variable type-converted to the alias
func ToDuration(in interface{}) Duration {
	switch at := in.(type) {
	case *time.Duration:
		return at
	case time.Duration:
		return &at
	}
	return nil
}

// ToFloat takes a variable of the builtin type the alias defined in this file either pointer or by value, and returns the variable type-converted to the alias
func ToFloat(in interface{}) Float {
	switch at := in.(type) {
	case *float64:
		return at
	case float64:
		return &at
	}
	return nil
}

// ToInt takes a variable of the builtin type the alias defined in this file either pointer or by value, and returns the variable type-converted to the alias
func ToInt(in interface{}) Int {
	switch at := in.(type) {
	case *int:
		return at
	case int:
		return &at
	}
	return nil
}

// ToString takes a variable of the builtin type the alias defined in this file either pointer or by value, and returns the variable type-converted to the alias
func ToString(in interface{}) String {
	switch at := in.(type) {
	case string:
		return &at
	case *string:
		return at
	}
	return nil
}

// ToURL takes a variable of the builtin type the alias defined in this file either pointer or by value, and returns the variable type-converted to the alias
func ToURL(in interface{}) URL {
	switch at := in.(type) {
	case string:
		return &at
	case *string:
		return at
	}
	return nil
}
