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

// An Encoding is a radix 64 encoding/decoding scheme, defined by a
// 64-character alphabet. The most common encoding is the "base64"
// encoding defined in RFC 4648 and used in MIME (RFC 2045) and PEM
// (RFC 1421).  RFC 4648 also defines an alternate encoding, which is
// the standard encoding with - and _ substituted for + and /.
type Encoding int

const (
	_MODE_URL  = 1 << 0
	_MODE_RAW  = 1 << 1
	_MODE_AVX2 = 1 << 2
	_MODE_JSON = 1 << 3
)

// StdEncoding is the standard base64 encoding, as defined in
// RFC 4648.
const StdEncoding Encoding = 0

// URLEncoding is the alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
const URLEncoding Encoding = _MODE_URL

// RawStdEncoding is the standard raw, unpadded base64 encoding,
// as defined in RFC 4648 section 3.2.
//
// This is the same as StdEncoding but omits padding characters.
const RawStdEncoding Encoding = _MODE_RAW

// RawURLEncoding is the unpadded alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
//
// This is the same as URLEncoding but omits padding characters.
const RawURLEncoding Encoding = _MODE_RAW | _MODE_URL

// JSONStdEncoding is the StdEncoding and encoded as JSON string as RFC 8259.
const JSONStdEncoding Encoding = _MODE_JSON

var (
	archFlags = 0
)
