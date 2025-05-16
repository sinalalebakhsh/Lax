package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	blacklistFile     = "blacklist.txt"
	wordListFile      = "word_list.txt"
	generated40File   = "generated_names.txt"
	initialLimit      = 40
)

var wordList []string
var generatedCount int

func loadWordsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("خطا در باز کردن فایل واژه:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}
	return words
}

func generateMeaningfulName() string {
	rand.Seed(time.Now().UnixNano())
	w1 := wordList[rand.Intn(len(wordList))]
	w2 := wordList[rand.Intn(len(wordList))]
	return strings.Title(w1) + strings.Title(w2)
}

func isNameInProject(name string, dir string) bool {
	found := false
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			content, _ := os.ReadFile(path)
			if strings.Contains(string(content), name) {
				found = true
				return filepath.SkipDir
			}
		}
		return nil
	})
	return found
}

func appendToBlacklist(name string) {
	f, err := os.OpenFile(blacklistFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString(name + "\n")
	}
}

func appendToGenerated(name string) {
	f, err := os.OpenFile(generated40File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString(name + "\n")
	}
}

func createGoFile(name string) {
	content := fmt.Sprintf(`package main

import "fmt"

func %s() {
	fmt.Println("Function %s executed.")
}
`, name, name)

	filename := name + ".go"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("خطا در ایجاد فایل:", err)
	}
}

func main() {
	wordList = loadWordsFromFile(wordListFile)
	projectDir := "."

	for {
		name := generateMeaningfulName()
		if isNameInProject(name, projectDir) {
			fmt.Println("تکراری بود:", name)
			appendToBlacklist(name)
			continue
		}

		createGoFile(name)
		fmt.Println("فایل ساخته شد با نام:", name)

		if generatedCount < initialLimit {
			appendToGenerated(name)
			generatedCount++
		}

		break
	}
}