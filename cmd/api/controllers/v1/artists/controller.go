package artists

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	dto "github.com/srntz/harp-api/internal/api/artist"
	"github.com/srntz/harp-api/internal/api/response"
	"github.com/srntz/harp-api/internal/domain/artist"
	"github.com/srntz/harp-api/internal/infrastructure/repositories"
)

type Controller struct {
	service artist.ArtistService
}

func NewController() *Controller {
	return &Controller{
		service: *artist.NewArtistService(&repositories.ArtistRepository{}),
	}
}

func (c *Controller) GetAlbumByID(ctx *gin.Context) {
	service := artist.NewArtistService(&repositories.ArtistRepository{})
	id := ctx.Param("id")

	// TODO create an RFC 7807 compliant error wrapper. set response body to a proper error
	artist, err := service.GetByID(id)
	if errors.Is(err, pgx.ErrNoRows) {
		ctx.JSON(404, gin.H{})
		return
	}

	// TODO create an RFC 7807 compliant error wrapper. set response body to a proper error
	if err != nil || artist == nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response.NewSuccess(dto.ResponseDTOFromDomainModel(artist)))
}
