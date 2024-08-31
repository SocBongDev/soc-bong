package students

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindStudent godoc
// @Summary Get list student details api
// @Description Get list student
// @Tags Student
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindStudentResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students [get]
func (h *StudentHandler) Find(c *fiber.Ctx) error {
	ctx, query := c.UserContext(), &StudentQuery{}
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "FindStudents.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "FindStudents request", "req", query)
	data, err := h.repo.Find(ctx, query)
	if err != nil {
		logger.ErrorContext(ctx, "FindStudents.All err", "err", err)
		return fiber.ErrInternalServerError
	}

	resp := FindStudentResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	return c.JSON(resp)
}
