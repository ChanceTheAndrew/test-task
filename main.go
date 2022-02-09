package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var enLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ruLetterRunes = []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
var numberOfFiles int = 10
var numberOfLinesInFile int = 100

type lineResult struct {
	randomDate           string
	randomLatinSymbols   string
	randomRussianSymbols string
	randomInt            int
	randomFloat          float64
}

func (a *lineResult) CreateRandDate() string {

	const layout = "2006-01-02"
	min := time.Date(2017, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	h := time.Unix(sec, 0)
	a.randomDate = h.Format(layout)
	return a.randomDate

}

func (a *lineResult) CreateRandomLatinSymbols() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = enLetterRunes[rand.Intn(len(enLetterRunes))]
	}
	a.randomLatinSymbols = string(b)
	return a.randomLatinSymbols
}

func (a *lineResult) CreateRandomRussianSymbols() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = ruLetterRunes[rand.Intn(len(ruLetterRunes))]
	}
	a.randomLatinSymbols = string(b)
	return a.randomLatinSymbols
}

func (a *lineResult) CreateRandomInt() int {

	number := rand.Intn(100000000)
	a.randomInt = number
	return a.randomInt
}

func (a *lineResult) CreateRandomFloat() float64 {

	var min float64 = 1
	var max float64 = 20
	f := rand.Float64()*(max-min) + min
	formatResult := strconv.FormatFloat(f, 'f', 8, 64)
	value, err := strconv.ParseFloat(formatResult, 64)
	if err != nil {
		panic("error")
	}
	a.randomFloat = value
	return a.randomFloat
}

func (a *lineResult) CreateLine() string {
	floatToString := fmt.Sprintf("%.8f", a.CreateRandomFloat())
	intToString := strconv.Itoa(a.CreateRandomInt())
	result := a.CreateRandDate() + "||" + a.CreateRandomRussianSymbols() + "||" + a.CreateRandomLatinSymbols() + "||" + intToString + "||" + floatToString
	return result
}

func (a *lineResult) CreateAndWriteFile() {
	var filename string
	var title string
	//CreateLine()

	for i := 1; i <= numberOfFiles; i++ {
		title = strconv.Itoa(i)
		filename = title + ".txt"
		//fmt.Println(filename)
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		for j := 0; j < numberOfLinesInFile; j++ {

			_, err2 := f.WriteString(a.CreateLine() + "\n")

			if err2 != nil {
				log.Fatal(err2)
			}
		}
		fmt.Println(filename, "done")
	}

}

//--------------------2------------------
func findAndDeleteLine() {
	fmt.Println("enter the characters whose lines you want to delete:")
	var a string
	count := 0
	sum := 0

	fmt.Scan(&a)
	for i := 1; i <= numberOfFiles; i++ {
		title := strconv.Itoa(i)
		filename := title + ".txt"
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var bs []byte
		buf := bytes.NewBuffer(bs)
		scanner := bufio.NewScanner(file)
		r, err := regexp.Compile(a) // this can also be a regex

		if err != nil {
			log.Fatal(err)
		}

		for scanner.Scan() {
			if r.MatchString(scanner.Text()) == false {
				_, err := buf.Write(scanner.Bytes())
				if err != nil {
					log.Fatal(err)
				}
				sum = sum + 1
				count = numberOfLinesInFile*numberOfFiles - sum

				_, err = buf.WriteString("\n")
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(filename, buf.Bytes(), 0666)
		if err != nil {
			log.Fatal(err)
		}

		//for scanner.Scan() {
		//	if r.MatchString(scanner.Text()) == true {
		//		count = count + 1
		//	}
		//}
	}
	fmt.Println("number of lines removed: ", count)
}

func combineFile() {
	findAndDeleteLine()
	var buf bytes.Buffer
	for i := 1; i <= 10; i++ {
		title := strconv.Itoa(i)
		filename := title + ".txt"
		b, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		buf.Write(b)
	}

	err := os.WriteFile("combineFile.txt", buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		title := strconv.Itoa(i)
		filename := title + ".txt"
		os.Remove(filename)
	}
}

func main() {
	a := lineResult{}
	a.CreateAndWriteFile()
	combineFile()
	fmt.Println(a)
}
