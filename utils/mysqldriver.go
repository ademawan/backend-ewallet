package utils

import (
	"backend-ewallet/configs"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	// connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
	// 	config.Database.Username,
	// 	config.Database.Password,
	// 	config.Database.Address,
	// 	config.Database.Port,
	// 	config.Database.Name,
	// )
	// fmt.Println(connectionString)
	// // "root:@tcp(127.0.0.1:3306)/be5db?charset=utf8&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	// if err != nil {
	// 	log.Info("failed to connect database :", err)
	// 	panic(err)
	// }

	// InitMigrate(db)

	//    postgres://tyenpjggpurdgd:4669642a7071c32c708ceaeaf99680615aed39bfc2712638df1985739b318ef2@ec2-34-194-158-176.compute-1.amazonaws.com:5432/d3514je8tibhi9
	DBURL := fmt.Sprintf("postgres://%s:%v@%s:%s/%s", config.Database.Username, config.Database.Password, config.Database.Address, config.Database.Port, config.Database.Name)
	fmt.Println(DBURL)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		fmt.Printf("Cannot connect to database ")
	} else {
		fmt.Printf("We are connected to the %s database", DBURL)
	}

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	//endPostgres

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	// db.Migrator().DropTable(&entities.Transaction{})
	// db.Migrator().DropTable(&entities.User{})
	// db.AutoMigrate(&entities.User{})

	// db.AutoMigrate(&entities.Transaction{})

}
