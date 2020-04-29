package main

import "fmt"

type  worker struct {
		person string `json:"name"`
		profession string `json:"profession"`
}
func main() {
	var a worker = worker{"Anton", "supply manager"}
	var b worker = worker{"Dima", "dancer"}
	var c worker = worker{"svetlana", "supervisor"}
	var d worker = worker{"Alina", "manager"}
	var f worker = worker{"Igor", "director"}
	fmt.Println(a.profession)
	fmt.Println(b.profession)
	fmt.Println(c.profession)
	fmt.Println(d.profession)	
	fmt.Println(f.profession)
}
