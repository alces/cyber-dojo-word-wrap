package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    for _, paragraph := range strings.Split(text, "\n") {
        if len(paragraph) <= width {
            result += paragraph + "\n"
        } else {
            line := ""
            for _, word := range strings.Split(paragraph, " ") {                
                if len(line) + len(word) + 1 < width {
                    if line != "" {
                        line += " "
                    }                   
                } else {
                    result += line + "\n"
                    line = ""
                }
                line += word
            }
            result += line + "\n"
        }    
    }    
    
    return
}

func wrapParagraph(text string, width int) (result string) {
    if len(text) <= width {
        return text
    }
    
    line := ""
    for _, word := range strings.Split(text, " ") {                
        if len(line) + len(word) + 1 < width {
            if line != "" {
                line += " "
            }                   
        } else {
            result += line + "\n"
            line = ""
        }
        line += word
    }
    result += line + "\n"
    
    return
}