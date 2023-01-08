package parser

import "xls-tools/app/reader"

const (
	HeaderBegin = 0 // 字段名
	HeaderType  = 1 // 字段类型
	HeaderDes   = 2 // 字段描述
	DataBegin   = 3 // 数据开始
)

type Parser []*Table

type Table struct {
	ExcelName  string
	SheetName  string
	StructName string
	Defines    []*HeaderDefine
	Data       [][]*Cell
}

type Cell struct {
	Row  int
	Col  int
	Data string
}

func NewParser(file *reader.File) Parser {
	excel := file.AllExcel()
	tables := make([]*Table, 0, len(excel))
	for _, e := range excel {
		for _, s := range e.GetSheetList() {
			t := &Table{
				ExcelName: e.Path,
				SheetName: s,
			}
			rows, err := e.GetRows(s)
			if err != nil {
				return nil
			}
			t.resolveHeader(rows)
			t.resolveData(rows)
			tables = append(tables, t)
		}
	}
	return tables
}

func (t *Table) resolveData(rows [][]string) {
	data := make([][]*Cell, len(rows))
	for i := DataBegin; i < len(rows); i++ {
		for j, c := range rows[i] {
			data[i] = append(data[i], &Cell{
				Row:  i,
				Col:  j,
				Data: c,
			})
		}
	}
	t.Data = data
}

func (t *Table) resolveHeader(rows [][]string) {
	data := make([]*HeaderDefine, 0, len(rows))
	for i := 0; i < len(rows[0]); i++ {
		h := new(HeaderDefine)
		h.resolve(rows[HeaderBegin][i],
			rows[HeaderType][i],
			rows[HeaderDes][i],
		)
		data = append(data, h)
	}
	t.Defines = data
}
