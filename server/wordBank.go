package main

import (
	"fmt"
	"time"
)

type wordBankEntry struct {
	genDate  time.Time
	wordBank []string
}

// ValidateWord Madafucka!
func ValidateWord(word string) bool {
	wordBank := getWordBank()
	for _, val := range wordBank {
		fmt.Printf("input word is %s wordbank word is %s\n", word, val)
		if val == word {
			return true
		}
	}
	return false

}

var internalWordBank []*wordBankEntry

func getWordBank() []string {
	var latestBank *wordBankEntry
	if len(internalWordBank) == 0 {
		latestBank = &wordBankEntry{time.Now(), genWordBank()}
		internalWordBank = append(internalWordBank, latestBank)
	} else if latestBank = internalWordBank[len(internalWordBank)-1]; latestBank.genDate.Day() != time.Now().Day() {
		latestBank = &wordBankEntry{time.Now(), genWordBank()}
		internalWordBank = append(internalWordBank, latestBank)
	}
	return latestBank.wordBank
}

func genWordBank() []string {
	return []string{"word", "single", "fine"}
}
