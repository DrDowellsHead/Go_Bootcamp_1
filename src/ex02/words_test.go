package main

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Тут проверяется сам алгоритм, без чтения из консоли.
func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name, words string
		k           int
		want        string
	}{
		{"README example", "aa bb cc aa cc cc cc aa ab ac bb", 3, "cc aa bb"},
		{"empty words", "", 3, ""},
		{"all spaces", "     ", 3, ""},
		{"K greater than unique", "cat dog cat", 10, "cat dog"},
		{"K equal to unique", "cat dog cat", 2, "cat dog"},
		{"one unique word", "same same same", 3, "same"},
		{"K zero", "a b c", 0, ""},
		{"K negative", "a b c", -1, ""},
		{"ties", "z a b z a b", 3, "a b z"},
		{"mixed frequencies and ties", "c b c a b z", 3, "b c a"},
		{"multiple spaces", "  cat   dog  cat  ", 2, "cat dog"},
		{"case sensitive", "a A a A", 2, "A a"},
		{"punctuation", "! ? ! ?", 2, "! ?"},
		{"unicode", "ёж кот ёж кот", 2, "кот ёж"},
		{"single word", "word", 1, "word"},
		{"tabs are not delimiters", "a\tb a\tb a", 2, "a\tb a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequent(tt.words, tt.k); got != tt.want {
				t.Errorf("TopKFrequent(%q, %d) = %q; want %q", tt.words, tt.k, got, tt.want)
			}
		})
	}
}

// А здесь уже собираем и запускаем всю программу.
// Так ловим ошибки main, которые вызовом TopKFrequent не обнаружить.
func TestConsoleInputOutput(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "ex02")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	// Собираем бинарник один раз для всех вариантов ввода.
	cmd := exec.Command("go", "build", "-o", binary, ".")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, output)
	}
	tests := []struct {
		name, input, want string
	}{
		{"README example", "aa bb cc aa cc cc cc aa ab ac bb\n3\n", "cc aa bb\n"},
		{"empty words and valid K", "\n3\n", "\n"},
		{"empty words and EOF", "\n", "\n"},
		{"empty words and empty K", "\n\n", "\n"},
		{"whitespace only", "    \n", "\n"},
		{"empty words and invalid K", "\nabc\n", "\n"},
		{"K larger than words", "dog cat dog\n9\n", "dog cat\n"},
		{"equal frequencies", "z a b z b a\n2\n", "a b\n"},
		{"zero K", "dog cat dog\n0\n", "\n"},
		{"negative K", "dog cat dog\n-1\n", "Invalid input\n"},
		{"non numeric K", "dog cat\nabc\n", "Invalid input\n"},
		{"extra K token", "dog cat\n1 abc\n", "Invalid input\n"},
		{"two K integers", "dog cat\n1 2\n", "Invalid input\n"},
		{"missing K", "dog cat\n", "Invalid input\n"},
		{"large line", strings.Repeat("cat ", 20000) + "\n1\n", "cat\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary)
			// Подменяем стандартный ввод данными конкретного теста.
			cmd.Stdin = strings.NewReader(tt.input)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("program failed: %v\n%s", err, output)
			}
			if got := string(output); got != tt.want {
				t.Errorf("output = %q; want %q", got, tt.want)
			}
		})
	}
}
