package main

import "fmt"

func main()  {
	dict := make(map[string]string)
	dict["go"] = "Golang"
	dict["cs"] = "C#"
	dict["rb"] = "Ruby"
	dict["py"] = "Python"
	dict["js"] = "JavaScript"

	for k, v := range dict {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	lan, ok := dict["fk"]
	fmt.Println(lan, ok)

	if lan, ok := dict["go"]; ok {
		fmt.Println(lan, ok)
	}
}