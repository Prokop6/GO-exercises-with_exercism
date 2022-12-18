package protein

import (
	"errors"
	"regexp"
)

var ErrStop error = errors.New("stop codon reached")
var ErrInvalidBase error = errors.New("invalid base")


var InverseRnaCodonTable map[string]string = map[string]string{ 
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
"UAA": "(Stop)",
"UAG": "(Stop)",
"UGU": "Cysteine",
"UGC": "Cysteine",
"UGA": "(Stop)",
"UGG": "Tryptophan",
"CUC": "Leucine",
"CUA": "Leucine",
"CUG": "Leucine",
"CCU": "Proline",
"CCC": "Proline",
"CCA": "Proline",
"CCG": "Proline",
"CAU": "Histidine",
"CAC": "Histidine",
"CAA": "Glutamine",
"CAG": "Glutamine",
"CGU": "Arginine",
"CGC": "Arginine",
"CGA": "Arginine",
"CGG": "Arginine",
"AUC": "Isoleucine",
"AUA": "Isoleucine",
"AUG": "Methionine",
"ACU": "Threonine",
"ACC": "Threonine",
"ACA": "Threonine",
"ACG": "Threonine",
"AAU": "Asparagine",
"AAC": "Asparagine",
"AAA": "Lysine",
"AAG": "Lysine",
"AGU": "Serine",
"AGC": "Serine",
"AGA": "Arginine",
"AGG": "Arginine",
"GUU": "Valine",
"GUC": "Valine",
"GUA": "Valine",
"GUG": "Valine",
"GCU": "Alanine",
"GCC": "Alanine",
"GCA": "Alanine",
"GCG": "Alanine",
"GAU": "Aspartic acid",
"GAC": "Aspartic acid",
"GAA": "Glutamic acid",
"GAG": "Glutamic acid",
"GGU": "Glycine",
"GGC": "Glycine",
"GGA": "Glycine",
"GGG": "Glycine",
}


var re = regexp.MustCompile(`((?:AUG)(?:[AGUC]{3})*)(:w
	:?UAA|UGA|UAG)?`)

func FromRNA(rna string) ([]string, error) {
	m := re.FindStringSubmatch(rna)
	var r []string

	if m[0] == "" {
		return r, ErrStop 
	}

	c_ar := splitStrand(m[0])
	
	for _, c := range c_ar {
		p, p_er := FromCodon(c)
		if p_er == ErrInvalidBase {
			return nil, p_er
		}
		if p_er == ErrStop {
			break
		} 
			r = append(r, p)
	}

	return r, nil
}

func splitStrand(strand string) ([]string) {
	var r []string
	for i := 0; i < len(strand)-2; i += 3 {
		r = append(r, strand[i:i+3])
	}

	return r
}

func FromCodon(codon string) (string, error) {
	p, ok := InverseRnaCodonTable[codon]

	if p == "(Stop)" {
		return "", ErrStop
	}
	if ok {
		return p, nil
	}

	return "", ErrInvalidBase

}
