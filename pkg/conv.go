package conv

import (
	//"fmt"
)

func Convert(text string) (string) {
	var key []string
	var data []string
	var jd string
	c := 0
	word := "\""

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

	for i:=0; i<c; i++ {
		jd += key[i]+data[i]
	}

	return jd
}