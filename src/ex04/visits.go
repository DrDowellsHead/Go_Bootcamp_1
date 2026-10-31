package main

import (
	"fmt"
	"time"
)

// Visit — одна запись: специальность врача и дата посещения.
type Visit struct {
	Specialty string
	Date      string
}

// Отдельный тип ошибки для ситуации, когда пациента нет в журнале.
type PatientNotFoundError struct{}

// Метод Error() делает этот тип реализацией стандартного интерфейса error.
func (PatientNotFoundError) Error() string {
	return "patient not found"
}

// Для каждого ФИО в map хранится отдельный срез посещений.
type Journal struct {
	patients map[string][]Visit
}

// Инициализируем map через make: в nil-map записывать новые ключи нельзя.
func NewJournal() *Journal {
	return &Journal{patients: make(map[string][]Visit)}
}

// Save проверяет данные и добавляет одну запись в историю пациента.
// Указательный получатель позволяет методам работать с тем же Journal.
func (j *Journal) Save(name, specialty, date string) error {
	// В Go шаблон даты — конкретное число 2006-01-02.
	// Parse проверяет и саму дату, например, не пропустит 30 февраля.
	parsed, err := time.Parse("2006-01-02", date)
	// Обратное форматирование проверяет, что дата записана строго как YYYY-MM-DD.
	// После ошибки Parse правая часть условия не вычисляется благодаря ||.
	if name == "" || specialty == "" || err != nil || parsed.Format("2006-01-02") != date {
		return fmt.Errorf("invalid visit data")
	}
	// Если пациента нет, map вернёт nil-срез, но append добавит в него первую запись.
	// Результат сохраняем обратно: append может изменить длину, вместимость и адрес массива.
	j.patients[name] = append(j.patients[name], Visit{Specialty: specialty, Date: date})
	return nil
}

// GetHistory возвращает копию посещений, чтобы снаружи не менять журнал.
func (j *Journal) GetHistory(name string) ([]Visit, error) {
	// Второе значение чтения map показывает, существует ли такой ключ.
	visits, exists := j.patients[name]
	if !exists {
		return nil, PatientNotFoundError{}
	}
	// Копируем элементы в новый срез; visits... передаёт их в append по одному.
	return append([]Visit(nil), visits...), nil
}

// Ищем не последний сохранённый визит, а самый поздний по дате.
func (j *Journal) GetLastVisit(name, specialty string) (string, error) {
	// Заодно получаем готовую ошибку PatientNotFoundError, если пациента нет.
	visits, err := j.GetHistory(name)
	if err != nil {
		return "", err
	}

	last := ""
	for _, visit := range visits {
		// Даты в формате YYYY-MM-DD можно сравнивать как строки:
		// лексикографический порядок совпадает с календарным.
		if visit.Specialty == specialty && visit.Date > last {
			last = visit.Date
		}
	}
	if last == "" {
		// Пациент есть, но к этому специалисту ни разу не обращался.
		return "", fmt.Errorf("visit not found")
	}
	return last, nil
}
