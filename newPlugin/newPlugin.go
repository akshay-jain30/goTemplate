package main

import (
	"fmt"
	common "github.com/akshay-jain30/goTemplate/structs"
)
func ReturnCommonStruct() common.SampleStruct {
	fmt.Println("generating a new struct with v11")
	return common.SampleStruct{
		Field1: 32,
		Field2: 455,
		Field3: 6109,
	}
}

//go build -buildmode=plugin -o newPlugin.so newPlugin.go