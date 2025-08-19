package main

import "fmt"

var points =[]int {29,11,233,11,22}


func sayHi(n string){
	fmt.Println("heyyy", n)
}




//maps

func menus(){

	menu := map[string]float64{
	"soup" : 234,
	"pizza" : 113,
	"salad" : 233,
}
 fmt.Println(menu["soup"])   ///234



 //looping thru the menu

 for k, v := range menu{
	fmt.Println(k, "-", v)
 }


 //ints as key type
 //variable := map [key type] valuetype

 phonebook := map[int]string{
	234343: "mario",
	345543: "abi",
	523332:"rue",
 }

 fmt.Println(phonebook)

 phonebook[345543] = "bowser"
  fmt.Println(phonebook)

}



//pass by value language