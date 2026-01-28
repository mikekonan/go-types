package country

import (
	"crypto/rand"
	"testing"
)

func BenchmarkAlpha2Code_UnmarshalJSON(b *testing.B) {
	var corpus = make([][]byte, 100)

	for i := 0; i < len(corpus); i++ {
		corpus[i] = make([]byte, 4)
		corpus[i][0] = '"'
		corpus[i][3] = '"'
		_, _ = rand.Read(corpus[i][1:3])
	}

	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var code Alpha2Code
		_ = code.UnmarshalJSON(corpus[i%len(corpus)])
	}
}
