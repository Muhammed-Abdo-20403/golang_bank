package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type human struct {
	name      string
	Expertise []string
}

type human2 struct {
	Name      string   `json:"name"`
	Expertise []string `json:"expertise"`
}

func main() {
	{
		// slice to JSON
		expertise := []string{"GoLang", "Ruby", "Mysql"}
		expertiseJSON, err := json.Marshal(expertise)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(expertiseJSON))
	}

	{
		expertise := map[string]string{"GoLang": "excellent", "Ruby": "Not Bad", "Mysql": "Good"}
		expertiseJSON, err := json.Marshal(expertise)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(expertiseJSON))
	}

	{
		lang := human{
			name:      "Mo Elsheikh",
			Expertise: []string{"GoLang", "Ruby", "Rails"},
		}
		langJSON, _ := json.Marshal(lang)
		fmt.Println(string(langJSON))
	}

	{
		lang := human2{
			Name:      "Mo Elshiekh",
			Expertise: []string{"GoLang", "Ruby", "Mysql"},
		}
		langJSON, _ := json.Marshal(lang)
		fmt.Println(string(langJSON))
	}

	{
		lang := &human2{
			Name:      "Mo Elshiekh",
			Expertise: []string{"GoLang", "Ruby", "Mysql"},
		}
		langJSON, _ := json.Marshal(lang)
		fmt.Println(string(langJSON))
	}

	{
		stn := `{"name":"Mo Elshiekh","expertise":["GoLang","Ruby","Mysql"]}`
		covt := []byte(stn)
		lang := human2{}
		fmt.Println(lang)
		err := json.Unmarshal(covt, &lang)
		if err != nil {
			panic(err)
		}
		fmt.Println(lang)
	}

	{
		stn := `{"name":"Mo Elshiekh","expertise":["GoLang","Ruby","Mysql"]}`
		covt := []byte(stn)
		lang := human{}
		fmt.Println(lang)
		err := json.Unmarshal(covt, &lang)
		if err != nil {
			panic(err)
		}
		fmt.Println(lang)
	}
	// Read Full File In The Memory
	{
		// bytes, err := ioutil.ReadFile("./files/test.json")
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Print(string(bytes))
	}

	{
		// // read file in chunks
		f, err := os.Open("./files/test.json")
		panicOnError(err)
		defer f.Close()

		reader := bufio.NewReader(f)
		b := make([]byte, 5)
		for {
			numberOfBytesRead, err := reader.Read(b)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Println("Error reading file:", err)
				}
				break
			}
			fmt.Println(string(b[0:numberOfBytesRead]))
			// fmt.Println(string(b))
		}
	}
	{
		f, err := os.Open("./files/test.json")
		panicOnError(err)
		defer f.Close()

		s := bufio.NewScanner(f)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		err = s.Err()
		panicOnError(err)
	}
}

func panicOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
