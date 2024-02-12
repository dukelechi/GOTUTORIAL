// we want to create a learncodeonline backend API, that users can Get, update and delete courses. Also we want to see if there's no unique id or title given to a course, then the course should not be added, we will create a kind of helper method for doing so. We will use slice as our fake database since we don't have a database in this section. further Down the road we will inject gorilla/mux so that every single route is being handled by a seperate method for all the CRUD operations

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course - file; usually goes into a separate file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // Author is of type Author struct below, with a pointer. We used a pointer instead of a reference because the Author type hasn't been defined below
}

// multiple models can interact with each other
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
} // as Author is included as a type (*Author) in Course struct above, all the properties in Author struct becomes available in Course struct

// CREATING A FAKE DATABASE
var courses []Course // courses is slice of type Course struct

// MIDDLEWARE or HELPER - file; means that helpers/middleware usually goes into a separate file
func (c *Course) Isempty() bool {
	// return c.CourseId == "" && c.CourseName == "" // ensures users fill them
	return c.CourseName == "" // we don't want to check for the courseId, we want the user to be allowed to move further if the courseId is not empty. We want to manually create the courseId, we don't want to rely on the user to pass us the unique Id (see decoder below)

}

func main() {
	fmt.Println("Welcome to API - learnCodeOnline.in")
	// creating a router
	r := mux.NewRouter()

	// seeding of data into our fake Database
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}}) // first seed
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})     // second seed

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":5000", r))
}

// CONTROLLERS - file

// serve home route; we don't want the homepage to be empty when a user opens it

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	// incase we want to set explicit headers
	w.Header().Set("Content-Type", "application/json")
	// throwing up all the info we have in our fake Database
	json.NewEncoder(w).Encode(courses) // whatever we pass in Encode will be thrown back as a json to whoever that's making the request. In this case, we want to throw back the courses in our fake DB
}

// GET ONE COURSE BASED ON REQUEST ID IN GOLANG

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// user that wants to request an info about a course will provide a courseId, so we will loop through the courses to search/compare the courses
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r) // mux has all the variables which will be extracted from the request(r) that is passed on

	// loop through the courses, find matching id and retrun the response
	for _, course := range courses { // we are not worried about the index "_"
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// incase the loop doesn't find anything
	json.NewEncoder(w).Encode("No Course found with the given id")
	return
}

// ADD A COURSE CONTROLLER IN GOLANG
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// when someone passes a data like this "{}"
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // it will obviously give us a return value, that's why we are using underscore
	if course.Isempty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	// generate unique id, convert to string
	// append course into courses

	// rand.Seed(time.Now().UnixNano())
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano())) // seed = time.Now().UnixNano()
	course.CourseId = strconv.Itoa(randomGen.Intn(100))          // courseId will be a random number rannging from 1 -100
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// UPDATE A COURSE CONTROLLER IN GOLANG
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	// someone needs to give us the id of the course to update and also,from the body itself we should be able to grab the values in JSON
	// two values is coming in, one from the request parameter(URL istself) and the other one is from the body
	params := mux.Vars(r)

	// loop, once we get the id, we need to remove that value from the index and add the value again back in the courses but this time it will be happening with my ID; LOOP, ID, REMOVE, ADD WITH my ID
	for index, course := range courses { // loop
		if course.CourseId == params["id"] { // id gotten
			courses = append(courses[:index], courses[index+1:]...) // REMOVING THE GOTTEN ID; courses[:index] means we don't have the initial value, keep checking till you find the index which is the "id", when found omit that particular index "courses[index+1:]" and we don't have the last value of the index which is indicated with "+1", and since it's a variadic operation, we put three dots
			// adding a new value based on the id; Add with my ID
			// getting the value to add from the request body
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			// since it's an update operation, the id should remain the same
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found
}

// DELETE A COURSE CONTROLLER IN GOLANG
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, get the ID, remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break // breaks the entire loop once the index has been retrieved
		}
	}

}
