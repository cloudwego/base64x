//go:build !amd64 || noasm || appengine
// +build !amd64 noasm appengine

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
	"testing"
)

func missingPaddingDecodeError() base64.CorruptInputError {
	return base64.CorruptInputError(4)
}

func TestUnquoteJSONBase64(t *testing.T) {
	src := []byte("VHdhcyBicmlsbGlnLCBhbmQgdGhlIHNsaXRoeSB0b3Zlcw==")
	got, err := unquoteJSONBase64(src)
	if err != nil {
		t.Fatalf("unquoteJSONBase64() error = %v", err)
	}
	if string(got) != string(src) {
		t.Fatalf("unquoteJSONBase64() = %q, want %q", got, src)
	}
}

func TestUnquoteJSONBase64RejectsInvalidEscapes(t *testing.T) {
	for _, src := range [][]byte{
		[]byte(`FPuc\x419l+`),
		[]byte(`FPuc\'A9l+`),
	} {
		if _, err := unquoteJSONBase64(src); err == nil {
			t.Fatalf("unquoteJSONBase64(%q) error = nil, want non-nil", src)
		}
	}
}

func TestFallbackDecodeMissingPaddingReturnsStdlibError(t *testing.T) {
	_, err := StdEncoding.DecodeString("123456")
	if err != base64.CorruptInputError(4) {
		t.Fatalf("DecodeString(%q) error = %v, want %v", "123456", err, base64.CorruptInputError(4))
	}
}
