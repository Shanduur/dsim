package fuse

import "fmt"

// FlagsArray is a custom type for flags
type FlagsArray []string

func (a FlagsArray) String() string {
	var argv []string

	for i := 0; i < len(a); i++ {
		argv = append(argv, a[i])
	}

	return fmt.Sprintf("%+v", argv)
}

// Set custom function for flags
func (a *FlagsArray) Set(value string) error {
	*a = append(*a, value)
	return nil
}
