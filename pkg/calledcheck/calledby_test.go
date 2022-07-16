package calledcheck_test

import (
	"testing"

	cc "github.com/xorvercom/util/pkg/calledcheck"
)

func TestCalledByFID(t *testing.T) {
	fid := cc.GetFunctionID()
	func() {
		if !cc.CalledByFID(fid) {
			t.Error("calledby error")
		}
		if cc.CalledByFID("xxx") {
			t.Error("calledby error")
		}
	}()
}

func TestCalledByPC(t *testing.T) {
	pc := cc.GetCallerPC()
	t.Log(pc)
	func() {
		if !cc.CalledByPC(pc) {
			t.Error("calledby error")
		}
		if cc.CalledByPC(cc.CalledPC(0)) {
			t.Error("calledby error")
		}
	}()
}
