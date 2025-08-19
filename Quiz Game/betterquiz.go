package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// prompt2 struct with JSON tags for decoding
type prompt2 struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

// Loads questions from a JSON file
func loadQuestions(filename string) ([]prompt2, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var questions []prompt2
	err = json.Unmarshal(data, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

// Ask a single question, returns true if answered correctly
func askQuestion2(p prompt2, reader *bufio.Reader) bool {
	fmt.Println("\n" + p.Question)

	optionLabels := []string{"a", "b", "c", "d"}
	for i, opt := range p.Options {
		if i < len(optionLabels) {
			fmt.Printf(" %s) %s\n", optionLabels[i], opt)
		}
	}

	fmt.Print("Your answer (a/b/c/d): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == strings.ToLower(p.Answer) {
		fmt.Println("✅ Correct!")
		return true
	}

	fmt.Printf("❌ Incorrect. The answer is: %s\n", strings.ToUpper(p.Answer))
	return false
}

// Reads string input from user with prompt22
func readStr2(reader *bufio.Reader, prompt2 string) string {
	fmt.Print(prompt2)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Reads int input from user with prompt22 and retries until valid number
func readInt2(reader *bufio.Reader, prompt2 string) int {
	for {
		fmt.Print(prompt2)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Invalid number, please try again.")
			continue
		}
		return num
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	questions, err := loadQuestions("questions.json")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		return
	}

	fmt.Println("🎉 Welcome to the Marlin Quiz Game!")

	name := readStr2(reader, "Heyy! What's your name? ")
	fmt.Printf("Heyy %v, welcome to the Marlin quiz game!\n", name)

	age := readInt2(reader, fmt.Sprintf("How old are you, %v? ", name))

	if age < 18 {
		fmt.Println("Sorry, you're too young for this :)")
		return
	}

	fmt.Println("Let's see what you've got!")

	score := 0
	for _, q := range questions {
		if askQuestion2(q, reader) {
			score++
		}
	}

	fmt.Printf("\n🎯 You got %d out of %d correct!\n", score, len(questions))
}
