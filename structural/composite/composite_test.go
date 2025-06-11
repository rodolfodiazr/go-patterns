package composite

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	// Save the original stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Close the writer and restore stdout
	_ = w.Close()
	os.Stdout = old

	// Read the output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func Test_File_Name(t *testing.T) {
	tCases := []struct {
		name             string
		file             *File
		expectedFilename string
	}{
		{
			name:             "blank filename",
			file:             &File{},
			expectedFilename: "",
		},
		{
			name:             "file with a set filename",
			file:             &File{filename: "README.md"},
			expectedFilename: "README.md",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			if filename := tc.file.Name(); filename != tc.expectedFilename {
				t.Errorf("expected filename '%s', got '%s'", tc.expectedFilename, filename)
			}
		})
	}
}

func Test_File_Display(t *testing.T) {
	tCases := []struct {
		name            string
		indent          int
		expectedDisplay string
	}{
		{
			name:            "2-space indentation",
			indent:          2,
			expectedDisplay: "  - File: file1.txt\n",
		},
		{
			name:            "4-space indentation",
			indent:          4,
			expectedDisplay: "    - File: file1.txt\n",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			f := &File{filename: "file1.txt"}

			display := captureOutput(func() {
				f.Display(tc.indent)
			})

			if display != tc.expectedDisplay {
				t.Errorf("expected:\n%q\ngot:\n%q", tc.expectedDisplay, display)
			}
		})
	}
}

func Test_Directory_Name(t *testing.T) {
	tCases := []struct {
		name            string
		directory       *Directory
		expectedDirName string
	}{
		{
			name:            "blank dirname",
			directory:       &Directory{},
			expectedDirName: "",
		},
		{
			name:            "directory with a set dirname",
			directory:       &Directory{dirname: "docs"},
			expectedDirName: "docs",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			if dirname := tc.directory.Name(); dirname != tc.expectedDirName {
				t.Errorf("expected directory name '%s', got '%s'", tc.expectedDirName, dirname)
			}
		})
	}
}

func TestDirectory_Display(t *testing.T) {
	tCases := []struct {
		name            string
		createDir       func() *Directory
		expectedDisplay []string
	}{
		{
			name: "empty directory",
			createDir: func() *Directory {
				return &Directory{dirname: "empty"}
			},
			expectedDisplay: []string{
				"+ Directory: empty",
			},
		},
		{
			name: "directory with 1 file",
			createDir: func() *Directory {
				root := &Directory{dirname: "root"}
				root.Add(&File{filename: "file1.txt"})
				return root
			},
			expectedDisplay: []string{
				"+ Directory: root",
				"  - File: file1.txt",
			},
		},
		{
			name: "directory with 1 file and 2 subdirectories",
			createDir: func() *Directory {
				root := &Directory{dirname: "root"}
				root.Add(&File{filename: "file1.txt"})

				assets := &Directory{dirname: "assets"}
				images := &Directory{dirname: "images"}
				images.Add(&File{filename: "img1.png"})

				assets.Add(images)
				root.Add(assets)
				return root
			},
			expectedDisplay: []string{
				"+ Directory: root",
				"  - File: file1.txt",
				"  + Directory: assets",
				"    + Directory: images",
				"      - File: img1.png",
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			output := captureOutput(func() {
				dir := tc.createDir()
				dir.Display(0)
			})

			lines := strings.Split(strings.TrimSpace(output), "\n")
			if len(lines) != len(tc.expectedDisplay) {
				t.Fatalf("expected %d lines, got %d", len(tc.expectedDisplay), len(lines))
			}

			for i, expectedLine := range tc.expectedDisplay {
				if lines[i] != expectedLine {
					t.Errorf("line %d mismatch:\nexpected: %q\ngot: %q", i, expectedLine, lines[i])
				}
			}
		})
	}
}
