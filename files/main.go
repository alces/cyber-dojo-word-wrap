package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    for _, p := range strings.Split(text, "\n") {
        if len(p) <= width {
            result += p + "\n"
        } else {
            line := ""
            for _, w := range strings.Split(p, " ") {                
                if len(line) + len(w) + 1 <= width {
                    if line != "" {
                        line += " "
                    }
                    line += w
                } else {
                    result += line +"\n"
                    line = w
                }
            }    
        }    
    }    
    
    return
}
