package el_test

import (
	"testing"

	el "github.com/runcom/go-el"
	"github.com/stretchr/testify/assert"
)

func TestLex(t *testing.T) {
	assert := assert.New(t)
	tok, err := el.Lex("abc.cc[1]")
	assert.Nil(err)
	assert.Len(tok, 6)
	assert.Equal("[", tok[3].Val)
	assert.Equal("]", tok[5].Val)
}
