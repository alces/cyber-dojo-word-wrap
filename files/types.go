package word

import (
    "strings"
)

type wrapper struct {
    width int
}

type paragraph struct {
    wrapper
    text string
}

func newWrapper(width int) *wrapper {
    return &wrapper{
        width: width,
    }
}

func (w *wrapper) addWord(text, line, word string) (string, string) {
    if line != "" {
        if len(line) + len(word) + 1 < w.width {
            line += " "                  
        } else {
            text += line + "\n"
            line = ""   
        }
    }
    
    return text, line + word
}

func (w *wrapper) newParagraph(text string) *paragraph {
    return &paragraph{
        wrapper: *w,
        text:    text,
    }
}

func (p *paragraph) wrap() (result string) {
    if len(p.text) <= p.width {
        return p.text
    }
    
    line := ""
    
    for _, word := range strings.Split(p.text, " ") {                
        result, line = p.addWord(result, line, word)
    }
    
    result += line
    
    return
}