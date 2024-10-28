package utils

import "testing"

func TestWildCardToRegexp(t *testing.T) {

	check := func(subject, expected string) {
		actual := WildCardToRegexp(subject)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("go-*", "^go-.*$")
	check("chris", "^chris$")
	check("chris*", "^chris.*$")
}
