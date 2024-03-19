package dictionary

type Dictionary map[string]string
type DictionaryError string; 

const (
  ErrWordNotFound = DictionaryError("Word was not found in dictionary");  
  ErrWordExists = DictionaryError("Word already exists in the dictionary"); 
  ErrUnchangedMeaning = DictionaryError("Word modification is unnecessary!"); 
  ErrWordNotExists = DictionaryError("Word doesn't exist and can't be modified"); 
)

func (err DictionaryError) Error() string {
  return string(err); 
}

func (d Dictionary) Search(word string) (string, error) {
	meaning, ok := d[word]
	if !ok {
		return "", ErrWordNotFound; 
	}
  return meaning, nil; 
}

func (d Dictionary) Add(word, meaning string) error {
  _, ok := d[word]; 
  if ok {
    return ErrWordExists;  
  }
  d[word] = meaning; 
  return nil 
}

func (d Dictionary) Modify(word, meaning string) error {
  old, err := d.Search(word); 

  switch err{
    case ErrWordNotFound:
      return ErrWordNotExists
    
    case nil: 
      if old == meaning {
         return ErrUnchangedMeaning 
      }
      fallthrough
    
    default: 
       d[word] = meaning; 
       return nil 
  }
}

func (d Dictionary) Delete(word string) {
  delete(d, word); 
}