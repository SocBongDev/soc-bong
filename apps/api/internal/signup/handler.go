package signup

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repo         UserSignUpRepository
	config       *config.Config
	tokenManager *common.Auth0TokenManager
}

func (h *UserHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/sign-up")

	r.Post(
		"/",
		h.Insert,
	)
}

func New(repo UserSignUpRepository, config *config.Config, clientID, clientSecret string) common.APIHandler {
	return &UserHandler{
		repo:         repo,
		config:       config,
		tokenManager: common.NewAuth0TokenManager(config, clientID, clientSecret),
	}

}
