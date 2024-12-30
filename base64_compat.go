//go:build !amd64 || !go1.17 || go1.24
// +build !amd64 !go1.17 go1.24

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

import (
	"encoding/base64"
)

/** Encoder Functions **/

// Encode encodes src using the specified encoding, writing
// EncodedLen(len(src)) bytes to out.
//
// The encoding pads the output to a multiple of 4 bytes,
// so Encode is not appropriate for use on individual blocks
// of a large data stream.
func (self Encoding) Encode(out []byte, src []byte) {
	switch self {
	case 0, _MODE_JSON:
		base64.StdEncoding.Encode(out, src)
	case _MODE_URL:
		base64.URLEncoding.Encode(out, src)
	case _MODE_RAW:
		base64.RawStdEncoding.Encode(out, src)
	case _MODE_RAW | _MODE_URL:
		base64.RawURLEncoding.Encode(out, src)
	default:
		base64.StdEncoding.Encode(out, src)
	}
}

// EncodeUnsafe behaves like Encode
func (self Encoding) EncodeUnsafe(out *[]byte, src []byte) {
	self.Encode(*out, src)
}

// EncodeToString returns the base64 encoding of src.
func (self Encoding) EncodeToString(src []byte) string {
	out := make([]byte, self.EncodedLen(len(src)))
	self.Encode(out, src)
	return mem2str(out)
}

// EncodedLen returns the length in bytes of the base64 encoding
// of an input buffer of length n.
func (self Encoding) EncodedLen(n int) int {
	switch self {
	case 0, _MODE_JSON:
		return base64.StdEncoding.EncodedLen(n)
	case _MODE_URL:
		return base64.URLEncoding.EncodedLen(n)
	case _MODE_RAW:
		return base64.RawStdEncoding.EncodedLen(n)
	case _MODE_RAW | _MODE_URL:
		return base64.RawURLEncoding.EncodedLen(n)
	default:
		return base64.StdEncoding.EncodedLen(n)
	}
}

/** Decoder Functions **/

// Decode decodes src using the encoding enc. It writes at most
// DecodedLen(len(src)) bytes to out and returns the number of bytes
// written. If src contains invalid base64 data, it will return the
// number of bytes successfully written and base64.CorruptInputError.
//
// New line characters (\r and \n) are ignored.
func (self Encoding) Decode(out []byte, src []byte) (int, error) {
	switch self {
	case 0, _MODE_JSON:
		return base64.StdEncoding.Decode(out, src)
	case _MODE_URL:
		return base64.URLEncoding.Decode(out, src)
	case _MODE_RAW:
		return base64.RawStdEncoding.Decode(out, src)
	case _MODE_RAW | _MODE_URL:
		return base64.RawURLEncoding.Decode(out, src)
	default:
		return base64.StdEncoding.Decode(out, src)
	}
}

// DecodeUnsafe behaves like Decode
func (self Encoding) DecodeUnsafe(out *[]byte, src []byte) (int, error) {
	return self.Decode(*out, src)
}

// DecodeString returns the bytes represented by the base64 string s.
func (self Encoding) DecodeString(s string) ([]byte, error) {
	switch self {
	case 0, _MODE_JSON:
		return base64.StdEncoding.DecodeString(s)
	case _MODE_URL:
		return base64.URLEncoding.DecodeString(s)
	case _MODE_RAW:
		return base64.RawStdEncoding.DecodeString(s)
	case _MODE_RAW | _MODE_URL:
		return base64.RawURLEncoding.DecodeString(s)
	default:
		return base64.StdEncoding.DecodeString(s)
	}
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base64-encoded data.
func (self Encoding) DecodedLen(n int) int {
	switch self {
	case 0, _MODE_JSON:
		return base64.StdEncoding.DecodedLen(n)
	case _MODE_URL:
		return base64.URLEncoding.DecodedLen(n)
	case _MODE_RAW:
		return base64.RawStdEncoding.DecodedLen(n)
	case _MODE_RAW | _MODE_URL:
		return base64.RawURLEncoding.DecodedLen(n)
	default:
		return base64.StdEncoding.DecodedLen(n)
	}
}
