package roles

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	repo         RoleRepository
	config       *config.Config
	tokenManager *common.Auth0TokenManager
}

func (h *RoleHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/roles")

	r.Get(
		"/",
		middlewares.ValidatePermissions([]string{"read:roles"}),
		h.Find,
	)

	r.Get(
		"/:id",
		middlewares.ValidatePermissions([]string{"read:roles"}),
		h.FindOne,
	)

	r.Post(
		"/",
		middlewares.ValidatePermissions([]string{"create:roles"}),
		h.Insert,
	)

	r.Put(
		"/:id",
		middlewares.ValidatePermissions([]string{"update:roles"}),
		h.Update,
	)
}

func New(repo RoleRepository, config *config.Config, clientID, clientSecret string) common.APIHandler {
	return &RoleHandler{
		repo:         repo,
		config:       config,
		tokenManager: common.NewAuth0TokenManager(config, clientID, clientSecret),
	}
}
