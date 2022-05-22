package main

import (
	"fmt"

	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
	"github.com/sebamgb/mypackage"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("seba")
	car := mypackage.CarPublic{
		Brand: "lamborgini",
		Year:  2022,
	}
	fmt.Println(car)
}
