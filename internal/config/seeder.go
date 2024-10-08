package config

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"simulation/internal/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	numUsers      = 10
	numAdmins     = 3
	numPostsMin   = 10
	numPostsMax   = 30
	numTags       = 10
	numCategories = 6
)

var categories = []string{
	"Technology",
	"Health",
	"Travel",
	"Food",
	"Lifestyle",
	"Finance",
}

var tags = []string{
	"AI",
	"Wellness",
	"Adventure",
	"Recipes",
	"Personal Finance",
	"Remote Work",
	"Fitness",
	"Travel Tips",
	"Productivity",
	"Tech Gadgets",
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func SeedDatabase(db *gorm.DB) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create categories
	var categoryModels []entity.Category
	for _, name := range categories {
		categoryModels = append(categoryModels, entity.Category{Name: name})
	}
	if err := db.Create(&categoryModels).Error; err != nil {
		return err
	}

	// Create tags
	var tagModels []entity.Tag
	for _, name := range tags {
		tagModels = append(tagModels, entity.Tag{Name: name})
	}
	if err := db.Create(&tagModels).Error; err != nil {
		return err
	}


	file, err := os.OpenFile("user.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()


	for i := 0; i < numUsers; i++ {
		email := "user" + strconv.Itoa(i+rand.Intn(1000)) + "@example.com"
		user := entity.User{
			Name:     "User " + strconv.Itoa(i),
			Username: "user" + strconv.Itoa(i),
			Email:    email,
			Password: "password", 
		}

		if i < numAdmins {
			admin := entity.Admin{
				User: user,
			}

			// Ensure unique email for admin
			if err := db.Where("email = ?", admin.Email).First(&entity.Admin{}).Error; err == gorm.ErrRecordNotFound {
	
				hashedPassword, err := HashPassword(admin.User.Password)
				if err != nil {
					return err
				}
				admin.User.Password = hashedPassword

				if err := db.Create(&admin).Error; err != nil {
					return err
				}

				userDetails := fmt.Sprintf("Admin: %s, Email: %s, Password: %s\n", admin.User.Username, admin.User.Email, admin.User.Password)
				if _, err := file.WriteString(userDetails); err != nil {
					return err
				}

	
				numPosts := r.Intn(numPostsMax-numPostsMin+1) + numPostsMin
				for j := 0; j < numPosts; j++ {
				
					numCategoriesToSelect := r.Intn(len(categoryModels)) + 1 
					startIndex := r.Intn(len(categoryModels) - numCategoriesToSelect + 1)

					blog := entity.Blog{
						Title:     "Blog Post " + strconv.Itoa(j) + " by " + admin.Username,
						Content:   "Content of blog post " + strconv.Itoa(j),
						ReadTime:  r.Intn(10) + 1,
						AdminID:   admin.ID,
						Categories: categoryModels[startIndex : startIndex+numCategoriesToSelect], 
						Tags:      getRandomTags(tagModels, 3, r), 
					}
					if err := db.Create(&blog).Error; err != nil {
						return err
					}
				}
			}
		} else {
			hashedPassword, err := HashPassword(user.Password) 
			if err != nil {
				return err
			}
			user.Password = hashedPassword 

			if err := db.Create(&user).Error; err != nil {
				return err
			}

			userDetails := fmt.Sprintf("User: %s, Email: %s, Password: %s\n", user.Username, user.Email, user.Password)
			if _, err := file.WriteString(userDetails); err != nil {
				return err
			}
		}
	}

	return nil
}


func getRandomTags(tags []entity.Tag, n int, r *rand.Rand) []entity.Tag {
	if len(tags) < n {
		return tags 
	}

	perm := r.Perm(len(tags))
	randomTags := make([]entity.Tag, n)
	for i := 0; i < n; i++ {
		randomTags[i] = tags[perm[i]]
	}
	return randomTags
}
