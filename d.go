package d

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Consts map[string]bool
type konst struct {
	*sync.Once
	b bool
}

var (
	consts = map[string]konst{}
	debugr = bufio.NewReader(os.Stdin)
)

func Inject(c Consts) {
	for Const, b := range c {
		_, found := consts[Const]
		if found {
			panic("constant collision (" + Const + ")")
		}
		consts[Const] = konst{b: b}
	}
}

func Bug(constant string) bool {
	if len(consts) == 0 {
		panic("must inject constants")
	}
	k, found := consts[constant]
	if !found {
		panic(constant + " not found")
	}
	if !k.b {
		return false
	}
	k.Do(func() {
		for {
			fmt.Print("WARNING: " + constant + " ON. Continue? (y/n)  ")
			input, _ := debugr.ReadString('\n')
			if strings.HasPrefix(input, "n") {
				log.Fatalln(constant + " rejected")
			} else if strings.HasPrefix(input, "y") {
				return
			}
		}
	})
	return true
}
