package bob

import (
	"fmt"
	"regexp"
	"strings"
)

func Hey(remark string) string {
	returnMessage := ""
	lastCharacter := remark[len(remark)-1:]
	boolVal, _ := regexp.MatchString("^[A-Z]+\\?", remark)
	boolVal2, _ := regexp.MatchString("^[A-Z]\\!", remark)
	fmt.Print(boolVal, boolVal2)
	if strings.HasSuffix(remark, "?") && strings.ToUpper(remark) == remark {
		returnMessage = "Calm down, I know what I'm doing!"
	} else if boolVal || boolVal2 {
		returnMessage = "Whoa, chill out!"
	} else if lastCharacter == "?" {
		returnMessage = "Sure."
	} else if remark == "" {
		returnMessage = "Fine. Be that way!"
	} else {
		returnMessage = "Whatever."
	}
	return returnMessage
}
