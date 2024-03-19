package combinations; 
 
    
func SequencesSimple(str string) [] string {
  var vector [] string; 
  for i := 0; i < len(str); i++{
    for j := len(str); j > i; j--{
      vector = append(vector, str[i:j]); 
    }
  }
  vector = append(vector, ""); 
  return vector; 
}

// func Permute(arr [] int, index int, vector [] int){
//   if index == len(arr){
//     fmt.Println(vector); 
//     return; 
//   }

//   Permute(arr, index+1, append(vector, arr[index])); 
//   Permute(arr, index+1, vector); 
// }

func Sequences(str string, index int, substr string, vector *[] string){

  if index == len(str){
    *vector = append(*vector, substr); 
    return 
  }
  
  Sequences(str, index+1, substr + string(str[index]), vector); 
  if len(substr) == 0{
    Sequences(str, index+1, substr, vector); 
  }else{
    *vector = append(*vector, substr); 
  }
}

func Subsequence(str string, index int, substr string, vector *[] string ) {
  if index == len(str){
    *vector = append(*vector, substr); 
    return 
  }
  
  Subsequence(str, index+1, substr + string(str[index]), vector); 
  Subsequence(str, index+1, substr, vector); 
}
