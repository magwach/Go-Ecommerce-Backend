package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func SecureNumericCode(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("n must be > 0")
	}
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil) // 10^n
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	
	format := fmt.Sprintf("%%0%dd", n)
	return fmt.Sprintf(format, num.Int64()), nil
}