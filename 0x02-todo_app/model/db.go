package model

imoort (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Setup()  {
	var err error
	db, err = sql.Open("postgres", "host=0.0.0.0 post=5432 user=admin password=admin dbname=todo sslmode=disable")

	if err!=nil{
		fmt.Println("Could not connect to db", err)
	}

	err = db.Ping()
	if err!=nil{
		fmt.Println("Could not ping db", err)
	}
}