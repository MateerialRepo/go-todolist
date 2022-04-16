package todo_list

import (
	"github.com/MateerialRepo/golang-todo/pkg/router"
	"github.com/MateerialRepo/golang-todo/pkg/utils"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"time"
)

var rendr *renderer.Render
var db *mgo.Database

//specifying database constant parameters
const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todolist"
	collectionName string = "todolist"
	port           string = ":8585"
)

//defining structure of the database
type (
	todoModel struct {
		ID          string    `bson:"_id omitempty"`
		Title       string    `bson:"title"`
		Completed   bool      `bson:"completed"`
		CreatedAt   time.Time `bson:"created_at"`
		UpdatedAt   time.Time `bson:"updated_at"`
		CompletedAt time.Time `bson:"completed_at"`
	}

	todo struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Completed   bool      `json:"completed"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CompletedAt time.Time `json:"completed_at"`
	}
)

//initialize the database connection
func dbInit() {
	rendr = renderer.New()
	sess, err := mgo.Dial(hostName)
	utils.CheckErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

func main() {
	r := router.Router()

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	//starting a go routine to create our server
	go func() {
		log.Println("The server is listening on port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()
}
