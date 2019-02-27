package transform

import (
	"bufio"
	"io"
)

type Echo struct{}

func (ec *Echo) Transform(r io.Reader) io.Reader {
	pr, pw := io.Pipe()
	go func() {
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			pw.Write(
				[]byte(line + "\n"),
			)
		}
	}()
	return pr
}
