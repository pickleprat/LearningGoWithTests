package mocking

import (
	"bytes"
	"fmt"
	"slices"
	"testing"
	"time"
)

func TestConfigurableSleeper(t *testing.T){
  duration := 5 * time.Second; 
  spy := &SpySleeper{}

  config := ConfigurableSleeper{
    ModifiedSleeper: spy, 
    Duration: duration, 
  }; 
  
  buffer := bytes.Buffer{}; 
  operations := [] Operation{
    Operation(1), 
    Operation(0), 
    Operation(1), 
    Operation(0), 
    Operation(1), 
    Operation(0), 
    Operation(1), 
  }
  
  Countdown(&buffer, &config); 

  if !slices.Equal(operations, config.OrderOfOperations){
    t.Errorf("Expected order of operations %v but recieved %v\n", config.OrderOfOperations, operations);  
  }  

  if spy.Calls / 4 != int(duration){
    t.Errorf("Expected number of calls to be %d but got %d", int(duration), spy.Calls/4); 
  } 
}

func TestCountdown(t *testing.T){
  buffer := bytes.Buffer{}; 
  spy := &SpySleeper{} 
  Countdown(&buffer, spy); 
  
  got := buffer.String(); 
  want := fmt.Sprintf("3\n2\n1\nBLAZINGLY FAST!\n"); 
  operations := [] Operation{
    Operation(1), 
    Operation(0), 
    Operation(1), 
    Operation(0), 
    Operation(1), 
    Operation(0), 
    Operation(1), 
  }
  
  if got != want{
    t.Errorf("Error: expected string %q but recieved %q\n", got, want);  
  } 

  if spy.Calls != endingNumber{
    t.Errorf("Error: expected number of calls to be %d but got %d\n" , endingNumber, spy.Calls)
  }

  if !slices.Equal(spy.OrderOfOperation, operations){
    t.Errorf("Expected Operations of order %v but recieved %v", operations, spy.OrderOfOperation); 
  }
}