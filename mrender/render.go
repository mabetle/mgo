package render

type Render interface {
	Rend(value interface{})
}

func Rend(value interface{}, render Render) {
	render.Rend(value)
}
