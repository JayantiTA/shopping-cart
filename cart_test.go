package cart

import "testing"

func TestCart(t *testing.T) {
	cart := NewCart()

	cart.AddProduct("Pisang Hijau", 2)

	cart.AddProduct("Semangka Kuning", 3)

	cart.AddProduct("Apel Merah", 1)
	cart.AddProduct("Apel Merah", 4)
	cart.AddProduct("Apel Merah", 2)

	cart.RemoveProduct("Semangka Kuning")
	cart.RemoveProduct("Semangka Merah")

	products := cart.GetProducts()
	expected := "Pisang Hijau (2)\nApel Merah (7)\n"
	if products != expected {
		t.Errorf("got %s, want %s", products, expected)
	}
}
