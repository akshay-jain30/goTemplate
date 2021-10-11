package main

import (
	common "github.com/akshay-jain30/goTemplate/structs"
)
func ReturnCommonStruct() common.SampleStruct {
	return common.SampleStruct{
		Field1: 1,
		Field2: 2,
		Field3: 3,
	}
}