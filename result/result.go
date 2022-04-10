// Result es una estructura que representa un resultado de una operación.
// puede ser un valor o un error y puede contener un mensaje adicional en
// caso de que se haya producido un error.
//
// Ademas puede gestionar ed una forma mas sencilla la propagación de errores
// a traves de las funciones `Error()`, `Catch()`, `Log()` y `Èxit()`
//
// Para el uso de esta estructura se debe utilizar la version de go 1.18+
package result

import (
	"log"
	"os"
)

type CatchCallback func(error)

// Logger es una variable global que contiene el logger a utiliza, por defecto el valor es igual al
// log por defecto en la libreria log.
var logger *log.Logger = log.Default()

type Result[T any] struct {
	value T
	err   error
	then  func(T)
	catch CatchCallback
}

func New[T any](value T, e error) *Result[T] {
	return &Result[T]{
		value: value,
		error: e,
		then:  nil,
		catch: nil,
	}
}

func SetLogger(l *log.Logger) {
	logger = l
}

func (r Result[T]) IsOk() bool {
	return r.error == nil
}

func (r Result[T]) IsErr() bool {
	return r.error != nil
}

func (r Result[T]) Then(fn func(T)) *Result[T] {
	r.then = fn
	return &r
}

func (r Result[T]) Catch(fn CatchCallback) *Result[T] {
	r.catch = fn
	return &r
}

func (r Result[T]) Value() T {
	return r.value
}

func (r Result[T]) Go() *Result[T] {
	if r.IsOk() {
		r.then(r.value)
	} else if r.catch != nil {
		r.catch(r.error)
	}
}

func (r Result[T]) Exit(code int, msg string, args ...any) *Result[T] {
	r.catch = func(err error) {
		logger.Printf(msg, args...)
		os.Exit(code)
	}
}

func (r Result[T]) Log(msg string, args ...any) *Result[T] {
	r.catch = func(err error) {
		logger.Printf(msg, args...)
	}
	return &r
}
