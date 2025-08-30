package main

import (
	"fmt"
	"os"
)

//custom type
//kinda like a class.
//struct describes a blueprint of some data. 

type bill struct{
	name string
	items map[string]float64
	tip float64
}


//make new bills
//return type-bill
func newBill(name string) bill{
	b := bill{
		name: name,
		items: map[string]float64{},
		tip:0,
	}

	return b
}

//adding methods using receiver function.

//format bill (b bill) so that it can only be called thru bill, it must accept bill first

func (b bill) format() string{
	fs := "Bill breakdown: \n"
	var total float64 = 0
	
	//list items
	for k,v := range b.items {
		fs += fmt.Sprintf("%-25v .....$%v \n", k+"", v)
		total +=v
	}
    //tip
	   fs += fmt.Sprintf("%-25v......$%v\n", "tip:", b.tip)

	//total
   fs += fmt.Sprintf("%-25v......$%0.2f", "total:", total + b.tip)
  return fs

}



//a receiver func to update tip

func (b *bill) updateTip(tip float64){
	b.tip = tip
}

//add item to bill

func (b *bill) addItem(name string, price float64){
	b.items[name]= price
}

//save bill

func (b *bill) save(){
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644) //permission

	if err!= nil{
		panic(err)
	}
	fmt.Println("bill saved to file")

}
