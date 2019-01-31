**Description**
Search algorithm in Go for finding the passwords, from a given list of words, that match two MD5 hashes

**Information provided**
- An anagram of the password: "poultry outwits ants" 
- MD5 Hashes of the secret passwords:

- `hash1 := "e4820b45d2277f3844eac66c903e84be"`
- `hash2 := "23170acc097c24edb98fc5488ab033fe"`

**Solution:**

*Step 1:*  
- Parse the alphabetically organized word file to create an array of 915,135 words

*Step 2:* 
- Generate a slice of valid words by removing words from the original list which contain caracters that are not within the given anagram. In addition to any repeated words.

*Methodology*
- The filtering is completed within the _generateValidWords()_ method. 
- If the examined word contains an invalid character or is a repeated word, then the inner loop is broken and it moves on to the next word in the array.
- If checks for invalid characters or repeated words pass, then the final check is to see if the word contains too many individual valid characters.
- This check is performed by looking at the frequency map of the provided anagram and comparing it to the frequency map created for the currently examined word. 
- If this check also passes then the word is added to the _validWords_ slice.

*Step 3:* 
- Once the _validWords_ slice is created, the _checkThreeWordAnagrams()_ method is called in order to create viable three-word anagrams and pass them to the _isTheSecret()_ method, which checks the MD5 hash of each viable three-word anagram against the provided hashes. 

*Methodology*
- The _checkThreeWordAnagrams()_ method has 3 _for_ loops. These loops create three-word anagrams and check to see if they match the length of the provided anagram and have a frequency map equal to that of the given anagram. 
- If these criteria are met, the potential password is passed to the _isTheSecret()_ method. 
- The _isTheSecret()_ method creates an MD5 hash of the three-word anagram and checks to see if it matches *hash1* or *hash2*.
- If there's a match it prints the password to the terminal.
- Once two matches have been found the loops break and the program terminates. 