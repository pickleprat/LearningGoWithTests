package main

import (
	"bytes"
	"fmt"
	"log"
	"main/injection"
	"main/mocking"
	"net/http"
	"os"
	"time"
)

func Server(){
  var arr [] int; 
  fmt.Println(arr) ; 
  buffer := bytes.Buffer{}
  Foolah(buffer); 
  fmt.Printf("OG buffer address : %p", &buffer); 
  port := ":6969" 
  mux := http.NewServeMux(); 
  mux.HandleFunc("/", GreetHandler); 
  fmt.Printf("Listening to port %v", port)
  err := http.ListenAndServe(port, mux); 
  if err != nil {
    log.Fatalf("Error %q cannot serve..", err)
  }
}

func Insert(nums [] int) {
  for i := 0; i < len(nums) - 1; i++{
    minNumber := nums[i]
    minIndex := i; 

    for j := i + 1; j < len(nums); j++{
      if nums[j] < minNumber{
        minNumber = nums[j]; 
        minIndex = j; 
      }
    }
    swap(nums, i, minIndex)
  }

}
func swap(nums [] int, i, j int){
  temp := nums[i]
  nums[i] = nums[j]; 
  nums[j] = temp; 
}

func Bubble(nums [] int){
  for i := len(nums); i > 0; i--{
   for j := 1; j < i; j++{
     if nums[j - 1] > nums[j]{
       swap(nums, j-1, j); 
     } 
   } 
  }
}

func doSomething(a [] int) (int, bool) {
  addr := &a; 
  fmt.Printf("The address inside the copy %p\n", addr)
  return -1, false; 
}

func Foo(a [2] int) {
  fmt.Printf("The address of the reference %p\n", &a); 
}

func Foolah(writer bytes.Buffer){
  fmt.Printf("The address of the buffer %p\n", &writer); 
}

func GreetHandler(w http.ResponseWriter, h *http.Request){
  injection.Greet(w, "Hello, Baka!"); 
}

func main() {
  // arr := [2] int {69, 420}; 
  // fmt.Printf("The address of the actual array %p\n", &arr); 
  // Foo(arr); 
  

  // a := [] int {2, 2, 2, 1, 3, 4, 3}; 
  // Insert(a); 
  // fmt.Println(a); 
  
  
} 