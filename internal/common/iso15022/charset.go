// Package iso15022 implements the character rules of the VSD FIN format (docs/03-ky-tu-va-chuyen-doi-tieng-viet.md):
// the valid character set and the reversible conversion of Vietnamese letters into it.
package iso15022

import "fmt"

// markerRune brackets every converted character: "Ó" becomes "?OS?".
const markerRune = '?'

// CharsetError reports the first character a text cannot carry.
type CharsetError struct {
	// Index is the rune offset of the offending character.
	Index int
	// Char is the offending character.
	Char rune
}

// Error implements the error interface.
func (e *CharsetError) Error() (message string) {
	return fmt.Sprintf("character %q at position %d is not allowed in a FIN message", e.Char, e.Index)
}

// IsValidChar reports whether char may appear directly in a FIN message: letters, digits, the
// punctuation / - ? : ( ) . , ' + , space, carriage return and line feed.
func IsValidChar(char rune) (valid bool) {
	switch {
	case char >= 'a' && char <= 'z', char >= 'A' && char <= 'Z', char >= '0' && char <= '9':
		return true
	}
	switch char {
	case '/', '-', '?', ':', '(', ')', '.', ',', '\'', '+', ' ', '\r', '\n':
		return true
	}
	return false
}

// ValidateCharset returns a *CharsetError for the first character outside the valid set, or nil.
// Use it on text that is already converted; it checks, it does not convert.
//
// Table 2 of the spec converts "/" to "?_?" and "&" to "?_38?", yet the valid set of section 2.1
// has no underscore. The converted forms must pass, so an underscore is accepted - but only between
// a pair of markers, where it is part of such a code.
func ValidateCharset(text string) (err error) {
	insideMarker := false
	index := 0
	for _, char := range text {
		if char == markerRune {
			insideMarker = !insideMarker
		}
		if !IsValidChar(char) && !(char == '_' && insideMarker) {
			return &CharsetError{Index: index, Char: char}
		}
		index++
	}
	return nil
}
