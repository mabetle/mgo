package demo

// TagModel include 5 validators:
// required,
// minlength,
// maxlength,
// min,
// max
type TagModel struct {
	Id    string `validator:"required"`
	Name  string `validator:"required,minlength=2,maxlength=10"`
	Email string `validator:"email"`
	Age   int    `validator:"min=3,max=35"`
}
