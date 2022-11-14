package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	f, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	idxLine := 1
	for scanner.Scan() {
		_, res, o := ProcessLine(scanner.Text(), old, new)
		occ += o
		lines = append(lines, idxLine)

		fmt.Println(res)
		idxLine += 1
	}
	return
}

// ProcessLine
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	found = false
	res = line
	line_split := strings.Split(line, " ")
	for i, v := range line_split {
		if strings.EqualFold(v, old) {
			occ += 1
			found = true
			line_split[i] = new
		}
	}
	return found, strings.Join(line_split, " "), occ
}

func main() {
	old := "Go"
	new := "Ruby"
	occ, lines, err := FindReplaceFile("wiki.txt", old, new)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("== Summary ==")
	defer fmt.Println("== End of Summary ==")
	fmt.Printf("Number of occurrences of %v: %v\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))

}
