package models

type User struct {
	Id        int    `binding:"required" msg:"Required"`
	Name      string `binding:"max=10" msg:"Maximum length is 10"`
	Email     string `binding:"email" msg:"Invalid email address"`
	Age       int    `binding:"min=1,max=100" msg:"Must between 1 and 100"`
	BirthDate string `binding:"datetime=01/02/2006" msg:"Invalid date format"`
}
