package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)]) //21:26
	leftOversLetter := string(runes[0 : len(letter)-key])         //0:21
	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOversLetter)
}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(hashLetter[letterPosition])
			return r
		}
		return r
	}
	strings.Map(findOne, plainText)
	return hashedString
}
func decrypt(key int, encrypttedText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(hashLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(originalLetter[letterPosition])
			return r
		}
		return r
	}
	strings.Map(findOne, encrypttedText)
	return hashedString
}

func main() {
	plainText := "ABCDE"
	fmt.Println("Plain Text", plainText)
	encrypted := encrypt(2, plainText)
	fmt.Println("Encrypted", encrypted)
	decrypted := decrypt(2, encrypted)
	fmt.Println("Decrypted", decrypted)
}
