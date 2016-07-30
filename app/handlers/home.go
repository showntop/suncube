package handlers

import (
	// "bytes"
	// "html/template"
	"net/http"
	// "os"

	"github.com/julienschmidt/httprouter"
	// "gopkg.in/authboss.v0"
	// "github.com/qor/widget"
	// "github.com/showntop/suncube/admin"
	"github.com/showntop/suncube/db"
	"github.com/showntop/suncube/models"
)

func IndexHomeHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//logic
	var videos []models.Video
	db.DB.Find(&videos)

	render.HTML(rw, http.StatusOK, "home/index", map[string]interface{}{
		"Result": map[string][]models.Video{"RecentVideos": videos},
	})
}
