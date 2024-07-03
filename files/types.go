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

func (p *paragraph) wrap() result string {
    return
}