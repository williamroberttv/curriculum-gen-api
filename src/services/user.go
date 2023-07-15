package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamroberttv/curriculum-gen-api/src/database"
	"github.com/williamroberttv/curriculum-gen-api/src/models"
	"github.com/williamroberttv/curriculum-gen-api/src/repositories"
)

type UserService struct {
	User *models.User
	UserRepository repositories.UserRepository
}

func CreateUser(c *fiber.Ctx) error {
		userRepo := repositories.NewUserRepositoryDb(database.DB)
		newUser := new(models.User)


		if err := c.BodyParser(newUser); err != nil {
				return err
		}

		user, _ := userRepo.FindUserByEmail(newUser.Email)

		if user != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Email already used.",
			})
		}

		user, err := userRepo.Insert(newUser)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "User created successfully",
			"user": user,
		})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id");
	
	userRepo := repositories.NewUserRepositoryDb(database.DB)
	user, err := userRepo.Find(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(user)

}
