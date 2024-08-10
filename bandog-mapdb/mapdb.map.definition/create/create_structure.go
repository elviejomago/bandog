package create

import (
	"errors"
	"fmt"
	"mapdb/mapdb.encrypt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const ENCRYPT_KEY = "a1s2d3f4a1s2d3f4"
const STRUCTURE_DICTIONARY = "STRUCTURE_DICTIONARY"

func createFolderInDisk(fullDataPath string, folderName string) (string, error) {
	encryptFolderName, errEncrypt := encrypt.Encrypt([]byte(folderName))
	pathWithFolderName := filepath.Join(fullDataPath, encryptFolderName)
	if errEncrypt != nil {
		return "", fmt.Errorf("Error encrypt folder name: %w", errEncrypt)
	}

	errMkdir := os.Mkdir(pathWithFolderName, 0755)
	if errMkdir != nil {
		return "", fmt.Errorf("Error creating directory: %w", errMkdir)
	}
	return encryptFolderName, nil
}

func createOrUpdateStructureDictionary(fullDataPath string, structureName string, structureType string) (string, error) {
	encryptStructureDictionary, errEncryptDictionary := encrypt.Encrypt([]byte(STRUCTURE_DICTIONARY))
	encryptStructureName, errEncrypt := encrypt.Encrypt([]byte(structureName))
	encryptStructureType, errEncryptType := encrypt.Encrypt([]byte(structureType))
	pathWithStructureDictionary := filepath.Join(fullDataPath, encryptStructureDictionary)

	if errEncrypt != nil || errEncryptType != nil || errEncryptDictionary != nil {
		return "", fmt.Errorf("Error encrypt process name %s: %w", structureType, errEncrypt)
	}

	var structureFile *os.File
	var errCreate error
	if _, err := os.Stat(pathWithStructureDictionary); err == nil {
		fmt.Println("Found file:", pathWithStructureDictionary)
		structureFile, errCreate = os.OpenFile(pathWithStructureDictionary, os.O_APPEND|os.O_WRONLY, 0644)
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("Error verificando el archivo: %w", err)
	} else {
		structureFile, errCreate = os.Create(pathWithStructureDictionary)
		if errCreate != nil {
			return "", fmt.Errorf("Error creating Structure Dictionary: %w", errCreate)
		}
	}

	defer structureFile.Close()

	rowStructureDictionary := fmt.Sprintf(
		"%s:%s:%s\n",
		encryptStructureName,
		encryptStructureType,
		strconv.FormatInt(time.Now().Unix(), 10),
	)

	_, errWrite := structureFile.WriteString(rowStructureDictionary)
	if errWrite != nil {
		return "", fmt.Errorf("Error writing Structure Dictionary file: %w", errWrite)
	}

	return encryptStructureName, nil
}
