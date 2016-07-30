package migrations

import (
	"github.com/qor/media_library"

	"github.com/showntop/suncube/db"
	"github.com/showntop/suncube/models"
)

func init() {
	db.DB.AutoMigrate(&media_library.AssetManager{})

	db.DB.AutoMigrate(&models.User{})

	//
	db.DB.AutoMigrate(&models.Quality{})
	db.DB.AutoMigrate(&models.Dlink{})

	//video
	db.DB.AutoMigrate(&models.Video{})
	db.DB.AutoMigrate(&models.AttachImage{})
	db.DB.AutoMigrate(&models.QualityVariation{})
	db.DB.AutoMigrate(&models.DlinkVariation{})
	//

}
