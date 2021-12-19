package htmltopdf

import (
	"os"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	err := Init()
	if err != nil {
		t.Error(err)
	}
	defer Destroy()

	reader := strings.NewReader("<html><body>Hello World</body></html>")
	object, err := NewObjectFromReader(reader)
	if err != nil {
		t.Error(err)
	}

	converter, err := NewConverter()
	if err != nil {
		t.Error(err)
	}
	defer converter.Destroy()

	converter.Add(object)
	converter.Title = "Sample document"
	converter.PaperSize = A4
	converter.Orientation = Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	outFile, err := os.Create("test.pdf")
	if err != nil {
		t.Error(err)
	}
	defer outFile.Close()

	converter.Run(outFile)
}
