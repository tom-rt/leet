package main

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i <= j {
		save := s[i]
		s[i] = s[j]
		s[j] = save
		i++
		j--
	}
}
