package transform

import (
	"bufio"
	"io"
	"log"

	gtr "github.com/aerokite/go-google-translate/pkg"
)

type Translator struct {
	From string
	To   string
}

func (tr *Translator) Transform(r io.Reader) io.Reader {
	pr, pw := io.Pipe()
	go func() {
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			// request struct
			req := &gtr.TranslateRequest{
				SourceLang: tr.From,
				TargetLang: tr.To,
				Text:       line,
			}
			// translate
			translated, err := gtr.Translate(req)
			if err != nil {
				log.Fatalln(err)
			}
			pw.Write(
				[]byte(translated + "\n"),
			)
		}
	}()
	return pr
}
