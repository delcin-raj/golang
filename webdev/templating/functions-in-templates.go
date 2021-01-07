package main
import (
	"check"
	"os"
	"text/template"
)
// If we want to pass functions to the templates then
// we should add the functions before parsing the template

var fm = template.FuncMap{
	"multiply":multiply,
}

func multiply(x []int) int {
	return x[0]*x[1]
}

func main() {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles(
		"templates/func-template.gohtml"))
	// in effect templates are created with functions	
	nf,err := os.Create("templates/parsedTemplates/func-template.html")
	defer nf.Close()
	check.Err(err)
	check.Err(tpl.ExecuteTemplate(nf,"func-template.gohtml",[]int{3,4}))
}
