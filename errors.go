package flagship

import ierr "github.com/LYH2263/go-flagship/internal/errors"

var (
	ErrClosed      = ierr.ErrClosed
	ErrInvalidKey  = ierr.ErrInvalidKey
	ErrInvalidRule = ierr.ErrInvalidRule
	ErrNotFound    = ierr.ErrNotFound
	ErrPersist     = ierr.ErrPersist
	ErrCanceled    = ierr.ErrCanceled
	ErrNilHasher   = ierr.ErrNilHasher
	ErrNilRollout  = ierr.ErrNilRollout
	ErrNilStore    = ierr.ErrNilStore
	ErrTooMany     = ierr.ErrTooMany
	ErrFlush       = ierr.ErrFlush
	ErrSync        = ierr.ErrSync
	ErrEmptyKey    = ierr.ErrEmptyKey
	ErrBadPercent  = ierr.ErrBadPercent
)
