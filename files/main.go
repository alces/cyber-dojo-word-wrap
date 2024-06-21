package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    for _, p := range strings.Split(text, "\n") {
        if len(p) <= width {
            result += p + "\n"
        } else {
            ls := 0
            for _, w := range strings.Split(p, " ") {
                ws := len(w) + 1
                if ls + ws <= width {
                    result += w
                    ls += ws
                } else {
                    result += "\n" + w
                    ls = 0
                }
            }    
        }    
    }    
    
    return
}
