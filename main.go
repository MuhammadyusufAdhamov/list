package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	PostgesHost      = "localhost"
	PotsgresUser     = "postgres"
	PostgresPassword = "7"
	PostgresPort     = 5432
	PostgresDatabase = "golang"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		PostgesHost,
		PostgresPort,
		PotsgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	dbManager := NewDBManager(db)

	//lists, err := dbManager.CreateList(&List{
	//	Title:       "Lorem ipsum1",
	//	Description: "Lorem ipsum is simply1",
	//	Assignee:    "lorem1",
	//	Status:      "ipsum1",
	//	Deadline:    "2022-11-02 00:00",
	//})
	//if err != nil {
	//	log.Fatalf("failed to create a list: %v", err)
	//}

	//lists, err := dbManager.GetList(5)
	//if err != nil {
	//	log.Fatalf("failed to get a list: %v", err)
	//}

	lists, err := dbManager.GetAllList(&GetListsQueyParam{
		Title: "Lorem",
		Limit: 3,
		Page:  1,
	})
	if err != nil {
		log.Fatalf("failed to get all list: %v", err)
	}
	for _, lists := range lists {
		PrintList(lists)
	}

	//sendList := List{
	//	Id:          1,
	//	Title:       "Muhammadyusuf",
	//	Description: "Adhamov",
	//	Assignee:    "assigne",
	//	Status:      "status",
	//	Deadline:    "2022-11-02",
	//}
	//lists, err := dbManager.UpdateList(&sendList)
	//if err != nil {
	//	log.Fatalf("failed to update to list: %v", err)
	//}

	//deleteList := List{
	//	Id: 5,
	//}
	//lists, err := dbManager.DeleteList(&deleteList)
	//if err != nil {
	//	log.Fatalf("failed to delete a list: %v", err)
	//}

	//list, err := dbManager.OrderByTitle()
	//if err != nil {
	//	log.Fatalf("failed to order by title a list: %v", err)
	//}
	//for _, lists := range list {
	//	PrintList(lists)
	//}
	//PrintList(lists)
}

func PrintList(list *List) {
	fmt.Println("-----------LIST----------")
	fmt.Println("Id:", list.Id)
	fmt.Println("Title:", list.Title)
	fmt.Println("Description:", list.Description)
	fmt.Println("Assignee:", list.Assignee)
	fmt.Println("Status:", list.Status)
	fmt.Println("Deadline:", list.Deadline)
}
