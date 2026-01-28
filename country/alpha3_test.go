package country

import (
	"crypto/rand"
	"testing"
)

func BenchmarkAlpha3Code_UnmarshalJSON(b *testing.B) {
	var corpus = make([][]byte, 100)

	for i := 0; i < len(corpus); i++ {
		corpus[i] = make([]byte, 3)
		_, _ = rand.Read(corpus[i])
	}

	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var code Alpha3Code
		_ = code.UnmarshalJSON(corpus[i%len(corpus)])
	}
}
