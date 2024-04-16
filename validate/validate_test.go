package validate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v "github.com/validate"
)

func TestValidate(t *testing.T) {
	t.Run("good case", func(t *testing.T) {
		tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.2OwfQX6peV5fddz3N1-klRvcTc-Fc6wYXA7WMMJQtqQ"
		secretKey := "ankita"

		err := v.Validate(tokenString, secretKey)

		assert.Equal(t, err, nil)

	})

	t.Run("bad case", func(t *testing.T) {
		tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.2OwfQX6peV5fddz3N1-klRvcTc-Fc6wYXA7WMMJQtqQ"
		secretKey := "dewari"

		err := v.Validate(tokenString, secretKey)

		assert.NotNil(t, err)
	})

}
