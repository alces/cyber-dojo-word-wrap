package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    for _, p := range strings.Split(text, "\n") {
        result += wrapParagraph(p, width) + "\n"
    }    
    
    return
}

func wrapParagraph(text string, width int) (result string) {
    if len(text) <= width {
        return text
    }
    
    line := ""
    for _, word := range strings.Split(text, " ") {                
        if line != "" {
            if len(line) + len(word) + 1 < width {
                line += " "                  
            } else {
                result += line + "\n"
                line = ""   
            }
        }
        line += word
    }
    result += line
    
    return
}