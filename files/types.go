package word

type wrapper struct {
    width int
}

func newWrapper(width int) *wrapper {
    return &wrapper{
        width: width,
    }
}