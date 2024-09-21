package blog

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type BlogHandler struct {
	storage BlogStorage
}

func NewBlogHandler(storage BlogStorage) *BlogHandler {
	return &BlogHandler{
		storage: storage,
	}
}

func (bh *BlogHandler) CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var data Blog
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Print(err.Error())
	}

	Err := ValidateBlogData(data)
	switch Err.Code {
	case 0:
		data.CreatedAt = time.Now().Format(time.DateTime)
		data.UpdatedAt = time.Now().Format(time.DateTime)
		bh.storage.CreateBlog(data)
		w.WriteHeader(http.StatusCreated)
	case 1:
		w.WriteHeader(http.StatusBadRequest)
	case 2:
		w.WriteHeader(http.StatusBadRequest)
	case 3:
		w.WriteHeader(http.StatusBadRequest)
	case 4:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bh *BlogHandler) UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {}

func (bh *BlogHandler) DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {}

func (bh *BlogHandler) GetAllBlogsHandler(w http.ResponseWriter, r *http.Request) {}

func (bh *BlogHandler) GetBlogHandler(w http.ResponseWriter, r *http.Request) {}
