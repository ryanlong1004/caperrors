package filters

import (
	"fmt"
	"os"
	"strings"
)

// Any holds a slice of values
type Any struct {
	values []string
}

// IsOK returns true if the value is in Any
func (a Any) IsOK(file os.FileInfo) bool {
	if (len(a.values)) <= 1 {
		return true
	}
	for _, needle := range a.values {
		if strings.Contains(absFilePath(file), needle) {
			fmt.Println(absFilePath(file))
			return true
		}
	}
	return false
}

// NewAny returns a new instance of Any
func NewAny(values []string) Any {
	return Any{
		values: values,
	}
}
