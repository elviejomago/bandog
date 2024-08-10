package main

import (
	"fmt"
	"mapdb/mapdb.constant/global.constant"
	"mapdb/mapdb.enum/definition.enum"
	"mapdb/mapdb.enum/structure.enum"
	"mapdb/mapdb.map.definition/create"
	"mapdb/mapdb.util/global.util"
	"strings"
)

func main() {
	sentence := "create MAP { name: LogDataResp2, key: integer, value: string }"
	sentence = GUtil.CleanAndFormatSentence(sentence)
	manageStructure(sentence)
}

func manageStructure(sentence string) {
	sentenceArray := strings.Split(sentence, GConstant.DELIMITER_SENTENCE)
	structure := strings.ToUpper(sentenceArray[1])
	switch structure {
	case EStructure.MAP:
		manageMapDefinition(sentence)
	}
}

func manageMapDefinition(sentence string) {
	sentenceArray := strings.Split(sentence, GConstant.DELIMITER_SENTENCE)
	definition := strings.ToUpper(sentenceArray[0])
	switch definition {
	case EDefinition.CREATE:
		res, err := create.CreateMap("D:\\desarrollo\\workspaces\\test\\testData", sentence)
		fmt.Println("Resp .... ", res)
		fmt.Println("Error .... ", err)
	}
}
