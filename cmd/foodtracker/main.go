package main

import (
	"net/http"
	// "os"

	"foodtracker/internal/api"

	// "github.com/supabase-community/supabase-go"
)



// var client *supabase.Client

func main() {
	// SUPABASE_URL := os.Getenv("SUPABASE_URL")
	// SUPABASE_PUBLIC_KEY := os.Getenv("SUPABASE_PUBLIC_KEY")
	// var err error
	// client, err = supabase.NewClient(SUPABASE_URL, SUPABASE_PUBLIC_KEY, &supabase.ClientOptions{})
	// if err != nil {
	// 	println("Error creating Supabase client: ", err)
	// }
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
