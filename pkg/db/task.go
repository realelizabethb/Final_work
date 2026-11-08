package db

import (
	"fmt"
	"strconv"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date,title,comment,repeat) VALUES (?,?,?,?)`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, fmt.Errorf("err with query: %w", err)
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("err with id: %w", err)
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	query := `SELECT id,date,title,comment,repeat FROM scheduler ORDER BY date ASC LIMIT ?`
	rows, err := DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]*Task, 0)
	for rows.Next() {
		t := &Task{}
		var id int64
		err := rows.Scan(&id, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		tasks = append(tasks, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
