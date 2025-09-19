package main

import (
	"fmt"
)

type Student struct{
	id int
	name string
	marks map[string]int
}

func main(){
	s:=make([]Student,5)
    
	for i:=0;i<len(s);i++{
        fmt.Println("Let ",(i+1)," student name is:")
		fmt.Scan(&s[i].name)
		fmt.Println("Let ",(i+1)," student name is:")
		fmt.Scan(s[i].)
	}
	saveToJson
}