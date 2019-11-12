// Package calledcheck は、現在の処理が特定の関数から呼ばれているのかをスタックトレースで調査します。
// 用途は再帰呼び出しのチェックのためです。
package calledcheck

import (
	"fmt"
	"runtime"
)

// FunctionID は関数を示す識別コードです。
type FunctionID string

// CalledPC は呼び出し元を示す識別コードです。
type CalledPC uintptr

// String はstringerです。
func (c CalledPC) String() string {
	pc := uintptr(c)
	fpc := runtime.FuncForPC(pc)
	n, l := fpc.FileLine(pc)
	return fmt.Sprintf("%s (%s:%d)", fpc.Name(), n, l)
}

// CalledByFID はスタックトレース上に FunctionID があるかを検査します。
func CalledByFID(fid FunctionID) bool {
	i := 1
	for {
		if pc, _, _, ok := runtime.Caller(i); ok {
			name := runtime.FuncForPC(pc).Name()
			if name == string(fid) {
				return true
			}
		} else {
			return false
		}
		i++
	}
}

// CalledByPC はスタックトレース上に pc があるかを検査します。
func CalledByPC(pc CalledPC) bool {
	ppc := uintptr(pc)
	i := 1
	for {
		if cpc, _, _, ok := runtime.Caller(i); ok {
			if ppc == cpc {
				return true
			}
		} else {
			return false
		}
		i++
	}
}

// GetFunctionID はそれを呼び出した関数の現在の FunctionID を取得します。
func GetFunctionID() FunctionID {
	/**
	if pc, _, _, ok := runtime.Caller(1); ok {
		return FunctionID(runtime.FuncForPC(pc).Name())
	}
	return ""
	*/
	pc, _, _, _ := runtime.Caller(1)
	return FunctionID(runtime.FuncForPC(pc).Name())
}

// GetCallerPC はそれを呼び出した関数の呼び出し元の PC を取得します。
func GetCallerPC() CalledPC {
	pc, _, _, _ := runtime.Caller(2)
	return CalledPC(pc)
}
