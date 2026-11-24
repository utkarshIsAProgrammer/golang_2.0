package main

import (
	"fmt"
	"sync"
)

type ProductStore struct {
	mu       sync.RWMutex
	products map[int]string
}

func (ps *ProductStore) addProduct(id int, name string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.products[id] = name
}

func (ps *ProductStore) getProduct(id int) string {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	return ps.products[id]
}

func (ps *ProductStore) getAllProducts() map[int]string {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	copiedProducts := make(map[int]string)
	for id, name := range ps.products {
		copiedProducts[id] = name

	}
	return copiedProducts

}

func main() {
	store := ProductStore{
		products: make(map[int]string),
	}

	store.addProduct(1, "Laptop")
	store.addProduct(2, "Keyboard")
	store.addProduct(3, "Mouse")

	fmt.Println("Product:", store.getProduct(1))

	allProducts := store.getAllProducts()
	fmt.Println("All Products:", allProducts)

	allProducts[1] = "Changed Laptop"

	fmt.Println("Modified Copy:", allProducts)
	fmt.Println("Original Store:", store.getAllProducts())
}
