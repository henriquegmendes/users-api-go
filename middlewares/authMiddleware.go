package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/models"
	"henrique.mendes/users-api/utils"
)

type usersRepository interface {
	Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB
	Create(user models.User) (models.User, error)
	FindByEmail(email string) models.User
	FindById(userId uint) models.User
	FindByNamePaginated(name string, page int, limit int) ([]models.User, int)
	UpdateUserById(userId uint64, data request.UserUpdateRequest) models.User
	DeleteUserById(userId uint64) error
}

type AuthMiddleware struct {
	repository usersRepository
}

func NewAuthMiddleware(repository usersRepository) *AuthMiddleware {
	return &AuthMiddleware{
		repository: repository,
	}
}

func (auth *AuthMiddleware) CheckAuthUserExists(c *fiber.Ctx) error {
	userId, error := utils.GetTokenInfo(c)
	if error != nil {
		c.Status(503).JSON(error.Error())
		return nil
	}

	user := auth.repository.FindById(uint(userId))
	if user.Id == 0 {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized user",
		})
		return nil
	}

	c.Next()
	return nil
}
