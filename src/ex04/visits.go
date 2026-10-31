package main

import (
	"fmt"
	"time"
)

type Visit struct {
	Specialty string
	Date      string
}

type PatientNotFoundError struct{}

func (PatientNotFoundError) Error() string {
	return "patient not found"
}

type Journal struct {
	patients map[string][]Visit
}

func NewJournal() *Journal {
	return &Journal{patients: make(map[string][]Visit)}
}

func (j *Journal) Save(name, specialty, date string) error {
	parsed, err := time.Parse("2006-01-02", date)
	if name == "" || specialty == "" || err != nil || parsed.Format("2006-01-02") != date {
		return fmt.Errorf("invalid visit data")
	}
	j.patients[name] = append(j.patients[name], Visit{Specialty: specialty, Date: date})
	return nil
}

func (j *Journal) GetHistory(name string) ([]Visit, error) {
	visits, exists := j.patients[name]
	if !exists {
		return nil, PatientNotFoundError{}
	}
	return append([]Visit(nil), visits...), nil
}

func (j *Journal) GetLastVisit(name, specialty string) (string, error) {
	visits, err := j.GetHistory(name)
	if err != nil {
		return "", err
	}

	last := ""
	for _, visit := range visits {
		if visit.Specialty == specialty && visit.Date > last {
			last = visit.Date
		}
	}
	if last == "" {
		return "", fmt.Errorf("visit not found")
	}
	return last, nil
}
