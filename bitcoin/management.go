package bitcoin

import "fmt"

type Bitcoin int; 

type Wallet struct{
  balance Bitcoin
}

func (btc Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", btc); 
}

func (w *Wallet) Spend(amount Bitcoin) error {
  if amount > w.balance {
    return fmt.Errorf("Insufficient Funds Error! Balance = %v", w.balance); 
  }
  w.balance -= amount; 
  return nil;  
}

func (w *Wallet) Deposit(amount Bitcoin){
  w.balance += amount; 
}

func (w *Wallet) Balance() Bitcoin{
  return w.balance; 
}