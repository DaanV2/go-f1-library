package c

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Strings(t *testing.T) {
	data := []byte{
		byte('h'),
		byte('e'),
		byte('y'),
		byte(0),
	}

	strc := String(data);
	assert.Len(t, strc, 3)
	assert.Equal(t, strc, "hey")
}