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
