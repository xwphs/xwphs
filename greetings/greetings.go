package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos return a map associates echo of named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map associate names with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v Welcome!",
		"Great to see you, %v",
		"Hello, %v! Well to meet",
	}
	return formats[rand.Intn(len(formats))]
}
