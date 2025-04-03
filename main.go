package main

import "fmt"
 
func main() {
// Slice
var score = []int{100,200,300,400}
fmt.Println(score)
score = append(score, 500,600)
fmt.Println(score, len(score), cap(score))
}




