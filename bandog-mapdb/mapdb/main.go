package main

import (
	create "mapdb/mapdb.map.definition"
)

func main() {
	create.IsValidCreateMapSentence("Create MAP { key: integer value: mapList };")
}
