# wellsql
Have the SQL Statement neatly in a template file in your Golang project
## Install
`go get github.com/slaxor/wellsql` or `govendor add github.com/slaxor/wellsql/`

## Use
####your_statements.sql.tmpl:
```sql
    -- SQLStatement: nameOfYourSQLInAnyWordChar
    SELECT OR INSERT OR ANY OTHER SQL ALSO with
    {{.Text}} {{.Template}} {{.Variability}} 
    -- SQLStatement: YourNextSQL
    ...
```    
####main.go
```go
    package main
    import (
        "log"
        "github.com/slaxor/wellsql"
    )
    
    func main() {
       sqlMap, err := wellsql.LoadFile("your_statements.sql.tmpl")
       	if err != nil {
		    log.Fatal(err)
	    }
	    log.Printf("%q", sqlMap["selectFromFoo"](nil))
	    c := struct{ Cond string }{"foo"}
	    log.Printf("%q", sqlMap["selectFromFoo"](c))
    }
```
