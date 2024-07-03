package word

type wrapper struct {
    width int
}

type paragraph struct {
    wrapper
    text string
}

func (w *wrapper) newParagraph(text string) *paragraph {
    return nil
}

func newWrapper(width int) *wrapper {
    return &wrapper{
        width: width,
    }
}