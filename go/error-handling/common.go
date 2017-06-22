package erratum

import (
	"errors"
	"io"
)

const testVersion = 2

// TransientError : temporary error
type TransientError struct {
	err error
}

func (e TransientError) Error() string {
	return e.err.Error()
}

// FrobError :
type FrobError struct {
	defrobTag string
	inner     error
}

func (e FrobError) Error() string {
	return e.inner.Error()
}

// Resource : interface
type Resource interface {
	io.Closer

	// Frob does something with the input string.
	// Because this is an incredibly badly designed system if there is an error
	// it panics.
	//
	// The paniced error may be a FrobError in which case Defrob should be called
	// with the defrobTag string.
	Frob(string)

	Defrob(string)
}

// ResourceOpener is a function that creates a resource.
//
// It may return a wrapped error of type TransientError. In this case the resource
// is temporarily unavailable and the caller should retry soon.
type ResourceOpener func() (Resource, error)

// Use : Main function
func Use(opener ResourceOpener, input string) (err error) {

	resource, err := opener()
	for err != nil {
		switch err.(type) {
		case TransientError:
			resource, err = opener()
		default:
			return err
		}
	}

	defer func() {
		if r := recover(); r != nil { //catch
			switch x := r.(type) {
			case FrobError:
				err = x
				resource.Defrob(x.defrobTag)
				resource.Close()
			case error:
				err = x
				resource.Close()
			default:
				err = errors.New("Unknown panic")
				resource.Close()
			}
		}
	}()
	resource.Frob(input)

	resource.Close()

	return err
}
