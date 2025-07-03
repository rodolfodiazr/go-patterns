package iterator

import "testing"

func Test_ProductIterator(t *testing.T) {
	products := []*Product{
		{Name: "Laptop", Price: 1200},
		{Name: "Phone", Price: 800},
		{Name: "Tablet", Price: 400},
	}

	collection := NewProductCollection()
	for _, p := range products {
		collection.Add(p)
	}

	it := collection.CreateIterator()
	for i, expected := range products {
		if !it.HasNext() {
			t.Fatalf("Expected HasNext() to be true at index %d", i)
		}

		got := it.Next()
		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	if it.HasNext() {
		t.Error("Expected HasNext() to be false at the end of iteration")
	}
}
