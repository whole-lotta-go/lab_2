package lab2

import (
	"fmt"
	"testing"
)

func TestEvalPostfix(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		// Simple expressions
		{"addition", "3 4 +", 7, false},
		{"subtraction", "5 3 -", 2, false},
		{"multiplication", "2 3 *", 6, false},
		{"division", "6 2 /", 3, false},
		{"exponentiation", "2 8 ^", 256, false},

		// Complex expressions
		{"complex expression 4", "2 3 ^ 4 5 + * 6 -", 66, false},          // (2 ^ 3) * (4 + 5) - 6
		{"complex expression 5", "10 2 ^ 3 / 4 2 ^ + 5 -", 44, false},     // (10 ^ 2) / 3 + (4 ^ 2) - 5
		{"complex expression 6", "3 2 ^ 4 * 5 2 ^ 2 / + 7 -", 41, false},  // (3 ^ 2) * 4 + (5 ^ 2) / 2 - 7
		{"complex expression 7", "2 3 ^ 4 2 ^ + 5 * 6 2 / -", 117, false}, // ((2 ^ 3) + (4 ^ 2)) * 5 - 6 / 2

		// Error cases
		{"empty string", "", 0, true},
		{"invalid characters", "a b +", 0, true},
		{"insufficient operands", "1 +", 0, true},
		{"too many operands", "1 2 3 +", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalPostfix(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalPostfix(\"%s\") error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EvalPostfix(\"%s\") = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func ExampleEvalPostfix() {
	result, err := EvalPostfix("3 4 +")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Output:
	// Result: 7
}
