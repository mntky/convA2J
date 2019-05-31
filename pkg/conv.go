package conv

import (
	//"fmt"
	//"strings"
)

func Convert(text string) (string) {
	var key []string
	var data []string
	c := 0
	word := "\""
	jd := "{ "

	//n is num , d = word(mozi code)
	for _, d := range text {
		if d == 61 {
			word += "\":" //:
			key = append(key, word)
			word = "\""
			continue
		}else if d == 38 {
			word += "\","
			data = append(data, word)
			word = "\""
			c += 1
			continue
		}
		word += string(d)
	}

	word += "\""
	data = append(data, word)

	for i:=0; i<=c; i++ {
		jd += key[i]+data[i]
	}
	jd += "}"

	return jd
}
