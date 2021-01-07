package main
import(
	"os"
	"text/template"
	"check"
)
//var tpl *template.Template
//func init_tpl() {
//	tpl,err := template.ParseFiles("templates/dynamic.gohtml")
//	check.Err(err)
//}
type person struct {
	Name string
	Age int
}
func main() {
	tpls,err := template.ParseFiles("templates/dynamic-struct.gohtml")
	check.Err(err)
	nfs,err := os.Create("templates/parsedTemplates/dynamic-struct.html")
	check.Err(err)
	defer nfs.Close()

//	check.Err(tpl.Execute(nf,
//		person{
//			"Delcin",
//			25,
//		}))
// does not work
	//names := []string{"Delcin","beuls","amazing"}
	delcin := person{
			"delcin",
			25,
		}
	check.Err(tpls.Execute(nfs,delcin))
}
