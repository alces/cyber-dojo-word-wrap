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
