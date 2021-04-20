package docs

import (
	"io/ioutil"
	"log"
)

var (
	markdownTmpl string
)

func init() {
	f, err := ioutil.ReadFile("/Users/t-oristania.nabasya/go/pkg/mod/github.com/tomato@v1.4.2/generate/docs/tmpl.md")
	if err != nil {
		log.Println("tmpl_markdown")
		panic(err)
	}
	markdownTmpl = string(f)
}
