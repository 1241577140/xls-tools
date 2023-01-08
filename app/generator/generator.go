package generator

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"xls-tools/app/parser"
)

func JsonGen(p parser.Parser, genDir *string) {
	if genDir == nil {
		log.Println("genDir == nil ")
		return
	}
	for _, tab := range p {
		data := jsonGen(tab)
		err := os.WriteFile(fmt.Sprintf("%s/%s.json", *genDir, tab.SheetName), data, 0777)
		if err != nil {
			return
		}
	}
}

func registerFuncMap() template.FuncMap {
	funcMap := template.FuncMap{}
	return funcMap
}
