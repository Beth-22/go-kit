package main

import (
	"fmt"
	//"os"
	"bufio"
	"strings"
	"strconv"
)


func readInt(reader *bufio.Reader, prompt string) int{
	for {
		fmt.Print(prompt)
		input2, _ :=reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)

		num, err :=strconv.Atoi(input2)
		if err != nil{
			fmt.Println("Invalid number, please try again")
			continue
		}
		return num
	}
}

func readStr(reader *bufio.Reader, prompt string) string{
	fmt.Print(prompt)
	input, _ :=reader.ReadString('\n')
	name := strings.TrimSpace(input)
	return name

}


type Prompt struct{
		question string
		options []string 
		answer string
}

func askQuestion(p Prompt, reader *bufio.Reader) bool{

	fmt.Println("\n" + p.question)

	optionLabels := []string {"a", "b", "c", "d"}
	for i, opt := range p.options{
		if i < len(optionLabels){
			fmt.Printf(" %s) %s\n", optionLabels[i], opt)
		}
	}
    
	fmt.Print("Your answer: ")
	input3, _ := reader.ReadString('\n')
	input3 = strings.TrimSpace(strings.ToLower(input3))


	if input3 == strings.ToLower(p.answer){
		fmt.Println("correct!")
		return true
	}else {
		fmt.Printf("Incorrect. The answer is : %s\n", strings.ToUpper(p.answer))
		return false
	}


} 

	




/*


func main(){
	
	
	reader := bufio.NewReader(os.Stdin)
	score :=0

	questions := []Prompt {
		{
			question: "What is the capital city of India?",
			options:  []string{"Delhi", "Mumbai", "Kolkata", "Chennai"},
			answer:   "a", // "Delhi"

		},
		{
            question: "What is 2 + 2?",
			options:  []string{"3", "4", "5", "22"},
			answer:   "b", // "4"
		},
		{
          
		question: "Which planet is known as the Red Planet?",
		options:  []string{"Earth", "Mars", "Jupiter", "Venus"},
		answer:   "b", // Mars
		},
		{
           question: "Which language is used to build Android apps in Jetpack Compose?",
		options:  []string{"Java", "Python", "Kotlin", "C++"},
		answer:   "c", // Kotlin
		},
	}

    fmt.Println("Welcome to the Marlin Quiz Game!")
	name :=readStr(reader, "Heyy! What's your name?")
	fmt.Printf("Heyy %v, welcome to the Marlin quiz game!\n", name)

	age := readInt(reader, fmt.Sprintf("How old are you %v?", name))
	

	if age >= 18{
		fmt.Println("Let's see what you've got")
		for _, q := range questions {
		if askQuestion(q, reader) {
			score++
		}
	}

	fmt.Printf("\n🎯 You got %d out of %d correct!\n", score, len(questions))


	}else{
		fmt.Println("sorry, you're too young for this :)")
	}

}*/