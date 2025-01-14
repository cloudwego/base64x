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
    `encoding/base64`
    `strings`
    `testing`
)

type TestPair struct {
    decoded string
    encoded string
}

type EncodingTest struct {
    enc  Encoding            // Encoding to test
    conv func(string) string // Reference string converter
}

var pairs = []TestPair{
    // RFC 3548 examples
    {"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
    {"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
    {"\x14\xfb\x9c\x03", "FPucAw=="},

    // RFC 4648 examples
    {"", ""},
    {"f", "Zg=="},
    {"fo", "Zm8="},
    {"foo", "Zm9v"},
    {"foob", "Zm9vYg=="},
    {"fooba", "Zm9vYmE="},
    {"foobar", "Zm9vYmFy"},

    // Wikipedia examples
    {"sure.", "c3VyZS4="},
    {"sure", "c3VyZQ=="},
    {"sur", "c3Vy"},
    {"su", "c3U="},
    {"leasure.", "bGVhc3VyZS4="},
    {"easure.", "ZWFzdXJlLg=="},
    {"asure.", "YXN1cmUu"},
    {"sure.", "c3VyZS4="},

    // Relatively long strings
    {
        "Twas brillig, and the slithy toves",
        "VHdhcyBicmlsbGlnLCBhbmQgdGhlIHNsaXRoeSB0b3Zlcw==",
    }, {
        "\x9dyH\xd2Y\x9e^e\x9e\xb1\x9a\x9a\x12\xfe\x8a\x07\xc7\x07\xcc\xe8l\x81" +
        "\xf2\xd9\xe3\x89\xb5\x98\xee\xbd\x8etQ`2>\\t:_\xd7w\xe6\xb5\x96\xc7\xff\x9c",
        "nXlI0lmeXmWesZqaEv6KB8cHzOhsgfLZ44m1mO69jnRRYDI+XHQ6X9d35rWWx/+c",
    },
}

var crlf_pairs = []TestPair{
    // RFC 3548 examples
    {"\x14\xfb\x9c\x03\xd9\x7e", "FPuc\r\nA9l+"},
    {"\x14\xfb\x9c\x03\xd9", "FP\r\r\r\rucA9k="},
    {"\x14\xfb\x9c\x03", "\r\nFPucAw=\r=\n"},

    // RFC 4648 examples
    {"", "\r"},
    {"f", "Zg\r\n=="},
    {"fo", "Zm\r\n8="},
    {"fooba", "Zm\r\n9vY\r\nmE="},

    // Wikipedia examples
    {"su", "c3U\r="},
    {"leasure.", "bGVhc3VyZ\nS4="},
    {"easure.", "ZW\r\nFzdXJlLg=\r=\r\n"},
    {"asure.", "YXN1cmUu"},
    {"sure.", "c3VyZ\r\nS4="},

    // Relatively long strings
    {
        "Twas brillig, and the slithy toves",
        "VHdhcyBicmlsbGlnLCBhbmQgdGhlIHNsaXRoeSB0b3Zlcw\r\n==\r\n",
    }, {
        "\x9dyH\xd2Y\x9e^e\x9e\xb1\x9a\x9a\x12\xfe\x8a\x07\xc7\x07\xcc\xe8l\x81" +
        "\xf2\xd9\xe3\x89\xb5\x98\xee\xbd\x8etQ`2>\\t:_\xd7w\xe6\xb5\x96\xc7\xff\x9c",
        "nXlI0lmeXmWesZqaEv6KB8cHzOhsg\r\nfLZ44m1mO69jnRRYDI+XH\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\nQ6X9d35rWWx/\r\n+c",
    },
}

var json_pairs = []TestPair{
    // RFC 3548 examples
    {"\x14\xfb\x9c\x03\xd9\x7e", `FPu\rcA9l+\n`},
    {"\x14\xfb\x9c\x03\xd9\x7e", `FPuc\u00419l+`},
    {"\x14\xfb\x9c\x03\xd9", `FPucA9k\u003d`},
    {"\x14\xfb\x9c\x03\xd9", `FPucA\u0039k\u003d`},
    {"\x14\xfb\x9c\x03", `FPucAw\u003d\u003d`},

    // RFC 4648 examples
    {"", ""},
    {"f", "Zg=="},
    {"fo", "Zm8="},
    {"foo", "Zm9v"},
    {"foob", "Zm9vYg=="},
    {"fooba", "Zm9vYmE="},
    {"foobar", "Zm9vYmFy"},

    // Wikipedia examples
    {"sure.", "c3VyZS4="},
    {"sure", "c3VyZQ=="},
    {"sur", "c3Vy"},
    {"su", "c3U="},
    {"leasure.", "bGVhc3VyZS4="},
    {"easure.", "ZWFzdXJlLg=="},
    {"asure.", "YXN1cmUu"},
    {"sure.", "c3VyZS4="},

    // Relatively long strings
    {
        "Twas brillig, and the slithy toves",
        "VHdhcyBicmlsbGlnLCBhbmQgdGhlIHNsaXRoeSB0b3Zlcw==",
    }, {
        "\x9dyH\xd2Y\x9e^e\x9e\xb1\x9a\x9a\x12\xfe\x8a\x07\xc7\x07\xcc\xe8l\x81" +
        "\xf2\xd9\xe3\x89\xb5\x98\xee\xbd\x8etQ`2>\\t:_\xd7w\xe6\xb5\x96\xc7\xff\x9c",
        `nXlI0lmeXmWesZqaEv6KB8cHzOhsgfLZ44m1mO\u0036\u0039jnRRYDI+XHQ6X9d35rWWx\/+c`,
    },
}

// Do nothing to a reference base64 string (leave in standard format)
func stdRef(ref string) string {
    return ref
}

// Convert a reference string to URL-encoding
func urlRef(ref string) string {
    ref = strings.ReplaceAll(ref, "+", "-")
    ref = strings.ReplaceAll(ref, "/", "_")
    return ref
}

// Convert a reference string to raw, unpadded format
func rawRef(ref string) string {
    return strings.ReplaceAll(ref, "=", "")
}

// Both URL and unpadding conversions
func rawURLRef(ref string) string {
    return rawRef(urlRef(ref))
}

var encodingTests = []EncodingTest{
    {StdEncoding, stdRef},
    {URLEncoding, urlRef},
    {RawStdEncoding, rawRef},
    {RawURLEncoding, rawURLRef},
}

func testEqual(t *testing.T, msg string, args ...interface{}) bool {
    t.Helper()
    if args[len(args) - 2] != args[len(args) - 1] {
        t.Errorf(msg, args...)
        return false
    }
    return true
}

func TestEncoder(t *testing.T) {
    for _, p := range pairs {
        for _, tt := range encodingTests {
            got := tt.enc.EncodeToString([]byte(p.decoded))
            testEqual(t, "Encode(%q) = %q, want %q", p.decoded, got, tt.conv(p.encoded))
        }
    }
}

func TestDecoder(t *testing.T) {
    for _, p := range pairs {
        for _, tt := range encodingTests {
            encoded := tt.conv(p.encoded)
            dbuf := make([]byte, tt.enc.DecodedLen(len(encoded)))
            count, err := tt.enc.Decode(dbuf, []byte(encoded))
            testEqual(t, "Decode(%q) = error %v, want %v", encoded, err, error(nil))
            testEqual(t, "Decode(%q) = length %v, want %v", encoded, count, len(p.decoded))
            testEqual(t, "Decode(%q) = %q, want %q", encoded, string(dbuf[0:count]), p.decoded)

            dbuf, err = tt.enc.DecodeString(encoded)
            testEqual(t, "DecodeString(%q) = error %v, want %v", encoded, err, error(nil))
            testEqual(t, "DecodeString(%q) = %q, want %q", encoded, string(dbuf), p.decoded)
        }
    }
}


func TestDecoderCRLF(t *testing.T) {
    for _, p := range crlf_pairs {
        for _, tt := range encodingTests {
            encoded := tt.conv(p.encoded)
            dbuf := make([]byte, tt.enc.DecodedLen(len(encoded)))
            count, err := tt.enc.Decode(dbuf, []byte(encoded))
            testEqual(t, "Decode(%q) = error %v, want %v", encoded, err, error(nil))
            testEqual(t, "Decode(%q) = length %v, want %v", encoded, count, len(p.decoded))
            testEqual(t, "Decode(%q) = %q, want %q", encoded, string(dbuf[0:count]), p.decoded)

            dbuf, err = tt.enc.DecodeString(encoded)
            testEqual(t, "DecodeString(%q) = error %v, want %v", encoded, err, error(nil))
            testEqual(t, "DecodeString(%q) = %q, want %q", encoded, string(dbuf), p.decoded)
        }
    }
}

func TestDecoderJSON(t *testing.T) {
    for _, p := range json_pairs {
        encoded := p.encoded
        dbuf := make([]byte, JSONStdEncoding.DecodedLen(len(encoded)))
        count, err := JSONStdEncoding.Decode(dbuf, []byte(encoded))
        testEqual(t, "Decode(%q) = error %v, want %v", encoded, err, error(nil))
        testEqual(t, "Decode(%q) = length %v, want %v", encoded, count, len(p.decoded))
        testEqual(t, "Decode(%q) = %q, want %q", encoded, string(dbuf[0:count]), p.decoded)

        dbuf, err = JSONStdEncoding.DecodeString(encoded)
        testEqual(t, "DecodeString(%q) = error %v, want %v", encoded, err, error(nil))
        testEqual(t, "DecodeString(%q) = %q, want %q", encoded, string(dbuf), p.decoded)
    }
}

func TestDecoderError(t *testing.T) {
    _, err := StdEncoding.DecodeString("!aGVsbG8sIHdvcmxk")
    if err != base64.CorruptInputError(0) {
        panic(err)
    }
    _, err = StdEncoding.DecodeString("aGVsbG8!sIHdvcmxk")
    if err != base64.CorruptInputError(7) {
        panic(err)
    }
    _, err = StdEncoding.DecodeString("123456")
    if err != base64.CorruptInputError(6) {
        panic(err)
    }
    _, err = StdEncoding.DecodeString("1234;6")
    if err != base64.CorruptInputError(4) {
        panic(err)
    }
    _, err = StdEncoding.DecodeString("F\xaa\xaa\xaa\xaaDDDDDDDDDDDDD//z")
    if err != base64.CorruptInputError(1) {
        panic(err)
    } 
}
