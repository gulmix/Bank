package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	TON = "TON"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, TON:
		return true
	}
	return false
}
