package composite

import "fmt"

// Component defines the interface for all objects in the hierarchy
type Component interface {
	Name() string
	Display(indent int)
}

// File is a leaf node
type File struct {
	filename string
}

func (f *File) Name() string {
	return f.filename
}

func (f *File) Display(indent int) {
	fmt.Printf("%s- File: %s\n", spaces(indent), f.filename)
}

// Directory is a composite node
type Directory struct {
	components []Component
	dirname    string
}

func (d *Directory) Name() string {
	return d.dirname
}

func (d *Directory) Add(component Component) {
	d.components = append(d.components, component)
}

func (d *Directory) Display(indent int) {
	fmt.Printf("%s+ Directory: %s\n", spaces(indent), d.dirname)
	for _, c := range d.components {
		c.Display(indent + 2)
	}
}

// Helper function to format indentation
func spaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

// Run demonstrates the Composite pattern
func Run() {
	root := &Directory{dirname: "root"}

	// Add files directly to root
	root.Add(&File{filename: "main.go"})
	root.Add(&File{filename: "README.md"})

	// Add a subdirectory with its own contents
	src := &Directory{dirname: "src"}
	src.Add(&File{filename: "app.go"})
	src.Add(&File{filename: "utils.go"})

	// Nested subdirectory
	tests := &Directory{dirname: "tests"}
	tests.Add(&File{filename: "app_test.go"})

	src.Add(tests)

	root.Add(src)

	root.Display(0)
}
