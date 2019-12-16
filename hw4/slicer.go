package hw4

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

var words = make(map[string]int)

// Slicer return slice of to most used words in file or string
func Slicer(s string, count int) ([]string, error) {
	_, err := os.Stat(s)
	if err != nil {
		processString(s)
	} else {
		_, err := processFile(s)
		if err != nil {
			return []string{}, err
		}
	}
	return sortWords(count), nil
}

// processString clear string and add words to map
func processString(s string) {
	clearString := clearText(s)
	calculateWordsFromString(clearString)
}

// clearText remove symbols and whitespaces
func clearText(s string) string {
	reg := regexp.MustCompile(`\.|,|'|:|\\|/|"|(\s—)|(\s-)|\d`)
	s = reg.ReplaceAllString(s, "")
	reg = regexp.MustCompile(`\s+`)
	s = reg.ReplaceAllString(s, " ")
	return strings.ToLower(s)
}

// calculateWordsFromString add every word from string to map
func calculateWordsFromString(s string) {
	temp := strings.Split(s, " ")
	for _, val := range temp {
		if len(val) > 0 {
			words[val]++
		}
	}
}

// sortWords sort words from map by their counts and return slice
func sortWords(count int) []string {
	type keyValue struct {
		Key   string
		Value int
	}

	var sortedStruct []keyValue

	for key, value := range words {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		return sortedStruct[i].Value > sortedStruct[j].Value
	})
	var result []string
	for _, keyValue := range sortedStruct[0:count] {
		result = append(result, keyValue.Key)
	}
	return result
}

// processFile will process every string from file
func processFile(s string) (bool, error) {
	file, err := os.Open(s)
	if err != nil {
		return false, err
	}
	reader := bufio.NewReader(file)
	defer file.Close()

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return false, err
			}
		}
		processString(line)
	}
	return true, nil
}
