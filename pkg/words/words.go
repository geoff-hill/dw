package words

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var wordMap = make(map[string][]string)

func WordMap() map[string][]string {
	if len(wordMap) == 0 {
		load()
	}
	return wordMap;
}

func load() {
	myfile, err := os.Open("/usr/share/dict/words")  //open the file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile)  //scan the contents of a file and print line by line
	
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		fields := strings.Fields(line)

		word := strings.ToLower(fields[0])

		var key = word;
		if len(key) > 3 {
			key = word[0:3]
		}

		var list []string
		list, ok := wordMap[key]
		if !ok {
			list = []string{}
		}
		list = append(list, word)
		wordMap[key] = list

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err) //print error if scanning is not done properly
	}
} 
