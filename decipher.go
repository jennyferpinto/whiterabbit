package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/pkg/profile"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

var anagramString = "poultry outwits ants"
var anagram = createAnagramSet("poultry outwits ants")

func createAnagramSet(anagram string) map[rune]int {
	set := make(map[rune]int)

	for i, ch := range anagram {
		if set[ch] == 0 {
			if unicode.IsSpace(ch) {
				// do nothing
			} else {
				set[ch] = 1
			}
		} else {
			i = set[ch]
			set[ch] = i + 1
		}
	}
	return set
}

func generateValidWords(words []string) []string {
	validWords := []string{}

	// looping through word list
	for i, w := range words {

		duplicate := false

		// check for duplicate words
		if i != len(words)-1 {
			if words[i] == words[i+1] {
				duplicate = true
			}
		}

		match := false

		// looping through chars in word
		for _, char := range w {
			match = false // need to reset the match after each iteration through anagram

			// dont add duplicates
			if duplicate == true {
				break
			}

			// checking for the char in the anagram
			for a, _ := range anagram {
				if char == a {
					match = true
					break
				}
			}
			if match == false { // break as soon as one char in word is not in anagram
				break
			}
		}
		if match == true {
			// if word doesn't have any non matching chars, then check for correct num of each character
			if tooManyLetters(createAnagramSet(w)) == false {
				validWords = append(validWords, w)
			}
		}
	}
	return validWords
}

func tooManyLetters(newWord map[rune]int) bool {

	for ch, _ := range newWord {
		if newWord[ch] > anagram[ch] {
			return true
		}
	}
	return false
}

func isEqual(newWord map[rune]int) bool {
	for ch, _ := range newWord {
		if newWord[ch] != anagram[ch] {
			return false
		}
	}
	return true
}

func isTheSecret(passcode string) bool {

	hash1 := "e4820b45d2277f3844eac66c903e84be"
	hash2 := "23170acc097c24edb98fc5488ab033fe"

	h := md5.New()
	io.WriteString(h, passcode)
	hashedPasscode := fmt.Sprintf("%x", h.Sum(nil))

	if fmt.Sprintf(hashedPasscode) == hash1 {
		fmt.Println("THE EASY SECRET CODE IS: " + passcode)
		return true
	}
	if fmt.Sprintf(hashedPasscode) == hash2 {
		fmt.Println("THE MEDIUM SECRET CODE IS: " + passcode)
		return true
	}

	return false
}

func checkThreeWordAnagrams(words []string) {
	correctLength := len(anagramString)
	count := 0

	for _, w1 := range words {
		for _, w2 := range words {

			if len(w1+" "+w2) >= correctLength {
				continue
			}
			if tooManyLetters(createAnagramSet(w1+" "+w2)) == true {
				continue
			}

			for _, w3 := range words {

				if len(w1+" "+w2+" "+w3) != correctLength {
					continue
				} else {

					if isEqual(createAnagramSet(w1+" "+w2+" "+w3)) == false {
						continue
					} else {
						if isTheSecret(w1+" "+w2+" "+w3) == true {
							count = count + 1
						}
						if count == 2 {
							return
						}
					}
				}
			}
		}
	}
}

func main() {

	defer profile.Start().Stop()

	//defer profile.Start(profile.MemProfile).Stop()

	words, err := ioutil.ReadFile("wordlist")
	if err != nil {
		fmt.Println("error reading file")
	}

	list := strings.Split(string(words), "\n")

	validWords := generateValidWords(list)

	fmt.Println("Now checking valid words for 3 word anagrams...")

	checkThreeWordAnagrams(validWords)
}
