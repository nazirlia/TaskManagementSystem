package main

import "fmt"

func main() {
	m := map[string]string{
		"Azerbaijan": "Baku",
		"USA":        "Washington",
		"Germany":    "Berlin",
		"Japan":      "Tokyo",
	}

	fmt.Println(capitalFinder(m, "Japans"))

	m["Latvia"] = "Riga"
	m["Lithuania"] = "Vilnius"

	for k, v := range m {
		fmt.Printf("m[%s] = %s\n", k, v)
	}
}

func capitalFinder(m map[string]string, s string) string {
	if v, ok := m[s]; ok {
		return "Capital of " + s + " is " + v
	}
	return "Country not found in map"
}
