package core

type BudDriveError struct {
	IsBudDriveError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewBudDriveError(code string, msg string, ctx *Context) *BudDriveError {
	return &BudDriveError{
		IsBudDriveError: true,
		Sdk:              "BudDrive",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *BudDriveError) Error() string {
	return e.Msg
}
