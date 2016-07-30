package admin

import (
	"github.com/qor/admin"
	"github.com/qor/media_library"
	"github.com/qor/qor"

	"github.com/showntop/suncube/db"
	"github.com/showntop/suncube/models"
)

var (
	Admin *admin.Admin
)

var (
	ctypes = []string{"视频", "电子书", "音乐"}
)

func init() {
	Admin = admin.New(&qor.Config{DB: db.DB.Set("publish:draft_mode", true)})
	Admin.SetSiteName("聚库")

	assetManager := Admin.AddResource(&media_library.AssetManager{}, &admin.Config{Invisible: true})

	// Add User
	user := Admin.AddResource(&models.User{})
	user.Meta(&admin.Meta{Name: "Gender", Config: &admin.SelectOneConfig{Collection: []string{"Male", "Female", "Unknown"}}})
	user.Meta(&admin.Meta{Name: "Role", Config: &admin.SelectOneConfig{Collection: []string{"Admin", "Maintainer", "Member"}}})
	user.Meta(&admin.Meta{Name: "Confirmed", Valuer: func(user interface{}, ctx *qor.Context) interface{} {
		if user.(*models.User).ID == 0 {
			return true
		}
		return user.(*models.User).Confirmed
	}})

	user.IndexAttrs("ID", "Email", "Name", "Gender", "Role")
	user.ShowAttrs(
		&admin.Section{
			Title: "基本信息",
			Rows: [][]string{
				{"Name"},
				{"Email", "Password"},
				{"Gender", "Role"},
				{"Confirmed"},
			}},
	)
	user.EditAttrs(user.ShowAttrs())

	//////////
	video := Admin.AddResource(&models.Video{}, &admin.Config{Menu: []string{"Video管理"}})
	video.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{AssetManager: assetManager}})

}
