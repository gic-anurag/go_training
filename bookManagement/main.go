package main

import (
	"BookManagement/dao"
	"BookManagement/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var bed = dao.BookDao{}

func getBookByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Method not allowed")
		return
	}

	name := strings.Split(r.URL.Path, "/")[2]

	books, err := bed.FindByBookName(name)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, books)
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid method")
		return
	}

	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := bed.Insert(book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	} else {
		respondWithJson(w, http.StatusAccepted, map[string]string{
			"message": "Record inserted successfully",
		})
	}

}

func deleteBookByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Method not allowed")
		return
	}
	var reqBody map[string]string

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	name := reqBody["name"]

	err := bed.DeleteBook(name)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{
		"message": "Record deleted successfully",
	})
}

func updateBookByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Method not allowed")
		return
	}
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	name := book.Name

	_ = bed.UpdateBook(name, book)

	respondWithJson(w, http.StatusOK, map[string]string{
		"message": "Record updated successfully",
	})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	bed.Server = "mongodb://localhost:27017/"
	bed.Database = "quickstart"
	bed.Collection = "book"
	bed.Connect()
}

func main() {
	http.HandleFunc("/add-book", createNewBook)
	http.HandleFunc("/get-book/", getBookByName)
	http.HandleFunc("/delete-book", deleteBookByName)
	http.HandleFunc("/update-book", updateBookByName)
	fmt.Println("starting server at 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
