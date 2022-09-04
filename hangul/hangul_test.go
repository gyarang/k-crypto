package hangul

import "testing"

func TestNewHangul(t *testing.T) {
	tests := []struct {
		input  rune
		expect Hangul
	}{
		{'가', Hangul{'ㄱ', 'ㅏ', 0}},
		{'긕', Hangul{'ㄱ', 'ㅢ', 'ㄱ'}},
		{'갛', Hangul{'ㄱ', 'ㅏ', 'ㅎ'}},
	}

	for _, tt := range tests {
		h, _ := NewHangul(tt.input)
		if h != tt.expect {
			t.Errorf("new hangul error, input: %s, expect: %v, get: %v", string(tt.input), tt.expect, h)
		}
	}
}
