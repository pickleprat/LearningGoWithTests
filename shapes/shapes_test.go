package shapes

import (
	"math"
	"testing"
)

type Shape interface {
  Area() float64
}

func TestArea(t *testing.T){
  tests := [] struct {
        Name string 
        Obj Shape 
        Expected float64
      } {
      { Name: "Circle", Obj: Circle { Radius: 1} , Expected: math.Pi }, 
      { Name: "Circle", Obj:  Circle { Radius:10.4}, Expected: 339.79466141227203 }, 
      { Name: "Circle", Obj: Circle { Radius: 69.69} , Expected: 15257.7607884782},
      { Name: "Rectangle", Obj: Rectangle{ Length: 2, Breadth: 2}, Expected: 4.0}, 
      { Name: "Rectangle", Obj: Rectangle{ Length: 5, Breadth: 3}, Expected: 15.0}, 
      { Name: "Rectangle", Obj: Rectangle{ Length: 69.0, Breadth: 2.0}, Expected: 138.0},
      { Name: "Triangle", Obj: Triangle { Base:4 , Height:5 }, Expected: 10}, 
      { Name: "Triangle", Obj: Triangle { Base:4 , Height:34.5 }, Expected: 69 }, 
    }

  // checking for shapes 
  for _, tcase := range tests{  
    t.Run(tcase.Name, func(t *testing.T){
      if tcase.Expected != tcase.Obj.Area() {
        t.Errorf("Expected value %g but recieved value %g\n", tcase.Expected, tcase.Obj.Area()); 
      }
    })
  }
      
  

  
}