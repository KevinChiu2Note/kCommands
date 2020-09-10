package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/kevinchiu2note/kcommands/kcat/cat"
	"os"
)

func main() {
	flag.Parse()
	// no command args
	if flag.NArg() == 0 {
		cat.Cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v", flag.Arg(i), err)
			continue
		}
		cat.Cat(bufio.NewReader(f))
	}
}
