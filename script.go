package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"script/models"
)

func main() {
	path := flag.String("path", "files/data.json", "The path to the file")

	flag.Parse()

	jsonFile, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dataArray []models.Data

	err = json.Unmarshal(byteValue, &dataArray)

	fmt.Println(err)

	for i := 0; i < len(dataArray); i++ {
		fmt.Println("Id: " + dataArray[i].Id)

		fmt.Println("Value: " + dataArray[i].Value)
	}
}
