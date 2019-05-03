package errlist

import (
	"strings"
)

// Error turns a list of errors into a single error implementation
// whose Error() function returns a nicely formatted message
func Error(errs []error) error {
	return &lst{errs: errs}
}

type lst struct {
	errs []error
}

func (l *lst) Error() string {
	strs := make([]string, len(l.errs))
	for i, err := range l.errs {
		strs[i] = err.Error()
	}
	return strings.Join(strs, "\n")
}
