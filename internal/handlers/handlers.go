package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/ojaswiii/MoMoney-Technical-Assignment/internal/driver"
	"github.com/ojaswiii/MoMoney-Technical-Assignment/internal/models"
)

var mutex sync.Mutex

func GetPost(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the request URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post models.Post

	// Check if the post is already in the database
	err = driver.FindPost(id)
	if err == nil {
		// If the post is already in the database, return it
		json.NewEncoder(w).Encode(post)
		log.Println("Post found in database, no need for api call")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Check again if the post is in the database
	err = driver.FindPost(id)
	if err == nil {
		// If the post is already in the database, return it
		json.NewEncoder(w).Encode(post)
		log.Println("Post found in database, no need for api call")
		return
	}

	// If the post is not in the database, fetch it from the external URL
	url := "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the response into a Post struct
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode post", http.StatusInternalServerError)
		return
	}

	// Save the post in database
	driver.SavePost(post)

	// Return the post to the user
	log.Println("Getting post from the given api")
	json.NewEncoder(w).Encode(post)
}
