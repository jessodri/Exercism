package strand

import "strings"

func letterSwitch(letter string) string {

	switch strings.ToUpper(letter) {
	case "G":
		return "C"
	case "C":
		return "G"
	case "T":
		return "A"
	case "A":
		return "U"
	default:
		return ""
	}
}

// ToRNA converts a DNA strand to RNA
func ToRNA(dna string) string {
	rna := ""

	for _, r := range dna {
		rna += letterSwitch(string(r))
	}
	return rna
}
