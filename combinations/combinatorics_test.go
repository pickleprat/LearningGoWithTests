package combinations; 

import (
  "testing"
)

func assertIncorrectLength(t *testing.T, want, got [] string){
    if len(want) != len(got){
      t.Errorf("Got length %v expected length %v", len(got),  len(want)) ;  
    }
}

func assertVectorMatching(t *testing.T, want, got [] string){
  var exists bool = false;  
  for _, subathome := range got {
    for _, subwant := range want{
      exists = false; 
      if subathome == subwant{
        exists = true; 
        break; 
      } 
    }

    if !exists {
       t.Errorf("Could not find a string in want %q, got %q instead.", want, got);  
    }
  }
}

func TestSequencesSimple(t *testing.T){
  // testing if sequences are correct 
  t.Run("Testing for accuracy of Subsequence", func(t *testing.T){
    var got [] string; 
    expected := [] string{"abc", "bc", "c", "a", "b", "ab", ""}; 
    Sequences("abc", 0, "", &got);

    // checking for length 
    assertIncorrectLength(t, expected, got);  
    // checking for values 
    assertVectorMatching(t, expected, got)
  })
  
  // testing for string pratyush
  t.Run("Testing for string pratyush", func(t *testing.T){
    var want [] string; 
    Sequences("pratyush", 0, "", &want); 
    got := SequencesSimple("pratyush"); 

    assertIncorrectLength(t, want, got); 
    assertVectorMatching(t, want, got) 
  })

  // testing for string farheena 
  t.Run("Testing for string Farheena", func(t *testing.T){
    var want [] string; 
    Sequences("farheena", 0, "", &want); 
    got := SequencesSimple("farheena"); 
    
    assertIncorrectLength(t, want, got);
    assertVectorMatching(t, want, got); 
  })
}


func BenchmarkSequences(b *testing.B){
  names := [] string {
    "pratyush", 
    "farheena", 
    "bianca", 
    "anam", 
    "faridi", 
  }
  
  var name string; 
  var vector []string; 
  for i := 0; i < b.N; i++{
    vector = [] string{}; 
    name = names[i%len(names)]
    Sequences(name, 0, "", &vector)
  }
}