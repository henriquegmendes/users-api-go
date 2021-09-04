package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/service"
	"henrique.mendes/users-api/utils"
)

type UsersController struct {
	service *service.UsersService
}

func NewUsersController(service *service.UsersService) *UsersController {
	return &UsersController{
		service: service,
	}
}

// Returns a list of Users. Response is paginated and requestor may filter by user's name.
// @Description Retrieves all users paginated.
// @Summary Retrieves all users paginated.
// @Tags Authenticated Routes
// @Produce json
// @Param name query string false "User Name"
// @Success 200 {array} response.UsersListResponse
// @Router /users [get]
func (contr UsersController) GetUsers(c *fiber.Ctx) error {
	name := c.Query("name", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	response := contr.service.FindByNamePaginated(name, page, limit)

	return c.JSON(response)
}

// Retrieves user based on its Id
// @Description Retrieves user based on its Id
// @Summary Retrieves user based on its Id
// @Tags Authenticated Routes
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} response.UserResponse
// @Router /users/{id} [get]
func (contr UsersController) GetUserById(c *fiber.Ctx) error {
	userId, error := strconv.ParseUint(c.Params("id"), 10, 64)
	if error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Id param must be a number",
		})
	}

	response := contr.service.FindById(uint(userId))
	if response.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(response)
}

// Update a User based on Id info present JWT token
// @Description Update a User based on Id info present JWT token
// @Summary Update a User based on Id info present JWT token
// @Tags Authenticated Routes
// @Produce json
// @Param data body request.UserUpdateRequest true "User Update Data"
// @Success 200 {object} response.UserResponse
// @Router /users [put]
func (contr UsersController) UpdateUser(c *fiber.Ctx) error {
	userId, error := utils.GetTokenInfo(c)
	if error != nil {
		return c.Status(503).JSON(error.Error())
	}

	data := new(request.UserUpdateRequest)
	if err := c.BodyParser(data); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	response := contr.service.UpdateUserById(userId, *data)
	if response.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(response)
}

// Delete a User based on Id info present JWT token
// @Description Delete a User based on Id info present JWT token
// @Summary Delete a User based on Id info present JWT token
// @Tags Authenticated Routes
// @Produce json
// @Success 204
// @Router /users [delete]
func (contr UsersController) DeleteUser(c *fiber.Ctx) error {
	userId, error := utils.GetTokenInfo(c)
	if error != nil {
		return c.Status(503).JSON(error.Error())
	}

	if error := contr.service.DeleteUserById(userId); error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": error.Error(),
		})
	}

	return c.SendStatus(204)
}
