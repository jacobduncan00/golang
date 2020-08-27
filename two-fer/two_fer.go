package twofer

import "fmt"

func ShareWith(name string) string {
	returnMessage := ""
	if name != "" {
		// returnMessage = "One for " + name + ", one for me."
		returnMessage = fmt.Sprintf("One for %s, one for me.", name)
	} else {
		returnMessage = "One for you, one for me."
	}
	return returnMessage
}
