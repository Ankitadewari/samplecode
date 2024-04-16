package create

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {

	expireIn := time.Hour * 24
	secretKey := "ankita"

	a, err := CreateToken(expireIn, secretKey)

	if err != nil {
		t.Errorf("Failed to create token: %v", err)
	}
	fmt.Print(a, "hi")
}
