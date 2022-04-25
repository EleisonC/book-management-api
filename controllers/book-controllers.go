package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/EleisonC/book-management-api.git/utils"
	"github.com/EleisonC/book-management-api.git/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error While Parsing")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	err := utils.ParseBody(r, CreateBook)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error While Parsing")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	bookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var UpdateBook = &models.Book{}
	err := utils.ParseBody(r, UpdateBook)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error While Parsing")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	booksDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != ""{
		booksDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != ""{
		booksDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != ""{
		booksDetails.Publication = UpdateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


