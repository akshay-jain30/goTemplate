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

func getFuncMap() template.FuncMap{
	var funMap = make(map[string]interface{})

	funMap["getFirstName"]= getFirstName
	funMap["getFamilyName"]= getFamilyName
	funMap["getCity"]= getCity
	funMap["formatAddress"]= formatAddress
	funMap["rearrangeContact"]= rearrangeContact
	funMap["getRandomString"]= getRandomString
	funMap["getRandomStringV2"]= getRandomStringV2

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

var tmplEmployee, err = template.New("employeeT").Funcs(getFuncMap()).ParseFiles("employeeT","addressT","addressFormatterT","metadataT")

func GoTemplate(e1 ComplexEmployee){
	var e2 SimpleEmployee

	if err != nil {
		fmt.Println(err)
	}

	buf := &bytes.Buffer{}
	if err != nil { panic(err) }
	err = tmplEmployee.Execute(buf, e1)
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

var tmplRecursive, _ = template.New("recursiveT").Funcs(getFuncMap()).ParseFiles("recursiveT")


func GoRecursiveTemplate(depth int){
	recTemplate := &RecursiveStruct{
		Depth: depth,
	}

	//recTJson , err := json.Marshal(recTemplate)
	buf := &bytes.Buffer{}
	if err != nil { panic(err) }
	err = tmplRecursive.Execute(buf, recTemplate)
	//fmt.Println(fmt.Printf("recTJson: %+v",string(recTJson)))
	//fmt.Println(fmt.Printf("templateOutput: %+v",buf))

	ioutil.WriteFile("filename", []byte(buf.String()), 0644)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(buf.String()),&recTemplate)
	if err != nil {
		fmt.Println(err)
	}
}



