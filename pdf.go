package htmltopdf

/*
#cgo LDFLAGS: -lwkhtmltox

#include <stdlib.h>
#include <wkhtmltox/pdf.h>
*/
import "C"
import "sync"

var sm *sync.Map

// Init initializes the library, allocating all necessary resources.
func Init() error {
	if sm != nil {
		return ErrLibraryAlereadyInitialized
	}

	if C.wkhtmltopdf_init(0) != 1 {
		return ErrLibraryNotInitialized
	}

	sm = new(sync.Map)
	return nil
}

// Version returns version of the wkhtmltopdf library.
func Version() string {
	return C.GoString(C.wkhtmltopdf_version())
}

// ExtendedQT returns true if the library is built against the
// wkhtmltopdf version of QT.
func ExtendedQT() bool {
	return C.wkhtmltopdf_extended_qt() != 0
}

// Destroy releases all the resources used by the library.
func Destroy() {
	C.wkhtmltopdf_deinit()
	sm = nil
}
