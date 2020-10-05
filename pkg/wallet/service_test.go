package wallet

import "testing"

func TestService_FindAccountByID_found(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+992001010000")
	if err != nil {
		t.Error(err)
	}
	_, err = svc.FindAccountByID(account.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestService_FindAccountByID_notFound(t *testing.T) {
	svc := &Service{}
	_, err := svc.RegisterAccount("+992001010000")
	if err != nil {
		t.Error(err)
	}
	var accountID int64 = 123
	_, err = svc.FindAccountByID(accountID)
	if err == nil {
		t.Errorf("User with ID %v found", accountID)
	}
}
