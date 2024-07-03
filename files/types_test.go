package word

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewParagraph(t *testing.T) {
    p := newWrapper(10).newParagraph("test")
    
    assert.NotNil(t, p)
    assert.Equal(t, "test", p.text)
    assert.Equal(t, 10, p.width)
}

func TestNewWrapper(t *testing.T) {
    w := newWrapper(5)
    
    assert.NotNil(t, w)
    assert.Equal(t, 5, w.width)
}