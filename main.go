package main

import (
	"fmt"
	"math"
	"time"
)

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func criminalSearch() {
	var people = map[string]Man{
		"Alexander": Man{"Alexander", "Pichushkin", 49, "Male", 49},
		"Mikhail":   Man{"Mikhail", "Popkov", 59, "Male", 78},
		"Andrei":    Man{"Andrei", "Chikatilo", 57, "Male", 52},
		"Javed":     Man{"Javed", "Iqbal", 40, "Male", 100},
		"Luis":      Man{"Luis", "Garavito", 66, "Male", 143},
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
		fmt.Printf("Person with highest criminal record: %v %v with proven %v victims\n", highestCriminal.Name, highestCriminal.LastName, highestCriminal.Crimes)
	} else {
		fmt.Println("No information about given suspects found in the database")
	}
}
func sliceWork() {
	someSlice := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		someSlice = append(someSlice, i)
	}
	fmt.Println(cap(someSlice), len(someSlice))

	y := someSlice[:]
	someSlice[0] = 666
	someSlice = append(someSlice, 212)
	y[1] = 666
	fmt.Println(someSlice, y)
	str := "N"
	notebook := make(map[string]string)
	notebook[str] = "N"
	if name, ok := notebook[str]; ok {
		fmt.Println(name, ok)
	}
	peopleAges := map[string]int{
		"John":   30,
		"Tony":   40,
		"Wolter": 49,
	}
	delete(peopleAges, "Tony")
	for k, v := range peopleAges {
		fmt.Println(k, v)
	}
}
func symbolBytes() {
	str := "⌘"
	fmt.Printf("Plain string: ")
	fmt.Printf("%s", str)
	fmt.Printf("\n")
	fmt.Println(fmt.Sprintf("%d", len(str)))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
}
func strToRunes() {
	name := "Sick"
	println(name[0:4])
	runes := []rune(name)
	fmt.Println(fmt.Sprintf("%#U", runes[0:4]))
}
func FizzBuzz() {
	for i := 1; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		}
	}
}
func simpleNumbers() {
	start := time.Now()
	counter := 1
	for i := 2; counter < 100000; i++ {
		c := false
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				c = true
			}
			if c {
				break
			}
		}
		if !c {
			//fmt.Println(i)
			counter++
		}
	}
	end := time.Now()
	fmt.Printf("Finished in %s\n", end.Sub(start))
}
func line(a1, a2, b1, b2 float64) {
	var x, y float64
	// y = aX*x + bX
	if a1 == a2 {
		if b1 == b2 {
			fmt.Println("Line 1 equals line 2")
			return
		}
		fmt.Println("Not crossing")
		return
	}
	x = (b2 - b1) / (a1 - a2)
	y = a1*x + b1
	fmt.Printf("x:%f, y:%f\n", x, y)
}
func Fact(n int64) int64 {
	res := int64(1)
	for i := int64(1); 1 <= n; i++ {
		res *= i
	}
	return res
}
func FactR(n int64) int64 {
	if n < 2 {
		return 1
	}
	return n * FactR(n-1)
}
func main() {
	line(2, 3, 2, 3)
}
