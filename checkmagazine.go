package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'checkMagazine' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY magazine
 *  2. STRING_ARRAY note
 */

func checkMagazine(magazine []string, note []string) { //magazine é o array de palavras da revista
	// Write your code here
	words := make(map[string]int)
	for _, word := range magazine {
		words[word]++
	}
	for _, noteWord := range note {
		ammount, exists := words[noteWord]
		if !exists || ammount <= 0 {
			fmt.Println("No")
			return
		}
		words[noteWord]--
	}
	fmt.Println("Yes")
}

func mainCheckMagazine() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLineCheckMagazine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkErrorCheckMagazine(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkErrorCheckMagazine(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(strings.TrimSpace(readLineCheckMagazine(reader)), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(strings.TrimSpace(readLineCheckMagazine(reader)), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
}

func readLineCheckMagazine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkErrorCheckMagazine(err error) {
	if err != nil {
		panic(err)
	}
}
