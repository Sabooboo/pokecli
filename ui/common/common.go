package common

type Common struct {
	Width  int
	Height int
}

func (c *Common) SetSize(width, height int) {
	c.Width = width
	c.Height = height
}
