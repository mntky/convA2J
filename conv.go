package main

import (
	"fmt"

)

func main() {
	testdata := "token=vFlRfdwGp9Nk5koVtSaaaa&team_id=teamaaaa&team_domain=xxxwork-club&service_id=65176051510&channel_id=JSJ7T8PM&channel_name=test&timestamp=1559196502400&user_id=43K5KD0&user_name=mntky&text=ojisan&trigger_word=ojisan"

	var key []string
	var data []string
	word := "\""
	c := 0

	//n is num , d = word(mozi code)
	for _, d := range testdata {
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
		fmt.Printf("%v%v\n",key[i],data[i])
	}

}
