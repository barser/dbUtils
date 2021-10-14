package internal

import (
	"dbUtils/internal/db"
	"dbUtils/internal/domain"
	"log"
	"sync"
)

func RunApp(wg *sync.WaitGroup, connString string) {
	defer wg.Done()

	conn := db.MakeNewConnector(connString)
	database := db.OpenDB(conn)
	defer database.Close()

	rows, err := database.Query("select login, id, coalesce(last_name,'') from sec_user order by last_name")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.Login, &u.Id, &u.LastName); err != nil {
			log.Println(err)
			return
		}
		users = append(users, u)
	}

	log.Println("users count", len(users))

	for i, user := range users {
		log.Printf("%d. %s (%s) - %s", i+1, user.LastName, user.Login, user.Id.String())
	}

}
