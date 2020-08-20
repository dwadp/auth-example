package models

type (
	User struct {
		ID       string `json:"id" bson:"_id"`
		Name     string `json:"name" bson:"name" binding:"required"`
		Email    string `json:"email" bson:"email" binding:"required,email"`
		Password string `json:"password" bson:"password" binding:"required,min=6,max=20"`
	}

	UserLogin struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}

	AuthClaims struct {
		ID string `json:"id"`
	}

	AuthUser struct {
		Token string `json:"token"`
		User  User   `json:"user"`
	}
)
