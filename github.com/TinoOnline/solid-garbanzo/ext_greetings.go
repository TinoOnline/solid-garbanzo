package main

import (
	"math/rand"
)

func ExtendedMessages(messages []string) []string  {
	extended := []string{}
	for _, message := range(messages){
		extended = append(extended, message + randomExtensions())
	}
	return extended
}

func randomExtensions() string {

	extensions := []string{
		"How are you doing?", "How was your day?", "What have you been up to?", "How has your week been going",
	}
	
	return extensions[rand.Intn(len(extensions))]
}