package main
import ("fmt"
"encoding/json"
)

type Person struct{
	Name string `json : "name"`
	
	Age int `json : "age"`
	
}
func Marshal(){
	p:= Person{
		Name : "Ganashree",
		Age :22,
    }

    data, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(data))
}
