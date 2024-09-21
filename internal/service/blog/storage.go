package blog

import (
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type BlogStorage struct {
	db *sql.DB
}

func NewBlogStorage(db *sql.DB) *BlogStorage {
	query := `CREATE TABLE IF NOT EXISTS blogs (
id SERIAL PRIMARY KEY,
title TEXT NOT NULL,
content TEXT NOT NULL,
category TEXT NOT NULL,
tags TEXT[] NOT NULL,
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`
	_, err := db.Exec(query)
	if err != nil {
		log.Print(err.Error())
	}
	return &BlogStorage{
		db: db,
	}
}

func (s *BlogStorage) CreateBlog(data Blog) error {
	query := `INSERT INTO blogs(title, content, category, tags, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(query, data.Title, data.Content, data.Category, pq.Array(data.Tags), data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogStorage) UpdateBlog(newData Blog, id string) error {
	query := `SELECT * FROM blogs WHERE id=$1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	var data Blog
	if rows.Next() {
		err = rows.Scan(&data.ID, &data.Title, &data.Content, &data.Category, pq.Array(&data.Tags), &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return err
		}
	} else {
		return errors.New("not found")
	}

	data.Title = newData.Title
	data.Content = newData.Content
	data.Category = newData.Category
	data.Tags = newData.Tags
	data.UpdatedAt = newData.UpdatedAt

	query = `UPDATE blogs SET title=$1, content=$2, category=$3, tags=$4, updated_at=$5 WHERE id=$6`

	_, err = s.db.Exec(query, data.Title, data.Content, data.Category, pq.Array(data.Tags), data.UpdatedAt, data.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BlogStorage) DeleteBlog(id string) error {
	query := `SELECT * FROM blogs WHERE id=$1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	if rows.Next() {
		query = `DELETE FROM blogs WHERE id=$1`
		_, err = s.db.Exec(query, id)
		if err != nil {
			return err
		}
	} else {
		return errors.New("not found")
	}

	return nil
}

func (s *BlogStorage) GetAllBlogs(term string) ([]Blog, error) {
	var dataSlice []Blog
	var query string
	var rows *sql.Rows
	var err error
	if term == "" {
		query = `SELECT * FROM blogs;`
		rows, err = s.db.Query(query)
	} else {
		query = `SELECT * FROM blogs WHERE title ILIKE '%' || $1 || '%' OR content ILIKE '%' || $1 || '%' OR category ILIKE '%' || $1 || '%' OR '$1' = ANY(tags);`
		rows, err = s.db.Query(query, term)
	}
	if err != nil {
		return dataSlice, err
	}
	for rows.Next() {
		var data Blog
		err = rows.Scan(&data.ID, &data.Title, &data.Content, &data.Category, pq.Array(&data.Tags), &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return dataSlice, err
		}
		dataSlice = append(dataSlice, data)
	}
	return dataSlice, nil
}

func (s *BlogStorage) GetBlog(id string) (Blog, error) {
	query := `SELECT * FROM blogs WHERE id=$1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return Blog{}, err
	}

	var data Blog
	if rows.Next() {
		err = rows.Scan(&data.ID, &data.Title, &data.Content, &data.Category, pq.Array(&data.Tags), &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return Blog{}, err
		}
	} else {
		return Blog{}, errors.New("not found")
	}
	return data, nil
}
