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

var paragraphWrapResults = []struct {
    text     string
    width    int
    expected string
} {
    {"abc", 5, "abc"},
    {"abcdefgh", 5, "abcdefgh"},
    {"That shows thee a weak slave", 11, "That shows\nthee a\nweak slave"},
}

func TestParagraphWrap(t *testing.T) {
    for _, r := range paragraphWrapResults {
        p := newWrapper(r.width).newParagraph(r.text)
        
        assert.Equal(t, r.expected, p.wrap())
    }
}
