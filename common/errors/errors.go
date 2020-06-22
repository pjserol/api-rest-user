package errors

// SpecificErr specific error
type SpecificErr struct {
	Err        error
	StatusCode int
	Status     string
}

func (e SpecificErr) Error() string {
	return e.Err.Error()
}

// NotFoundErr not found error
type NotFoundErr struct {
	Err error
}

func (e NotFoundErr) Error() string {
	return e.Err.Error()
}

// BadRequestErr bad request error
type BadRequestErr struct {
	Err error
}

func (e BadRequestErr) Error() string {
	return e.Err.Error()
}

// ConflictErr conflict error
type ConflictErr struct {
	Err error
}

func (e ConflictErr) Error() string {
	return e.Err.Error()
}
