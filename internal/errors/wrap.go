package errors

import "fmt"

func Wrap(sentinel error, msg string) error {
	if sentinel == nil {
		return fmt.Errorf("%s", msg)
	}
	return fmt.Errorf("%w: %s", sentinel, msg)
}

func WrapErr(sentinel, err error) error {
	if sentinel == nil {
		return err
	}
	if err == nil {
		return sentinel
	}
	return fmt.Errorf("%w: %w", sentinel, err)
}
