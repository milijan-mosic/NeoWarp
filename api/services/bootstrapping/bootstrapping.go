package bootstrapping

import (
	"github.com/gin-gonic/gin"

	lib "github.com/milijan-mosic/NeoWarp/lib"
	models "github.com/milijan-mosic/NeoWarp/models"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Bootstrap() {
	db := lib.GetDBConnection()

	// db.AutoMigrate(&User{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Bus{})
	db.AutoMigrate(&models.Station{})
}

// EXAMPLES

func CreateUser(c *gin.Context) {
	db := lib.GetDBConnection()

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(200, user)
}

func GetUsers(c *gin.Context) {
	db := lib.GetDBConnection()

	var users []User
	db.Find(&users)
	c.JSON(200, users)
}
