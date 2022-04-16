package todo_list

import (
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
)

var rendr *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todolist"
	collectionName string = "todolist"
	port           string = ":8585"
)

func main() {

}
