package blog

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
		w.WriteHeader(http.StatusInternalServerError)
	}

	Err := ValidateBlogData(data)
	switch Err.Code {
	case 0:
		data.CreatedAt = time.Now().Format(time.DateTime)
		data.UpdatedAt = time.Now().Format(time.DateTime)
		err := bh.storage.CreateBlog(data)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
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

func (bh *BlogHandler) UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var data Blog
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	parm := mux.Vars(r)

	Err := ValidateBlogData(data)
	switch Err.Code {
	case 0:
		data.UpdatedAt = time.Now().Format(time.DateTime)
		err := bh.storage.UpdateBlog(data, parm["id"])
		if err != nil {
			if err.Error() == "not found" {
				w.WriteHeader(http.StatusNotFound)
				break
			}
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			break
		}
		w.WriteHeader(http.StatusOK)
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

func (bh *BlogHandler) DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	parm := mux.Vars(r)

	err := bh.storage.DeleteBlog(parm["id"])
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (bh *BlogHandler) GetAllBlogsHandler(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")

	data, err := bh.storage.GetAllBlogs(term)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		responseData, err := json.Marshal(data)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(responseData)
	}
}

func (bh *BlogHandler) GetBlogHandler(w http.ResponseWriter, r *http.Request) {
	parm := mux.Vars(r)

	data, err := bh.storage.GetBlog(parm["id"])
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		responseData, err := json.Marshal(data)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(responseData)
	}
}
