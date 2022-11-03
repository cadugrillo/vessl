package users

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/xid"
)

type Users struct {
	Users []User `json:"Users"`
}

type User struct {
	ID                string     `json:"ID"`
	Username          string     `json:"Username"`
	Password          string     `json:"Password"`
	Role              string     `json:"Role"` // Guest, User, Expert, Admin
	FullName          string     `json:"FullName"`
	Email             string     `json:"Email"`
	Telephone         string     `json:"Telephone"`
	AppsRepositoryUrl string     `json:"AppsRepositoryUrl"`
	ApiKey            string     `json:"ApiKey"`
	ApiKeyTs          int64      `json:"ApiKeyTs"`
	Permissions       Permission `json:"Permissions"`
}

type Permission struct {
	Dashboard      bool `json:"Dashboard"`
	Apps           bool `json:"Apps"`
	AppsRepository bool `json:"AppsRepository"`
	Users          bool `json:"Users"`
	Settings       bool `json:"Settings"`
	System         bool `json:"System"`
	Images         bool `json:"Images"`
	AppLauncher    bool `json:"AppLauncher"`
	Volumes        bool `json:"Volumes"`
	Networks       bool `json:"Networks"`
	Placeholder4   bool `json:"Placeholder4"`
}

type ApiKeyData struct {
	ApiKey string
	Ts     int64
}

var (
	once sync.Once
	db   *sql.DB
	//err  error
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

		db, err = sql.Open("sqlite3", "./database/vessl-database.db")
		if err != nil {
			log.Println("Opening:", err.Error())
			return
		} else {
			log.Println("SQL connection with Vessl database opened successfully")
		}

		addTable, err := db.Prepare("CREATE TABLE IF NOT EXISTS cgUsers(ID varchar(255) primary key, Username varchar(255), Password varchar(255), Role varchar(255), FullName varchar(255), Email varchar(255), Telephone varchar(255), AppsRepositoryUrl varchar(255), ApiKey varchar(255), ApiKeyTs integer, Dashboard boolean, Apps boolean, AppsRepository boolean, Users boolean, Settings boolean, `System` boolean, Images boolean, AppLauncher boolean, Volumes boolean, Networks boolean)")
		if err != nil {
			log.Println("Preparing:", err.Error())
			return
		}
		_, err = addTable.Exec()
		if err != nil {
			log.Println("Executing:", err.Error())
			return
		}
		log.Println("Table cgUsers created successfully")

		var masterUser User
		masterUser.ID = "master"
		masterUser.Username = "master"
		masterUser.FullName = "master user account"
		masterUser.Password = "cgMaster@3306"
		masterUser.Role = "Admin"
		masterUser.AppsRepositoryUrl = "https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/templates-1.0.json"
		masterUser.ApiKey = ""
		masterUser.ApiKeyTs = 0
		masterUser.Permissions.Apps = true
		masterUser.Permissions.AppsRepository = true
		masterUser.Permissions.Dashboard = true
		masterUser.Permissions.Images = true
		masterUser.Permissions.Settings = true
		masterUser.Permissions.System = true
		masterUser.Permissions.Users = true
		masterUser.Permissions.AppLauncher = true
		masterUser.Permissions.Volumes = true
		masterUser.Permissions.Networks = true

		r, err := db.Prepare("INSERT OR REPLACE INTO cgUsers (ID, Username, Password, Role, FullName, Email, Telephone, AppsRepositoryUrl, ApiKey, ApiKeyTs, Dashboard, Apps, AppsRepository, Users, Settings, `System`, Images, AppLauncher, Volumes, Networks) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = r.Exec(masterUser.ID, masterUser.Username, masterUser.Password, masterUser.Role, masterUser.FullName, "", "", masterUser.AppsRepositoryUrl, masterUser.ApiKey, masterUser.ApiKeyTs, masterUser.Permissions.Dashboard, masterUser.Permissions.Apps, masterUser.Permissions.AppsRepository, masterUser.Permissions.Users, masterUser.Permissions.Settings, masterUser.Permissions.System, masterUser.Permissions.Images, masterUser.Permissions.AppLauncher, masterUser.Permissions.Volumes, masterUser.Permissions.Networks)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("Master user created!")
	} else {

		db, err = sql.Open("sqlite3", "./database/vessl-database.db")
		if err != nil {
			log.Println("Opening:", err.Error())
			return
		} else {
			log.Println("SQL connection with Vessl database opened successfully")
		}
	}
}

func GetUsers() Users {

	var users Users
	var user User
	var Password string

	rows, err := db.Query("SELECT * FROM cgUsers")
	if err != nil {
		log.Println(err.Error())
		return Users{}
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &Password, &user.Role, &user.FullName, &user.Email, &user.Telephone, &user.AppsRepositoryUrl, &user.ApiKey, &user.ApiKeyTs, &user.Permissions.Dashboard, &user.Permissions.Apps, &user.Permissions.AppsRepository, &user.Permissions.Users, &user.Permissions.Settings, &user.Permissions.System, &user.Permissions.Images, &user.Permissions.AppLauncher, &user.Permissions.Volumes, &user.Permissions.Networks)
		user.ApiKey = ""
		user.ApiKeyTs = 0
		users.Users = append(users.Users, user)
	}

	return users
}

