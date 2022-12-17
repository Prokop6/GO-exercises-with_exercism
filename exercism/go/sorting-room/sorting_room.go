package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	number := nb.Number()

	return fmt.Sprintf("This is a box containing the number %.1f", float32(number))
}


type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	_, isOK := fnb.(FancyNumber)
	if !isOK { return 0 }
	
	num, err :=	strconv.Atoi(fnb.Value())
	if err != nil { return 0 }
	
	return num
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	
	 return fmt.Sprintf("This is a fancy box containing the number %.1f", float32(ExtractFancyNumber(fnb)))
	
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch val := i.(type) {
	case int:
		return DescribeNumber(float64(val))
	case float64:
		return DescribeNumber(val)	
	case NumberBox:
		return DescribeNumberBox(val)
	case FancyNumberBox:
		return DescribeFancyNumberBox(val)
	default:
		return "Return to sender"
	}
}
