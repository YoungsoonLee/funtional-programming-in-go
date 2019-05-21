package main

import (
	"fmt"
	"strings"
)

type WordSize int

const (
	ZERO WordSize = 6 * iota
	SMALL
	MEDIUM
	LARGE
	XLARGE
	XXLARGE   WordSize = 50
	SEPARATOR          = ", "
)

type ChainLink struct {
	Data []string
}

func (v *ChainLink) Value() []string {
	return v.Data
}

type stringFunc func(s string) (result string) // !!!

func (v *ChainLink) Map(fn stringFunc) *ChainLink {
	var mapped []string
	orig := *v
	for _, s := range orig.Data {
		mapped = append(mapped, fn(s))
	}
	v.Data = mapped
	return v
}

func (v *ChainLink) Filter(max WordSize) *ChainLink {
	filterd := []string{}
	orig := *v
	for _, s := range orig.Data {
		if len(s) <= int(max) {
			filterd = append(filterd, s)
		}
	}
	v.Data = filterd
	return v
}

func main() {
	constants := `
** Constants ***
ZERO: %v
SMALL: %d
MEDIUM: %d
LARGE: %d
XLARGE: %d
XXLARGE: %d
`
	fmt.Printf(constants, ZERO, SMALL, MEDIUM, LARGE, XLARGE, XXLARGE)

	words := []string{
		"tiny",
		"marathon",
		"philanthropinist",
		"supercalifragilisticexpialidocious",
	}

	data := ChainLink{words}
	fmt.Printf("unfiltered: %#v\n", data.Value())

	filterd := data.Filter(SMALL)
	fmt.Printf("filterd: %#v\n", filterd)

	fmt.Printf("filterd and mapped (<= SMALL sized words): %#v\n", filterd.Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filterd and mapped (<= MEDIUM sized words): %#v\n", data.Filter(MEDIUM).Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filterd and mapped (<= LARGE sized words): %#v\n", data.Filter(XLARGE).Map(strings.ToUpper).Filter(LARGE).Value())

	data = ChainLink{words}
	fmt.Printf("filterd and mapped (<= LARGE sized words): %#v\n", data.Filter(XLARGE).Map(strings.ToUpper).Filter(LARGE).Value())

	// fmt.Printf("norig_data: %v\n", u.Join(orig_data, SEPARATOR))
}
