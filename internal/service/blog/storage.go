package blog

import (
	"database/sql"
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

func (s *BlogStorage) CreateBlog(data Blog) {
	query := `INSERT INTO blogs(title, content, category, tags, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(query, data.Title, data.Content, data.Category, pq.Array(data.Tags), data.CreatedAt, data.UpdatedAt)
	if err != nil {
		log.Print(err.Error())
	}
}

func (s *BlogStorage) UpdateBlog() {}

func (s *BlogStorage) DeleteBlog() {}
