package word

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewParagraph(t *testing.T) {
    assert.NotNil(t, newWrapper(10).newParagraph("test"))
}

func TestNewWrapper(t *testing.T) {
    w := newWrapper(5)
    
    assert.NotNil(t, w)
    assert.Equal(t, 5, w.width)
}