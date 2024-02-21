package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	man1 := Man{"Alexander", "Pichushkin", 49, "Male", 49}
	man2 := Man{"Mikhail", "Popkov", 59, "Male", 78}
	man3 := Man{"Andrei", "Chikatilo", 57, "Male", 52}
	man4 := Man{"Javed", "Iqbal", 40, "Male", 100}
	man5 := Man{"Luis", "Garavito", 66, "Male", 143}

	var people = map[string]Man{
		"Alexander": man1,
		"Mikhail":   man2,
		"Andrei":    man3,
		"Javed":     man4,
		"Luis":      man5,
	}
	var highestCriminal Man
	var highestCriminalFound bool
	suspects := []string{"Andrei", "Mikhail", "Alexander", "Luis"}

	for _, name := range suspects {
		person, ok := people[name]
		if !ok {
			continue
		}

		if person.Crimes > highestCriminal.Crimes {
			highestCriminal = person
			highestCriminalFound = true
		}
	}
	if highestCriminalFound {
		fmt.Printf("Person with highest criminal record: %v %v with proven %v victims", highestCriminal.Name, highestCriminal.LastName, highestCriminal.Crimes)
	} else {
		fmt.Println("No information about given suspects found in the database")
	}
}
