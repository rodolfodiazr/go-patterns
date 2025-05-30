package iterator

import "fmt"

// Product represents an item in the collection
type Product struct {
	Name  string
	Price float64
}

// Iterator defines the behavior of an iterator
type Iterator interface {
	HasNext() bool
	Next() *Product
}

// ProductCollection holds a list of products
type ProductCollection struct {
	products []*Product
}

// NewProductCollection creates a new product collection
func NewProductCollection() *ProductCollection {
	return &ProductCollection{
		products: []*Product{},
	}
}

// Add adds a product to the collection
func (c *ProductCollection) Add(p *Product) {
	c.products = append(c.products, p)
}

// ProductIterator implements the Iterator interface
type ProductIterator struct {
	collection *ProductCollection
	index      int
}

func (it *ProductIterator) HasNext() bool {
	return it.index < len(it.collection.products)
}

func (it *ProductIterator) Next() *Product {
	if !it.HasNext() {
		return nil
	}

	product := it.collection.products[it.index]
	it.index++
	return product
}

// CreateIterator returns an iterator for the collection
func (c *ProductCollection) CreateIterator() Iterator {
	return &ProductIterator{
		collection: c,
		index:      0,
	}
}

// Run demonstrates the Iterator pattern
func Run() {
	collection := NewProductCollection()
	collection.Add(&Product{Name: "Laptop", Price: 1200})
	collection.Add(&Product{Name: "Phone", Price: 800})
	collection.Add(&Product{Name: "Tablet", Price: 400})

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		product := iterator.Next()
		fmt.Printf("Product: %s - $%.2f\n", product.Name, product.Price)
	}
}
