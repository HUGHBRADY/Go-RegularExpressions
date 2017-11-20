// Go RegExp Problem Sheet - Hugh Brady 

// Code adapted from https://stackoverflow.com/questions/33994677/pick-a-random-value-from-a-go-slice

package main

import (
	"math/rand"
	"time"
	"fmt"
	"regexp"
	"strings"
)

func elizaResponse(input string) string {

	// Array of strings 
	responses := []string{
		"Why do you say that?",
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	
	// Checks if "father" is a word in the string
	// if input contains "father"; return response
	// (?i) flags the condition as case insensitive
	if dad, _ := regexp.MatchString(`(?i)\bfather\b`, input); dad {
		return "Why don’t you tell me more about your father?"
	}

	// Compile a regular expression into a RegExp object used for matching text
	// Checks if input begins with variants “I am" and capture part of the input
	iam := regexp.MustCompile(`(?i)i(?:'| a)?m([^.!?]*)[.!?]?`)
	if iam.MatchString(input) {
		return (iam.ReplaceAllString(input, "How do you know you are$1?"))		
	}

	// Will return a random string from responses array otherwise
	return responses[rand.Intn(len(responses))]
}

func elizaReflection(input string) string {
	// List the pronouns to switch
	pronouns := [][]string{
		{`am`, `are`},
		{`I`, `you`},
		{`me`, `you`},
		{`your`, `my`},
		{`you`, `I`},
		{`my`, `your`},
	}

	// Split input into values
	boundaries := regexp.MustCompile(`\b`)

	values := boundaries.Split(input, -1)

	//Loop through the range of values and reflect the pronoun if it finds a match
	for i, token := range values {
		for _, reflection := range pronouns {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				
				values[i] = reflection[1]
				break
			}
		}
	}
	
	//Join the string of values back together
	answer := strings.Join(values, ``)

	return ("Output: How do you know that " + answer)
}

func main() {
	// Seed used to randomize Eliza response
	rand.Seed(time.Now().UTC().UnixNano())

	// Print the input to console
	// Then, print Eliza's response to the string you pass to her
	fmt.Println("Input: " + " People say I look like both my mother and father.")
	fmt.Println("Output: " + elizaResponse("People say I look like both my mother and father.")) 

	fmt.Println("\nInput: " + " Father was a teacher.")
	fmt.Println("Output: " + elizaResponse("Father was a teacher."))

	fmt.Println("\nInput: " + " I was my father’s favourite.")
	fmt.Println("Output: " + elizaResponse("I was my father’s favourite."))

	fmt.Println("\nInput: " + " I'm looking forward to the weekend.")
	fmt.Println("Output: " + elizaResponse("I'm looking forward to the weekend."))

	fmt.Println("\nInput: " + " My grandfather was French!")
	fmt.Println("Output: " + elizaResponse("My grandfather was French!"))

	fmt.Println("\nInput: " + "I AM happy.")
	fmt.Println("Output: " + elizaResponse("I AM happy."))
	
	fmt.Println("\nInput: " + "I am not happy with your responses.")
	fmt.Println("Output: " + elizaResponse("I am not happy with your responses."))

	fmt.Println("\nInput: " + "im not sure that you understand the effect that your questions are having on me.")
	fmt.Println("Output: " + elizaResponse("im not sure that you understand the effect that your questions are having on me."))

	fmt.Println("\nInput: " + "i am supposed to just take what you’re saying at face value?")	
	fmt.Println("Output: " + elizaResponse("i am supposed to just take what you’re saying at face value?"))
		
	fmt.Println("\nInput: " + "I am not sure that you understand the effect your questions are having on me.")	
	fmt.Println("Output: " + elizaResponse("I am not sure that you understand the effect your questions are having on me."))
}