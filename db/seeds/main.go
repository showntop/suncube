// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jinzhu/now"

	"github.com/showntop/suncube/db"
	"github.com/showntop/suncube/db/seeds"
	"github.com/showntop/suncube/models"
)

/* How to upload file
 * $ brew install s3cmd
 * $ s3cmd --configure (Refer https://github.com/theplant/qor-example)
 * $ s3cmd put local_file_path s3://qor3/
 */

var (
	fake           = seeds.Fake
	truncateTables = seeds.TruncateTables

	Seeds  = seeds.Seeds
	Tables = []interface{}{
		&models.User{},
		&models.Quality{},
		&models.Video{}, &models.AttachImage{}, &models.QualityVariation{}, &models.DlinkVariation{},
	}
)

func main() {
	truncateTables(Tables...)
	createRecords()
}

func createRecords() {
	fmt.Println("Start create sample data...")

	createAdminUsers()
	fmt.Println("--> Created admin users.")

	createUsers()
	fmt.Println("--> Created users.")

	createQualities()
	fmt.Println("--> Created qualities.")

	createVideos()
	fmt.Println("--> Created videos.")

	fmt.Println("--> Done!")
}

func createAdminUsers() {
	user := models.User{}
	user.Email = "dev@getqor.com"
	user.Password = "$2a$10$a8AXd1q6J1lL.JQZfzXUY.pznG1tms8o.PK.tYD.Tkdfc3q7UrNX." // Password: testing
	user.Confirmed = true
	user.Name = "QOR Admin"
	user.Role = "Admin"
	db.DB.Create(&user)
}

func createUsers() {
	totalCount := 10
	for i := 0; i < totalCount; i++ {
		user := models.User{}
		user.Email = fake.Email()
		user.Name = fake.Name()
		user.Gender = []string{"Female", "Male"}[i%2]
		if err := db.DB.Create(&user).Error; err != nil {
			log.Fatalf("create user (%v) failure, got err %v", user, err)
		}

		day := (-14 + i/45)
		user.CreatedAt = now.EndOfDay().Add(time.Duration(day*rand.Intn(24)) * time.Hour)
		if user.CreatedAt.After(time.Now()) {
			user.CreatedAt = time.Now()
		}
		if err := db.DB.Save(&user).Error; err != nil {
			log.Fatalf("Save user (%v) failure, got err %v", user, err)
		}
	}
}

func createQualities() {
	for _, s := range Seeds.Qualities {
		quality := models.Quality{}
		quality.Name = s.Name
		if err := db.DB.Create(&quality).Error; err != nil {
			log.Fatalf("create quality (%v) failure, got err %v", quality, err)
		}
	}
}

func createVideos() {
	for _, p := range Seeds.Videos {
		// category := findCategoryByName(p.CategoryName)

		video := models.Video{}
		// video.CategoryID = category.ID
		video.Name = p.Name
		// video.Code = p.Code
		video.Description = p.Description
		video.MadePlace = p.MadePlace
		// for _, c := range p.Collections {
		// collection := findCollectionByName(c.Name)
		// video.Collections = append(video.Collections, *collection)
		// }

		if err := db.DB.Create(&video).Error; err != nil {
			log.Fatalf("create video (%v) failure, got err %v", video, err)
		}

		for _, i := range p.Images {
			image := models.AttachImage{}
			if file, err := openFileByURL(i.URL); err != nil {
				fmt.Printf("open file (%q) failure, got err %v", i.URL, err)
			} else {
				defer file.Close()
				image.Image.Scan(file)
			}
			image.VideoId = video.ID
			if err := db.DB.Create(&image).Error; err != nil {
				log.Fatalf("create quality_variation_image (%v) failure, got err %v", image, err)
			}
		}

		for _, cv := range p.QualityVariations {
			quality := findQualityByName(cv.QualityName)

			qualityVariation := models.QualityVariation{}
			qualityVariation.VideoID = video.ID
			qualityVariation.QualityID = quality.ID
			if err := db.DB.Create(&qualityVariation).Error; err != nil {
				log.Fatalf("create quality_variation (%v) failure, got err %v", qualityVariation, err)
			}
		}
	}

}

func findQualityByName(name string) *models.Quality {
	quality := &models.Quality{}
	if err := db.DB.Where(&models.Quality{Name: name}).First(quality).Error; err != nil {
		log.Fatalf("can't find quality with name = %q, got err %v", name, err)
	}
	return quality
}

func openFileByURL(rawURL string) (*os.File, error) {
	if fileURL, err := url.Parse(rawURL); err != nil {
		return nil, err
	} else {
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName := segments[len(segments)-1]

		filePath := filepath.Join("/tmp", fileName)

		if _, err := os.Stat(filePath); err == nil {
			return os.Open(filePath)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return file, err
		}

		check := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := check.Get(rawURL) // add a filter to check redirect
		if err != nil {
			return file, err
		}
		defer resp.Body.Close()
		fmt.Printf("----> Downloaded %v\n", rawURL)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return file, err
		}
		return file, nil
	}
}
