package main

import (
	"fmt"
	"net/http"
	"os"

	// "os"

	"foodtracker/internal/api"

	// "github.com/supabase-community/supabase-go"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/joho/godotenv"
)

var (
	appwriteClient    client.Client
	todoDatabase      *models.Database
	todoCollection    *models.Collection
	appwriteDatabases *databases.Databases
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(os.Getenv("APPWRITE_PROJECT_ID") + " " + os.Getenv("APPWRITE_API_KEY"))
	appwriteClient = appwrite.NewClient(
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)
	// SUPABASE_URL := os.Getenv("SUPABASE_URL")
	// SUPABASE_PUBLIC_KEY := os.Getenv("SUPABASE_PUBLIC_KEY")
	// var err error
	// client, err = supabase.NewClient(SUPABASE_URL, SUPABASE_PUBLIC_KEY, &supabase.ClientOptions{})
	// if err != nil {
	// 	println("Error creating Supabase client: ", err)
	// }

	prepareDatabase()
	seedDatabase()
	getTodos()
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}

func prepareDatabase() {
	appwriteDatabases = appwrite.NewDatabases(appwriteClient)

	var err error
	todoDatabase, err = appwriteDatabases.Create(
		id.Unique(),
		"TodosDB",
	)
	if err != nil {
		panic(fmt.Sprintf("Error creating database: %v", err))
	}

	todoCollection, err = appwriteDatabases.CreateCollection(
		todoDatabase.Id,
		id.Unique(),
		"Todos",
	)
	if err != nil {
		panic(fmt.Sprintf("Error creating collection: %v", err))
	}

	appwriteDatabases.CreateStringAttribute(
		todoDatabase.Id,
		todoCollection.Id,
		"title",
		255,
		true,
	)

	appwriteDatabases.CreateStringAttribute(
		todoDatabase.Id,
		todoCollection.Id,
		"description",
		255,
		false,
	)

	appwriteDatabases.CreateBooleanAttribute(
		todoDatabase.Id,
		todoCollection.Id,
		"isComplete",
		true,
	)
}

func seedDatabase() {
	testTodo1 := map[string]interface{}{
		"title":       "Buy apples",
		"description": "At least 2KGs",
		"isComplete":  true,
	}

	testTodo2 := map[string]interface{}{
		"title":      "Wash the apples",
		"isComplete": true,
	}

	testTodo3 := map[string]interface{}{
		"title":       "Cut the apples",
		"description": "Don't forget to pack them in a box",
		"isComplete":  false,
	}

	appwriteDatabases.CreateDocument(
		todoDatabase.Id,
		todoCollection.Id,
		id.Unique(),
		testTodo1,
	)

	appwriteDatabases.CreateDocument(
		todoDatabase.Id,
		todoCollection.Id,
		id.Unique(),
		testTodo2,
	)

	appwriteDatabases.CreateDocument(
		todoDatabase.Id,
		todoCollection.Id,
		id.Unique(),
		testTodo3,
	)
}

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

type TodoList struct {
	*models.DocumentList
	Documents []Todo `json:"documents"`
}

func getTodos() {
	todoResponse, _ := appwriteDatabases.ListDocuments(
		todoDatabase.Id,
		todoCollection.Id,
	)

	var todos TodoList
	todoResponse.Decode(&todos)

	for _, todo := range todos.Documents {
		fmt.Printf("Title: %s\nDescription: %s\nIs Todo Complete: %t\n\n", todo.Title, todo.Description, todo.IsComplete)
	}
}
