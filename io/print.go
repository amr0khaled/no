package print

import (
	"fmt"
	"os"
	"strings"
)

type String string

type Line struct {
	Offset int
	Size   int
	Text   string
}

type File struct {
	Name  string
	Size  int // size of all lines ->> line.size + line.size
	Lines []Line
}

func ReadFile(filename string) *File {
	x := strings.Split(filename, "/")
	f := &File{
		Name: x[len(x)-1],
	}

	buf, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	code := string(buf)
	fmt.Println(code)
	lines := strings.Split(code, "\n")
	var flines []Line
	var totalSize int = 0
	for i, line := range lines {
		lsize := len(line)
		if lsize == 0 {
			continue
		}
		totalSize += lsize
		fline := &Line{
			Offset: i,
			Size:   lsize,
			Text:   line,
		}
		flines = append(flines, *fline)
	}
	f.Size = totalSize
	f.Lines = flines
	return f
}
