package user

type RegisterUserInput struct {
	Name       string `binding:"required"`
	Email      string `binding:"required,email"`
	Occupation string `binding:"required"`
	Password   string `binding:"required"`
}
