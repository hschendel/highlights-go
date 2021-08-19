package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []personalDataRecord {
		{"Hulk", "Hogan"},
		{"Bret", "Hart"},
		{"Tito", "Santana"},
		{"Eddie", "Guerrero"},
	}
	sortedData := sortedPersonalDataRecordsBySurname(data)
	sort.Sort(sortedData)

	for _, record := range sortedData {
		fmt.Printf("%s, %s\n", record.Surname, record.FirstName)
	}
}

type personalDataRecord struct {
	FirstName string
	Surname string
}

// sortedPersonalDataRecordsBySurname implements sort.Interface
type sortedPersonalDataRecordsBySurname []personalDataRecord

func (s sortedPersonalDataRecordsBySurname) Len() int {
	return len(s)
}

func (s sortedPersonalDataRecordsBySurname) Less(i, j int) bool {
	lt := s[i].Surname < s[j].Surname
	if lt {
		return true
	}
	if s[i].Surname != s[j].Surname {
		return false
	}
	return s[i].FirstName < s[j].FirstName
}

func (s sortedPersonalDataRecordsBySurname) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}
