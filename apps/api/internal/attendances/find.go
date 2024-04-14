package attendances

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type FindAttendanceResp common.FindResponse[Attendance]

// @FindAttendance godoc
// @Summary Get list attendance details api
// @Description Get list attendance
// @Tags Attendance
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindAttendanceResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /attendances [get]
func (h *AttendanceHandler) Find(c *fiber.Ctx) error {
	query := &AttendanceQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindAttendances.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindAttendances request: %#v\n", query)

	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindAttendances.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindAttendanceResp{Data: data}
	log.Printf("FindAttendances success. Response: %#v\n", resp)

	return c.JSON(resp)
}
