package hangul

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestHangul_Rune(t *testing.T) {
	tests := []struct {
		name   string
		input  Hangul
		expect rune
	}{
		{"받침 없음", Hangul{'ㄱ', 'ㅏ', 0}, '가'},
		{"받침 있음", Hangul{'ㄱ', 'ㅏ', 'ㄱ'}, '각'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.input.Rune()
			assert.Equal(t, r, tt.expect)
		})
	}
}
