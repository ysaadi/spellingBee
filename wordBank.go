package main

import "time"

type wordBankEntry struct {
	genDate  time.Time
	wordBank []string
}

func ValidateWord(word string) bool {
	wordBank := getWordBank()
	for _, val := range wordBank {
		if val == word {
			return true
		}
	}
	return false

}

var internalWordBank []*wordBankEntry

func getWordBank() []string {
	if len(internalWordBank) == 0 {
		latestBank := &wordBankEntry{time.Now(), genWordBank()}
		internalWordBank = append(internalWordBank, latestBank)
		return latestBank.wordBank
	}
	latestBank := internalWordBank[len(internalWordBank)-1]
	latestTime := latestBank.genDate
	if latestTime.Day() != time.Now().Day() {
		latestBank = &wordBankEntry{time.Now(), genWordBank()}
		internalWordBank = append(internalWordBank, latestBank)
		return latestBank.wordBank
	}
	return latestBank.wordBank
}

func genWordBank() []string {
	return []string{"word", "single", "fine"}
}
