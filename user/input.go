package user

type RegisterUserInput struct {
	Name       string `binding:"required"`
	Email      string `binding:"required,email"`
	Occupation string `binding:"required"`
	Password   string `binding:"required"`
}

type LoginUserInput struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}
