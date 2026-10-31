package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

func main() {
	journal := NewJournal()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 64*1024), 10*1024*1024)

	for scanner.Scan() {
		switch strings.TrimSpace(scanner.Text()) {
		case "Save":
			name, ok := readLine(scanner)
			if !ok {
				return
			}
			specialty, ok := readLine(scanner)
			if !ok {
				return
			}
			date, ok := readLine(scanner)
			if !ok {
				return
			}
			if err := journal.Save(name, specialty, date); err != nil {
				fmt.Println("Invalid input")
			}

		case "GetHistory":
			name, ok := readLine(scanner)
			if !ok {
				return
			}
			visits, err := journal.GetHistory(name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, visit := range visits {
				fmt.Println(visit.Specialty, visit.Date)
			}

		case "GetLastVisit":
			name, ok := readLine(scanner)
			if !ok {
				return
			}
			specialty, ok := readLine(scanner)
			if !ok {
				return
			}
			date, err := journal.GetLastVisit(name, specialty)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(date)

		default:
			fmt.Println("Invalid input")
		}
	}
}
