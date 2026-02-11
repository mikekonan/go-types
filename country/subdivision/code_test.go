package subdivision

import (
	"crypto/rand"
	"testing"
)

func BenchmarkCode_UnmarshalJSON(b *testing.B) {
	var corpus = make([][]byte, 100)

	// generate random codes in format "XX-XX" (quoted JSON strings)
	for i := 0; i < len(corpus); i++ {
		corpus[i] = make([]byte, 7)
		corpus[i][0] = '"'
		corpus[i][6] = '"'
		_, _ = rand.Read(corpus[i][1:3])
		corpus[i][3] = '-'
		_, _ = rand.Read(corpus[i][4:6])
	}

	b.SetBytes(5)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var code Code
		_ = code.UnmarshalJSON(corpus[i%len(corpus)])
	}
}

func BenchmarkCode_ValidateForCountry(b *testing.B) {
	// collect all valid codes
	codes := make([]Code, 0, len(SubdivisionByCode))
	for _, sub := range SubdivisionByCode {
		codes = append(codes, sub.code)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sub := SubdivisionByCode[string(codes[i%len(codes)])]
		_ = codes[i%len(codes)].ValidateForCountry(sub.countryCode)
	}
}

func BenchmarkByCode(b *testing.B) {
	codes := make([]Code, 0, len(SubdivisionByCode))
	for _, sub := range SubdivisionByCode {
		codes = append(codes, sub.code)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ByCode(codes[i%len(codes)])
	}
}

func BenchmarkByCountryCode(b *testing.B) {
	countries := make([]string, 0, len(SubdivisionsByCountry))
	for cc := range SubdivisionsByCountry {
		countries = append(countries, cc)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ByCountryCodeStr(countries[i%len(countries)])
	}
}
