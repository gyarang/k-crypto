package table

import "testing"

func TestNewTable(t *testing.T) {
	tb, _ := NewTable("금강산")
	for _, tt := range tb {
		t.Error(tt)
	}
}
