package word

import (
    "strings"
)

type wrapper struct {
    width int
}

func Wrap(text string, width int) (result string) {
    w := wrapper{
        width: width,
    }
    
    for _, p := range strings.Split(text, "\n") {
        result += w.wrapParagraph(p) + "\n"
    }    
    
    return
}

func addWord(text, line, word string, width int) (string, string) {
    if line != "" {
        if len(line) + len(word) + 1 < width {
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
        result, line = addWord(result, line, word, w.width)
    }
    
    result += line
    
    return
}