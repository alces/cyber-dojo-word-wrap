package word

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