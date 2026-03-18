package reading_writing

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	data := `name,age,grade
Alice,20,85
Bob,22,55
Carol,21,70`
	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = os.WriteFile("data.csv", []byte(data), 0666)
	if err != nil {
		fmt.Println(err)
	}

	// Part 2

	reader, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(reader))
	lines := strings.Split(string(reader), "\n")
	var students []string
	for i, line := range lines {
		splitLine := strings.Split(line, ",")
		if i == 0 {
			continue
		}
		name := splitLine[0]
		grade, err := strconv.Atoi(splitLine[2])
		if err != nil {
			fmt.Println(err)
		}
		if grade > 60 {
			students = append(students, name)
		}
	}
	fmt.Println("Students with passing grades:")
	for _, student := range students {
		fmt.Printf("- %s\n", student)
	}
}
