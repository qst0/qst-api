package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Pixel Smithy")
	if err != nil {
		panic(err)
	}
}

func Motd(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Thank you for playing! Join our discord and help us build an awesome game!")
	if err != nil {
		panic(err)
	}
}
//func TodoIndex(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//	if err := json.NewEncoder(w).Encode(todos); err != nil {
//		panic(err)
//	}
//}

//func TodoShow(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	var todoId int
//	var err error
//	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
//		panic(err)
//	}
//	todo := RepoFindTodo(todoId)
//	if todo.Id > 0 {
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(http.StatusOK)
//		if err := json.NewEncoder(w).Encode(todo); err != nil {
//			panic(err)
//		}
//		return
//	}
//
//	// If we didn't find it, 404
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusNotFound)
//	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
//		panic(err)
//	}
//
//}
//
///*
//Test with this curl command:
//curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
//*/
//func TodoCreate(w http.ResponseWriter, r *http.Request) {
//	var todo Todo
//	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//	if err != nil {
//		panic(err)
//	}
//	if err := r.Body.Close(); err != nil {
//		panic(err)
//	}
//	if err := json.Unmarshal(body, &todo); err != nil {
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(422) // unprocessable entity
//		if err := json.NewEncoder(w).Encode(err); err != nil {
//			panic(err)
//		}
//	}
//
//	t := RepoCreateTodo(todo)
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusCreated)
//	if err := json.NewEncoder(w).Encode(t); err != nil {
//		panic(err)
//	}
//}
