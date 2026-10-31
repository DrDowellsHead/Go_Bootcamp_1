package main

import (
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// Проверяем порядок, повторы, пустые списки и отрицательные числа.
func TestIntersection(t *testing.T) {
	tests := []struct {
		name   string
		first  []int
		second []int
		want   []int
	}{
		{"README example", []int{5, 3, 4, 2, 1, 6}, []int{6, 4, 2, 4}, []int{4, 2, 6}},
		{"no common numbers", []int{1, 2}, []int{3, 4}, nil},
		{"empty first", nil, []int{1, 2}, nil},
		{"empty second", []int{1, 2}, nil, nil},
		{"both empty", nil, nil, nil},
		{"first order", []int{9, 3, 5}, []int{5, 9, 3}, []int{9, 3, 5}},
		{"duplicates in first", []int{2, 2, 1, 2}, []int{2, 1}, []int{2, 1}},
		{"duplicates in second", []int{2, 1}, []int{2, 2, 1}, []int{2, 1}},
		{"negative numbers", []int{-3, 5, -1}, []int{-1, -3}, []int{-3, -1}},
		{"zero", []int{0, 2, 1}, []int{1, 0}, []int{0, 1}},
		{"one common", []int{4, 1, 2}, []int{2}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection(tt.first, tt.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection(%v, %v) = %v; want %v", tt.first, tt.second, got, tt.want)
			}
		})
	}
}

// Ошибки преобразования проверяем отдельно от поиска пересечения.
func TestParseNumbers(t *testing.T) {
	tests := []struct {
		line string
		want []int
		bad  bool
	}{
		{"1 2 3", []int{1, 2, 3}, false},
		{" -2   0 +3  ", []int{-2, 0, 3}, false},
		{"\t2\t3", []int{2, 3}, false},
		{"", nil, false},
		{"   ", nil, false},
		{"3 abc", nil, true},
		{"2.0 1", nil, true},
		{"999999999999999999999999999999999999999", nil, true},
	}
	for _, tt := range tests {
		got, err := parseNumbers(tt.line)
		if (err != nil) != tt.bad {
			t.Errorf("parseNumbers(%q) error = %v, want invalid %v", tt.line, err, tt.bad)
			continue
		}
		if !tt.bad && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("parseNumbers(%q) = %v, want %v", tt.line, got, tt.want)
		}
	}
}

// Здесь проверяем реальный stdout готовой программы, включая формат вывода.
func TestConsoleInputOutput(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "ex03")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	build := exec.Command("go", "build", "-o", binary, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	tests := []struct{ name, input, want string }{
		{"README example", "5 3 4 2 1 6\n6 4 2 4\n", "4 2 6\n"},
		{"no common", "1 2\n3 4\n", "Empty intersection\n"},
		{"empty first", "\n1 2\n", "Empty intersection\n"},
		{"empty second", "1 2\n\n", "Empty intersection\n"},
		{"both empty", "\n\n", "Empty intersection\n"},
		{"invalid first", "1 abc\n1 2\n", "Invalid input\n"},
		{"invalid second", "1 2\n3.14\n", "Invalid input\n"},
		{"overflow int", "1 2\n999999999999999999999999999999\n", "Invalid input\n"},
		{"negative and zero", "0 -1 8\n-1 0\n", "0 -1\n"},
		{"repeated first", "3 3 1 3\n3 1\n", "3 1\n"},
		{"first order", "7 2 1\n1 7 2\n", "7 2 1\n"},
		{"extra spaces", "   7    2  \n 2   7  \n", "7 2\n"},
		{"first missing", "", "Invalid input\n"},
		{"second missing", "2 3\n", "Invalid input\n"},
		{"long line", strings.Repeat("2 ", 20000) + "\n2\n", "2\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary)
			cmd.Stdin = strings.NewReader(tt.input)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("program error: %v\n%s", err, out)
			}
			if string(out) != tt.want {
				t.Errorf("output = %q; want %q", out, tt.want)
			}
		})
	}
}
