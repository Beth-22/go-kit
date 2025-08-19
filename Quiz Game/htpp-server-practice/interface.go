package main

//import "fmt"

type Book interface{
	countPages() int  //pages is a method
	//if i were to add another method here and didnt implement it as a return type on the function i wanna mplement the interface on, it would be an error
}



type book1 struct{
	firstPage, lastPage int

}

type book2 struct{
	middlePage int
}

//our compiler doesnt yet know book1 and book2 accept the Book interface. no keyword, we're gonna implement it

func (b1 book1) countPages() int{
	return b1.lastPage- b1.firstPage
}
func (b2 book2) countPages() int{
	return b2.middlePage * 2
}


func calculatePage(p Book) int{
	return p.countPages()

}

/*

func main(){
	fmt.Println("heyyy")
	booky1 := book1{firstPage:1, lastPage: 240}
	booky2 := book2{middlePage: 150}


	fmt.Println("Book 1 page:", calculatePage(booky1))
	fmt.Println("Book2 page : ", calculatePage(booky2))


}

*/

