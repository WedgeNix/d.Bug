package d

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Consts map[string]bool

var (
	consts = Consts{}
	debugr = bufio.NewReader(os.Stdin)
)

func Inject(c Consts) {
	for Const, b := range c {
		_, found := consts[Const]
		if found {
			panic("constant collision (" + Const + ")")
		}
		consts[Const] = b
	}
}

func Bug(constant string) bool {
	if len(consts) == 0 {
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
		log.Print("WARNING: " + constant + " ON. Continue? (y/n)  ")
		input, _ := debugr.ReadString('\n')
		if strings.HasPrefix(input, "n") {
			log.Fatalln(constant + " rejected")
		} else if strings.HasPrefix(input, "y") {
			return true
		}
	}
}
