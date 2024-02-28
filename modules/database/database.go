package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/rs/xid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	Template          []Template `json:"Template"`
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

type Template struct {
	ID             uint           `gorm:"primarykey" json:"ID,omitempty"`
	Type           int            `json:"type"`
	Title          string         `json:"title"`
	Name           string         `json:"name"`
	Hostname       string         `json:"hostname"`
	Description    string         `json:"description"`
	Info_url       string         `json:"info_url"`
	Categories     pq.StringArray `gorm:"type:text[]" json:"categories"`
	Platform       string         `json:"platform"`
	Logo           string         `json:"logo"`
	Image          string         `json:"image"`
	Restart_policy string         `json:"restart_policy"`
	Network        string         `json:"network"`
	Run            pq.StringArray `gorm:"type:text[]" json:"run"`
	Env            pq.StringArray `gorm:"type:text[]" json:"env"`
	Cmd            pq.StringArray `gorm:"type:text[]" json:"cmd"`
	Ports          pq.StringArray `gorm:"type:text[]" json:"ports"`
	Volumes        pq.StringArray `gorm:"type:text[]" json:"volumes"`
	UserID         uint           `json:"UserID,omitempty"`
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
			log.Fatal(err.Error())
		}
		log.Println("SQLite connection with Vessl-database opened successfully")

	} else {

		db, err = gorm.Open(sqlite.Open("./database/vessl-database.db"), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error(), "failed to open database")
			log.Fatal(err.Error())
		}
		log.Println("SQLite connection with Vessl-database opened successfully")
	}

	err := db.AutoMigrate(&User{}, &Template{})
	if err != nil {
		fmt.Println(err.Error(), "failed to open database")
		log.Fatal(err.Error())
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

func GetUsersValidation() (users Users) {
	db.Preload("Template").Find(&users.Users)
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
	db.Select(clause.Associations).Delete(&user)
	return "User successfully deleted"
}

// func (s *User) Update() string {
// 	db.Preload("Permission").Save(s)
// 	return "User successfully updated!"
// }

func UpdateUser(user User) string {
	db.Save(&user)
	return "User successfully updated!"
}

func SaveTemplate(user User, appTemplate Template) string {
	appTemplate.UserID = user.ID
	db.Preload("Template").Save(&appTemplate)
	return "Template successfully saved!"
}

func DeleteTemplate(user User, appTemplate Template) string {
	appTemplate.UserID = user.ID
	db.Preload("Template").Delete(&appTemplate)
	return "Template successfully deleted!"
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

func GetAppsTemplates(Id string) (string, []Template) {
	user := User{}
	db.Preload("Template").Where(User{UUID: Id}).Find(&user)
	return user.AppsRepositoryUrl, user.Template
}

func (s *User) UpdateApiKey() *User {
	ApiKey := uuid.NewString()
	db.Model(&s).Update("ApiKey", ApiKey)
	db.Model(&s).Update("ApiKeyTs", time.Now().UnixMilli())

	s.ApiKey = ApiKey
	s.Password = ""

	return s
}
