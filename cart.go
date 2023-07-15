package cart

import "strconv"

type Cart struct {
	products map[string]int
}

func NewCart() *Cart {
	return &Cart{products: make(map[string]int)}
}

func (c *Cart) AddProduct(product string, quantity int) {
	c.products[product] += quantity
}

func (c *Cart) RemoveProduct(product string) {
	delete(c.products, product)
}

func (c *Cart) GetProducts() string {
	var products string
	for product, quantity := range c.products {
		products += product + " (" + strconv.Itoa(quantity) + ")\n"
	}
	return products
}
