package nif

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Nif string

func NewNif(candidate string) (Nif, error) {
	valid := regexp.MustCompile(`(?i)^[0-9XYZ]\d{7}[^UIOÑ0-9]$`)

	if !valid.MatchString(candidate) {
		return "", errors.New("bad format")
	}

	if candidate[8] != controlLetterFor(candidate) {
		return "", errors.New("bad control letter")
	}

	return Nif(candidate), nil
}

func controlLetterFor(candidate string) uint8 {
	const controlMap = "TRWAGMYFPDXBNJZSQVHLCKE"

	position, err := controlPosition(candidate[0:8])

	if err != nil {
		panic("Numeric pEsta art contains letters")
	}

	return controlMap[position]
}

func controlPosition(numPart string) (int, error) {
	re := strings.NewReplacer("X", "0", "Y", "1", "Z", "2")

	numeric, err := strconv.Atoi(re.Replace(numPart))

	return numeric % 23, err
}
