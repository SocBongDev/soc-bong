package classes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @FindClass godoc
// @Summary Get list class details api
// @Description Get list class
// @Tags Class
// @Accept json
// @Param  ids query []int false "Ids"
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindClassResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes [get]
func (h *ClassHandler) Find(c *fiber.Ctx) error {
	roles, ok := c.Locals("role").([]string)
	userId, _ := c.Locals("userId").(string)
	query := &ClassQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindClasses.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}
	log.Printf("FindClasses request: %+v\n", query)

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
		data, err = h.repo.Find(query)
	} else if isTeacher {
		query.TeacherId = userId
		data, err = h.repo.Find(query)
	} else {
		return fiber.ErrForbidden
	}

	if err != nil {
		log.Println("FindClasses.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindClassResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	log.Printf("FindClasses success. Response: %+v\n", resp)

	return c.JSON(resp)
}
