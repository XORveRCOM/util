// Package calledcheck は、現在の処理が特定の関数から呼ばれているのかをスタックトレースで調査します。
// 用途は再帰呼び出しのチェックのためです。
package calledcheck

import (
	"runtime"
)

// FunctionID は関数を示す識別コードです。
type FunctionID string

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
func CalledByPC(pc uintptr) bool {
	i := 1
	for {
		if cpc, _, _, ok := runtime.Caller(i); ok {
			if pc == cpc {
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
func GetCallerPC() uintptr {
	pc, _, _, _ := runtime.Caller(2)
	return pc
}
