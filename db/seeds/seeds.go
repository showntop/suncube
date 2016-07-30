package seeds

import (
	"math/rand"
	"path/filepath"
	"time"

	"github.com/azumads/faker"
	"github.com/jinzhu/configor"

	"github.com/showntop/suncube/db"
)

var Fake *faker.Faker

var Seeds = struct {
	Videos []struct {
		Name        string
		Description string
		MadePlace   string
		Tags        []struct {
			Name string
		}
		Images []struct {
			URL string
		}
		QualityVariations []struct {
			QualityName string
		}
	}
	Qualities []struct {
		Name string
	}
}{}

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob("db/seeds/data/*.yml")
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}
}

func TruncateTables(tables ...interface{}) {
	for _, table := range tables {
		if err := db.DB.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}

		db.DB.AutoMigrate(table)
	}
}
