package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
)

var appName = flag.String("a", "", "app name")

func main() {
	flag.Parse()

	// app
	WirteTemplate("./cmd/generator/app.go.template", "./internal/app/"+*appName, "app.go")

	// cmd
	WirteTemplate("./cmd/generator/cmd.go.template", "./cmd/"+*appName, "main.go")

	// build
	WirteTemplate("./cmd/generator/Dockerfile.template", "./build/"+*appName, "Dockerfile")

	// config
	WirteTemplate("./cmd/generator/config.yml.template", "./configs", *appName+".yml.example")
	WirteTemplate("./cmd/generator/config.yml.template", "./configs", *appName+".yml")
}

func WirteTemplate(templatePath, targetDir, targetFileName string) {
	placeholder := []byte("$$$$$")

	template, err := ioutil.ReadFile(templatePath)
	if err != nil {
		panic(err)
	}
	template = bytes.ReplaceAll(template, placeholder, []byte(*appName))

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		os.Mkdir(targetDir, 0755)
	}
	err = ioutil.WriteFile(targetDir+"/"+targetFileName, template, 0644)
	if err != nil {
		panic(err)
	}
}
