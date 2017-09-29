package d

import (
	"bufio"
	"os"
	"strings"
)

type Consts map[string]bool

var (
	consts Consts
	debugr = bufio.NewReader(os.Stdin)
)

func Inject(c Consts) {
	consts = c
}

func Bug(constant string) bool {
	if consts == nil {
		panic("must inject constants")
	}
	b, found := consts[constant]
	if !found {
		panic(constant + " not found")
	}
	if !b {
		return false
	}
	for {
		println("WARNING: " + constant + " ON. Continue? (y/n)")
		input, _ := debugr.ReadString('\n')
		if strings.HasPrefix(input, "n") {
			panic(constant + " rejected")
		} else if strings.HasPrefix(input, "y") {
			return true
		}
	}
}
