package definition

import (
	"fmt"
	"regexp"
	"strings"
)

func contains(lista []string, elemento string) bool {
	for _, item := range lista {
		if strings.ToUpper(item) == elemento {
			return true
		}
	}
	return false
}

func IsValidCreateMapSentence(sentence string) bool {
	isValid := false

	// remove especial characters
	regex := regexp.MustCompile(REGEX_REMOVE_SPECIAL_CHARS)
	sentenceReduced := regex.ReplaceAllString(sentence, NO_SPACE)

	// replace blank spaces by delimiter
	regex = regexp.MustCompile(REGEX_REMOVE_BLANK_SPACES)
	sentenceReduced = regex.ReplaceAllString(sentenceReduced, DELIMITER_SENTENCE)

	sentenceArray := strings.Split(sentenceReduced, DELIMITER_SENTENCE)

	if len(sentenceArray) == 10 {
		isValid = true
	}

	if isValid &&
		strings.ToUpper(sentenceArray[0]) == "CREATE" &&
		strings.ToUpper(sentenceArray[1]) == "MAP" &&
		strings.ToUpper(sentenceArray[2]) == "{" &&
		strings.ToUpper(sentenceArray[9]) == "}" {
		isValid = true
	} else {
		isValid = false
	}

	if isValid &&
		contains(sentenceArray, "NAME") &&
		contains(sentenceArray, "KEY") &&
		contains(sentenceArray, "VALUE") {
		isValid = true
	} else {
		isValid = false
	}

	fmt.Println(sentenceArray)
	fmt.Println(len(sentenceArray))
	fmt.Println(isValid)
	return true
}
