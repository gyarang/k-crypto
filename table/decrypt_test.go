package table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable_Decrypt(t *testing.T) {
	tb, err := NewTable("금강산")
	assert.Nil(t, err)

	tests := []struct {
		name      string
		input     string
		expect    []rune
		expectErr error
	}{
		{"example", "7355614782 97486861544147516448", []rune{'ㅅ', 'ㅗ', 'ㄱ', 'ㅎ', 'ㅣ', ' ', 'ㅅ', 'ㅏ', 'ㅇ', 'ㄱ', 'ㅕ', 'ㅇ', 'ㅎ', 'ㅏ', 'ㄹ', 'ㅏ'}, nil},
		{"invalid input 0", "00", []rune{}, DecryptNotValidInput},
		{"invalid input odd length", "111", []rune{}, DecryptNotValidInput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := tb.Decrypt(tt.input)
			if tt.expectErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, d, tt.expect)
			} else {
				assert.Equal(t, err, DecryptNotValidInput)
			}
		})
	}
}
