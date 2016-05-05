package mmath

type DivideZeroError struct{

}

func NewDivideZeroError()DivideZeroError{
	return DivideZeroError{}
}

func (e DivideZeroError)Error()string{
	return "Divide by 0 error."
}

