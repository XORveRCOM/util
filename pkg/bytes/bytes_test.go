package bytes_test

import (
	"reflect"
	"testing"

	"github.com/xorvercom/util/pkg/bytes"
)

func Test_clearBytes(t *testing.T) {
	type args struct {
		size int
		init byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "size zero", args: args{size: 0, init: 0}, want: []byte{}},
		{name: "init one", args: args{size: 1, init: 1}, want: []byte{1}},
		{name: "init two", args: args{size: 2, init: 2}, want: []byte{2, 2}},
		{name: "init three", args: args{size: 3, init: 3}, want: []byte{3, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bytes.ClearBytes(tt.args.size, tt.args.init)
			t.Logf("got:%v", got)
			t.Logf("tt.want:%v", tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clearBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transpose(t *testing.T) {
	type args struct {
		src []byte
	}
	src1 := []byte{1 << 0, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7}
	want1 := [][]byte{{1 << 0}, {1 << 1}, {1 << 2}, {1 << 3}, {1 << 4}, {1 << 5}, {1 << 6}, {1 << 7}}
	src2 := []byte{1 << 7, 1 << 6, 1 << 5, 1 << 4, 1 << 3, 1 << 2, 1 << 1, 1 << 0}
	want2 := [][]byte{{1 << 7}, {1 << 6}, {1 << 5}, {1 << 4}, {1 << 3}, {1 << 2}, {1 << 1}, {1 << 0}}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{name: "test1", args: args{src: src1}, want: want1},
		{name: "test2", args: args{src: src2}, want: want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bytes.TransposeBits(tt.args.src)
			t.Logf("got:%v", got)
			t.Logf("tt.want:%v", tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
