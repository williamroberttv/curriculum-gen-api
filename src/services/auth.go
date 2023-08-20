package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/williamroberttv/curriculum-gen-api/src/database"
	"github.com/williamroberttv/curriculum-gen-api/src/repositories"
	"github.com/williamroberttv/curriculum-gen-api/src/utils"
)

type AuthenticaterUser struct {
	Email string `valid:"email" json:"email" gorm:"type:varchar(255); unique_index; not null;"`
	Password string `valid:"notnull" json:"password" gorm:"type:varchar(255); not null;"`
}

func Authenticate(c *fiber.Ctx) error {
	authUser := new(AuthenticaterUser)
	if err := c.BodyParser(authUser); err != nil {
		return err
	}

	userRepo := repositories.NewUserRepositoryDb(database.DB)
	user, _ := userRepo.FindUserByEmail(authUser.Email)

	comparePassword := utils.CheckPasswordHash(string(user.Password), string(authUser.Password))

	if comparePassword != nil{
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid credentials.",
		})
	}

	if user == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found.",
		})
	}

	jwtToken, err := generateToken(user.ID)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(jwtToken)
}

func RefreshToken(c *fiber.Ctx) error {
	authorizationHeader := strings.Split(c.Get("Authorization"), " ")[1]
	key := []byte(os.Getenv("SECRET_KEY"))
	
	token, err := validateToken(authorizationHeader, key)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token.",
		})
	}

	id := token.Claims.(jwt.MapClaims)["id"].(string)

	userRepo := repositories.NewUserRepositoryDb(database.DB)
	user, _ := userRepo.Find(id)

	if user == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found.",
		})
	}

	jwtToken, err := generateToken(user.ID)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(jwtToken)
}

func generateToken(id string) (map[string]string, error){
	var(
		key = []byte(os.Getenv("SECRET_KEY"))
		t *jwt.Token
		err error
	)

	refreshToken, err := generateRefreshToken([]byte(key), id)
	if err != nil {
		return nil, err
	}

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		}) 
	t.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()
	jwtToken, err := t.SignedString(key)

	response := map[string]string{
		"jwtToken":     jwtToken,
		"refreshToken": refreshToken,
	}

	return response, err
}

func generateRefreshToken(secretKey []byte, subject string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = subject
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
			return "", err
	}
	return refreshToken, nil
}

func validateToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }
    return token, nil
}