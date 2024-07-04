package main

import (
	"fmt"
	"mapdb/mapdb.map.definition/create"
)

func main() {
	res, err := create.CreateMap("D:\\desarrollo\\workspaces\\test\\testData", "LogData3")
	fmt.Println("Resp .... ", res)
	fmt.Println("Error .... ", err)
}
