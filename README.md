# convA2J
---
type application/x-www-form-urlencoded convert to JSON

## Setting

`
go get -u github.com/mntky/convA2J/pkg
`

`
import "github.com/mntky/convA2J/pkg"
`

```
test := test1=a&test2=b&test3=c&test4=d
resp := conv.Convert(test)

fmt.Println(resp)
//
{"test1":"a","test2":"b","test3":"c","test4":"d"}
```
