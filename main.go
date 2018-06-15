package main

import "os"

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_URL"),
	)
	app.Run(":8080")
	defer app.DB.Close()
}
