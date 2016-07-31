package wellsql

import (
	"reflect"
	"testing"
)

var testNames = []string{"empty", "lastInFile", "simple", "multiLine"}
var testFileName = "testdata/template_file.sql.tmpl"

func TestLoadFile(t *testing.T) {
	fm, err := LoadFile(testFileName)
	if err != nil {
		t.Errorf("%v", err)
	}
	if fm == nil {
		t.Errorf("fm is emtpy")
	}
}

func TestMap(t *testing.T) {
	fm, err := LoadFile(testFileName)
	if err != nil {
		t.Error(err)
	}

	for _, k := range testNames {
		result := reflect.TypeOf(fm[k]).String()
		expect := "func(interface {}) string"
		if result != expect {
			t.Errorf("\"%v\" must be a \"%v\" but is \"%v\"", k, expect, result)
		}
	}
}

// type TestTable struct {
// In struct {
// Filter, SortBy, V1, V3 string
// }
// Out string
// }

func TestFuncs(t *testing.T) {
	fm, err := LoadFile(testFileName)
	if err != nil {
		t.Error(err)
	}
	type TemplateData struct {
		Filter, SortBy, V1, V3 string
		Fields                 []string
	}
	in := make(map[string]TemplateData)
	in["simple"] = TemplateData{Filter: "bar=1"}
	in["multiLine"] = TemplateData{Filter: "bar=1", SortBy: "name ASC"}
	in["withRange"] = TemplateData{Fields: []string{"foo,", "bar,", "baz"}}
	// in["withRange"] = TemplateData{}
	in["empty"] = TemplateData{}
	in["lastInFile"] = TemplateData{V1: "Val1", V3: "Val3"}

	out := make(map[string]string)
	out["simple"] = "SELECT * FROM foo WHERE bar=1\n\n"
	out["multiLine"] = "SELECT * FROM foo WHERE bar=1\nSORT BY name ASC\n\n"
	out["empty"] = "\n"
	out["withRange"] = "SELECT foo, bar, baz  FROM foo\n\n"
	out["lastInFile"] = "INSERT INTO foo (v1, v2, v3) VALUES (\n'Val1', 'Val2', 'Val3')\n\n"

	for name, fun := range fm {
		result := fun(in[name])
		expect := out[name]
		if result != expect {
			t.Errorf("%q must equal %q but is %q", name, expect, result)
		}
	}
}
