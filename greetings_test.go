package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking for a valid return value
func TestHelloName(t *testing.T) {
	name := "XWP"
	wants := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !wants.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("XWP") = %q, %v, want match for %#q`, msg, err, wants)
	}
}

// TestHelloEmpty calls greetings.Hello with a string, checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, error`, msg, err)
	}
}
