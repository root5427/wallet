package types

// Money represents a monetary amount in minimal units (cents, kopeykas, dirams, etc.).
type Money int64

// PaymentCategory represents processed payment category
type PaymentCategory string

// PaymentStatus represents payment status
type PaymentStatus string

// Predefined payment statuses
const (
	StatusOk         PaymentStatus = "OK"
	StatusFail       PaymentStatus = "FAIL"
	StatusInProgress PaymentStatus = "INPROGRESS"
)

type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
