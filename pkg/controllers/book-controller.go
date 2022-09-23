package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pushkar/go-bookstore/pkg/models"
	"github.com/pushkar/go-bookstore/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing bookId")
	}
	book, _ := models.GetbookById(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	requestbook := models.Book{}
	utils.ParseBody(r, requestbook)
	b := requestbook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["bookId"]
	bId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing bookId")
	}
	book := models.DeleteBook(bId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookToUpdate := models.Book{}
	utils.ParseBody(r, bookToUpdate)
	param := mux.Vars(r)
	bookId := param["bookId"]
	bId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing bookId")
	}
	bookDetails, db := models.GetbookById(bId)

	if bookToUpdate.Name != "" {
		bookDetails.Name = bookToUpdate.Name
	}
	if bookToUpdate.Author != "" {
		bookDetails.Author = bookToUpdate.Author
	}
	if bookToUpdate.Publication != "" {
		bookDetails.Publication = bookToUpdate.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
