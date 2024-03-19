package slicer 

func Sum(vector [] int) int {
  sum := 0; 
  for _, elem := range vector {
    sum += elem;  
  }
  return sum; 
}

func SumAllTails(vectors ...[] int) [] int {
  additions := SumAll(vectors...); 

  for idx, adds := range additions {
    if len(vectors[idx]) > 0 {
      additions[idx] = adds - vectors[idx][0]; 
    }
  }

  return additions; 
}

func SumAll(vectors ...[] int) [] int{
  var additions [] int; 
  for _, vector := range vectors{
     additions = append(additions, Sum(vector)) ; 
  }

  return additions; 
}