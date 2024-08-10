package GUtil

import (
	"mapdb/mapdb.constant/global.constant"
	"regexp"
)

func CleanAndFormatSentence(text string) string {
	cleanAndFormatSentence := removeSemicolon(text)
	return replaceBlankSpaceByDelimiter(cleanAndFormatSentence)
}

func removeSemicolon(text string) string {
	regex := regexp.MustCompile(GConstant.REGEX_REMOVE_SEMICOLON)
	return regex.ReplaceAllString(text, GConstant.NO_SPACE)
}

func replaceBlankSpaceByDelimiter(text string) string {
	regex := regexp.MustCompile(GConstant.REGEX_REMOVE_BLANK_SPACES)
	return regex.ReplaceAllString(text, GConstant.DELIMITER_SENTENCE)
}
