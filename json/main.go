package main

import (
	"encoding/json"
	"fmt"
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

}
