package controllers

import ( 
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindSongs godoc
// @Summary find songs
// @Schemes
// @Description fetch all songs data
// @Tags songs
// @Accept json
// @Produce json
// @Success 200 {string} song data
// @Router /songs [get]
func FindSongs(c *gin.Context) {
	var songs []models.Song
	models.DB.Find(&songs)

	c.JSON(http.StatusOK, gin.H{"data": songs})
}

// CreateSong godoc
// @Summary create song
// @Schemes
// @Description create song entry with name and artist
// @Tags songs
// @Accept json
// @Produce json
// @Success 201 {string} song data
// @Router /songs [post]
func CreateSong(c *gin.Context) {
	var input models.CreateSong

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song := models.Song{Name: input.Name, Artist: input.Artist}

	models.DB.Create(&song)

	c.JSON(http.StatusCreated, gin.H{"data": song})
}

// FindSong godoc
// @Summary find song
// @Schemes
// @Description find song entry by id
// @Tags songs
// @Accept json
// @Produce json
// @Success 200 {string} song data
// @Router /songs/{id} [get]
func FindSong(c *gin.Context) {
	var song models.Song

	if err := models.DB.Where("id = ?", c.Param("id")).First(&song).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "song not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": song})
}

// UpdateSong godoc
// @Summary update song
// @Schemes
// @Description update song entry by id
// @Tags songs
// @Accept json
// @Produce json
// @Success 200 {string} song data
// @Router /songs/{id} [put]
func UpdateSong(c *gin.Context) {
	var song models.Song
	var input models.UpdateSong

	if err := models.DB.Where("id = ?", c.Param("id")).First(&song).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "song not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&song).Updates(models.Song{Name: input.Name, Artist: input.Artist})

	c.JSON(http.StatusOK, gin.H{"data": song})
}

// DeleteSong godoc
// @Summary delete song
// @Schemes
// @Description delete song entry by id
// @Tags songs
// @Accept json
// @Produce json
// @Success 204 {string} empty content
// @Router /songs/{id} [delete]
func DeleteSong(c *gin.Context) {
	var song models.Song

	if err := models.DB.Where("id = ?", c.Param("id")).First(&song).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "song not found"})
		return
	}

	models.DB.Delete(&song)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}