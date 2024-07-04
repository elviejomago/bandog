package create

import (
	"mapdb/mapdb.map.definition"
	"regexp"
	"strings"
)

func arrayContains(lista []string, elemento string) bool {
	for _, item := range lista {
		if strings.ToUpper(item) == elemento {
			return true
		}
	}
	return false
}

func removeSemicolon(text string) string {
	regex := regexp.MustCompile(definition.REGEX_REMOVE_SEMICOLON)
	return regex.ReplaceAllString(text, definition.NO_SPACE)
}

func replaceBlankSpaceByDelimiter(text string) string {
	regex := regexp.MustCompile(definition.REGEX_REMOVE_BLANK_SPACES)
	return regex.ReplaceAllString(text, definition.DELIMITER_SENTENCE)
}

func isValidSentenceLen(sentenceArray []string) bool {
	if len(sentenceArray) == CREATE_SENTENCE_LEN {
		return true
	} else {
		return false
	}
}

func areValidSpecialWords(sentenceArray []string, prevValid bool) bool {
	if prevValid &&
		strings.ToUpper(sentenceArray[0]) == CREATE &&
		strings.ToUpper(sentenceArray[1]) == MAP &&
		strings.ToUpper(sentenceArray[2]) == OPEN_BRACKET &&
		strings.ToUpper(sentenceArray[9]) == CLOSE_BRACKET {
		return true
	} else {
		return false
	}
}

func areValidDefinitionWords(sentenceArray []string, prevValid bool) bool {
	if prevValid &&
		arrayContains(sentenceArray, NAME) &&
		arrayContains(sentenceArray, KEY) &&
		arrayContains(sentenceArray, VALUE) {
		return true
	} else {
		return false
	}
}
