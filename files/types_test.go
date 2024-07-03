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

var addWordResults = []struct {
    text         string
    line         string
    word         string
    width        int
    expectedText string
    expectedLine string
} {
    {"", "", "abc", 5, "", "abc"},
    {"", "abc", "xyz", 10, "", "abc xyz"},
    {"", "abc xyz", "qwe", 7, "abc xyz\n", "qwe"},
}

func TestAddWord(t *testing.T) {
    for _, r := range addWordResults {
        w := newWrapper(r.width)
        
        text, line := w.addWord(r.text, r.line, r.word)
        assert.Equal(t, r.expectedText, text, "Unexpected text")
        assert.Equal(t, r.expectedLine, line, "Unexpected line")
    }
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
