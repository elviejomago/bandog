package create

import (
	"fmt"
)

func CreateMap(fullDataPath string, sentence string) (bool, error) {
	isValidCreate, createMapName, createInvalidMessage := IsValidCreateMapSentence(sentence)
	if !isValidCreate {
		return false, fmt.Errorf("Invalid Sentence: %w %w %w", createInvalidMessage, " detail: ", createInvalidMessage)
	}
	_, errCreateFolder := createFolderInDisk(fullDataPath, createMapName)
	if errCreateFolder != nil {
		return false, fmt.Errorf("Error creating folder for map: %w", errCreateFolder)
	}
	_, errDictionary := createOrUpdateStructureDictionary(fullDataPath, createMapName, "MAP")
	if errDictionary != nil {
		return false, fmt.Errorf("Error creating or update map dictionary for map: %w", errDictionary)
	}
	return true, nil
}
