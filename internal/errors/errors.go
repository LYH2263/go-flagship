package errors

import "errors"

var (
	ErrClosed      = errors.New("flagship: closed")
	ErrInvalidKey  = errors.New("flagship: invalid key")
	ErrInvalidRule = errors.New("flagship: invalid rule")
	ErrNotFound    = errors.New("flagship: not found")
	ErrPersist     = errors.New("flagship: persist failed")
	ErrCanceled    = errors.New("flagship: canceled")
	ErrNilHasher   = errors.New("flagship: nil hasher")
	ErrNilRollout  = errors.New("flagship: nil rollout")
	ErrNilStore    = errors.New("flagship: nil store")
	ErrTooMany     = errors.New("flagship: too many flags")
	ErrFlush       = errors.New("flagship: flush failed")
	ErrSync        = errors.New("flagship: sync failed")
	ErrEmptyKey    = errors.New("flagship: empty key")
	ErrBadPercent  = errors.New("flagship: percent out of range")
)
