package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"text/template"
)

var getFirstName = func(name string) string {
	return name
}

var getFamilyName = func(name string) string {
	return name
}

var getCity = func(address Address) string {
	return address.City
}

var formatAddress = func(address Address) Address {
	address.StreetName = "*****"
	return address
}

var rearrangeContact = func(contact []string) []string {
	sort.Strings(contact)
	return contact
}

var getRandomString  = func() string {
	if rand.Intn(10)%2==1 {
		return "odd"
	}else {
		return "even"
	}
}

var getRandomStringV2  = func() string {
	if rand.Intn(10)%3==0 {
		return "divisibleby3"
	}else {
		return "notdivisibleby3"
	}
}

var invokeDynamicTemplate = func(r *RecursiveStruct) string {
	r.Depth--
	if r.Depth == 0 {
		return ""
	}else if r.Depth%2==1 {
		output,_ := executeTemplate("oddT",r)
		return output
	}else{
		output,_ := executeTemplate("evenT",r)
		return output
	}
}

func getFuncMap() template.FuncMap{
	var funMap = make(map[string]interface{})

	funMap["getFirstName"]= getFirstName
	funMap["getFamilyName"]= getFamilyName
	funMap["getCity"]= getCity
	funMap["formatAddress"]= formatAddress
	funMap["rearrangeContact"]= rearrangeContact
	funMap["getRandomString"]= getRandomString
	funMap["getRandomStringV2"]= getRandomStringV2
	funMap["invokeDynamicTemplate"]= invokeDynamicTemplate

	return funMap
}

func conversionDirect(e1 ComplexEmployee){
	var e2 SimpleEmployee

	e2.ID = e1.ID
	e2.FirstName = getFirstName(e1.Name)
	e2.FamilyName = getFirstName(e1.Name)
	e2.Address = formatAddress(e1.PrimaryAddress)
	e2.Contact = rearrangeContact(e1.Contact)

	e2.Metadata.metadata_1 = getRandomString()
	e2.Metadata.metadata_2 = getRandomString()
	e2.Metadata.metadata_3 = getRandomString()
	e2.Metadata.metadata_4 = getRandomString()
	e2.Metadata.metadata_5 = getRandomString()
	e2.Metadata.metadata_6 = getRandomString()
	e2.Metadata.metadata_7 = getRandomString()
	e2.Metadata.metadata_8 = getRandomString()
	e2.Metadata.metadata_9 = getRandomString()
	e2.Metadata.metadata_10 = getRandomString()

	//fmt.Println(fmt.Printf("employee: %+v",e2))
}

func GoTemplate(e1 ComplexEmployee){
	var e2 SimpleEmployee

	buf := &bytes.Buffer{}
	err := templatesMap[pathPrefix+"employeeT"].Execute(buf, e1)
	//fmt.Println(fmt.Printf("employee: %+v",buf))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(buf.String()), &e2)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(fmt.Printf("employee: %+v",e2))
}


func GoRecursiveTemplate(depth int, isStatic bool){
	recTemplate := &RecursiveStruct{
		Depth: depth,
	}

	buf := &bytes.Buffer{}
	var templateName string
	if isStatic {
		templateName = "recursiveStaticT"
	}else{
		templateName = "recursiveDynamicT"
	}
	err := templatesMap[pathPrefix+templateName].Execute(buf, recTemplate)

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("templateOutput buffer: %+v \n",buf)

	//Code to dump the json to file
	ioutil.WriteFile("filename", []byte(buf.String()), 0644)

	err = json.Unmarshal([]byte(buf.String()),&recTemplate)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("templateOutput json: %+v \n",recTemplate)
}

func executeTemplate(templateName string, data interface{}) (string,error){
	buf := &bytes.Buffer{}
	err := templatesMap[pathPrefix+templateName].Execute(buf, data)
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	return buf.String(),nil
}

func bindThroughPlugin(depth int){


	output := f.(func(int) string)(depth)

	ioutil.WriteFile("filename", []byte(output), 0644)

}



