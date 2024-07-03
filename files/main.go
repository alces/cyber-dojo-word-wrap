package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    w := newWrapper(width)
    
    for _, p := range strings.Split(text, "\n") {
        result += w.wrapParagraph(p) + "\n"
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

func (w *wrapper) wrapParagraph(text string) (result string) {
    if len(text) <= w.width {
        return text
    }
    
    line := ""
    
    for _, word := range strings.Split(text, " ") {                
        result, line = w.addWord(result, line, word)
    }
    
    result += line
    
    return
}