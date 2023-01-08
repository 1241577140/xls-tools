package parser

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"xls-tools/utils/str"
)

type Type int

const (
	TypeBase  Type = iota + 1
	TypeSlice Type = iota + 1
)

var typeIndex = map[string]Type{
	"int":      TypeBase,
	"string":   TypeBase,
	"bool":     TypeBase,
	"[]int":    TypeSlice,
	"[]string": TypeSlice,
	"[]bool":   TypeSlice,
}

type HeaderDefine struct {
	TypeName string // 名
	Type     string // 类型
	Describe string // 备注
}

func (h *HeaderDefine) resolve(header, t, des string) {
	if _, ok := typeIndex[t]; !ok {
		log.Fatalf("not allow type %s", t)
		return
	}
	h.TypeName = header
	h.Type = t
	h.Describe = des
}

func (h *HeaderDefine) ParseJson(s string) string {
	switch h.Type {
	case "int":
		return fmt.Sprintf(`"%s":%s`, h.TypeName, s)
	case "string":
		return fmt.Sprintf(`"%s":"%s"`, h.TypeName, s)
	case "bool":
		return fmt.Sprintf(`"%s":%v`, h.TypeName, str.ParseBool(s))
	case "[]int", "[]string", "[]bool":
		return h.sliceSeparate(s)
	default:
		return ""
	}
}

func (h *HeaderDefine) sliceSeparate(s string) string {
	ss := strings.Split(s, ",")
	l := len(ss)
	bf := bytes.Buffer{}
	bf.WriteString(fmt.Sprintf(`"%s":`, h.TypeName))
	bf.WriteString("[")
	sfunc := func(s string) string {
		return fmt.Sprintf("%s", s)
	}
	switch h.Type {
	case "[]int":
	case "[]string":
		sfunc = func(s string) string {
			return fmt.Sprintf(`"%s"`, s)
		}
	case "[]bool":
	}
	for i, d := range ss {
		bf.WriteString(sfunc(d))
		if i != l-1 {
			bf.WriteString(",")
		}
	}
	bf.WriteString("]")
	return bf.String()
}
