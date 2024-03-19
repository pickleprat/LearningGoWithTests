package slicer 

import (
  "testing"
  "slices"
)

// func assertSliceInequality(want, got [] int, t *testing.T){
//   if !slices.Equal(want, got){
//       t.Errorf("The function doesn't work correctly, wanted %v but recieved %v", want, got); 
//     }
// }

func TestSumAll(t *testing.T){
  vectors := [2][] int {
    {1, 2, 3, 4, 5, 6}, 
    {3, 4, 2, 1, 4}, 
  }

  want := [] int {21, 14}; 
  got := SumAll(vectors[0], vectors[1]); 

  if !slices.Equal(want, got){
    t.Errorf("Expected %v but recieved %v", want, got); 
  }
}

func TestSumAllTails(t *testing.T){
  assertSliceInequality := func(want, got [] int, t testing.TB){
    t.Helper(); 
    if !slices.Equal(want, got){
       t.Errorf("Wanted vector %v but recieved %v\n", want, got);  
    } 
  }
  t.Run("Testing for non null values", func(t *testing.T){
    a := [] int {1, 2, 3, 4, 5, 6}; 
    b := [] int {3, 4, 2, 1, 4}; 

    want := [] int {20, 11}; 
    got := SumAllTails(a, b); 

    assertSliceInequality(want, got, t); 
  })

  t.Run("Testing for the null cases", func(t *testing.T){
     a := [] int {2, 1, 3, 4, 5, 6}; 
     b := [] int {4, 3, 2, 1, 4}; 
     c := [] int{}; 

    want := [] int{19, 10, 0}; 
    got := SumAllTails(a, b, c); 
    assertSliceInequality(want, got, t); 
    
  })
}