package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_pic"`
}

const userFile = "users.json"

func readUsersFromFile() ([]user, error) {
	file, err := os.Open(userFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var users []user
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func writeUsersToFile(users []user) error {
	file, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}

func isUsernameTaken(users []user, username string) bool {
	for _, u := range users {
		if u.Username == username {
			return true
		}
	}
	return false
}

// GET /users
func getUsers(c *gin.Context) {
	users, err := readUsersFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load users from file"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

// POST /signup
func createUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	newUser.ProfilePic = "uploads/default.jpg"
	users, err := readUsersFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load users from file"})
		return
	}
	if isUsernameTaken(users, newUser.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "username already taken"})
		return
	}

	users = append(users, newUser)
	if err := writeUsersToFile(users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func signInUser(c *gin.Context) {
	var loginDetails user
	if err := c.BindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	users, err := readUsersFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read users"})
		return
	}
	for _, u := range users {
		if u.Username == loginDetails.Username && u.Password == loginDetails.Password {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Login successful",
				"username":   u.Username,
				"profilePic": u.ProfilePic,
			})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})

}

func editProfileUser(c *gin.Context) {
	var userDetails user
	if err := c.BindJSON(&userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	users, err := readUsersFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read users"})
		return
	}
	var userPtr *user
	for i := range users {
		if users[i].Username == userDetails.Username {
			userPtr = &users[i]
			break
		}
	}
	if userPtr == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// make sure path to file exist
	if _, err := os.Stat(userDetails.ProfilePic); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File does not exist in the uploads folder"})
		return
	}
	userPtr.ProfilePic = userDetails.ProfilePic
	if err := writeUsersToFile(users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "Profile updated successfully",
		"username":   userPtr.Username,
		"profilePic": userPtr.ProfilePic,
	})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/signup", createUser)
	router.POST("/signin", signInUser)
	router.POST("/editProfile", editProfileUser)
	router.Run("localhost:8080")
}
