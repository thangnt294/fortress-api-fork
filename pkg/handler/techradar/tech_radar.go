// please edit this file only with approval from hnh
package techradar

import (
	"html"
	"net/http"
	"regexp"
	"strings"

	"github.com/dstotijn/go-notion"
	"github.com/dwarvesf/fortress-api/pkg/config"
	"github.com/dwarvesf/fortress-api/pkg/logger"
	"github.com/dwarvesf/fortress-api/pkg/model"
	"github.com/dwarvesf/fortress-api/pkg/service"
	"github.com/dwarvesf/fortress-api/pkg/store"
	"github.com/dwarvesf/fortress-api/pkg/view"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type handler struct {
	store   *store.Store
	service *service.Service
	logger  logger.Logger
	repo    store.DBRepo
	config  *config.Config
}

// New returns a handler
func New(store *store.Store, repo store.DBRepo, service *service.Service, logger logger.Logger, cfg *config.Config) IHandler {
	return &handler{
		store:   store,
		repo:    repo,
		service: service,
		logger:  logger,
		config:  cfg,
	}
}

// List godoc
// @Summary Get list items from DF TechRadar
// @Description Get list items from DF TechRadar
// @Tags TechRadar
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.TechRadar
// @Failure 400 {object} view.ErrorResponse
func (h *handler) List(c *gin.Context) {
	filter := &notion.DatabaseQueryFilter{}

	rings := []string{"Adopt", "Trial", "Assess", "Hold"}
	filterRings := rings

	if c.Query("name") != "" && len(c.Query("name")) < 2 {
		c.JSON(http.StatusBadRequest, view.CreateResponse[any](nil, nil, nil, nil, "name must be at least 2 characters"))
		return
	}

	if len(c.Request.URL.Query()["ring"]) != 0 {
		filterRings = c.Request.URL.Query()["ring"]
	}

	for _, r := range filterRings {
		if !funk.Contains(rings, r) {
			continue
		}
		filter.Or = append(filter.Or, notion.DatabaseQueryFilter{
			Property: "Ring",
			DatabaseQueryPropertyFilter: notion.DatabaseQueryPropertyFilter{
				Select: &notion.SelectDatabaseQueryFilter{
					Equals: r,
				},
			},
		})
	}

	resp, err := h.service.Notion.GetDatabase(h.config.Notion.Databases.TechRadar, filter, nil, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.CreateResponse[any](nil, nil, err, nil, "can't get items tech radar from notion"))
		return
	}

	var techs = []model.TechRadar{}
	for _, r := range resp.Results {
		props := r.Properties.(notion.DatabasePageProperties)

		name := props["Name"].Title[0].Text.Content
		if c.Query("name") != "" {
			input := c.Query("name")
			escaped := html.EscapeString(input)
			matched, err := regexp.MatchString(".*"+strings.ToLower(escaped)+".*", html.EscapeString(strings.ToLower(name)))
			if err != nil || !matched {
				continue
			}
		}

		if r.Icon != nil && r.Icon.Emoji != nil {
			name = *r.Icon.Emoji + " " + props["Name"].Title[0].Text.Content
		}

		assign := ""
		if len(props["Assign"].People) > 0 {
			assign = props["Assign"].People[0].Name
		}
		quadrant := ""
		if props["Quadrant"].Select != nil {
			quadrant = props["Quadrant"].Select.Name
		}
		ring := ""
		if props["Ring"].Select != nil {
			ring = props["Ring"].Select.Name
		}
		categories := []string{}
		for _, c := range props["Categories"].MultiSelect {
			categories = append(categories, c.Name)
		}
		tags := []string{}
		for _, t := range props["Tag"].MultiSelect {
			tags = append(tags, t.Name)
		}

		techs = append(techs, model.TechRadar{
			ID:         r.ID,
			Name:       name,
			Assign:     assign,
			Quadrant:   quadrant,
			Categories: categories,
			Ring:       ring,
			Tags:       tags,
		})
	}

	c.JSON(http.StatusOK, view.CreateResponse[any](techs, nil, nil, nil, "get list earn items successfully"))
}
