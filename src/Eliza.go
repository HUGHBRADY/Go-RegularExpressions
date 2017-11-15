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
		"Why do you say that?",
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	// Compile a regular expression into a RegExp object used for matching text

	// Checks if "father" is a word in the string
	// if input contains "father", return
	// (?i) flags the condition as case insensitive
	if dad, _ := regexp.MatchString(`(?i)\\bfather\\b`, input); dad {
		return "Why don’t you tell me more about your father?"
	}

	// Checks if input begins with “I am" and capture part of the input
	iam := regexp.MustCompile(`I am ([^.!?]*)[.!?]?`)
	if iam.MatchString(input) {
		return (iam.ReplaceAllString(input, "How do you know you are $1?"))		
	}

	// Will return a random string from responses array otherwise
	return responses[rand.Intn(len(responses))]
	
}

func main() {
	// Seed used to randomize Eliza response
	rand.Seed(time.Now().UTC().UnixNano())

	// Print the input to console
	// Then, print Eliza's response to the string you pass to her
	fmt.Println("Input: " + " People say I look like both my mother and father.")
	fmt.Println("Output: " + ElizaResponse("People say I look like both my mother and father.")) 

	fmt.Println("\nInput: " + " Father was a teacher.")
	fmt.Println("Output: " + ElizaResponse("Father was a teacher."))

	fmt.Println("\nInput: " + " I was my father’s favourite.")
	fmt.Println("Output: " + ElizaResponse("I was my father’s favourite."))

	fmt.Println("\nInput: " + " I am looking forward to the weekend.")
	fmt.Println("Output: " + ElizaResponse("I'm looking forward to the weekend."))

	fmt.Println("\nInput: " + " My grandfather was French!")
	fmt.Println("Output: " + ElizaResponse("My grandfather was French!"))

	fmt.Println("\nInput: " + "I am happy.")
	fmt.Println("Output: " + ElizaResponse("I am happy."))
	
	fmt.Println("\nInput: " + "I am not happy with your responses.")
	fmt.Println("Output: " + ElizaResponse("I am not happy with your responses."))

	fmt.Println("\nInput: " + "I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println("Output: " + ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))

	fmt.Println("\nInput: " + "I am supposed to just take what you’re saying at face value?")	
	fmt.Println("Output: " + ElizaResponse("I am supposed to just take what you’re saying at face value?"))
		
}