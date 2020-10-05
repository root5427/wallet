package wallet

import (
	"testing"

	"github.com/google/uuid"
)

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

func TestService_FindPaymentByID_found(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+992001010000")
	svc.Deposit(account.ID, 20_000_00)
	if err != nil {
		t.Error(err)
	}
	payment, err := svc.Pay(account.ID, 10_000_00, "Mobile")
	if err != nil {
		t.Error(err)
	}
	_, err = svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestService_FindPaymentByID_notFound(t *testing.T) {
	svc := &Service{}
	paymentID := uuid.New().String()
	_, err := svc.FindPaymentByID(paymentID)
	if err == nil {
		t.Error(err)
	}
}
