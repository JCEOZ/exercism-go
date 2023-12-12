package sorting

import "strconv"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return "This is the number " + strconv.FormatFloat(f, 'f', 1, 64)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return "This is a box containing the number " + strconv.FormatFloat(float64(nb.Number()), 'f', 1, 64)
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
	switch box := fnb.(type) {
	case FancyNumber:
		value, _ := strconv.Atoi(box.Value())
		return value
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return "This is a fancy box containing the number " + strconv.FormatFloat(float64(ExtractFancyNumber(fnb)), 'f', 1, 64)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch value := i.(type) {
	case int:
		return DescribeNumber(float64(value))
	case float64:
		return DescribeNumber(value)
	case NumberBox:
		return DescribeNumberBox(value)
	case FancyNumberBox:
		return DescribeFancyNumberBox(value)
	default:
		return "Return to sender"
	}
}
