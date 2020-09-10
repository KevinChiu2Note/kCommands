package cat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		_, _ = fmt.Fprintf(os.Stdout, "%s", buf)
	}
}
