package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const coursePath = "courses"
const apiBasePath = "/api"

type Course struct {
	CourseID   int     `json: "id"`
	CourseName string  `json: "coursename"`
	Price      float64 `json: "price"`
	instructor string  `json: "instructor"`
}

func setUpDB() {
	var err error
	DB, err := sql.Open("mysql", "root:abc123@tcp(127.0.0.1:3305)/todojojo")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect Success ", DB)
	DB.SetConnMaxIdleTime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}

func getCourseList() ([]Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	results, err := DB.QueryContext(ctx, `SELECT 
	id, 
	coursename, 
	price, 
	instructor 
	FROM onlinecourse`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	courses := make([]Course, 0)
	for results.Next() {
		var course Course
		results.Scan(&course.id,
			&course.coursename,
			&course.price,
			&course.instructor)

		courses = append(courses, course)
	}
	return courses, nil

}

func handlerCourses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		courseList, err := getCourseList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(courseList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var course Course
		err := json.NewDecoder(r.Body).Decode(&course)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		CourseID, err := insertProduct(course)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"id": %d}`, CourseID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func insertProduct(course Course) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := DB.ExecContext(ctx, `INSERT INTO onlinecourse
	(coursename, price, instructor)
	VALUES (?, ?, ?)`,
		course.CourseName,
		course.Price,
		course.instructor)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	Insertid, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(Insertid), nil
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPath[len(urlPath)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("no course with id %d", ID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		cousreJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(cousreJSON)

	case http.MethodPut:
		var updateCourse Course
		Bodybyte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, &updateCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updateCourse.ID != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updateCourse
		CourseList[listItemIndex] = *course
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func enableCorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() {
	coursesHandler := http.HandlerFunc(handlerCourses)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, coursePath), enableCorsMiddleware(coursesHandler))
}

func main() {

	setUpDB()
	SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
