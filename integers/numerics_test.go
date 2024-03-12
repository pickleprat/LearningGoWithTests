package integers

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleAdd(){
  var sum int64; 
  sum = Add(4, 10); 
  fmt.Println(sum); 
  // Output: 14
}

func TestAddition(t *testing.T){
  // testing the add function 
  t.Run("Testing the add function", func( t *testing.T){
     var x, y int64; 
     var want, got int64; 
     for i := 0; i < 10; i ++{
       x = rand.Int63(); 
       y = rand.Int63(); 
       want = x +y; 
       got = Add(x, y) ; 
       
       if want != got{
         t.Errorf("Expected %d but recieved %d", want, got); 
         break; 
       }
     }  
  })
}

