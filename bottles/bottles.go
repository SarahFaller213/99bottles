package bottles

import (
	"bytes"
	"fmt"
)

type Generator struct {
}

func NewGenerator() Generator {
	return Generator{}
}

func (g *Generator) Verses(start, end int) string {
	//use SB
	var buffer bytes.Buffer
	for i := start; i > end-1; i-- {
		str := g.Verse(i)
		buffer.WriteString(str)
		if i != end {
			buffer.WriteString("\n\n")
		}
	}
	return buffer.String()
}

func (g *Generator) Verse(num int) string {
	// could simplfy the swith cstatements to the frist num 
	if num == 0 {
		return fmt.Sprintf("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.")
	}
	str1 := "bottles"
	if num < 2 {
		str1 = "bottle"
	}
	num2 := num - 1
	str2 := ""
	switch num2 {
	case 0:
		str2 = "Take it down and pass it around, no more bottles of beer on the wall."
	case 1:
		str2 = "Take one down and pass it around, 1 bottle of beer on the wall."
	default:
		str2 = fmt.Sprintf("Take one down and pass it around, %d bottles of beer on the wall.", num2)

	}
	return fmt.Sprintf("%d %s of beer on the wall, %d %s of beer.\n%s", num, str1, num, str1, str2)
}

func (g *Generator) Song() string {
	return g.Verses(99, 0)
}
