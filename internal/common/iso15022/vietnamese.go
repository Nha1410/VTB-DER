package iso15022

import (
	"fmt"
	"strings"
)

// toneKeys are the telex tone keys in the order the letter groups below list their runes:
// huyền, sắc, hỏi, ngã, nặng.
const toneKeys = "fsrxj"

// toneless maps the letters that carry no tone mark to their telex code (lower case).
var toneless = map[rune]string{
	'ă': "aw", 'ơ': "ow", 'ư': "uw", 'â': "aa", 'ô': "oo", 'ê': "ee", 'đ': "dd",
}

// toned lists, per telex base, the five toned lower-case runes in toneKeys order. Table 2 of the
// spec is exactly this telex rule, so the table is derived instead of written out 130 times;
// TestEncode_MatchesSpecTable proves the result equals the spec table row for row.
var toned = []struct {
	base  string
	runes string
}{
	{"a", "àáảãạ"}, {"aw", "ằắẳẵặ"}, {"aa", "ầấẩẫậ"},
	{"e", "èéẻẽẹ"}, {"ee", "ềếểễệ"},
	{"i", "ìíỉĩị"},
	{"o", "òóỏõọ"}, {"oo", "ồốổỗộ"}, {"ow", "ờớởỡợ"},
	{"u", "ùúủũụ"}, {"uw", "ừứửữự"},
	{"y", "ỳýỷỹỵ"},
}

// specials are the non-Vietnamese characters Table 2 also converts; the number is the ASCII code.
var specials = map[rune]string{
	'/': "_", '&': "_38", '#': "_35", '%': "_37", '\\': "_92",
}

var (
	// encodeCodes maps a character to its code between the markers, e.g. 'Ó' → "OS".
	encodeCodes = map[rune]string{}
	// decodeChars is the reverse: "OS" → 'Ó'.
	decodeChars = map[string]rune{}
)

func init() {
	register := func(char rune, code string) {
		encodeCodes[char] = code
		decodeChars[code] = char
	}
	for char, code := range toneless {
		register(char, code)
		register(toUpperRune(char), strings.ToUpper(code))
	}
	for _, group := range toned {
		for index, char := range []rune(group.runes) {
			code := group.base + string(toneKeys[index])
			register(char, code)
			register(toUpperRune(char), strings.ToUpper(code))
		}
	}
	for char, code := range specials {
		register(char, code)
	}
}

// toUpperRune upper-cases one rune.
func toUpperRune(char rune) (upper rune) {
	return []rune(strings.ToUpper(string(char)))[0]
}

// Encode converts free text into the FIN character set: every Vietnamese letter and the characters
// / & # % \ become "?code?" (e.g. "KHÓA" → "KH?OS?A"). Text must be in NFC (precomposed) form.
// It returns a *CharsetError for any other character the format cannot carry - including a literal
// "?", which would be read back as a marker, and "@", which the spec says to write as "(at)".
//
// Apply it to field values only, never to the structural characters of a line (":", "//", "/" inside
// a tag), or the structure itself would be converted.
func Encode(text string) (encoded string, err error) {
	var builder strings.Builder
	index := 0
	for _, char := range text {
		if code, ok := encodeCodes[char]; ok {
			builder.WriteRune(markerRune)
			builder.WriteString(code)
			builder.WriteRune(markerRune)
		} else if char != markerRune && IsValidChar(char) {
			builder.WriteRune(char)
		} else {
			return "", &CharsetError{Index: index, Char: char}
		}
		index++
	}
	return builder.String(), nil
}

// Decode reverses Encode: every "?code?" becomes its original character. A "?" that opens no known
// code - unterminated, empty or unknown - is an error, because the sender must always bracket a
// conversion.
func Decode(encoded string) (text string, err error) {
	var builder strings.Builder
	for position := 0; position < len(encoded); {
		if encoded[position] != markerRune {
			builder.WriteByte(encoded[position])
			position++
			continue
		}
		closing := strings.IndexByte(encoded[position+1:], markerRune)
		if closing < 0 {
			return "", fmt.Errorf("unterminated conversion at byte %d", position)
		}
		code := encoded[position+1 : position+1+closing]
		char, ok := decodeChars[code]
		if !ok {
			return "", fmt.Errorf("unknown conversion %q at byte %d", code, position)
		}
		builder.WriteRune(char)
		position += closing + 2
	}
	return builder.String(), nil
}
