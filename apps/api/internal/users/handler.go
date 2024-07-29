package users

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repo         UserRepository
	config       *config.Config
	tokenManager *common.Auth0TokenManager
}

func (h *UserHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/users")

	r.Get(
		"/",
		h.Find,
	)

	r.Get(
		"/:id<int,min(1)>",
		h.FindOne,
	)

	r.Post(
		"/",
		h.Insert,
	)
}

func New(repo UserRepository, config *config.Config, clientID, clientSecret string) common.APIHandler {
	return &UserHandler{
		repo:         repo,
		config:       config,
		tokenManager: common.NewAuth0TokenManager(config, clientID, clientSecret),
	}
}
