package table

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable_Encrypt(t *testing.T) {
	tb, err := NewTable("금강산")
	assert.Nil(t, err)

	encoded, err := tb.Encrypt("속히 상경하라")
	assert.Nil(t, err)
	fmt.Println(encoded)
}
