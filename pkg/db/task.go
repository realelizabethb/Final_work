package db

import (
	"errors"
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

func GetTask(id string) (*Task, error) {
	query := `SELECT id,date,title,comment,repeat FROM scheduler WHERE id=?`
	var task Task
	err := DB.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date=?, title=?,comment=?,repeat=? WHERE id=?`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	kol, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if kol == 0 {
		return errors.New("incorrect query id")
	}
	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id=?`
	res, err := DB.Exec(query, id)
	if err != nil {
		return err
	}
	kol, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if kol == 0 {
		return errors.New("incorrect query id")
	}
	return nil
}

func UpdateDate(date string, id string) error {
	query := `UPDATE scheduler SET date=? WHERE id=?`
	res, err := DB.Exec(query, date, id)
	if err != nil {
		return err
	}
	kol, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if kol == 0 {
		return errors.New("incorrect query id")
	}
	return nil
}
