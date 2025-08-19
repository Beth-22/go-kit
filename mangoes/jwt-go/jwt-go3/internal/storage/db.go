package storage

import(
	//"errors"
	//"sync"
	//"JWT-GO3/internal/models"
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)


/*

//db coy
var (
	users = make(map[string]models.User)
	mu     sync.RWMutex
)
func CreateUser(user models.User) error{
	mu.Lock()
	defer mu.Unlock() //always unlock after locking, and defer helps to do that when this function is done running

	if _, exists := users[user.Username];exists{
		return errors.New("user already exists")
	}
	users[user.Username]=user
    return nil


}


func GetUser(username string) (models.User, error){
	mu.RLock()
	defer mu.RUnlock()
    
	user, exists := users[username]
	if !exists {
		return models.User{}, errors.New("user Not Found")
	}
	return user,nil

}

func UpdateUser(user models.User) error{
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[user.Username]; !exists {
		return errors.New("user not found")
	}
	users[user.Username] = user
	return nil

}

*/


//using postgres db

var Pool *pgxpool.Pool

func ConnectDB() error{
	err := godotenv.Load()
	if err != nil{
		return fmt.Errorf("error loading .env file: %w",err)
	}

	dbURL := os.Getenv("DB_URL")

	Pool, err = pgxpool.New(context.Background(), dbURL)

	if err != nil{
		return fmt.Errorf("unable to create connection pool: %w",err)

	}

	err = Pool.Ping(context.Background())
	if err != nil{
		return fmt.Errorf("unable to connect to db: %w", err)
	}
	fmt.Println("connected to Postgres db succesfuly!")
	return nil
}