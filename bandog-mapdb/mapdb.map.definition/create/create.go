package create

import (
	"fmt"
	"mapdb/mapdb.map.definition"
	"strings"
)

func IsValidCreateMapSentence(sentence string) bool {
	isValid := false

	sentenceReduced := sentence
	sentenceReduced = removeSemicolon(sentenceReduced)
	sentenceReduced = replaceBlankSpaceByDelimiter(sentenceReduced)

	sentenceArray := strings.Split(sentenceReduced, definition.DELIMITER_SENTENCE)

	isValid = isValidSentenceLen(sentenceArray)
	isValid = areValidSpecialWords(sentenceArray, isValid)
	isValid = areValidDefinitionWords(sentenceArray, isValid)

	fmt.Println(sentenceArray)
	fmt.Println(len(sentenceArray))
	fmt.Println(isValid)
	return isValid
}
