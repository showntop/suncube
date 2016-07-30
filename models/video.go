package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
)

type Video struct {
	Name              string
	Description       string `sql:"size:2000"`
	MadePlace         string
	PlayUrl           string
	Images            []AttachImage
	QualityVariations []QualityVariation
	DlinkVariations   []DlinkVariation
	ReleaseDate       time.Time
	// Creator     User
	LikeNum  int
	WatchNum int
	Mins     int //video size b

	gorm.Model
}

type QualityVariation struct {
	VideoID   uint
	Video     Video
	QualityID uint
	Quality   Quality

	gorm.Model
}

type DlinkVariation struct {
	VideoID uint
	Video   Video
	DlinkID uint
	Dlink   Dlink

	gorm.Model
}

type AttachImage struct {
	gorm.Model
	VideoId uint
	Image   AttachImageStorage `sql:"type:varchar(4096)"`
}

type AttachImageStorage struct{ media_library.FileSystem }

func (video Video) Cover() string {
	imageURL := "/images/default_video.png"
	if len(video.Images) > 0 {
		imageURL = video.Images[0].Image.URL()
	}
	return imageURL
}

func (c Video) DefaultUrl() string {
	defaultUrl := fmt.Sprintf("/videos/%v", c.ID)
	return defaultUrl
}

func (c Video) FormatedReleaseDate() string {
	return c.ReleaseDate.Format("2006年01月02")
}

func (AttachImageStorage) GetSizes() map[string]media_library.Size {
	return map[string]media_library.Size{
		"small":  {Width: 320, Height: 320},
		"middle": {Width: 640, Height: 640},
		"big":    {Width: 1280, Height: 1280},
	}
}
