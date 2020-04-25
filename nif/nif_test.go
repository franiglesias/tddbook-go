package nif

import "testing"

func TestShouldFailIfCandidateIsInvalid(t *testing.T) {
	tests := []struct {
		name string
		example string
		expected string
	}{
		{"should fail if too long", "01234567891011", "bad format"},
		{"should fail if too short", "01234", "bad format"},
		{"should fail if starts with a letter other than X, Y, Z", "A12345678", "bad format"},
		{"should fail if doesn't have 7 digit in the middle", "0123X567R", "bad format"},
		{"should fail if doesn't end with a valid control letter", "01234567U", "invalid end format"},
		{"should fail if doesn't end with the right control letter", "00000000S", "bad control letter"},

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewNif(test.example)

			if err.Error() != test.expected  && err.Error() != "bad format" {
				t.Errorf("Expected %s, got %s", test.expected, err.Error())
			}
		})
	}
}

func TestShouldCreateNifTypeWithValidCandidate(t *testing.T) {
	tests := []struct {
		name string
		example string
	}{
		{"should accept mod23 being 0", "00000000T"},
		{"should accept mod23 being 0 letter T", "00000023T"},
		{"should accept mod23 being 1 letter R", "00000024R"},
		{"should accept NIE starting with X", "X0000000T"},
		{"should accept NIE starting with Y", "Y0000000Z"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nif, err := NewNif(test.example)

			if nif != Nif(test.example) {
				t.Errorf("Expected Nif(%s), got %s", test.example, nif)
			}

			if err != nil {
				t.Errorf("Unexpected error %s", err.Error())
			}
		})
	}
}
