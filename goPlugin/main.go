package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"plugin"
)

type RecursiveStruct struct {
	Depth int
	Property_1 string
	Property_2 string
	Property_3 string
	Property_4 string
	Property_5 string
	Property_6 string
	Property_7 string
	Property_8 string
	Property_9 string
	Property_10 string
	Next *RecursiveStruct
	DynamicTemplateChecker string
}

var getRandomString  = func() string {
	if rand.Intn(10)%2==1 {
		return "odd"
	}else {
		return "even"
	}
}



func Bind(depth int) string {
	depth--
	r := &RecursiveStruct{}

	r.Depth = depth
	r.Property_1 = getRandomString()
	r.Property_2 = getRandomString()
	r.Property_3 = getRandomString()
	r.Property_4 = getRandomString()
	r.Property_5 = getRandomString()
	r.Property_6 = getRandomString()
	r.Property_7 = getRandomString()
	r.Property_8 = getRandomString()
	r.Property_9 = getRandomString()
	r.Property_10 = getRandomString()
	if depth!=0 {
		output := bindThroughPlugin(depth)
		var v RecursiveStruct
		json.Unmarshal([]byte(output),&v)
		r.Next = &v
	}

	output,_ := json.Marshal(r)
	return string(output)

}

func bindThroughPlugin(depth int) string{
	p, err := plugin.Open("goPlugin/goPlugin.so")

	if err!=nil {
		fmt.Printf("unable to initialise shared file err=%+v \n",err)
	}


	f, err  := p.Lookup("Bind")

	if err!=nil {
		fmt.Printf("unable to lookup desired function, err=%+v \n",err)
	}

	output := f.(func(int) string)(depth)
	return output
}

func main() {
	//do nothing
}
//go build -buildmode=plugin -o goPlugin.so main.go

