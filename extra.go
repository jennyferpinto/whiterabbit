package main

// func isPotentialAnagram(set map[rune]int) bool {

// 	eq := reflect.DeepEqual(set, anagram)
// 	if eq {
// 		return true
// 	}
// 	return false
// }

// func concatStringVariadic(words ...string) string {

// 	var strBuilder strings.Builder

// 	for i, str := range words {
// 		if i == len(words)-1 {
// 			strBuilder.WriteString(str)
// 		} else {
// 			strBuilder.WriteString(str)
// 			strBuilder.WriteString(" ")
// 		}

// 	}
// 	return strBuilder.String()
// }

// func generatePossibleFourWordAnagrams(words []string) []string {
// 	validAnagrams := []string{}
// 	for _, w1 := range words {
// 		for _, w2 := range words {
// 			twoLen := len(w1) + len(w2) + 1

// 			if twoLen >= lenAnagram {
// 				continue
// 			}

// 			concatedTwoWord := concatStringVariadic(w1, w2)
// 			if tooManyLetters(createAnagramSet(concatedTwoWord)) == true {
// 				continue
// 			}
// 			for _, w3 := range words {
// 				threeLen := len(w1) + len(w2) + len(w3) + 2

// 				if threeLen >= lenAnagram {
// 					continue
// 				}

// 				concatedThreeWord := concatString(w1, w2, w3)
// 				if tooManyLetters(createAnagramSet(concatedThreeWord)) == true {
// 					continue
// 				}

// 				for _, w4 := range words {
// 					fourLen := len(w1) + len(w2) + len(w3) + len(w4) + 3

// 					if fourLen == lenAnagram+1 {
// 						concatedFourWord := concatStringVariadic(w1, w2, w3, w4)
// 						wordHash := createAnagramSet(concatedFourWord)

// 						if tooManyLetters(wordHash) == false {
// 							validAnagrams = append(validAnagrams, concatedFourWord)
// 							// if (isTheHardSecret(concatedFourWord)) == true {
// 							// 	return
// 							// }
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return validAnagrams
// }

// func findAllKLength(words []string, k int) {
// 	n := len(words)
// 	finalAllRec(words, "", n, k)
// }

// func finalAllRec(words []string, prefix string, n int, k int) {

// 	if k == 0 {
// 		trimmed := strings.TrimSpace(prefix)

// 		if len(trimmed) == 20 {

// 			if tooManyLetters(createAnagramSet(trimmed)) == false {
// 				// fmt.Println(trimmed)
// 				// fmt.Println(len(trimmed))

// 				validAnagrams = append(validAnagrams, trimmed)

// 				// if isTheSecret(prefix) == true {
// 				// 	fmt.Println("THE PHRASE IS " + prefix)
// 				// }
// 			}
// 		}
// 		return
// 	}

// 	for i := 0; i < n; i++ {

// 		var s1 strings.Builder

// 		s1.WriteString(prefix)
// 		s1.WriteString(" ")
// 		s1.WriteString(words[i])

// 		newPrefix := s1.String()

// 		// newPrefix := prefix + " " + words[i]
// 		finalAllRec(words, newPrefix, n, k-1)
// 	}
// }
