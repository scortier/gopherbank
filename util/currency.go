package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	INR = "INR"
)

// IsSupportCurrency checks if currency is supported
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, INR:
		return true
	}
	return false
}
