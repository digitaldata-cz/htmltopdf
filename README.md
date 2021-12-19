# htmltopdf

Package htmltopdf implements wkhtmltopdf Go bindings. It can be used to convert HTML documents to PDF files.
The package does not use the wkhtmltopdf binary. Instead, it uses the wkhtmltox library directly.

## Usage

```go
package main

import (
  "log"
  "os"

  "github.com/digitaldata-cz/htmltopdf"
)

func main() {
  htmltopdf.Init()
  defer htmltopdf.Destroy()

  // Create object from file.
  object1, err := htmltopdf.NewObject("sample1.html")
  if err != nil {
    log.Fatal(err)
  }
  object1.Header.ContentCenter = "[title]"
  object1.Header.DisplaySeparator = true

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
  converter.Add(object1)
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
```

For more information see <http://wkhtmltopdf.org/usage/wkhtmltopdf.txt>

## LICENSE

  [MIT](LICENSE)
  