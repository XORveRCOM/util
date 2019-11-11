package calledcheck

import (
	"testing"
)

func TestCalledByFID(t *testing.T) {
	fid := GetFunctionID()
	func() {
		if !CalledByFID(fid) {
			t.Error("calledby error")
		}
		if CalledByFID("xxx") {
			t.Error("calledby error")
		}
	}()
}

func TestCalledByPC(t *testing.T) {
	pc := GetCallerPC()
	func() {
		if !CalledByPC(pc) {
			t.Error("calledby error")
		}
		if CalledByPC(uintptr(0)) {
			t.Error("calledby error")
		}
	}()
}
