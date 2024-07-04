package create

import (
	"fmt"
)

func CreateMap(fullDataPath string, mapName string) (bool, error) {
	_, errCreateFolder := createFolderInDisk(fullDataPath, mapName)
	if errCreateFolder != nil {
		return false, fmt.Errorf("Error creating folder for map: %w", errCreateFolder)
	}
	_, errDictionary := createOrUpdateStructureDictionary(fullDataPath, mapName, "MAP")
	if errDictionary != nil {
		return false, fmt.Errorf("Error creating or update map dictionary for map: %w", errDictionary)
	}
	return true, nil
}
