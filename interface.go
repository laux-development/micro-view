package view

import "io"

// ViewProvider Interface for all the views
type ViewProvider interface {
	Home(io.Writer)
	Login(io.Writer)
}
