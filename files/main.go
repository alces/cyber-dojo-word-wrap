package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    w := newWrapper(width)
    
    for _, p := range strings.Split(text, "\n") {
        result += w.newParagraph(p).wrap() + "\n"
    }    
    
    return
}


