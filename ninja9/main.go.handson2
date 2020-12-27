package main

import "fmt"

type person struct{}

func (*person) speak(){
	fmt.Println("I'm Alive")
}

type human interface{
	speak()
}

func saySomething(h human){
	h.speak()
}

func main () {
p := person{}
saySomething(&p)

}