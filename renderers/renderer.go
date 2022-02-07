package renderers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}


func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc , err := CreateTemplateCache()
	if err != nil{
		fmt.Println("error occured while creating template cache",err)
		return
	}
	parsed_template,template_exists:=tc[tmpl]
	if !template_exists {
		fmt.Println("Template does not exist")
		return
	}
	buf:=new(bytes.Buffer)
	parsed_template.Execute(buf,nil)
	buf.WriteTo(w)
	
}

func CreateTemplateCache() (map[string]*template.Template,error) {
	myCache:=map[string]*template.Template{}
	pages,err:=filepath.Glob("./templates/*.page.tmpl")
	if err!=nil{
		return myCache,err
	}
	for _,page:= range pages{
		name:=filepath.Base(page)
		ts , err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return myCache,err
		}

		mathches_layout, err:= filepath.Glob("./templates/*.layout.tmpl")
		if err != nil{
			return myCache,err
		}

		if len(mathches_layout)>0{
			ts , err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err !=nil {
				return myCache,err
			}
		}
		myCache[name]=ts
	}
	return myCache,err

}