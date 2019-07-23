package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {

	fileserver := http.FileServer(http.Dir("assets/"))
	//fungsi StripPrefix membungkus fileserver supaya rute tidak terlalu panjang
	stripprefix := http.StripPrefix("/static/", fileserver)
	//menghandle semua request rute yang diawali dengan /static/ akan diarahkan ke directory assets
	http.Handle("/static/", stripprefix)

	//pangil semua isi folder views
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Mr. Manday", "title": "belajar web golang"}
		err = tmpl.ExecuteTemplate(w, "_header", data)
		err = tmpl.ExecuteTemplate(w, "index", data)
		err = tmpl.ExecuteTemplate(w, "_footer", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe("localhost:9000", nil)
}
