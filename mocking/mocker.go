package mocking

import (
	"fmt"
	"io"
	"time"
)

type Operation int; 

const (
  endingNumber = 3; 
  startNumber = 1; 
  SleepOperation = Operation(0); 
  WriteOperation = Operation(1) ; 
) 

type Sleeper interface {
  Sleep(); 
  Write(); 
}

type ConfigurableSleeper struct {
  ModifiedSleeper Sleeper 
  Duration time.Duration 
  OrderOfOperations [] Operation
}

func (config *ConfigurableSleeper) Sleep(){
  for i := 0; i < int(config.Duration); i++{
    config.ModifiedSleeper.Sleep(); 
  }
  config.OrderOfOperations = append(config.OrderOfOperations, SleepOperation)
}

func (config *ConfigurableSleeper) Write() {
  config.OrderOfOperations = append(config.OrderOfOperations, WriteOperation); 
}

type SpySleeper struct {
  Calls int 
  OrderOfOperation [] Operation  
  Duration time.Duration 
}

func (spy *SpySleeper) Write(){
  spy.OrderOfOperation = append(spy.OrderOfOperation, WriteOperation); 
}

func (spy *SpySleeper) Sleep(){
  spy.Calls++; 
  spy.OrderOfOperation = append(spy.OrderOfOperation, SleepOperation); 
}

type OGSleeper struct {
  OrderOfOperation [] Operation 
  Duration time.Duration 
}

func (og *OGSleeper) Sleep(){
  time.Sleep( 1 * time.Second); 
  og.OrderOfOperation = append(og.OrderOfOperation, SleepOperation); 
}

func (og *OGSleeper) Write(){
  og.OrderOfOperation = append(og.OrderOfOperation, WriteOperation); 
}

func Countdown(buffer io.Writer, sleeper Sleeper){
  for i := endingNumber; i >= startNumber; i-- {
    sleeper.Write(); 
    fmt.Fprintln(buffer, i); 
    
    sleeper.Sleep(); 
  } 
  sleeper.Write(); 
  fmt.Fprintln(buffer, "BLAZINGLY FAST!"); 
}