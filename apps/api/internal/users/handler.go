package users

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repo         Repository
	config       *config.Config
	tokenManager *common.Auth0TokenManager
}

func (h *UserHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/users")

	r.Get(
		"/",
		middlewares.ValidatePermissions([]string{"read:users"}),
		h.Find,
	)

	r.Get(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"read:users"}),
		h.FindOne,
	)

	r.Post(
		"/",
		middlewares.ValidatePermissions([]string{"create:users"}),
		h.Insert,
	)

	r.Put(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:users"}),
		h.Update,
	)

	route := router.Group("/sign-up")
	route.Post(
		"/",
		h.SignUp,
	)
}

func New(repo Repository, config *config.Config, clientID, clientSecret string) common.APIHandler {
	return &UserHandler{
		repo:         repo,
		config:       config,
		tokenManager: common.NewAuth0TokenManager(config, clientID, clientSecret),
	}
}
