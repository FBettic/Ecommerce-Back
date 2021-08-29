package rest

import (
	"ecommerce-back/internal/database"
	"ecommerce-back/internal/logs"
	"ecommerce-back/products"
	"fmt"
	"log"
	"os"

	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

// Cargar del archivo llamado ".env"
var _ = godotenv.Load("rest/.env")

// Extraemos los valores almacenados en el .env
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

var productsHandler = products.ProductHandler{}

func Start(port string) {

	log.Println("Connecting to database")

	db := database.GetDB(ConnectionString)
	productsHandler = products.GetProductHandler(db)

	// Listo, aquí ya podemos usar a db!
	log.Println("successful connection")

	doMigrate(db, "prod_db")

	log.Println("Starting Router")

	Router(port)
}

func doMigrate(client *database.MySQL, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}
