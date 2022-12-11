package strand

var dnaToRna = map[rune]string{
	'A': "U",
	'G': "C",
	'C': "G",
	'T': "A",
}

func ToRNA(dna string) string {
	rna := ""

	for _, n := range dna {
		rna = rna + dnaToRna[n]
	}

	return rna
}
