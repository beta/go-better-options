// Copyright (c) 2018 Beta Kuang
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import "fmt"

// Logger is a demo logger type.
type Logger struct {
	verbosity int
	name      string
}

// Option sets the options specified.
// It returns an option to revert the modifications to the options.
func (l *Logger) Option(opts ...option) (previous option) {
	for i, opt := range opts {
		if i == 0 {
			previous = opt(l)
		} else {
			previous = previous.merge(opt(l))
		}
	}
	return previous
}

type option func(*Logger) option

// Merge merges two options o and opt.
func (o option) merge(opt option) option {
	return func(l *Logger) option {
		previous := o(l)
		return previous.merge(opt(l))
	}
}

// Verbosity sets Foo's verbosity level to v.
func Verbosity(v int) option {
	return func(l *Logger) option {
		previous := l.verbosity
		l.verbosity = v
		return Verbosity(previous)
	}
}

// Name sets Foo's name to v.
func Name(v string) option {
	return func(l *Logger) option {
		previous := l.name
		l.name = v
		return Name(previous)
	}
}

func main() {
	logger := &Logger{
		verbosity: 0,
		name:      "Artoria Pendragon",
	}
	fmt.Printf("Verbosity: %d, name: %s\n", logger.verbosity, logger.name)

	// Set options.
	previous := logger.Option(Verbosity(1), Name("Astolfo"))
	fmt.Printf("Verbosity: %d, name: %s\n", logger.verbosity, logger.name)

	// Restore previous options.
	logger.Option(previous)
	fmt.Printf("Verbosity: %d, name: %s\n", logger.verbosity, logger.name)
}
