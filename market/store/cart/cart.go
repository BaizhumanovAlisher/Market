package cart

import "market/models"

type Line struct {
	models.Product
	Quantity int
}

func (cl *Line) GetLineTotal() float64 {
	return cl.Price * float64(cl.Quantity)
}

type Cart interface {
	AddProduct(models.Product)
	GetLines() []*Line
	RemoveLineForProduct(id int)
	GetItemCount() int
	GetTotal() float64

	Reset()
}

type BasicCart struct {
	lines []*Line
}

func (cart *BasicCart) AddProduct(p models.Product) {
	for _, line := range cart.lines {
		if line.Product.ID == p.ID {
			line.Quantity++
			return
		}
	}

	cart.lines = append(cart.lines, &Line{
		Product: p, Quantity: 1,
	})
}

func (cart *BasicCart) GetLines() []*Line {
	return cart.lines
}

func (cart *BasicCart) RemoveLineForProduct(id int) {
	for index, line := range cart.lines {
		if line.Product.ID == id {
			cart.lines = append(cart.lines[0:index], cart.lines[index+1:]...)
		}
	}
}

func (cart *BasicCart) GetItemCount() (total int) {
	for _, l := range cart.lines {
		total += l.Quantity
	}

	return
}

func (cart *BasicCart) GetTotal() (total float64) {
	for _, line := range cart.lines {
		total += float64(line.Quantity) * line.Product.Price
	}

	return
}

func (cart *BasicCart) Reset() {
	cart.lines = []*Line{}
}
