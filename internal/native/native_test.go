/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package native

import (
	`reflect`
	`testing`
	`unsafe`
)

func TestEncoderRecover(t *testing.T) {
    t.Run("nil dst", func(t *testing.T) {
        in := []byte("abc")
        defer func(){
            if v := recover(); v != nil {
                println("recover:", v)
            } else {
                t.Fatal("not recover")
            }
        }()
        B64Encode(nil, &in, 0)
    })
    t.Run("nil src", func(t *testing.T) {
        in := []byte("abc")
        (*reflect.SliceHeader)(unsafe.Pointer(&in)).Data = uintptr(0)
        out := make([]byte, 0, 10)
        defer func(){
            if v := recover(); v != nil {
                println("recover:", v)
            } else {
                t.Fatal("not recover")
            }
        }()
        B64Encode(&out, &in, 0)
    })
}


func TestDecoderRecover(t *testing.T) {
    t.Run("nil dst", func(t *testing.T) {
        in := []byte("abc")
        defer func(){
            if v := recover(); v != nil {
                println("recover:", v)
            } else {
                t.Fatal("not recover")
            }
        }()
        B64Decode(nil, unsafe.Pointer(&in[0]), len(in), 0)
    })
    t.Run("nil src", func(t *testing.T) {
        out := make([]byte, 0, 10)
        defer func(){
            if v := recover(); v != nil {
                println("recover:", v)
            } else {
                t.Fatal("not recover")
            }
        }()
        B64Decode(&out, nil, 5, 0)
    })
}