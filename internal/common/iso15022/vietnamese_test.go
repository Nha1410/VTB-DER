package iso15022

import (
	"errors"
	"os"
	"regexp"
	"strings"
	"testing"
)

// specTablePath is Table 2 of the spec as transcribed in docs/05-bang-chuyen-doi-ky-tu.md.
const specTablePath = "../../../docs/05-bang-chuyen-doi-ky-tu.md"

// markdownUnescape undoes the backslash escapes the docs generator added (\_, \\, \*).
var markdownUnescape = regexp.MustCompile(`\\(.)`)

func TestEncode_MatchesSpecTable(t *testing.T) {
	raw, err := os.ReadFile(specTablePath)
	if err != nil {
		t.Fatalf("read spec table: %v", err)
	}
	rowCount := 0
	for _, line := range strings.Split(string(raw), "\n") {
		cells := strings.Split(strings.Trim(strings.TrimSpace(line), "|"), "|")
		if len(cells) != 2 || !strings.Contains(cells[1], "?") {
			continue
		}
		char := markdownUnescape.ReplaceAllString(strings.TrimSpace(cells[0]), "$1")
		wantCode := markdownUnescape.ReplaceAllString(strings.TrimSpace(cells[1]), "$1")

		gotCode, encodeErr := Encode(char)
		if encodeErr != nil || gotCode != wantCode {
			t.Errorf("Encode(%q) = %q, %v; want %q", char, gotCode, encodeErr, wantCode)
		}
		rowCount++
	}
	const specRows = 139
	if rowCount != specRows {
		t.Errorf("checked %d spec rows, want %d - the table parser or the doc changed", rowCount, specRows)
	}
}

func TestEncode_SpecExamples(t *testing.T) {
	tests := []struct {
		name         string
		inputText    string
		wantEncoded  string
		wantEncodedN int
	}{
		{name: "VD1 from spec", inputText: "KHÓA", wantEncoded: "KH?OS?A", wantEncodedN: 7},
		{name: "VD2 from spec", inputText: "CÔNG TY SỮA", wantEncoded: "C?OO?NG TY S?UWX?A"},
		{name: "plain ASCII is untouched", inputText: "Nguyen Van A (HN), 12-3", wantEncoded: "Nguyen Van A (HN), 12-3"},
		{name: "slash and ampersand are converted", inputText: "A/B&C", wantEncoded: "A?_?B?_38?C"},
		{name: "empty", inputText: "", wantEncoded: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncoded, err := Encode(tt.inputText)
			if err != nil {
				t.Fatalf("Encode(%q) error: %v", tt.inputText, err)
			}
			if gotEncoded != tt.wantEncoded {
				t.Errorf("Encode(%q) = %q, want %q", tt.inputText, gotEncoded, tt.wantEncoded)
			}
			if tt.wantEncodedN != 0 && len(gotEncoded) != tt.wantEncodedN {
				t.Errorf("length after conversion = %d, want %d", len(gotEncoded), tt.wantEncodedN)
			}
		})
	}
}

func TestEncode_RejectsUncarryableCharacters(t *testing.T) {
	tests := []struct {
		name      string
		inputText string
		wantChar  rune
		wantIndex int
	}{
		{name: "at sign must be written (at)", inputText: "a@b", wantChar: '@', wantIndex: 1},
		{name: "literal question mark would read back as a marker", inputText: "Why?", wantChar: '?', wantIndex: 3},
		{name: "unmapped letter", inputText: "Zoë", wantChar: 'ë', wantIndex: 2},
		{name: "emoji", inputText: "ok 😀", wantChar: '😀', wantIndex: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Encode(tt.inputText)
			var charsetErr *CharsetError
			if !errors.As(err, &charsetErr) {
				t.Fatalf("Encode(%q) error = %v, want *CharsetError", tt.inputText, err)
			}
			if charsetErr.Char != tt.wantChar || charsetErr.Index != tt.wantIndex {
				t.Errorf("got char %q at %d, want %q at %d", charsetErr.Char, charsetErr.Index, tt.wantChar, tt.wantIndex)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		inputText   string
		wantText    string
		wantFailure bool
	}{
		{name: "VD1", inputText: "KH?OS?A", wantText: "KHÓA"},
		{name: "VD2", inputText: "C?OO?NG TY S?UWX?A", wantText: "CÔNG TY SỮA"},
		{name: "special characters", inputText: "?_?, ?_38?, ?_35?, ?_37?, ?_92?", wantText: "/, &, #, %, \\"},
		{name: "unterminated marker", inputText: "KH?OS", wantFailure: true},
		{name: "unknown code", inputText: "A?ZZ?B", wantFailure: true},
		{name: "empty code", inputText: "A??B", wantFailure: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotText, err := Decode(tt.inputText)
			if (err != nil) != tt.wantFailure {
				t.Fatalf("Decode(%q) error = %v, wantFailure %v", tt.inputText, err, tt.wantFailure)
			}
			if gotText != tt.wantText {
				t.Errorf("Decode(%q) = %q, want %q", tt.inputText, gotText, tt.wantText)
			}
		})
	}
}

func TestEncodeDecode_RoundTripsEveryConvertedCharacter(t *testing.T) {
	for char := range encodeCodes {
		inputText := "x" + string(char) + "y"
		encoded, err := Encode(inputText)
		if err != nil {
			t.Fatalf("Encode(%q): %v", inputText, err)
		}
		if err = ValidateCharset(encoded); err != nil {
			t.Errorf("Encode(%q) = %q is not in the valid character set: %v", inputText, encoded, err)
		}
		gotText, err := Decode(encoded)
		if err != nil || gotText != inputText {
			t.Errorf("round trip of %q = %q, %v", inputText, gotText, err)
		}
	}
}
