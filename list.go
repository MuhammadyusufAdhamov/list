package main

import (
	"database/sql"
	"fmt"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{
		db: db,
	}
}

type List struct {
	Id          int
	Title       string
	Description string
	Assignee    string
	Status      string
	Deadline    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type GetListsQueyParam struct {
	Title    string
	Assignee string
	Page     int32
	Limit    int32
}

func (c *DBManager) CreateList(list *List) (*List, error) {

	query := `
		insert into lists(
		                  Title,
		                  Description,
		                  Assignee,
		                  Status,
		                  Deadline
		) values ($1,$2,$3,$4,$5) 
			returning Id,Title,Description,Assignee,Status,Deadline
	`

	row := c.db.QueryRow(
		query,
		list.Title,
		list.Description,
		list.Assignee,
		list.Status,
		list.Deadline,
	)

	var result List
	err := row.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Assignee,
		&result.Status,
		&result.Deadline,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *DBManager) GetList(id int) (*List, error) {
	var result List

	query := `
		select 
			Id,
			Title,
			Description,
			Assignee,
			Status,
			Deadline,
			Created_at
		from lists
		where id = $1 and deleted_at is null
	`
	row := c.db.QueryRow(query, id)
	err := row.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Assignee,
		&result.Assignee,
		&result.Deadline,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *DBManager) GetAllList(params *GetListsQueyParam) ([]*List, error) {
	var lists []*List

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" limit %d offset %d", params.Limit, offset)

	filter := " where true "
	if params.Title != "" {
		filter += "and title ilike '%" + params.Title + " %' "
	}
	if params.Assignee != "" {
		filter += "and assignee ilike '%" + params.Assignee + " %' "
	}

	query := `
		select 
			Id,
			Title,
			Description,
			Assignee,
			Status,
			Deadline,
			Created_at
		from lists
	` + filter + `and deleted_at is null
		order by created_at desc
	` + limit

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var list List
		err := rows.Scan(
			&list.Id,
			&list.Title,
			&list.Description,
			&list.Assignee,
			&list.Status,
			&list.Deadline,
			&list.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		lists = append(lists, &list)
	}
	return lists, nil
}

func (c *DBManager) UpdateList(list *List) (*List, error) {
	query := `
		update lists set 
			Title=$1,
			Description=$2,
			Assignee=$3,
			Status=$4,
			Deadline=$5,
			Updated_at=$6
		where id = $7
		returning Id,Title,Description,Assignee,Status,Deadline,Created_at,Updated_at
	`
	updateAt := time.Now()
	row := c.db.QueryRow(
		query,
		list.Title,
		list.Description,
		list.Assignee,
		list.Status,
		list.Deadline,
		updateAt,
		list.Id,
	)

	var result List
	err := row.Scan(
		&list.Id,
		&list.Title,
		&list.Description,
		&list.Assignee,
		&list.Status,
		&list.Deadline,
		&list.CreatedAt,
		&list.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *DBManager) DeleteList(list *List) (*List, error) {
	query := `update lists set 
                 Deleted_at=$1 
             where Id = $2
			returning Id,Deleted_at
	`
	delTime := time.Now()
	row := c.db.QueryRow(
		query,
		delTime,
		list.Id,
	)

	var result List

	err := row.Scan(
		&list.Id,
		&list.DeletedAt,
	)

	if err != nil {
		return nil, err
	}
	return &result, err
}
