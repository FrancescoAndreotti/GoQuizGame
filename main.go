package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	path := SetPath()
	records := readCsvFile(path)
	StartQuiz(records)
}

func StartTimer(t time.Duration) {
	timer := time.NewTimer(t * time.Second)
	<-timer.C
	fmt.Println("\nThe time is up")
	os.Exit(1)
}

func StartQuiz(records [][]string) {
	reader := bufio.NewReader(os.Stdin)
	var lines []Problem
	lines = CsvIntoStruct(records)
	fmt.Print("Click enter to start the timer")
	_, _ = reader.ReadString('\n')
	go StartTimer(30)
	var counter int
	for _, i := range lines {
		fmt.Print(i.Question + " = ")
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)
		if answer == i.Answer {
			counter++
		}
	}
	fmt.Printf("Your score is: %d/%d", counter, len(lines))
}

func SetPath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("The path is set to default to problems.csv")
	fmt.Print("Enter the file path: ")
	path, _ := reader.ReadString('\n')
	path = strings.Replace(path, "\n", "", -1)
	if len(path) == 0 {
		path = "problems.csv"
	}
	return path
}

func CsvIntoStruct(records [][]string) []Problem {
	ret := make([]Problem, len(records))

	for i, l := range records {
		ret[i] = Problem{
			Question: l[0],
			Answer:   l[1],
		}
	}
	return ret
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

type Problem struct {
	Question string
	Answer   string
}
