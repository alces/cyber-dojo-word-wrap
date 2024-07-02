package word

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var unwrappedText = `SAMPSON A dog of the house of Montague moves me.
GREGORY To move is to stir, and to be valiant is to stand. Therefore if thou art moved thou runn'st away.
SAMPSON A dog of that house shall move me to stand. I will take the wall of any man or maid of Montague's.
GREGORY That shows thee a weak slave, for the weakest goes to the wall.
SAMPSON 'Tis true, and therefore women, being the weaker vessels, are ever thrust to the wall. Therefore I will push Montague's men from the wall and thrust his maids to the wall.
GREGORY The quarrel is between our masters and us their men.
SAMPSON 'Tis all one. I will show myself a tyrant. When I have fought with the men, I will be civil with the maids; I will cut off their heads.`

var wrapped50Text = `SAMPSON A dog of the house of Montague moves me.
GREGORY To move is to stir, and to be valiant is
to stand. Therefore if thou art moved thou
runn'st away.
SAMPSON A dog of that house shall move me to
stand. I will take the wall of any man or maid of
Montague's.
GREGORY That shows thee a weak slave, for the
weakest goes to the wall.
SAMPSON 'Tis true, and therefore women, being the
weaker vessels, are ever thrust to the wall.
Therefore I will push Montague's men from the
wall and thrust his maids to the wall.
GREGORY The quarrel is between our masters and us
their men.
SAMPSON 'Tis all one. I will show myself a
tyrant. When I have fought with the men, I will
be civil with the maids; I will cut off their
heads.
`

func TestWordWrap(t *testing.T) {
    assert.Equal(t, wrapped50Text, Wrap(unwrappedText, 50))
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
        w := wrapper{
            width: r.width,
        }
        
        text, line := w.addWord(r.text, r.line, r.word, r.width)
        assert.Equal(t, r.expectedText, text, "Unexpected text")
        assert.Equal(t, r.expectedLine, line, "Unexpected line")
    }
}

var wrapParagraphResults = []struct {
    text     string
    width    int
    expected string
} {
    {"abc", 5, "abc"},
    {"abcdefgh", 5, "abcdefgh"},
    {"That shows thee a weak slave", 11, "That shows\nthee a\nweak slave"},
}

func TestWrapParagraph(t *testing.T) {
    for _, r := range wrapParagraphResults {
        w := wrapper{
            width: r.width,
        }
        
        assert.Equal(t, r.expected, w.wrapParagraph(r.text))
    }
}