package json

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	cfg, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(cfg, &data)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range data {
		fmt.Printf("%s: %v\n", k, v)
	}
}
