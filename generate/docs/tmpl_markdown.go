package docs

import (
	"io/ioutil"
	"log"
)

var (
	markdownTmpl string
)

func init() {
	f, err := ioutil.ReadFile("/Users/nasyaoristania/go/pkg/mod/github.com/tomato/generate/docs/tmpl.md")
	if err != nil {
		log.Println("tmpl_markdown")
		panic(err)
	}
	markdownTmpl = string(f)
}
