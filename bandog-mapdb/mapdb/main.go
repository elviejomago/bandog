package main

import (
	"fmt"
	"mapdb/mapdb.encrypt"
	"mapdb/mapdb.map.definition/create"
)

func main() {
	create.IsValidCreateMapSentence("Create MAP { name: logData key: integer value: mapList };")
	enc, _ := encrypt.Encrypt([]byte("LogData"), []byte("a1s2d3f4a1s2d3f4"))
	fmt.Println("Encriptado .... " + enc)
	des, _ := encrypt.Decrypt(enc, []byte("a1s2d3f4a1s2d3f4"))
	fmt.Println("Desencriptado .... ", des)
}
