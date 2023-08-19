package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Course struct {
	ID         int    `json: "id"`
	Name       string `json: "name"`
	Price      string `json: "price"`
	Instructor string `json: "instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
		{
			"id": 1,
			"name": "Golang",
			"price": "2000",
			"instructor": "DevLab"
		},
		{
			"id": 2,
			"name": "JavaScript",
			"price": "1000",
			"instructor": "DevLab"
		},
		{
			"id": 3,
			"name": "Flutter",
			"price": "1500",
			"instructor": "DevLab"
		}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		Bodybyte, nil := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(Bodybyte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newCourse.ID = getNextID()
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return

	}
}

func getNextID() int {
	highestID := -1
	for _, course := range CourseList {
		if course.ID > highestID {
			highestID = course.ID
		}
	}
	return highestID + 1
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
