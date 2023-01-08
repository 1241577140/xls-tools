package main

import (
	"flag"
	"xls-tools/app/generator"
	"xls-tools/app/parser"
	"xls-tools/app/reader"
)

var (
	fileName   = flag.String("n", "", "file name,separator ','")
	fileDir    = flag.String("d", "", "file dir")
	genJsonDir = flag.String("gen_json", "", "gen json dir")
)

func main() {
	flag.Parse()
	file := reader.NewFile()
	file.SetName(fileName).SetDir(fileDir)
	err := file.Load()
	if err != nil {
		return
	}

	newParser := parser.NewParser(file)
	generator.JsonGen(newParser, genJsonDir)
}
