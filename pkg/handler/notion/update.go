// Package notion please edit this file only with approval from hnh
package notion

import (
	"net/http"

	"github.com/dstotijn/go-notion"
	"github.com/gin-gonic/gin"

	"github.com/dwarvesf/fortress-api/pkg/model"
	"github.com/dwarvesf/fortress-api/pkg/view"
)

// ListUpdates godoc
// @Summary Get list updates from DF Updates
// @Description Get list updates from DF Updates
// @Tags Notion
// @Accept  json
// @Produce  json
// @Success 200 {object} view.MessageResponse
// @Failure 400 {object} view.ErrorResponse
// @Router /notion/update [get]
func (h *handler) ListUpdates(c *gin.Context) {
	resp, err := h.service.Notion.GetDatabase(h.config.Notion.Databases.Updates, nil, []notion.DatabaseQuerySort{
		{
			Property:  "Created at",
			Direction: notion.SortDirDesc,
		},
	}, 5)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.CreateResponse[any](nil, nil, err, nil, "can't get updates from notion"))
		return
	}

	var updates []model.NotionUpdate

	for _, r := range resp.Results {
		props := r.Properties.(notion.DatabasePageProperties)

		name := props["Name"].Title[0].Text.Content
		if r.Icon != nil && r.Icon.Emoji != nil {
			name = *r.Icon.Emoji + " " + props["Name"].Title[0].Text.Content
		}

		audience := ""
		if len(props["Audience"].MultiSelect) > 0 {
			audience = props["Audience"].MultiSelect[0].Name
		}

		updates = append(updates, model.NotionUpdate{
			ID:        r.ID,
			Name:      name,
			CreatedAt: props["Created at"].Date.Start.Time,
			Audience:  audience,
		})
	}

	c.JSON(http.StatusOK, view.CreateResponse[any](updates, nil, nil, nil, "get list updates successfully"))
}
