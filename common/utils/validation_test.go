package utils_test

import (
	"testing"

	"github.com/pjserol/api-rest-user/common/utils"
)

func Test_ValidateEmail(t *testing.T) {

	tests := []struct {
		name           string
		email          string
		expectedOutput bool
	}{
		{
			name:           "test happy path",
			email:          "user@test.com",
			expectedOutput: true,
		},
		{
			name:           "test email wrong format",
			email:          "@test.com",
			expectedOutput: false,
		},
		{
			name:           "test email wrong format",
			email:          "user@",
			expectedOutput: false,
		},
		{
			name:           "test email wrong format",
			email:          "test@.com",
			expectedOutput: false,
		},
		{
			name:           "test email wrong format",
			email:          "user@test",
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		res := utils.ValidateEmail(test.email)

		if test.expectedOutput != res {
			t.Errorf("for %s, expected %v, but got %v", test.name, test.expectedOutput, res)
		}
	}
}
