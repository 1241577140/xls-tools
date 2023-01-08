package generator

import (
	"bytes"
	"xls-tools/app/parser"
)

func jsonGen(tab *parser.Table) []byte {
	bf := bytes.Buffer{}
	bf.WriteString("[")
	for i := parser.DataBegin; i < len(tab.Data); i++ {
		row := tab.Data[i]
		bf.WriteString("{")
		for i, c := range row {
			headerDefine := tab.Defines[c.Col]
			bf.WriteString(headerDefine.ParseJson(c.Data))
			if i != len(row)-1 {
				bf.WriteString(",")
			}
		}
		bf.WriteString("}")
		if i != len(tab.Data)-1 {
			bf.WriteString(",")
		}
	}
	bf.WriteString("]")
	return bf.Bytes()
}
