package main

import (
	"net/http"
	"os"

	"github.com/supabase-community/supabase-go"
)



var client *supabase.Client

func main() {
	SUPABASE_URL := os.Getenv("SUPABASE_URL")
	SUPABASE_PUBLIC_KEY := os.Getenv("SUPABASE_PUBLIC_KEY")
	var err error
	client, err = supabase.NewClient(SUPABASE_URL, SUPABASE_PUBLIC_KEY, &supabase.ClientOptions{})
	if err != nil {
		println("Error creating Supabase client: ", err)
	}

	
	
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	// Get all todos
	todos, count, err := client.From("todos").Select("*", "exact", false).Execute()
	if err != nil {
		println("Error getting todos: ", err)
	}
	println("Todos: ", todos)
	println("Count: ", count)

	bytes := []byte{}
	for _, todo := range todos {
		// byte string
		bytes = append(bytes, []byte(todo["title"].(string))...)
	}
	w.Write(bytes)

}