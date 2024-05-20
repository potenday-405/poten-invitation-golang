package external

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"time"
)

func testInitializer() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}
}

func Test_Godotenv(t *testing.T) {
	testInitializer()
	account := os.Getenv("MYSQL_ACCOUNT")
	t.Log(account)
}

func Test_GormConnection(t *testing.T) {
	testInitializer()
	type User struct {
		UserId       string    `json:"user_id"`
		Email        string    `json:"email"`
		UserPassword string    `json:"user_password"`
		UserName     string    `json:"user_name"`
		Phone        string    `json:"phone"`
		UserStatus   string    `json:"user_status"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	user := &User{
		UserId:       "test2",
		Email:        "test2@test1.com",
		UserPassword: "testpassword2",
		UserName:     "tester2",
		Phone:        "01011112223",
		UserStatus:   "act",
	}
	NewDB().Table("users").Create(user)

	var users []*User
	NewDB().Table("users").Select("user_id,email,user_password,user_name,phone,user_status,created_at,updated_at").Scan(&users)
	t.Log(users)
}
