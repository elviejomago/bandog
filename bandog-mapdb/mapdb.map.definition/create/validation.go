package create

import (
	"mapdb/mapdb.constant/global.constant"
	"strings"
)

func IsValidCreateMapSentence(sentence string) (bool, string, string) {
	sentenceArray := strings.Split(sentence, GConstant.DELIMITER_SENTENCE)
	isValid, invalidMessage := isAllValidCreateMapSentence(sentenceArray)
	if !isValid {
		return isValid, "", invalidMessage
	}
	createMapName := findCreateMapName(sentenceArray)

	return isValid, createMapName, invalidMessage
}

func isAllValidCreateMapSentence(sentenceArray []string) (bool, string) {
	isValidLen := isValidSentenceLen(sentenceArray)
	isValidSpecialWords := areValidSpecialWords(sentenceArray, isValidLen)
	isValidDefinitionWords := areValidDefinitionWords(sentenceArray, isValidSpecialWords)

	var invalidMessage string
	if isValidLen {
		invalidMessage = "Invalid Length"
	}
	if isValidSpecialWords {
		invalidMessage = "Invalid Special Words"
	}
	if isValidDefinitionWords {
		invalidMessage = "Invalid Definition Words"
	}
	isValid := isValidLen && isValidSpecialWords && isValidDefinitionWords
	return isValid, invalidMessage
}
