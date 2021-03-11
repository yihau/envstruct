package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl = `package envstruct

import (
	"os"
	"testing"
)
{{range $i, $block := .Blocks}}
func TestFillIn{{.Type | Title}}(t *testing.T) {
	c := struct {
		UnsetVar    {{$block.Type}}   ` + "`env:\"UNSET_VAR\"`" + `
		SetVar      {{$block.Type}}   ` + "`env:\"SET_VAR\"`" + `
		UnsetPtr    *{{$block.Type}}  ` + "`env:\"UNSET_PTR\"`" + `
		SetPtr      *{{$block.Type}}  ` + "`env:\"SET_PTR\"`" + `
		UnsetPtrPtr **{{$block.Type}} ` + "`env:\"UNSET_PTR_PTR\"`" + `
		SetPtrPtr   **{{$block.Type}} ` + "`env:\"SET_PTR_PTR\"`" + `
	}{}

	os.Setenv("SET_VAR", "{{index $block.EnvValue 0}}")
	os.Setenv("SET_PTR", "{{index $block.EnvValue 1}}")
	os.Setenv("SET_PTR_PTR", "{{index $block.EnvValue 2}}")

	var setVarExpected {{.Type}} = {{index $block.ExpectedValue 0}}
	var setPtrExpected {{.Type}} = {{index $block.ExpectedValue 1}}
	var setPtrPtrExpected {{.Type}} = {{index $block.ExpectedValue 2}}

	var varZeroValue {{.Type}} = {{$block.ZeroValue}}

	err := FillIn(&c)
	if err != nil {
		t.Errorf("FillIn error, err: %v", err)
	}

	if c.UnsetVar != varZeroValue {
		t.Errorf("c.UnsetVar expected: %v, got: %v", varZeroValue, c.UnsetVar)
	}
	if c.SetVar != setVarExpected {
		t.Errorf("c.SetVar expected: %v, got: %v", setVarExpected, c.SetVar)
	}
	if c.UnsetPtr != nil {
		t.Errorf("c.UnsetPtr expected nil")
	}
	if c.SetPtr == nil {
		t.Errorf("c.SetPtr expected not nil")
		if *c.SetPtr != setPtrExpected {
			t.Errorf("*c.SetPtr expected: %v, got: %v", setPtrExpected, *c.SetPtr)
		}
	}
	if c.UnsetPtrPtr != nil {
		t.Errorf("c.UnsetPtrPtr expected nil")
	}
	if c.SetPtrPtr == nil {
		t.Errorf("c.SetPtrPtr expected not nil")
		if *c.SetPtrPtr == nil {
			t.Errorf("*c.SetPtrPtr expected not nil")
			if **c.SetPtrPtr != setPtrPtrExpected {
				t.Errorf("**c.SetPtrPtr expected: %v, got: %v", setPtrPtrExpected, **c.SetPtrPtr)
			}
		}
	}
}
{{end}}
`

func main() {
	type Block struct {
		Type          string
		ZeroValue     string
		EnvValue      []string
		ExpectedValue []string
	}
	type Data struct {
		Blocks []Block
	}
	funcMap := template.FuncMap{
		"Title": strings.Title,
	}
	tmpl, err := template.New("test").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./decode_gen_test.go")
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = tmpl.Execute(f, Data{
		Blocks: []Block{
			{
				Type:          "bool",
				ZeroValue:     "false",
				EnvValue:      []string{"T", "true", "1"},
				ExpectedValue: []string{"true", "true", "true"},
			},
			{
				Type:          "int8",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "int16",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "int32",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "int64",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "int",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "uint8",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "uint16",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "uint32",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "uint64",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "uint",
				ZeroValue:     "0",
				EnvValue:      []string{"1", "2", "3"},
				ExpectedValue: []string{"1", "2", "3"},
			},
			{
				Type:          "float32",
				ZeroValue:     "0.0",
				EnvValue:      []string{"1.1", "2.2", "3.3"},
				ExpectedValue: []string{"1.1", "2.2", "3.3"},
			},
			{
				Type:          "float64",
				ZeroValue:     "0.0",
				EnvValue:      []string{"1.1", "2.2", "3.3"},
				ExpectedValue: []string{"1.1", "2.2", "3.3"},
			},
			{
				Type:          "string",
				ZeroValue:     `""`,
				EnvValue:      []string{"Hello", "World", "!!!"},
				ExpectedValue: []string{`"Hello"`, `"World"`, `"!!!"`},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
