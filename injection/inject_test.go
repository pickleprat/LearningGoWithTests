package injection

import (
	"bytes"
	"testing"
)

func TestInjection(t *testing.T){
  buffer := bytes.Buffer{}
  want := "Hello Baka!"; 
  Greet(&buffer, want); 
  got := buffer.String(); 

  if want != got {
    t.Errorf("Unexpected error, message written to buffer %q does not equal message wanted %q", got, want); 
  }
}