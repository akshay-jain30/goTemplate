package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"plugin"
	"text/template"
)

var templatesMap map[string]*template.Template
var pathPrefix string

func initialiseTemplates(){
	wd,_ := os.Getwd()
	pathPrefix = wd+"/templates/"

	templatesMap = make(map[string]*template.Template)

	var templateFileList, _ = ioutil.ReadDir("./templates/")
	var templatesFileNameList []string

	for _,template := range templateFileList {
		templatesFileNameList = append(templatesFileNameList,pathPrefix + template.Name())
	}

	for index,templateName := range templatesFileNameList {
		//fmt.Printf("initialising template %s \n %v \n",templateName, templatesFileNameList)
		newTemplate,err := template.New(templateFileList[index].Name()).Funcs(getFuncMap()).ParseFiles(templatesFileNameList...)
		if err!=nil {
			fmt.Println(err)
			panic(err)
		}
		//fmt.Printf("new template %s %+v \n",templateName, newTemplate)
		templatesMap[templateName]= newTemplate
	}
}

var p, _ = plugin.Open("goPlugin/goPlugin.so")
var f, _ = p.Lookup("Bind")

var pluginDict map[string] *plugin.Plugin
var pluginFunc map[string] plugin.Symbol


func main(){
	//initialiseTemplates()
	//GoRecursiveTemplate(2,true)
	bindThroughPlugin(2)
	//checkVersionConflict()
	go initiateSOFiles()
	checkHotReload()
}

func initVersionTest() {

}

