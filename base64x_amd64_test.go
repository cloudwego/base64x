package base64x

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestEncoderRecover(t *testing.T) {
	t.Run("nil dst", func(t *testing.T) {
		in := []byte("abc")
		defer func() {
			if v := recover(); v != nil {
				println("recover:", v)
			} else {
				t.Fatal("not recover")
			}
		}()
		b64encode(nil, &in, int(StdEncoding))
	})
	t.Run("nil src", func(t *testing.T) {
		in := []byte("abc")
		(*reflect.SliceHeader)(unsafe.Pointer(&in)).Data = uintptr(0)
		out := make([]byte, 0, 10)
		defer func() {
			if v := recover(); v != nil {
				println("recover:", v)
			} else {
				t.Fatal("not recover")
			}
		}()
		b64encode(&out, &in, int(StdEncoding))
	})
}

func TestDecoderRecover(t *testing.T) {
	t.Run("nil dst", func(t *testing.T) {
		in := []byte("abc")
		defer func() {
			if v := recover(); v != nil {
				println("recover:", v)
			} else {
				t.Fatal("not recover")
			}
		}()
		b64decode(nil, unsafe.Pointer(&in[0]), len(in), int(StdEncoding))
	})
	t.Run("nil src", func(t *testing.T) {
		out := make([]byte, 0, 10)
		defer func() {
			if v := recover(); v != nil {
				println("recover:", v)
			} else {
				t.Fatal("not recover")
			}
		}()
		b64decode(&out, nil, 5, int(StdEncoding))
	})
}
