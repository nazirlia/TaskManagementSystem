package reading_writing

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Part 1
	data := `Once upon a time, there was a brave knight.
He fought dragons and saved villages.
The people loved him dearly.`

	file, err := os.OpenFile("story.txt", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buf := bufio.NewWriter(file)
	_, err = buf.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}
	err = buf.Flush()
	if err != nil {
		fmt.Println(err)
	}

	file.Seek(0, 0) // when we write to the file, cursor's being at the end of file, so when we try to read file it returns EOF. this resets cursor's place.
	reader := bufio.NewReader(file)
	if err != nil {
		fmt.Println(err)
	}
	var lineCount int
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if line != "" {
				fmt.Print(line)
				lineCount++
			}
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(line)
		lineCount++
	}
	fmt.Printf("\nTotal lines in file: %d\n", lineCount)
}
