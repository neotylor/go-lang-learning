package validators

import "testing"

func TestValidateUserInput(t *testing.T) {
	tests := []struct {
		username string
		password string
		wantErr  bool
	}{
		{"ab", "pass123!", true},
		{"validuser", "short", true},
		{"validuser", "longbutno123", true},
		{"validuser", "NoSymbol123", true},
		{"validuser", "StrongPass1!", false},
	}

	for _, tt := range tests {
		err := ValidateUserInput(tt.username, tt.password)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateUserInput(%q, %q) = %v, wantErr %v", tt.username, tt.password, err, tt.wantErr)
		}
	}
}
