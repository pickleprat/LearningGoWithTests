package dictionary

import (
	"testing"
) 

func assertNilError(t testing.TB, err error){
  t.Helper(); 
  if err != nil {
      t.Errorf("Error should be nil but recieved %q", err); 
  }
}

func assertDefinition(t testing.TB, definition, meaning string){
  t.Helper(); 
  if definition != meaning {
    t.Errorf("Error word definition %q != %q", definition, meaning); 
  }
}

func assertUnexpectedError(t testing.TB, want, got error){
  t.Helper(); 
  if want != got {
    t.Errorf("Recieved Error %q however wanted %q", got, want); 
  }
}

func TestDelete(t *testing.T){
  var dictionary Dictionary = Dictionary {
    "pickle": "saucy cuisine example: deez nutz", 
  }
  word := "pickle"; 
  dictionary.Delete(word); 

  _, err := dictionary.Search(word); 
  if err != ErrWordNotFound {
    t.Errorf("Word Not deleted successfully! recieved error %q", err); 
  }
  
}

func TestModify(t *testing.T){
  var dictionary Dictionary = Dictionary {
    "pickle": "saucy cuisine example: deez nutz", 
  }
  
  t.Run("Testing proper modification", func(t *testing.T){
    word := "pickle"; 
    meaning := "a man's shlong."
    
    err := dictionary.Modify(word, meaning)
    assertNilError(t, err); 

    definition, err := dictionary.Search(word); 
    assertNilError(t, err); 
    assertDefinition(t, definition, meaning); 
  })


  t.Run("Checking for Failed Modifications", func(t *testing.T){
    word := "pickle"; 
    meaning := "a man's shlong."; 
    
    err := dictionary.Modify(word, meaning); 
    assertUnexpectedError(t, ErrUnchangedMeaning, err); 

    err = dictionary.Modify("pickleboy", "A guy gives his shlong to other people"); 
    assertUnexpectedError(t, ErrWordNotExists, err); 
    
  })
}
  


func TestAdd(t *testing.T){
  // Testing if adding to the dictionary works. 
  var dictionary Dictionary = Dictionary {
    "pickle": "saucy cuisine example: deez nutz", 
  }
  
  t.Run("Successful Addition", func(t *testing.T){
    word := "soy-boy" 
    meaning := "A untrustworthy incompetent man with nothing to offer. Example: Pratyush Rao"
    err := dictionary.Add(word, meaning); 
    assertNilError(t, err); 

    definition, err := dictionary.Search(word); 
    assertNilError(t, err); 
    assertDefinition(t, definition, meaning); 
    
    
  })

  t.Run("Unsuccessful Addition", func(t *testing.T){
    word := "pickle" 
    meaning := "saucy cuisine example: deez nutz"

    err := dictionary.Add(word, meaning); 
    assertUnexpectedError(t, ErrWordExists, err); 
  })
  
}

func TestSearch(t *testing.T) {
  // testing to see if the word found is appropriate or not
  var dictionary Dictionary = Dictionary {
    "pickle": "saucy cuisine example: deez nutz", 
  }
  
  t.Run("FindingWord", func(t *testing.T){
    // testing for finding the word. 
    got, err := dictionary.Search("pickle"); 
    want := "saucy cuisine example: deez nutz"; 

    if got != want {
      t.Errorf("Incorrect meaning got: %q when wanted: %q", got, want); 
    } 

    assertNilError(t, err);  

    t.Run("Getting proper error", func(t *testing.T){
      // test for getting the appros error 
      word := "transformer"
      got, err := dictionary.Search(word); 

      if got != "" {
        t.Errorf("Word %q should not exist but recieved %q", word, got)
      }
      assertUnexpectedError(t, ErrWordNotFound, err); 
    })

  })
}