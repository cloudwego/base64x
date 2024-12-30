/*
 * Copyright 2024 CloudWeGo Authors
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

package base64x

import "encoding/base64"

/** Encoder Functions **/

// Encode encodes src using the specified encoding, writing
// EncodedLen(len(src)) bytes to out.
//
// The encoding pads the output to a multiple of 4 bytes,
// so Encode is not appropriate for use on individual blocks
// of a large data stream.
//
// If out is not large enough to contain the encoded result,
// it will panic.
func (self Encoding) Encode(out []byte, src []byte) {
	if len(src) != 0 {
		if buf := out[:0:len(out)]; self.EncodedLen(len(src)) <= len(out) {
			self.EncodeUnsafe(&buf, src)
		} else {
			panic("encoder output buffer is too small")
		}
	}
}

// EncodeUnsafe behaves like Encode, except it does NOT check if
// out is large enough to contain the encoded result.
//
// It will also update the length of out.
func (self Encoding) EncodeUnsafe(out *[]byte, src []byte) {
	b64encode(out, &src, int(self)|archFlags)
}

// EncodeToString returns the base64 encoding of src.
func (self Encoding) EncodeToString(src []byte) string {
	nbs := len(src)
	ret := make([]byte, 0, self.EncodedLen(nbs))

	/* encode in native code */
	self.EncodeUnsafe(&ret, src)
	return mem2str(ret)
}

// EncodedLen returns the length in bytes of the base64 encoding
// of an input buffer of length n.
func (self Encoding) EncodedLen(n int) int {
	if (self & _MODE_RAW) == 0 {
		return (n + 2) / 3 * 4
	} else {
		return (n*8 + 5) / 6
	}
}

/** Decoder Functions **/

// Decode decodes src using the encoding enc. It writes at most
// DecodedLen(len(src)) bytes to out and returns the number of bytes
// written. If src contains invalid base64 data, it will return the
// number of bytes successfully written and base64.CorruptInputError.
//
// New line characters (\r and \n) are ignored.
//
// If out is not large enough to contain the encoded result,
// it will panic.
func (self Encoding) Decode(out []byte, src []byte) (int, error) {
	if len(src) == 0 {
		return 0, nil
	} else if buf := out[:0:len(out)]; self.DecodedLen(len(src)) <= len(out) {
		return self.DecodeUnsafe(&buf, src)
	} else {
		panic("decoder output buffer is too small")
	}
}

// DecodeUnsafe behaves like Decode, except it does NOT check if
// out is large enough to contain the decoded result.
//
// It will also update the length of out.
func (self Encoding) DecodeUnsafe(out *[]byte, src []byte) (int, error) {
	if n := b64decode(out, mem2addr(src), len(src), int(self)|archFlags); n >= 0 {
		return n, nil
	} else {
		return 0, base64.CorruptInputError(-n - 1)
	}
}

// DecodeString returns the bytes represented by the base64 string s.
func (self Encoding) DecodeString(s string) ([]byte, error) {
	src := str2mem(s)
	ret := make([]byte, 0, self.DecodedLen(len(s)))

	/* decode into the allocated buffer */
	if _, err := self.DecodeUnsafe(&ret, src); err != nil {
		return nil, err
	} else {
		return ret, nil
	}
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base64-encoded data.
func (self Encoding) DecodedLen(n int) int {
	if (self & _MODE_RAW) == 0 {
		return n / 4 * 3
	} else {
		return n * 6 / 8
	}
}
