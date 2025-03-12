package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func main() {
	type Items struct {
		Type        string
		Description string
		Value       string
	}

	type Sections struct {
		Name  string
		Items []Items
	}

	type Content struct {
		Sections []Sections
	}

	content := Content{}

	contentFile, err := os.ReadFile("content.json")
	if err != nil {
		log.Printf("failed to read the content file: %s\n", err)
		return
	}

	err = json.Unmarshal(contentFile, &content)
	if err != nil {
		log.Printf("failed to parse the content file: %s\n", err)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("failed to load template index.html:", err)
		return
	}

	file, err := os.Create("memo.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, content)
	if err != nil {
		panic(err)
	}
}
