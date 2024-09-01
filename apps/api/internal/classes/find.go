package classes

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindClass godoc
// @Summary Get list class details api
// @Description Get list class
// @Tags Class
// @Accept json
// @Param  query query ClassQuery false "Query parameters"
// @Success 200 {object} FindClassResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes [get]
func (h *ClassHandler) Find(c *fiber.Ctx) error {
	ctx := c.UserContext()
	roles, ok := c.Locals("role").([]string)
	userId, _ := c.Locals("userId").(string)
	query := &ClassQuery{}
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "FindClasses.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}
	logger.InfoContext(ctx, "FindClasses request", "req", query)

	isAdmin := false
	isTeacher := false

	if ok {
		for _, role := range roles {
			if role == "admin" {
				isAdmin = true
				break
			}

			if role == "teacher" {
				isTeacher = true
			}
		}
	}

	var data []*Class
	var err error

	if isAdmin {
		data, err = h.repo.Find(ctx, query)
	} else if isTeacher {
		query.TeacherId = userId
		data, err = h.repo.Find(ctx, query)
	} else {
		return fiber.ErrForbidden
	}

	if err != nil {
		logger.ErrorContext(ctx, "FindClasses.All err", "err", err)
		return fiber.ErrInternalServerError
	}

	resp := FindClassResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	return c.JSON(resp)
}
