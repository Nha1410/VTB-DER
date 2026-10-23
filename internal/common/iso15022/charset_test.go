package iso15022

import (
	"errors"
	"testing"
)

func TestValidateCharset(t *testing.T) {
	tests := []struct {
		name      string
		inputText string
		wantIndex int // -1 means valid
		wantChar  rune
	}{
		{name: "every allowed punctuation", inputText: "aZ09/-?:().,'+ \r\n", wantIndex: -1},
		{name: "vietnamese letter", inputText: "Nguyễn", wantIndex: 4, wantChar: 'ễ'},
		{name: "brace would break the block structure", inputText: "A}", wantIndex: 1, wantChar: '}'},
		{name: "tab", inputText: "A\tB", wantIndex: 1, wantChar: '\t'},
		{name: "underscore inside a converted slash", inputText: "A?_?B ?_38?", wantIndex: -1},
		{name: "bare underscore", inputText: "A_B", wantIndex: 1, wantChar: '_'},
		{name: "underscore after a closed marker", inputText: "?OS?_", wantIndex: 4, wantChar: '_'},
		{name: "empty", inputText: "", wantIndex: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCharset(tt.inputText)
			if tt.wantIndex == -1 {
				if err != nil {
					t.Fatalf("ValidateCharset(%q) = %v, want nil", tt.inputText, err)
				}
				return
			}
			var charsetErr *CharsetError
			if !errors.As(err, &charsetErr) {
				t.Fatalf("ValidateCharset(%q) = %v, want *CharsetError", tt.inputText, err)
			}
			if charsetErr.Index != tt.wantIndex || charsetErr.Char != tt.wantChar {
				t.Errorf("got %q at %d, want %q at %d", charsetErr.Char, charsetErr.Index, tt.wantChar, tt.wantIndex)
			}
		})
	}
}
