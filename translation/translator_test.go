package translation_test

import (
	"testing"

	"github.com/fseba/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word     string
		language string
		want     string
	}{
		{
			name:     "hello in English",
			word:     "hello",
			language: "english",
			want:     "hello",
		},
		{
			name:     "hello in German",
			word:     "hello",
			language: "german",
			want:     "hallo",
		},
		{
			name:     "hello in Finish",
			word:     "hello",
			language: "finnish",
			want:     "hei",
		},
		{
			name:     "unknown language returns empty string",
			word:     "unknown",
			language: "unknown",
			want:     "",
		},
		{
			name:     "unsupported word returns empty string",
			word:     "goodbye",
			language: "german",
			want:     "",
		},
		{
			name:     "capitalized language returns correct translation",
			word:     "hello",
			language: "German",
			want:     "hallo",
		},
		{
			name:     "capitalized word returns correct translation",
			word:     "Hello",
			language: "german",
			want:     "hallo",
		},
		{
			name:     "space padded word returns correct translation",
			word:     "hello ",
			language: "german",
			want:     "hallo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translation.Translate(tt.word, tt.language)
			if got != tt.want {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
