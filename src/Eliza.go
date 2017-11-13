// Go RegExp Problem Sheet - Hugh Brady 

// Code adapted from https://stackoverflow.com/questions/33994677/pick-a-random-value-from-a-go-slice

package main

import "math/rand"
import "time"
import "fmt"
import "regexp"

func ElizaResponse(input string) string {

	// Array of strings 
	responses := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	// Checks if "father" is a word in the string
	dad, _ := regexp.MatchString("(?i)\\bfather\\b", input)

	// if input contains "father", return
	if (dad) {
		return("Why don’t you tell me more about your father?")
	}

	// Will return a random string from responses array otherwise
	return responses[rand.Intn(len(responses))]
	
}

func main() {
	// Seed used to randomize Eliza response
	rand.Seed(time.Now().UTC().UnixNano())

	// Print the input to console
	// Then, print Eliza's response to the string you pass to her
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father.")) 

	fmt.Println("\nFather was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))

	fmt.Println("\nI was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))

	fmt.Println("\nI'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))

	fmt.Println("\nMy grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
}