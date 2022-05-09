package main

import (
	"log"
	"os"
	"runtime"

	"github.com/digitaldata-cz/htmltopdf"
)

func init() {
	// Set main function to run on the main thread.
	runtime.LockOSThread()

	// Initialize library.
	if err := htmltopdf.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer htmltopdf.Destroy()

	// Create object from file.
	object, err := htmltopdf.NewObject("sample1.html")
	if err != nil {
		log.Fatal(err)
	}
	object.Header.ContentCenter = "[title]"
	object.Header.DisplaySeparator = true

	// Create object from URL.
	object2, err := htmltopdf.NewObject("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	object2.Footer.ContentLeft = "[date]"
	object2.Footer.ContentCenter = "Sample footer information"
	object2.Footer.ContentRight = "[page]"
	object2.Footer.DisplaySeparator = true

	// Create object from reader.
	inFile, err := os.Open("sample2.html")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	object3, err := htmltopdf.NewObjectFromReader(inFile)
	if err != nil {
		log.Fatal(err)
	}
	object3.Zoom = 1.5
	object3.TOC.Title = "Table of Contents"

	// Create converter.
	converter, err := htmltopdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()

	// Add created objects to the converter.
	converter.Add(object)
	converter.Add(object2)
	converter.Add(object3)

	// Set converter options.
	converter.Title = "Sample document"
	converter.PaperSize = htmltopdf.A4
	converter.Orientation = htmltopdf.Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	// Convert objects and save the output PDF document.
	outFile, err := os.Create("out.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := converter.Run(outFile); err != nil {
		log.Fatal(err)
	}
}