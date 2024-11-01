package luhn

import (
	"testing"
)

func TestLuhnAlgorithm(t *testing.T) {
	tests := []struct {
		name         string
		input        int64
		wantFull     int64
		wantChecksum int
		wantIsValid  bool
		expectError  bool
	}{
		{
			name:         "valid credit card (Visa)",
			input:        400000000000000,
			wantFull:     4000000000000002,
			wantChecksum: 2,
			wantIsValid:  true,
		},
		{
			name:         "valid credit card (Mastercard)",
			input:        555555555555444,
			wantFull:     5555555555554444,
			wantChecksum: 4,
			wantIsValid:  true,
		},
		{
			name:         "valid Canadian SIN",
			input:        46838723,
			wantFull:     468387238,
			wantChecksum: 8,
			wantIsValid:  true,
		},
		{
			name:         "invalid credit card",
			input:        400000000000001,
			wantFull:     4000000000000018,
			wantChecksum: 8,
			wantIsValid:  false,
		},
		{
			name:         "zero input",
			input:        0,
			wantFull:     0,
			wantChecksum: 0,
			wantIsValid:  false,
			expectError:  true,
		},
		{
			name:         "single digit",
			input:        7,
			wantFull:     75,
			wantChecksum: 5,
			wantIsValid:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test GetFullNumber
			gotFull, err := FullNumber(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("GetFullNumber() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if gotFull != tt.wantFull && tt.wantIsValid {
				t.Errorf("GetFullNumber() = %v, want %v", gotFull, tt.wantFull)
			}

			// Test GetCheckDigit
			gotChecksum, err := CheckDigit(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("GetCheckDigit() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if gotChecksum != tt.wantChecksum && tt.wantIsValid {
				t.Errorf("GetCheckDigit() = %v, want %v", gotChecksum, tt.wantChecksum)
			}

			// Test IsValidNumber
			if gotIsValid := IsValid(tt.wantFull); gotIsValid != tt.wantIsValid {
				t.Errorf("IsValidNumber() = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}
