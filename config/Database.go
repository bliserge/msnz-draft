package config

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var SESSION *gocql.Session

func ConnectDb() {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = os.Getenv("database")
	cluster.Consistency = gocql.One
	cluster.ConnectTimeout = time.Second * 5
	cluster.Timeout = time.Second * 5
	session, err := cluster.CreateSession()
	if err != nil {
		panic("Database connection error " + err.Error())
	}
	SESSION = session
}
