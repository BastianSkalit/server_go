package main

func main() {
	a := App{}
	// konfigurasi database disini
	a.Initialize("DB_USERNAME", "DB_PASSWORD", "db_go")

	a.Run(":8080")
}
