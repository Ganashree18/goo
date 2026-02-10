package main
import ("fmt"
"encoding/json"
)

type Person1 struct{
	Name string
	Age int
}
func Unmarshal(){
	p:= `{
		"Name" : "Ganashree",
		"Age" :22
    }`
	var a Person1 

   err := json.Unmarshal([]byte(p) , &a)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(a)
}

 