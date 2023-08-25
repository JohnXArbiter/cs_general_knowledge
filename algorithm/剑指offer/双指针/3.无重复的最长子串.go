package main

func maxRepeating(sequence string, word string) (ans int) {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] == word[0] {
			j, count := i, 0
			for j+len(word) <= len(sequence) && sequence[j:j+len(word)] == word {
				count++
				j += len(word)
			}
			if ans < count {
				ans = count
			}
		}
	}
	return
}
