package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Users struct {
	Users []User `json:"Users"`
}

type User struct {
	ID                uint       `gorm:"primarykey" json:"ID"`
	CreatedAt         time.Time  `json:"-"`
	UpdatedAt         time.Time  `json:"-"`
	UUID              string     `json:"UUID"`
	Username          string     `json:"Username"`
	Password          string     `json:"Password"`
	Role              string     `json:"Role"` // Guest, User, Vessl-User, Expert, Admin
	FullName          string     `json:"FullName"`
	Email             string     `json:"Email"`
	Telephone         string     `json:"Telephone"`
	AppsRepositoryUrl string     `json:"AppsRepositoryUrl"`
	ApiKey            string     `json:"ApiKey"`
	ApiKeyTs          int64      `json:"ApiKeyTs"`
	Permission        Permission `gorm:"embedded;embeddedPrefix:perm_" json:"Permission"`
}

type Permission struct {
	Apps           bool `json:"Apps"`
	AppsRepository bool `json:"AppsRepository"`
	AppLauncher    bool `json:"AppLauncher"`
	Images         bool `json:"Images"`
	Volumes        bool `json:"Volumes"`
	Networks       bool `json:"Networks"`
	System         bool `json:"System"`
	Users          bool `json:"Users"`
	HostSettings   bool `json:"HostSettings"`
	HostStats      bool `json:"HostStats"`
}

var (
	once           sync.Once
	db             *gorm.DB
	loginRetries   int
	loginRetriesTs int64
)

func init() {
	once.Do(initialiseDBconn)
}

func initialiseDBconn() {

	if _, err := os.Stat("./database/vessl-database.db"); err != nil {

		log.Println("Creating vessl-database.db...")
		file, err := os.Create("./database/vessl-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("vessl-database.db created")

		db, err = gorm.Open(sqlite.Open("./database/vessl-database.db"), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error(), "failed to connect database")
			panic(err.Error())
		}
		log.Println("SQLite connection with Vessl-database opened successfully")

	} else {

		db, err = gorm.Open(sqlite.Open("./database/vessl-database.db"), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error(), "failed to open database")
			panic(err.Error())
		}
		log.Println("SQLite connection with Vessl-database opened successfully")
	}

	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err.Error(), "failed to open database")
		panic(err.Error())
	}
	fmt.Println("Vessl-database ORM Models created")

	masterUser := User{
		UUID:              "master",
		Username:          "master",
		FullName:          "master user account",
		Password:          "cgMaster@3306",
		Role:              "Admin",
		AppsRepositoryUrl: "https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/templates-1.0.json",
		ApiKey:            "",
		ApiKeyTs:          0,
		Permission: Permission{
			Apps:           true,
			AppsRepository: true,
			HostStats:      true,
			Images:         true,
			HostSettings:   true,
			System:         true,
			Users:          true,
			AppLauncher:    true,
			Volumes:        true,
			Networks:       true,
		},
	}

	user := User{}
	result := db.Where(User{UUID: "master"}).Attrs(masterUser).FirstOrCreate(&user)
	if result.RowsAffected != 0 {
		log.Println("Master user created!")
	}
}

func NewUser() (user *User) {

	return &User{
		UUID:              xid.New().String(),
		Username:          "user-" + xid.New().String(),
		FullName:          "",
		Password:          "",
		Role:              "User",
		AppsRepositoryUrl: "https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/templates-1.0.json",
		ApiKey:            "",
		ApiKeyTs:          0,
		Permission: Permission{
			Apps:           true,
			AppsRepository: false,
			HostStats:      false,
			Images:         false,
			HostSettings:   false,
			System:         false,
			Users:          false,
			AppLauncher:    false,
			Volumes:        false,
			Networks:       false,
		},
	}
}

func GetUsers() (users Users) {
	db.Find(&users.Users)
	// for i := 0; i < len(users.Users); i++ {
	// 	users.Users[i].Password = "**************"
	// 	users.Users[i].ApiKey = ""
	// 	users.Users[i].ApiKeyTs = 0
	// }
	return
}

// func GetModelID(user User) (UserID uint, PermissionID uint) {
// 	db.Find(&user)
// 	UserID = user.Model.ID
// 	PermissionID = 0 //user.Permission.Model.ID
// 	return
// }

func GetUsersValidation() (users Users) {
	db.Find(&users.Users)
	// for _, User := range users.Users {
	// 	User.ApiKeyTs = 0
	// }
	return
}

func AddUser() string {
	db.Create(NewUser())
	log.Println("New user successfully added")
	return "New user successfully added"
}

func DeleteUser(user User) string {
	if user.UUID == "master" {
		return "master user can't be deleted"
	}
	db.Delete(&user)
	return "User successfully deleted"
}

// func (s *User) Update() string {
// 	db.Preload("Permission").Save(s)
// 	return "User successfully updated!"
// }

func UpdateUser(user User) string {
	//UserID, _ := GetModelID(user)
	//user.Model.ID = UserID
	//user.Permission.Model.ID = PermissionID
	db.Debug().Save(&user)
	return "User successfully updated!"
}

func (s *User) Validate() User {
	none := User{}
	none.Username = "invalid"

	if loginRetries >= 4 {
		if time.Now().UnixMilli()-loginRetriesTs < 300000 {
			none.Username = "locked"
			return none
		}
		loginRetries = 0
	}

	Users := GetUsersValidation()
	for i := 0; i < len(Users.Users); i++ {
		if s.Username == Users.Users[i].Username {
			if s.Password == Users.Users[i].Password {
				loginRetries = 0
				return *Users.Users[i].UpdateApiKey()
			}
		}
	}
	loginRetries++
	loginRetriesTs = time.Now().UnixMilli()
	none.Password = fmt.Sprintf("Failed Attemps %d of 5", loginRetries)
	return none
}

func GetAppsRepositoryUrl(Id string) string {
	user := User{}
	db.Where(User{UUID: Id}).Find(&user)
	return user.AppsRepositoryUrl
}

func (s *User) UpdateApiKey() *User {
	ApiKey := uuid.NewString()
	db.Model(&s).Update("ApiKey", ApiKey)
	db.Model(&s).Update("ApiKeyTs", time.Now().UnixMilli())

	s.ApiKey = ApiKey
	s.Password = ""

	return s
}
