package injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, message string){
  fmt.Fprintf(writer, message); 
}