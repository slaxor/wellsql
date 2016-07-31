package wellsql

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"text/template"
)

// the file is expected to be in the form of
//-- SQLStatement: nameOfYourSQLInAnyWordChar
//SELECT OR INSERT OR ANY OTHER SQL ALSO with {{.Text}}
//{{.Template}} {{.Variability}}
//-- SQLStatement: YourNextSQL
//...
// You get rewarded with a map of functions which may
// be used like this:
// fnmap["nameOfYourSQLInAnyWordChar"](yourTmpltDataStruct)
// which returns the actual SQL as a string
func LoadFile(tmplFile string) (map[string]func(interface{}) string, error) {
	f, err := ioutil.ReadFile(tmplFile)
	if err != nil {
		return nil, err
	}
	tmpl := string(f)
	re := regexp.MustCompile("(?U)SQLStatement: (\\w+)\n((?sm).+)(--|\\z)")
	fm := make(map[string]func(interface{}) string)
	for _, match := range re.FindAllStringSubmatch(tmpl, -1) {
		name := match[1]
		stmnt := match[2]
		t := template.Must(template.New(name).Parse(stmnt))
		fm[name] = func(input interface{}) string {
			var doc bytes.Buffer
			t.Execute(&doc, input)
			return doc.String()
		}
	}

	return fm, nil
}
