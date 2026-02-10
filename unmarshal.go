package main
import ("fmt"
"encoding/json"
)

type Person struct{
	Name string
	Age int
}
func main(){
	p:= `{
		"Name" : "Ganashree",
		"Age" :22
    }`
	var a Person 

   err := json.Unmarshal([]byte(p) , &a)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(a)
}

 