package parents

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type FindParentResp common.FindResponse[Parent]

// @FindParent godoc
// @Summary Get list parent details api
// @Description Get list parent
// @Tags Parent
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindParentResp
// @Failure 500 {string} string
// @Router /parents [get]
func (h *ParentHandler) Find(c *fiber.Ctx) error {
	query := &ParentQuery{}

	if err := c.QueryParser(query); err != nil {
		log.Println("FindParents.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindParents request: %#v\n", query)

	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindParents.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindParentResp{Data: data, Page: query.GetPage()}
	log.Printf("FindParents success. Response: %#v\n", resp)

	return c.JSON(resp)
}
