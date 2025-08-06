package main

import "fmt"

func main(){
    // Код является високосным, если он кратен 400, либо кратен 4, но при этом не кратен 100
    var year int
    fmt.Scan(&year)
    
    if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0){
        fmt.Println("YES")
    }else{
        fmt.Println("NO")}


}
