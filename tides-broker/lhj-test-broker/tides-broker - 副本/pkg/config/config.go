package config

import (
	"fmt"
	"os"
	"tides-broker/pkg/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	config *Config
	serverConfig *Config
	db     *gorm.DB
	serverdb *gorm.DB
	err    error
)

const (
	URLSuffix    = "cloudtides.vthink.cloud"
	ROLE_HIGHEST = "SITE_ADMIN"
	ORG_HIGHEST  = "SITE"
)

func init() {
	initConfig()
}

func initConfig() {
	godotenv.Load("../.env")

	// connect to server and create a database for broker if is does not exist
	serverConfig = &Config{}
	serverConfig.PostgresHost = os.Getenv("POSTGRES_HOST")
	serverConfig.PostgresPort = os.Getenv("POSTGRES_PORT")
	serverConfig.PostgresUser = os.Getenv("POSTGRES_USER")
	serverConfig.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	serverConfig.PostgresDB = os.Getenv("POSTGRES_DB")

    var serverdbinfo string
	serverdbinfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
                                serverConfig.PostgresHost,
                                serverConfig.PostgresPort,
                                serverConfig.PostgresUser,
                                serverConfig.PostgresPassword,
                                serverConfig.PostgresDB)
	serverdb, err = gorm.Open(postgres.Open(serverdbinfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// check if borker db exists
    rs := serverdb.Exec(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = %s;", os.Getenv("BROKER_DB")))
    //     if rs.Error != nil {
    //         panic(err)
    //     }

    // create broker db if not exists
    var rec = make(map[string]interface{})
    if rs.Find(rec); len(rec) == 0 {
        stmt := fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("BROKER_DB"))
        serverdb.Exec(stmt)
    }

    // close serverdb connection
    sql, err := serverdb.DB()
    defer func() {
        _ = sql.Close()
    }()
    if err != nil {
        panic(err)
    }

    // load in the configuration of broker db
	config = &Config{}
	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	config.PostgresUser = os.Getenv("POSTGRES_USER")
	config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	config.PostgresDB = os.Getenv("BROKER_DB")
	StartDB()
}

// GetConfig returns a pointer to the current config.
func GetConfig() *Config {
	return config
}

// GetDB returns a pointer to the database
func GetDB() *gorm.DB {
	return db
}

// StartDB initiates database with user configuration, also migrates db schema
func StartDB() {
	var dbinfo string
	
	dbinfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDB)
	fmt.Println(dbinfo)
	db, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	
	// defer db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Template{})
	db.AutoMigrate(&models.Org{})
	db.AutoMigrate(&models.Log{})
	db.AutoMigrate(&models.Host{})


	fmt.Println("DB connection success")
	CreateAdmin()
	TemplateSetup()
}

// CreateAdmin sets up an admin account
func CreateAdmin() {
	db := GetDB()
	var adm models.User
	if db.Where("username = ?", config.AdminUser).First(&adm).RowsAffected == 0 {
		admin := models.User{
			Username: config.AdminUser,
			Password: config.AdminPassword,
			Priority: models.UserPriorityHigh,
			Role:     ROLE_HIGHEST,
			OrgName:  ORG_HIGHEST,
			PwReset:  true,
		}
		db.Create(&admin)
	}

	db.Create(&models.Org{
		OrgName: ORG_HIGHEST,
	})
}

// TemplateSetup sets up a VM template instance
func TemplateSetup() {
	db := GetDB()
	var tem models.Template
	if db.Where("name = ?", "tides-boinc-attached").First(&tem).RowsAffected == 0 {
		newTem := models.Template{
			GuestOS:          "Ubuntu-18.04",
			MemorySize:       8,
			Name:             "tides-boinc-attached",
			ProvisionedSpace: 16,
			VMName:           "tides-gromacs",
		}
		db.Create(&newTem)
	}
}
