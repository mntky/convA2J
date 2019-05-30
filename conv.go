package main

import (
	"fmt"

)

func main() {
	testdata := "token=vFlRfdwGp9Nk5koVtSrbM1Qw&team_id=T63EC9465&team_domain=ecc-network-club&service_id=652176051510&channel_id=CJSJ7T8PM&channel_name=test&timestamp=1559196598.002400&user_id=U643K5KD0&user_name=minatoo&text=ojisan&trigger_word=ojisan"

	var key []string
	var data []string
	var word string
	c := 0

	//n is num , d = word(mozi code)
	for _, d := range testdata {
		if d == 61 {
			word += ":" //:
			key = append(key, word)
			word = ""
			continue
		}else if d == 38 {
			word += ","
			data = append(data, word)
			word = ""
			c += 1
			continue
		}
		word += string(d)
	}

	for i:=0; i<c; i++ {
		fmt.Printf("%v%v\n",key[i],data[i])
	}

}
