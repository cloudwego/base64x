package bench

import (
	"testing"
	"encoding/base64"
	`crypto/rand`
	`io`

	. "github.com/cloudwego/base64x"
	cris "github.com/cristalhq/base64"
)

func benchmarkStdlibDecoder(b *testing.B, v string) {
    src := []byte(v)
    dst := make([]byte, base64.StdEncoding.DecodedLen(len(v)))
    b.SetBytes(int64(len(v)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = base64.StdEncoding.Decode(dst, src)
        }
    })
}

func benchmarkBase64xDecoder(b *testing.B, v string) {
    src := []byte(v)
    dst := make([]byte, StdEncoding.DecodedLen(len(v)))
    b.SetBytes(int64(len(v)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = StdEncoding.Decode(dst, src)
        }
    })
}


func benchmarkStdlibWithSize(b *testing.B, nb int) {
    buf := make([]byte, nb)
    dst := make([]byte, base64.StdEncoding.EncodedLen(nb))
    _, _ = io.ReadFull(rand.Reader, buf)
    b.SetBytes(int64(nb))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            base64.StdEncoding.Encode(dst, buf)
        }
    })
}

func benchmarkBase64xWithSize(b *testing.B, nb int) {
    buf := make([]byte, nb)
    dst := make([]byte, StdEncoding.EncodedLen(nb))
    _, _ = io.ReadFull(rand.Reader, buf)
    b.SetBytes(int64(nb))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            StdEncoding.Encode(dst, buf)
        }
    })
}


func benchmarkCrisWithSize(b *testing.B, nb int) {
    buf := make([]byte, nb)
    dst := make([]byte, cris.StdEncoding.EncodedLen(nb))
    _, _ = io.ReadFull(rand.Reader, buf)
    b.SetBytes(int64(nb))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            cris.StdEncoding.Encode(dst, buf)
        }
    })
}

func BenchmarkEncoderStdlib_16B    (b *testing.B) { benchmarkStdlibWithSize(b, 16) }
func BenchmarkEncoderStdlib_56B    (b *testing.B) { benchmarkStdlibWithSize(b, 56) }
func BenchmarkEncoderStdlib_128B   (b *testing.B) { benchmarkStdlibWithSize(b, 128) }
func BenchmarkEncoderStdlib_4kB    (b *testing.B) { benchmarkStdlibWithSize(b, 4 * 1024) }
func BenchmarkEncoderStdlib_256kB  (b *testing.B) { benchmarkStdlibWithSize(b, 256 * 1024) }
func BenchmarkEncoderStdlib_1MB    (b *testing.B) { benchmarkStdlibWithSize(b, 1024 * 1024) }

func BenchmarkEncoderBase64x_16B   (b *testing.B) { benchmarkBase64xWithSize(b, 16) }
func BenchmarkEncoderBase64x_56B   (b *testing.B) { benchmarkBase64xWithSize(b, 56) }
func BenchmarkEncoderBase64x_128B  (b *testing.B) { benchmarkBase64xWithSize(b, 128) }
func BenchmarkEncoderBase64x_4kB   (b *testing.B) { benchmarkBase64xWithSize(b, 4 * 1024) }
func BenchmarkEncoderBase64x_256kB (b *testing.B) { benchmarkBase64xWithSize(b, 256 * 1024) }
func BenchmarkEncoderBase64x_1MB   (b *testing.B) { benchmarkBase64xWithSize(b, 1024 * 1024) }

func BenchmarkEncoderCris_16B    (b *testing.B) { benchmarkCrisWithSize(b, 16) }
func BenchmarkEncoderCris_56B    (b *testing.B) { benchmarkCrisWithSize(b, 56) }
func BenchmarkEncoderCris_128B   (b *testing.B) { benchmarkCrisWithSize(b, 128) }
func BenchmarkEncoderCris_4kB    (b *testing.B) { benchmarkCrisWithSize(b, 4 * 1024) }
func BenchmarkEncoderCris_256kB  (b *testing.B) { benchmarkCrisWithSize(b, 256 * 1024) }
func BenchmarkEncoderCris_1MB    (b *testing.B) { benchmarkCrisWithSize(b, 1024 * 1024) }

var data = `////////////////////////////////////////////////////////////////`
func BenchmarkDecoderStdLib  (b *testing.B) { benchmarkStdlibDecoder(b, data) }
func BenchmarkDecoderBase64x (b *testing.B) { benchmarkBase64xDecoder(b, data) }
