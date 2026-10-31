package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLine возвращает аргумент команды и признак успешного чтения.
// Если ввод оборвался, продолжать команду не из чего.
func readLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

func main() {
	// Журнал создаём один раз: все команды должны работать с одной памятью.
	journal := NewJournal()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 64*1024), 10*1024*1024)

	// Команды приходят отдельными строками. После обработки ждём следующую.
	for scanner.Scan() {
		switch strings.TrimSpace(scanner.Text()) {
		// Save последовательно считывает ФИО, специальность и дату.
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

		// Историю печатаем построчно. continue возвращает к ожиданию новой команды.
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

		// Выводим одну самую позднюю дату для указанного врача.
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
