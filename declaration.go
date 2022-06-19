package imbue

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

// declaration is an interface that describes how to build a value of a specific
// type.
type declaration interface {
	// GetType returns the type of the value constructed by this declaration.
	GetType() reflect.Type

	// Location returns the location of the declaration in code.
	Location() (string, int)

	// IsDependency returns true if other declarations depend upon this one.
	IsDependency() bool

	// Dependencies returns the declarations that this declaration depends upon,
	// sorted by type.
	Dependencies() []declaration

	// dependsOn returns true if d depends on t, whether directly or indirectly.
	dependsOn(t declaration, cycle []declaration) ([]declaration, bool)

	// markAsDependency marks the declaration as a dependency. That is, other
	// declarations depend upon this one.
	markAsDependency()
}

// constructor is a function that constructs values of type T.
type constructor[T any] func(*Context) (T, error)

// declarationOf describes how to build values of type T.
type declarationOf[T any] struct {
	m sync.Mutex

	file string
	line int

	isConstructed bool
	deps          map[reflect.Type]declaration
	isDep         bool
	construct     constructor[T]
	value         T
}

// Declare declares a constructor for values of type T.
func (d *declarationOf[T]) Declare(
	decl func() (constructor[T], error),
) error {
	file, line := findLocation()

	d.m.Lock()
	d.file = file
	d.line = line
	d.m.Unlock()

	c, err := decl()
	if err != nil {
		return err
	}

	d.m.Lock()
	d.construct = c
	d.m.Unlock()

	return nil
}

// AddDependency marks t as a dependency of d.
func (d *declarationOf[T]) AddDependency(t declaration) error {
	if t.GetType() == d.GetType() {
		file, line := d.Location()

		return fmt.Errorf(
			"constructor for %s (%s:%d) depends on itself",
			d.GetType(),
			filepath.Base(file),
			line,
		)
	}

	if cycle, ok := t.dependsOn(d, nil); ok {
		message := fmt.Sprintf(
			"constructor for %s introduces a cyclic dependency:",
			d.GetType(),
		)

		for i := len(cycle) - 1; i >= 0; i-- {
			dep := cycle[i]
			file, line := dep.Location()

			message += fmt.Sprintf(
				"\n\t-> %s (%s:%d)",
				dep.GetType(),
				filepath.Base(file),
				line,
			)
		}

		return errors.New(message)
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.deps == nil {
		d.deps = map[reflect.Type]declaration{}
	}

	d.deps[t.GetType()] = t
	t.markAsDependency()

	return nil
}

// Resolve returns the value constructed by this declaration.
//
// The constructor is called only once. Subsequent calls to Resolve() return the
// same value.
func (d *declarationOf[T]) Resolve(ctx *Context) (T, error) {
	d.m.Lock()
	defer d.m.Unlock()

	if d.isConstructed {
		return d.value, nil
	}

	if d.construct == nil {
		panic(fmt.Sprintf(
			"no constructor is declared for %s",
			d.GetType(),
		))
	}

	v, err := d.construct(ctx)
	if err != nil {
		return d.value, fmt.Errorf(
			"constructor for %s (%s:%d) failed: %w",
			d.GetType(),
			filepath.Base(d.file),
			d.line,
			err,
		)
	}

	d.isConstructed = true
	d.construct = nil
	d.value = v

	return v, nil
}

// GetType returns the type of the value constructed by this declaration.
func (d *declarationOf[T]) GetType() reflect.Type {
	return typeOf[T]()
}

// Location returns the location of the declaration in code.
func (d *declarationOf[T]) Location() (string, int) {
	d.m.Lock()
	defer d.m.Unlock()

	return d.file, d.line
}

// IsDependency returns true if other declarations depend upon this one.
func (d *declarationOf[T]) IsDependency() bool {
	d.m.Lock()
	defer d.m.Unlock()

	return d.isDep
}

// Dependencies returns the declarations that this declaration depends upon,
// sorted by type.
func (d *declarationOf[T]) Dependencies() []declaration {
	d.m.Lock()
	defer d.m.Unlock()

	return sortDeclarations(d.deps)
}

// dependsOn returns true if d depends on t, whether directly or indirectly.
func (d *declarationOf[T]) dependsOn(t declaration, cycle []declaration) ([]declaration, bool) {
	if t.GetType() == d.GetType() {
		return append(cycle, d), true
	}

	d.m.Lock()
	defer d.m.Unlock()

	for _, dep := range d.deps {
		if cycle, ok := dep.dependsOn(t, cycle); ok {
			return append(cycle, d), true
		}
	}

	return nil, false
}

// markAsDependency marks the declaration as a dependency. That is, other
// declarations depend upon this one.
func (d *declarationOf[T]) markAsDependency() {
	d.m.Lock()
	defer d.m.Unlock()

	d.isDep = true
}

// findLocation returns the file and line number of the first frame in the
// current goroutine's stack that is NOT part of the imbue package.
func findLocation() (string, int) {
	var pointers [8]uintptr
	skip := 2 // Always skip runtime.Callers() and findLocation().

	for {
		count := runtime.Callers(skip, pointers[:])
		iter := runtime.CallersFrames(pointers[:count])
		skip += count

		for {
			fr, more := iter.Next()

			if !more || !strings.HasPrefix(fr.Function, "github.com/dogmatiq/imbue.") {
				return fr.File, fr.Line
			}
		}
	}
}
