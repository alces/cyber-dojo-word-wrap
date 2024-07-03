package word

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewWrapper(t *testing.T) {
    w := newWrapper(5)
    
    assert.NotNil(t, w)
    assert.Equal(t, 5, w.width)
}