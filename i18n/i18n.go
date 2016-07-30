package i18n

import (
	"os"
	"path/filepath"

	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/i18n/backends/yaml"

	"github.com/showntop/suncube/db"
)

var I18n *i18n.I18n

func init() {
	RootPath := os.Getenv("GOPATH") + "/src/github.com/showntop/suncube"

	I18n = i18n.New(database.New(db.DB), yaml.New(filepath.Join(RootPath, "config/locales")))
}
