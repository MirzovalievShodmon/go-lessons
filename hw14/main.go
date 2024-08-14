package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1
	// length, err := countCharacters("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 2
	// numOfLines, err := countLines("test2.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("Number of lines: %d\n", numOfLines)

	// 3
	// file, err := os.Open("test2.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// defer file.Close()

	// data, _ := io.ReadAll(file)
	// wordCount := countWords(string(data))
	// fmt.Println(wordCount)

	// fmt.Println(length)

	// 4
	// err := stringToFile("test2.txt", "testing")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 5
	// content, _ := readFirstLine("test2.txt")
	// fmt.Println(content)

	// 6
	// err := copyFile("test2.txt", "test4.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//7
	// err := readAndWriteToFile("test5.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 8
	// str, err := reverseReadFile("test2.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(str)

	// 9
	// err := concatenateFiles("file1.txt", "file2.txt", "conct.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 10
	// fmt.Println(fileExists("file1.txt"))

	// 11
	// c, err := countUniqueWords("file1.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(c)

	// 12
	// replaceWordInFile("file1.txt", "world", "new world")

	// 13
	// f, err := wordFrequency("file1.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%#v", f)

	// 14
	// err := reverseLines("file1.txt", "file3.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 15
	removeEmptyLines("file2.txt", "file5.txt")
}

func countCharacters(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, fmt.Errorf("Couldn't open the file: %v", err)
	}

	defer file.Close()

	data, _ := io.ReadAll(file)
	return utf8.RuneCount(data), nil

	// 2 method
	// var length int
	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanRunes)
	// for scanner.Scan() {
	// 	length++
	// }

	// return length, nil
}

func countLines(fileName string) (int, error) {
	file, err := os.Open(fileName)
	var lines int

	if err != nil {
		return 0, fmt.Errorf("Couldn't open the file %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines++
	}
	return lines, nil
}

func countWords(s string) int {
	return len(strings.Fields(s))
}

func stringToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		return fmt.Errorf("Couldn't open to %v file. %s\n", filename, err)
	}

	defer file.Close()

	n, err := file.Write([]byte(content))

	if err != nil {
		return fmt.Errorf("Coudln't write to file %s\n", err)
	}

	fmt.Println(n)
	return nil
}

func readFirstLine(fileName string) (string, error) {
	file, err := os.Open(fileName)
	var firstLine string

	if err != nil {
		return "", fmt.Errorf("Couldn't open to %v file. %s\n", fileName, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		firstLine = scanner.Text()
		break
	}
	return firstLine, nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)

	if err != nil {
		return fmt.Errorf("Couldn't open the source file %v: %s", src, err)
	}

	defer sourceFile.Close()

	destFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		return fmt.Errorf("Something went wrong with destionation file: %s", err)
	}

	_, err = io.Copy(destFile, sourceFile)

	if err != nil {
		return fmt.Errorf("Failed to copy file")
	}

	fmt.Printf("Successfully copied content of %v to %v\n", src, dst)
	return nil
}

func readAndWriteToFile(fileName string) error {
	fmt.Println("Type anything you want to save to a file(Type 'quit' to exit): ")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)

	if err != nil {
		return fmt.Errorf("Couldn't open the file: %s\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "quit" {
			break
		}
		_, err := file.Write([]byte(scanner.Text() + "\n"))

		if err != nil {
			return fmt.Errorf("Error writing to file: %s", err)
		}
	}

	fmt.Println("Successfully written to file from console.")
	return nil
}

func reverseReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return "", fmt.Errorf("Error opening the file: %s", err)
	}

	defer file.Close()
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	return strings.Join(lines, "\n"), nil
}

func concatenateFiles(file1, file2, outputFile string) error {
	f1, err := os.Open(file1)

	if err != nil {
		return err
	}

	f2, err := os.Open(file2)

	if err != nil {
		return err
	}

	defer f1.Close()
	defer f2.Close()

	outFile, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	_, err = io.Copy(outFile, f1)

	if err != nil {
		return err
	}

	_, err = io.Copy(outFile, f2)

	if err != nil {
		return err
	}

	fmt.Printf("Copied %v, %v contents to %v\n", file1, file2, outputFile)

	return nil
}

func fileExists(fileName string) bool {
	file, err := os.Open(fileName)

	defer file.Close()

	return err == nil
}

func countUniqueWords(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	words := make(map[string]int)
	var uniqueWordCount int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if _, exists := words[scanner.Text()]; !exists {
			words[scanner.Text()]++
			uniqueWordCount++
		}
	}
	return uniqueWordCount, nil
}

func replaceWordInFile(fileName, oldWord, newWord string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	var content string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	re := regexp.MustCompile("(?i)" + oldWord)
	content = re.ReplaceAllString(content, newWord)
	file.Truncate(0)
	file.Write([]byte(content))
	return nil
}

func wordFrequency(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	words := make(map[string]int)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[strings.ToLower(scanner.Text())]++

	}
	return words, nil

}

func reverseLines(fileName, outputFile string) error {
	reversedContent, err := reverseReadFile(fileName)

	if err != nil {
		return err
	}

	dst, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	dst.Write([]byte(reversedContent))
	return nil
}

func removeEmptyLines(fileName, outputFile string) error {
	file, err := os.Open(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	dst, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, 0644)

	defer dst.Close()
	var content string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) > 0 {
			content += scanner.Text() + "\n"
		}
	}

	dst.Write([]byte(content))
	return nil
}
