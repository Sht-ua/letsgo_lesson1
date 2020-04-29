package main

import (
	"fmt"
)

type  Profession struct {
	Doctor string
	Driver string
	Policeman string
	Teacher string
	Fireman string
}
func main() {
	var a = Profession{Teacher: "Учитель младших класов"}
	var b = Profession{Doctor: "Педиатр"}
	var c = Profession{Fireman: "Старший пожарник"}
	var d = Profession{Policeman: "Полицейский ДПС"}
	var f = Profession{Driver: "Водитель автобуса"}

	fmt.Println(a,b,c,d,f)


}


