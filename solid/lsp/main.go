package main

type IRectangle interface {
	setWidth(int)
	setHeight(int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) setWidth(w int) {
	r.width = w
}

func (r *Rectangle) setHeight(h int) {
	r.height = h
}

type Square struct {
	side int
}

func (r *Square) setWidth(w int) {
	r.side = w
}

func (r *Square) setHeight(h int) {
	r.side = h
}

func main() {

	var _ IRectangle = (*Rectangle)(nil)
	var _ IRectangle = (*Square)(nil)

}
