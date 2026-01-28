package country

import (
	"crypto/rand"
	"testing"
)

func BenchmarkAlpha3Code_UnmarshalJSON(b *testing.B) {
	var corpus = make([][]byte, 100)

	for i := 0; i < len(corpus); i++ {
		corpus[i] = make([]byte, 5)
		corpus[i][0] = '"'
		corpus[i][4] = '"'
		_, _ = rand.Read(corpus[i][1:4])
	}

	b.SetBytes(3)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var code Alpha3Code
		_ = code.UnmarshalJSON(corpus[i%len(corpus)])
	}
}
