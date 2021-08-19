package main

import "fmt"

func main() {
	map1 := map[int]int{1: 100, 2: 200, 3: 300}
	// random order on purpose!!!
	for key, value := range map1 {
		fmt.Printf("%d: %d\n", key, value)
	}
	fmt.Println()
	for key, value := range map1 {
		fmt.Printf("%d: %d\n", key, value)
	}
	fmt.Println()

	delete(map1, 2)
	map1[4] = 400
	for key, value := range map1 {
		fmt.Printf("%d: %d\n", key, value)
	}
	fmt.Println()

	translations := make(map[locale]texts, 20)
	var nilMap map[int]int = nil
	fmt.Printf("len(translations) = %d, len(nilMap) = %d\n\n", len(translations), len(nilMap))

	translations[locale{"de", "DE"}] = texts{"Gemüse", "Kartoffeln und Tomaten"}
	translations[locale{"de", "AT"}] = texts{"Gemüse", "Erdäpfel und Paradeiser"}
	translations[locale{"en", "IE"}] = texts{"Vegetables", "Potatoes and tomatoes"}

	for l, t := range translations {
		fmt.Printf("%s: %s\n", l, t)
		if l.Language == "de" {
			t.Title = "Grünzeug" // WRONG, t is not a reference !!!
		}
	}
	for l, t := range translations {
		fmt.Printf("%s: %s\n", l, t)
	}
	fmt.Println()

	for l, t := range translations {
		fmt.Printf("%s: %s\n", l, t)
		if l.Language == "de" {
			// This is invalid
			// translations[l].Title = "Grünzeug"
			// Instead, write back t:
			t.Title = "Grünzeug"
			translations[l] = t
		}
	}
	for l, t := range translations {
		fmt.Printf("%s: %s\n", l, t)
	}
	fmt.Println()

	if polishTexts, found := translations[locale{"pl", "PL"}]; found {
		fmt.Printf("Polish: %s\n", polishTexts)
	} else {
		fmt.Println("No Polish translation found.")
	}
	fmt.Println()
}

type locale struct {
	Language string
	Country  string
}

func (l *locale) String() string {
	return fmt.Sprintf("%s_%s", l.Language, l.Country)
}

type texts struct {
	Title string
	Body  string
}

func (t *texts) String() string {
	return fmt.Sprintf("%s -> %s", t.Title, t.Body)
}
