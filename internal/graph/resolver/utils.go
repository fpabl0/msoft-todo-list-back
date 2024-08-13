package resolver

import (
	"github.com/msoft-g1/todo-list-backend/internal/errs"
)

func makeAppError(err error) *errs.Error {
	aperr, ok := err.(*errs.Error)
	if !ok {
		return errs.NewWithError(errs.CodeUnexpected, err)
	}
	return aperr
}
