package word

import (
    "strings"
)

func Wrap(text string, width int) (result string) {
    for _, p := range strings.Split(text, "\n") {
        if len(p) <= width {
            result += p + "\n"
        }
    }    
    
    return
}
