package main

import (
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestSaveAndHistory(t *testing.T) {
	journal := NewJournal()
	if err := journal.Save("Иванов Иван Иванович", "ортопед", "2024-04-13"); err != nil {
		t.Fatal(err)
	}
	if err := journal.Save("Иванов Иван Иванович", "невролог", "2024-05-24"); err != nil {
		t.Fatal(err)
	}
	want := []Visit{{Specialty: "ортопед", Date: "2024-04-13"}, {Specialty: "невролог", Date: "2024-05-24"}}
	got, err := journal.GetHistory("Иванов Иван Иванович")
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetHistory = %v, %v; want %v", got, err, want)
	}
}

func TestGetLastVisit(t *testing.T) {
	journal := NewJournal()
	for _, date := range []string{"2024-05-20", "2023-11-10", "2024-06-01", "2024-05-01"} {
		if err := journal.Save("Иванов", "ортопед", date); err != nil {
			t.Fatal(err)
		}
	}
	if err := journal.Save("Иванов", "невролог", "2025-02-01"); err != nil {
		t.Fatal(err)
	}
	got, err := journal.GetLastVisit("Иванов", "ортопед")
	if err != nil || got != "2024-06-01" {
		t.Errorf("GetLastVisit = %q, %v; want 2024-06-01", got, err)
	}
}

func TestPatientNotFoundError(t *testing.T) {
	journal := NewJournal()
	_, err := journal.GetHistory("Неизвестный")
	if _, ok := err.(PatientNotFoundError); !ok {
		t.Errorf("GetHistory error type = %T; want PatientNotFoundError", err)
	}
	if err == nil || err.Error() != "patient not found" {
		t.Errorf("GetHistory error = %v; want patient not found", err)
	}
	_, err = journal.GetLastVisit("Неизвестный", "ортопед")
	if _, ok := err.(PatientNotFoundError); !ok {
		t.Errorf("GetLastVisit error type = %T; want PatientNotFoundError", err)
	}
}

func TestMissingSpecialty(t *testing.T) {
	journal := NewJournal()
	if err := journal.Save("Иванов", "ортопед", "2024-01-01"); err != nil {
		t.Fatal(err)
	}
	_, err := journal.GetLastVisit("Иванов", "невролог")
	if err == nil || err.Error() != "visit not found" {
		t.Errorf("missing specialty error = %v; want visit not found", err)
	}
}

func TestInvalidVisitData(t *testing.T) {
	tests := []struct{ name, specialty, date string }{
		{"", "ортопед", "2024-01-01"},
		{"Иванов", "", "2024-01-01"},
		{"Иванов", "ортопед", "2024-02-30"},
		{"Иванов", "ортопед", "2024-1-1"},
		{"Иванов", "ортопед", "2024/01/01"},
		{"Иванов", "ортопед", "wrong"},
	}
	journal := NewJournal()
	for _, tt := range tests {
		if err := journal.Save(tt.name, tt.specialty, tt.date); err == nil {
			t.Errorf("Save(%q, %q, %q) unexpectedly succeeded", tt.name, tt.specialty, tt.date)
		}
	}
	if len(journal.patients) != 0 {
		t.Errorf("invalid visits changed journal: %+v", journal.patients)
	}
}

func TestLeapYear(t *testing.T) {
	journal := NewJournal()
	if err := journal.Save("Иванов", "ортопед", "2024-02-29"); err != nil {
		t.Fatalf("valid leap date rejected: %v", err)
	}
	if err := journal.Save("Иванов", "ортопед", "2023-02-29"); err == nil {
		t.Fatal("invalid leap date accepted")
	}
}

func TestHistoryIsCopy(t *testing.T) {
	journal := NewJournal()
	if err := journal.Save("Иванов", "ортопед", "2024-01-01"); err != nil {
		t.Fatal(err)
	}
	visits, err := journal.GetHistory("Иванов")
	if err != nil {
		t.Fatal(err)
	}
	visits[0].Specialty = "изменено"
	original, err := journal.GetHistory("Иванов")
	if err != nil || original[0].Specialty != "ортопед" {
		t.Errorf("journal history modified through returned slice: %v; %v", original, err)
	}
}

func TestPatientIsolation(t *testing.T) {
	journal := NewJournal()
	for _, name := range []string{"Иванов", "Петров"} {
		if err := journal.Save(name, "ортопед", "2024-01-01"); err != nil {
			t.Fatal(err)
		}
	}
	if err := journal.Save("Иванов", "невролог", "2024-01-02"); err != nil {
		t.Fatal(err)
	}
	first, _ := journal.GetHistory("Иванов")
	second, _ := journal.GetHistory("Петров")
	if len(first) != 2 || len(second) != 1 {
		t.Errorf("patient histories mixed: Ivanov %v, Petrov %v", first, second)
	}
}

func TestConsoleInputOutput(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "ex04")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	build := exec.Command("go", "build", "-o", binary, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	tests := []struct{ name, input, want string }{
		{"README scenario", "Save\nИванов Иван Иванович\nортопед\n2024-04-13\nSave\nИванов Иван Иванович\nневролог\n2024-05-24\nGetHistory\nИванов Иван Иванович\nGetLastVisit\nИванов Иван Иванович\nортопед\n", "ортопед 2024-04-13\nневролог 2024-05-24\n2024-04-13\n"},
		{"unknown patient history", "GetHistory\nНеизвестный\n", "patient not found\n"},
		{"unknown patient latest visit", "GetLastVisit\nНеизвестный\nортопед\n", "patient not found\n"},
		{"patient exists but specialty not found", "Save\nИванов\nортопед\n2024-01-01\nGetLastVisit\nИванов\nневролог\n", "visit not found\n"},
		{"invalid date and then continue", "Save\nИванов\nортопед\n2024-02-30\nGetHistory\nИванов\n", "Invalid input\npatient not found\n"},
		{"empty patient name", "Save\n\nортопед\n2024-01-01\n", "Invalid input\n"},
		{"unknown command then successful save", "Wrong\nSave\nИванов\nортопед\n2024-01-01\nGetHistory\nИванов\n", "Invalid input\nортопед 2024-01-01\n"},
		{"out of order dates", "Save\nИванов\nортопед\n2025-01-01\nSave\nИванов\nортопед\n2023-01-01\nGetLastVisit\nИванов\nортопед\n", "2025-01-01\n"},
		{"two patients", "Save\nИванов\nортопед\n2024-01-01\nSave\nПетров\nневролог\n2024-02-02\nGetHistory\nПетров\n", "невролог 2024-02-02\n"},
		{"empty input", "", ""},
		{"unfinished save", "Save\nИванов\n", ""},
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