func AddUser() string {
	r, err := db.Prepare("INSERT INTO cgUsers (ID, Username, Password, Role, FullName, Email, Telephone, AppsRepositoryUrl, ApiKey, ApiKeyTs, Dashboard, Apps, AppsRepository, Users, Settings, `System`, Images, AppLauncher, Volumes, Networks) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	_, err = r.Exec(xid.New().String(), xid.New().String(), "", "User", "", "", "", "https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/templates-1.0.json", "", 0, true, false, false, false, false, false, false, false, false, false)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	log.Println("New user successfully added")

	return "New user successfully added"
}

func UpdateUser(User User) string {

	if User.Password == "" {
		r, err := db.Prepare("UPDATE cgUsers SET Username=?, Role=?, FullName=?, Email=?, Telephone=?, AppsRepositoryUrl=?, Dashboard=?, Apps=?, AppsRepository=?, Users=?, Settings=?, `System`=?, Images=?, AppLauncher=?, Volumes=?, Networks=? WHERE ID=?")
		if err != nil {
			log.Println(err.Error())
			return err.Error()
		}
		_, err = r.Exec(User.Username, User.Role, User.FullName, User.Email, User.Telephone, User.AppsRepositoryUrl, true, User.Permissions.Apps, User.Permissions.AppsRepository, User.Permissions.Users, User.Permissions.Settings, User.Permissions.System, User.Permissions.Images, User.Permissions.AppLauncher, User.Permissions.Volumes, User.Permissions.Networks, User.ID)
		if err != nil {
			log.Println(err.Error())
			return err.Error()
		}
		log.Println("User successfully updated")

		return "User successfully updated!"
	}

	r, err := db.Prepare("UPDATE cgUsers SET Username=?, Password=?, Role=?, FullName=?, Email=?, Telephone=?, AppsRepositoryUrl=?, Dashboard=?, Apps=?, AppsRepository=?, Users=?, Settings=?, `System`=?, Images=?, AppLauncher=?, Volumes=?, Networks=? WHERE ID=?")
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	_, err = r.Exec(User.Username, User.Password, User.Role, User.FullName, User.Email, User.Telephone, User.AppsRepositoryUrl, true, User.Permissions.Apps, User.Permissions.AppsRepository, User.Permissions.Users, User.Permissions.Settings, User.Permissions.System, User.Permissions.Images, User.Permissions.AppLauncher, User.Permissions.Volumes, User.Permissions.Networks, User.ID)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	log.Println("User successfully updated")

	return "User successfully updated!"
}

func DeleteUser(Id string) string {
	if Id == "master" {
		return "master user can't be deleted"
	}
	r, err := db.Prepare("DELETE FROM cgUsers WHERE ID=?")
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	_, err = r.Exec(Id)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	log.Println("User successfully deleted")

	return "User successfully deleted"
}

func Validate(User User) User {

	none := User
	none.Username = "invalid"

	Users := GetUsersValidation()
	for i := 0; i < len(Users.Users); i++ {
		if User.Username == Users.Users[i].Username {
			if User.Password == Users.Users[i].Password {
				return UpdateApiKey(Users.Users[i])
			}
		}
	}
	return none
}

func GetUsersValidation() Users {

	var users Users
	var user User

	rows, err := db.Query("SELECT * FROM cgUsers")
	if err != nil {
		log.Println(err.Error())
		return Users{}
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.FullName, &user.Email, &user.Telephone, &user.AppsRepositoryUrl, &user.ApiKey, &user.ApiKeyTs, &user.Permissions.Dashboard, &user.Permissions.Apps, &user.Permissions.AppsRepository, &user.Permissions.Users, &user.Permissions.Settings, &user.Permissions.System, &user.Permissions.Images, &user.Permissions.AppLauncher, &user.Permissions.Volumes, &user.Permissions.Networks)
		user.ApiKeyTs = 0
		users.Users = append(users.Users, user)
	}

	return users
}

func GetAppsRepositoryUrl(UserId string) string {

	var url string

	rows, err := db.Query("SELECT AppsRepositoryUrl FROM cgUsers WHERE ID=?", UserId)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&url)
	}

	return url
}

func GetApiKeys() []ApiKeyData {

	var apiKeysData []ApiKeyData
	var apiKeyData ApiKeyData

	rows, err := db.Query("SELECT ApiKey, ApiKeyTs FROM cgUsers")
	if err != nil {
		log.Println(err.Error())
		return []ApiKeyData{}
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&apiKeyData.ApiKey, &apiKeyData.Ts)
		apiKeysData = append(apiKeysData, apiKeyData)
	}

	return apiKeysData
}

func UpdateApiKey(User User) User {

	User.Password = ""
	User.ApiKey = uuid.NewString()

	r, err := db.Prepare("UPDATE cgUsers SET ApiKey=?, ApiKeyTs=? WHERE Username=?")
	if err != nil {
		log.Println(err.Error())
	}
	_, err = r.Exec(User.ApiKey, time.Now().UnixMilli(), User.Username)
	if err != nil {
		log.Println(err.Error())
	}
	return User
}
