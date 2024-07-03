package word

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewWrapper(t *testing.T) {
    assert.NotNil(t, newWrapper(5))
}