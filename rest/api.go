package rest

import (
	"log"
	"os"

	"github.com/fbettic/ecommerce-back/internal/database"
	"github.com/fbettic/ecommerce-back/internal/logs"
	"github.com/fbettic/ecommerce-back/products"

	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/golang-migrate/migrate"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

var productsHandler = products.ProductHandler{}

func Start() {

	db := database.GetDB()
	productsHandler = products.GetProductHandler(db)

	// Listo, aquí ya podemos usar a db!
	log.Println("Successful connection")

	doMigrate(db, os.Getenv("DATABASE_NAME"))

	log.Println("Starting Router")

	Router(os.Getenv("PORT"))
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
	logs.Log().Infof("Current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("No migration needed")
	}
}
