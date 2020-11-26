package protein

import (
	"errors"
)

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

var aminoAcids = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(str string) (string, error) {

	val, ok := aminoAcids[str]

	if !ok {
		return "", ErrInvalidBase
	}

	if val == "STOP" {
		return "", ErrStop
	}

	return val, nil
}

func FromRNA(str string) ([]string, error) {
	buffer := ""
	var ret []string

	for _, elem := range str {
		buffer += string(elem)

		if len(buffer) == 3 {
			val, err := FromCodon(buffer)

			if err == ErrInvalidBase {
				return ret, ErrInvalidBase
			}

			if err == ErrStop {
				break
			}

			ret = append(ret, val)

			buffer = ""
		}
	}

	return ret, nil
}
