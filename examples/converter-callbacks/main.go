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

	// Create converter.
	converter, err := htmltopdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()

	// Set up converter callbacks.
	converter.Warning = func(msg string) {
		log.Printf("warning: %s\n", msg)
	}
	converter.Error = func(msg string) {
		log.Printf("error: %s\n", msg)
	}
	converter.PhaseChanged = func(phase int) {
		log.Printf("phase #%d: %s\n", phase, converter.PhaseDescription(phase))
	}
	converter.ProgressChanged = func(percent int) {
		log.Printf("progress: %d%%\n", percent)
	}
	converter.Finished = func(success bool) {
		log.Printf("finished: %t\n", success)
	}

	// Create objects.
	object, err := htmltopdf.NewObject("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	object.Header.ContentCenter = "[title]"
	object.Header.DisplaySeparator = true
	object.Footer.ContentLeft = "[date]"
	object.Footer.ContentCenter = "Sample footer information"
	object.Footer.ContentRight = "[page]"
	object.Footer.DisplaySeparator = true

	converter.Add(object)

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
