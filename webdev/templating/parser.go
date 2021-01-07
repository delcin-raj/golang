package main
import (
	"os"
	"log"
	"text/template"
)

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	tpl,err := template.ParseFiles("templates/index.gohtml")
	handleErr(err)
	nf,err := os.Create("templates/parsedTemplates/index.html")
	handleErr(err)
	defer nf.Close()
	err = tpl.Execute(nf,nil)
	handleErr(err)

	// using glob and working with multiple templates at once
	tplg,err := template.ParseGlob("templates/*.gohtml")
	handleErr(err)
	nfg,err := os.Create("templates/parsedTemplates/gollebed.html")
	handleErr(err)
	defer nfg.Close()
//	for i := 0; i < 4; i++ {
//		err = tplg.Execute(nfg,nil) // curious to look at the order of execution
//		// The above statement only executes the first in order
//		handleErr(err)
//	}
	// The loop logic also do not work
	// so the only knowledge using which I can finish this is
	handleErr(tplg.ExecuteTemplate(nfg,"index.gohtml",nil))
	handleErr(tplg.ExecuteTemplate(nfg,"two.gohtml",nil))
	handleErr(tplg.ExecuteTemplate(nfg,"one.gohtml",nil))

}


