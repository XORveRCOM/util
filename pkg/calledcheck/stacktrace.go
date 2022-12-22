package calledcheck

import (
	"fmt"
	"runtime"
)

func StackTrace() []string {
	result := []string{}
	i := 1
	for {
		if pc, _, _, ok := runtime.Caller(i); ok {
			fpc := runtime.FuncForPC(pc)
			n, l := fpc.FileLine(pc)
			result = append(result, fmt.Sprintf("%s (%s:%d)", fpc.Name(), n, l))
		} else {
			return result
		}
		i++
	}
}
