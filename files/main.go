package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    for _, paragraph := range strings.Split(text, "\n") {
        if len(p) <= width {
            result += paragraph + "\n"
        } else {
            line := ""
            for _, w := range strings.Split(paragraph, " ") {                
                if len(line) + len(w) + 1 < width {
                    if line != "" {
                        line += " "
                    }                   
                } else {
                    result += line + "\n"
                    line = ""
                }
                line += w
            }
            result += line + "\n"
        }    
    }    
    
    return
}
