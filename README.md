Search algorithm in Go for finding the passwords which match two given MD5 hashes

*Information provided*

- An anagram of the password: "poultry outwits ants" 
- The MD5 Hashes:

	`hash1 := "e4820b45d2277f3844eac66c903e84be"`
	`hash2 := "23170acc097c24edb98fc5488ab033fe"`


*Step 1:*  
- Parse the alphabetically organized word file to create an array of 915,135 words

*Step 2:* 
- Generate a slice of valid words by removing words from the original list which contain caracters that are not within the given anagram. In addition to any repeated words.
- The filtering is completed within the _generateValidWords()_ method. 
- If the examined word contains an invalid character or is a repeated word then the inner loop is broken and it moves on to the next word in the array. - - If checks for invalid characters or repeated words pass then the final check is to see if the word contains too many individual valid characters.
- This check is performed by looking at the frequency map of the provided anagram and comparing it to the current words map
 - If this check also passes then the word is added to the _validWords_ slice.

*Step 3:* 
- Once the _validWords_ slice is created, the _checkThreeWordAnagrams()_ method is called with this slice. 
- This method has 3 *for* loops. These loops create three-word anagrams and check to see if they match the length of the provided anagram and have a frequency map that is equal to the frequency map of the given anagram. 
- If so, the potential password is passed to the _isTheSecret()_ method. This method creates an MD5 hash of the 3 word anagram and checks to see if it matches *hash1* or *hash2*.
- If there's a match it prints the password to the terminal and once two matches are found the loops break and the program terminates. 