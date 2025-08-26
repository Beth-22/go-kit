package main

import (
	"fmt"
	"math"
	"strings"
	/*"sort"
	"strings"*/)

//functions

func sayHello(n string){
	fmt.Printf("Hello %v \n", n)
}

func sayBye(n string){
	fmt.Printf("byee %v \n", n)
}

//a slice and a function that takes in strings-parameters
func cycleNames(n []string, f func(string)){
	for _, v := range n{
		f(v)
	}
}




//return
                        //return type
func circleArea(r float64) float64{
	return math.Pi * r * r
}


func getInitials (n string)(string, string){
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names{
		initials = append(initials,v[:1])
	}

	if len(initials)>1{
		return initials[0], initials[1]
	}

	return initials[0], "_"

}











func test(){
	fn, sn := getInitials("lola lockhart")
	fmt.Println(fn,sn)
	sayHello("mario")
	sayBye("luigi")
	cycleNames([]string{"cloud", "barret", "tulla"}, sayHello)
	a1 := circleArea(10.5)
	fmt.Printf("the area of the circle is %0.3f", a1)
	//greeting := "hello there friendss!"
   /*
	//strings package
	fmt.Println(strings.Contains(greeting,"hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "heyy"))  //a copy, doesnt replace it, in fact none of these functions do. they just return what u want they dont replace the original values
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting,"ll")) //where ll is found
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45,23,55,4,26,775,33,89}

	//sort package
	sort.Ints(ages)
	fmt.Println(ages)

	index :=sort.SearchInts(ages,26)
	fmt.Println(index)

	friends := []string{"luigi", "mario", "bowser", "yoshi"}

	sort.Strings(friends)
	fmt.Println(friends)
	
	fmt.Println(sort.SearchStrings(friends,"luigi"))

*/

fmt.Println("loops")


//loops
/*
	x := 0
     //while x is less than 5
	for x < 5{
		fmt.Println("value of x is ", x)
		x++
	}

*/
/*
for i := 0; i<5; i++{
	fmt.Println("value of i is ", i)

}
*/

//loop thru a slice of strings

//friends := []string{"luigi", "mario", "bowser", "yoshi"}
/*
for i := 0; i<len(friends); i++{
	fmt.Println(friends[i])
}
*/
//kinda like for in

/*
for index, value := range friends{
	fmt.Printf("the value at index %v is %v \n", index, value)

}
*/

//booleans and conditionals


/*
age := 45

fmt.Println(age <=50)
fmt.Println(age ==50)

if age < 30{
	fmt.Println("age is less than 30")
}else if age <40{
	fmt.Println("age is less than 40")
}else{
	fmt.Println("age is not less than 45")
}

//nesting loops and ifs

for index, value := range friends{
	if index == 1{
		fmt.Println("continuing at pos", index)
		continue
	}
	fmt.Printf("the value at pos %v is %v", index, value)
}

*/



}

