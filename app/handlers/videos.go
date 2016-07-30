package handlers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/suncube/db"
	"github.com/showntop/suncube/models"
)

func IndexVideoHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Write([]byte("name"))
}

func ShowVideoHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var video = models.Video{}
	var videos []models.Video

	id, _ := strconv.Atoi(ps.ByName("id"))
	db.DB.Where(id).First(&video)
	db.DB.Find(&videos).Limit(9)

	render.HTML(rw, http.StatusOK, "videos/show", map[string]interface{}{
		"Result": map[string]interface{}{
			"CurrentVideo":  video,
			"RelatedVideos": videos,
		},
	})
}
