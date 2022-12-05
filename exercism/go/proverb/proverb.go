package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var	proverb []string

	for i := 0; i < len(rhyme); i++ {
		var verse string

		if i == len(rhyme) - 1{
			verse = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			verse = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}

		proverb = append(proverb, verse)
	}

	return proverb

}
