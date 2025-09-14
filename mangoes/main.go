package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
) //package-for formatting strings and printing stuff to the console. fmt.Capitalmethod, fmt.Println


func getInput(prompt string, r *bufio.Reader) (string,error){
    fmt.Print(prompt)
    input, err:= r.ReadString('\n')
    return strings.TrimSpace(input), err

}

func createBill() bill{
    reader := bufio.NewReader(os.Stdin)

    /*fmt.Print("create a new bill name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)*/ //refactored above by another function
      
    name, _  := getInput("create a new bill name:", reader)
    b := newBill(name)
    fmt.Println("Created the bill -", b.name)

    return b

}


func promptOptions(b bill){
    reader := bufio.NewReader(os.Stdin)
    opt , _  := getInput("choose option \n a- add item \n s- save bill \n t- add tip :", reader)

    switch opt {
    case "a" :
        name, _ := getInput("Item name: ", reader)
        price, _ := getInput("item price:", reader)

       //parsing floats
        p, err := strconv.ParseFloat(price, 64)

        if err != nil{
            fmt.Println("the price must be a number")
            promptOptions(b)
        }
        b.addItem(name,p)
        fmt.Println("item added - ", name, price)
        promptOptions(b)

    case "t" :
        tip, _ := getInput("enter tip amount ($):" , reader)

        t,err :=strconv.ParseFloat(tip, 64)
        if err != nil{
            fmt.Println("the tip must be a number")
            promptOptions(b)
        }
        b.updateTip(t)
         fmt.Println("tip added - ", tip)
        promptOptions(b)

    case "s" :
        b.save()
        fmt.Println("saved the file", b.name) 
    default:
        fmt.Println("not a valid option.")
        promptOptions(b)

    }
    
}





func main() {
    mybill := createBill()
    promptOptions(mybill)

    fmt.Println(mybill)
   



   // menus()
    //sayHi("ruby")
    //test()
    //fmt.Println("Hello, world!")

/*
    //strings
    var nameOne string = "mario"
    var nameTwo = "luigi"

    var nameThree string
    nameThree = "mangionie"
    nameOne ="peach"

    fmt.Println(nameOne,nameTwo, nameThree)


    nameFour := "yoshi" //only when initializing, also can only be declared inside a function

    fmt.Println(nameFour)



    //integers

    var ageOne int = 20
    var ageTwo = 30
    ageThree := 40


    //bits & memory
    var numOne int8 = 25  //8bit   -128 to 128
    var numTwo int8 = 127  

   //unsigned int - cant have a -ve number, uint8-can go beyond 128 cause -ves arent included so 0-255
    var numThree uint = 25


    //float
    var scoreOne float32 = 23.98
    var scoreTwo float64 = 23.98
    scoreThree := 24.2222


    fmt.Println(ageOne,ageTwo, ageThree, numOne, numTwo, numThree, scoreOne, scoreTwo, scoreThree)

    */


    //fmt
   
   /* fmt.Print("hello, ")  //doesnt add a new line
    fmt.Print("world! ")
    fmt.Print("heyyy \nworld")

    fmt.Println("hello ninjas!")
    fmt.Println("bye ninjas")    

    age := 25
    name := "shaun"

    fmt.Println("my age is", age, "and my name is", name)

    //formatted strings -Printf
    //%vor %_ the underscore can be many different letter(e.g u can use %q to add quotations), -> format specifier

    fmt.Printf("my age is %v and my name is %v \n", age, name)

    fmt.Printf("age is of type %T", age)
    // %f is for float, %0.1f  -round to 1 decimal point

    //sprintf (save formatted strings)

    var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
    fmt.Println("the saved string is: ", str)

   //arrays and slices

   //array-fixed length and cant change

   var ages [3]int = [3]int{20,25,30}
   var nums = [3]int{20,24,30}  //or 
   names := [3]string{"yoshi", "mario", "peach"}

   fmt.Println(ages,nums, names)
   fmt.Println(len(nums))  //gives us lenght of nums array
    
   //slices-typical array from other languages, is an array under the hood tho

   var scores = []int{100,23,33}
   scores[2] = 34
   scores = append(scores,55)

   fmt.Println(scores)

   //slice ranges - get a range from an array/slice

   rangeOne := names[1:2] //1 ibclusive, 2 not
   rangeTwo := names[2:] //from 2 till end
   rangeThree := names[:2] // from start to 2 not including 2
   fmt.Println(rangeOne, rangeTwo, rangeThree) 
   
   //we can append ranges
   rangeOne = append(rangeOne, "cooper")
*/
   


}


//variables-store informational data like strings, numbers arrays, etc
