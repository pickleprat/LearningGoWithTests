package bitcoin

import (
	"fmt"
	"testing"
)

func TestManagement(t *testing.T){
  
  checkBalance := func(w Wallet, want Bitcoin, tb testing.TB){
    tb.Helper(); 
    if w.Balance() != want {
      t.Errorf("The expected balance was %s but got %s", want, w.Balance())
    }
  }

  t.Run("TestDeposit", func (t *testing.T){
    w := Wallet{}; 
    w.Deposit(Bitcoin(100));
    checkBalance(w, Bitcoin(100), t); 

    err := w.Spend(Bitcoin(31)); 
    
    if err != nil {
      t.Errorf("Error incurred while withdrawl %s", err)
    }
    
    checkBalance(w, Bitcoin(69), t); 
  })

  t.Run("TestSpend", func(t *testing.T){
    w := Wallet{}; 
    w.Deposit(Bitcoin(1000)); 
    err := w.Spend(Bitcoin(31)); 
    if err != nil {
      t.Errorf("Error incurred while withdrawl %s", err); 
    }
    checkBalance(w, Bitcoin(969), t); 
    
    w.Deposit(Bitcoin(200)); 
    checkBalance(w, Bitcoin(1169), t); 
  })

  t.Run("TestError", func(t *testing.T){
    w := Wallet{}; 
    w.Deposit(100); 
    err := w.Spend(200); 
    if err == nil {
      t.Errorf("System crash suspected, More amount than existing deducted"); 
    }

    if err.Error() != fmt.Sprintf("Insufficient Funds Error! Balance = %v", w.Balance()){
      t.Errorf("Recieved an Error but not what was expected!"); 
    } 
    checkBalance(w, Bitcoin(100), t); 
  })
}