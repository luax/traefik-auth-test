package main

import (
	"log"
	"net/http"
	"os"
	"traefik-auth-test/lib/utils"
	"traefik-auth-test/services/postgres"
	"traefik-auth-test/services/postgres/schema"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	db := postgres.NewPostgres(dbUrl)
	defer db.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.PrintRequest(r)
		query := r.FormValue("query")
		rows, err := db.Query("SELECT * FROM messages WHERE lower(message) LIKE '%' || lower($1) || '%'", query)
		if err != nil {
			log.Println(err)
			utils.Error(w, http.StatusForbidden, "DB error")
			return
		}
		defer rows.Close()
		messages := []schema.Message{}
		for rows.Next() {
			message := schema.Message{}
			if err := rows.Scan(&message.ID, &message.Message, &message.CreatedAt); err == nil {
				messages = append(messages, message)
			}
		}
		utils.Ok(w, messages)
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
