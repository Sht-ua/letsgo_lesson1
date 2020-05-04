package main

import "fmt"

type human struct {
	Name string
}
type worker struct {
	human
	Company string
}
type  doctor struct {
	worker
	Stack[]string
}
type  Driver struct {
	worker
	DrivingExperience string
}
type  Policeman struct {
	worker
	Rank string
}
type  Teacher struct {
	worker
	Degree string
}
type  Firefighter struct {
	worker
	Rank2 string
}

func main(){
	user := Firefighter{Rank2: "Дежурный"}
	user1 := worker{Company: "Рога и копыта"}
	user3 := human{Name: "Ираклий"}
	x:= Driver{DrivingExperience: "Опытный"}
	v := human{"Иван"}
	p:=&v
	p.Name = "Анатолий"
	//user11 := &v
	
	
	fmt.Println(v.Name)
	fmt.Println(user.Name)
	fmt.Println(user.Rank2)
	fmt.Println(user1.Company)
	fmt.Println(user3.Name)
	fmt.Println(x.DrivingExperience)


}
