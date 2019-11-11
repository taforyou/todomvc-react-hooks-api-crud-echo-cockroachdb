package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

// ประกาศหน้าตา mockTodos
// https://stackoverflow.com/questions/33892599/how-to-initialize-values-for-nested-struct-array-in-golang
// https://stackoverflow.com/questions/24809235/initialize-a-nested-struct
var (
	mockTodos = &TodoResponses{
		TodoResponses: []TodoResponse{
			{Title: "hey1", Id: 1111, Completed: false, Url: "http://localhost:8080/todos/1111"},
			{Title: "hey2", Id: 1112, Completed: false, Url: "http://localhost:8080/todos/1112"},
			{Title: "hey3", Id: 1113, Completed: false, Url: "http://localhost:8080/todos/1113"},
		},
	}
	pointIndex = -1

	// ประกาศแบบนี้อ่อนมาก เพราะว่ามันไม่ใช่หน้าตา Struct จริงๆ
	// mockTodo = &TodoResponse{Title: "hey", Id: 8866, Completed: false, Url: "http://localhost:9000/todos/8866"}
)

type TodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoResponses struct {
	TodoResponses []TodoResponse
}

type TodoResponse struct {
	Title     string `json:"title"`
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Url       string `json:"url"`
}

func main() {

	//pconn := postgresConnection()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))

	// Routes
	e.GET("api/todos", GetTodos)
	e.POST("api/todos/", AddTodos)
	e.PATCH("api/todos/:id", PatchTodos)
	e.DELETE("api/todos/:id", DeleteTodos)

	e.Static("/", "build")
	e.Logger.Fatal(e.Start(":8080"))

	// defer pconn.Close()

}

func DeleteTodos(c echo.Context) (err error) {

	// ####### RESPONSED NEEDED ########
	// {}

	id, _ := strconv.Atoi(c.Param("id"))
	u := &TodoRequest{}
	if err := c.Bind(u); err != nil {
		return err
	}

	// กากมากนะที่ทำแบบนี้ เพราะนี้เท่ากับว่าไปวนหา Index ใหม่ทั้งหมด แต่ mock เฉยๆ เลยไม่ซีเรียสก็ได้เพราะเดี๋ยวก็ไปแก้ที่ DB อีกอยู่ดี
	for i, s := range mockTodos.TodoResponses {
		if id == s.Id {
			pointIndex = i // assign index ที่ทำการลบ
		}
	}

	// https://yourbasic.org/golang/last-item-in-slice/
	// https://stackoverflow.com/questions/29005825/how-to-remove-element-of-struct-array-in-loop-in-golang
	// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang

	// Tips Remove element array of struct
	// a = append(a[:i], a[i+1:]...)
	mockTodos.TodoResponses = append(mockTodos.TodoResponses[:pointIndex], mockTodos.TodoResponses[pointIndex+1:]...)
	fmt.Println(mockTodos.TodoResponses)

	return c.JSON(http.StatusOK, "{}")
}

func PatchTodos(c echo.Context) (err error) {

	// ####### RESPONSED NEEDED ########
	// {
	// 	"title": "1288",
	// 	"completed": "false",
	// 	"id": 1288,
	// 	"url": "http://todo-backend-node-koa.herokuapp.com/todos/1288"
	// }

	id, _ := strconv.Atoi(c.Param("id"))
	u := &TodoRequest{}
	if err := c.Bind(u); err != nil {
		return err
	}

	// กากมากนะที่ทำแบบนี้ เพราะนี้เท่ากับว่าไปวนหา Index ใหม่ทั้งหมด แต่ mock เฉยๆ เลยไม่ซีเรียสก็ได้เพราะเดี๋ยวก็ไปแก้ที่ DB อีกอยู่ดี
	for i, s := range mockTodos.TodoResponses {
		if id == s.Id {
			mockTodos.TodoResponses[i].Completed = !s.Completed
			pointIndex = i // assign index ที่ทำการแก้ไข
		}
	}

	// mockTodo.Completed = !mockTodo.Completed
	// todoResponse := TodoResponse{}
	// todoResponse.Title = u.Title
	// todoResponse.Id = 8888 // Mock
	// todoResponse.Completed = !u.Completed
	// todoResponse.Url = "http://localhost:9000/todos/8888" // Mock

	// return TodoResponses[index ที่ทำการแก้ไข]
	return c.JSON(http.StatusOK, mockTodos.TodoResponses[pointIndex])
}

func AddTodos(c echo.Context) (err error) {

	// ####### RESPONSED NEEDED ########
	// {
	// 	"title": "testpostman",
	// 	"completed": false,
	// 	"id": 1284,
	// 	"url": "http://todo-backend-node-koa.herokuapp.com/todos/1284"
	// }

	// ####### start.v1 map แบบกากๆ ########
	// var u map[string]interface{}
	// if err := c.Bind(&u); err != nil {
	// 	return err
	// }

	// mapString := make(map[string]string)

	// for key, value := range u {
	// 	strKey := fmt.Sprintf("%v", key)
	// 	strValue := fmt.Sprintf("%v", value)

	// 	mapString[strKey] = strValue
	// }

	// fmt.Println(mapString["title"])
	// fmt.Println(mapString["completed"])
	// ####### end.v1 map แบบกากๆ ########

	// ####### start.v2 map แบบเท่ห์ ########
	u := &TodoRequest{}
	if err := c.Bind(u); err != nil {
		return err
	}
	// ####### end.v2 map แบบเท่ห์ ########

	// ####### start.v1 Add แบบอ่อนๆ ########
	// todoResponse := TodoResponse{}
	// todoResponse.Title = u.Title
	// todoResponse.Id = 9999
	// todoResponse.Completed = u.Completed
	// todoResponse.Url = "http://todo-backend-node-koa.herokuapp.com/todos/1279"
	// ####### end.v1 Add แบบอ่อนๆ ########

	// ####### start.v2 Add แบบเทพๆ ########
	// เอา Index มาก่อน
	// arr := []int{2, 3, 5, 7, 11, 13, 1}
	// sort.Ints(arr[:])
	// fmt.Println("last index : ", arr[len(arr)-1])

	idIndex := []int{}
	for _, s := range mockTodos.TodoResponses {
		idIndex = append(idIndex, s.Id)
	}
	sort.Ints(idIndex[:])
	fmt.Println("last index : ", idIndex[len(idIndex)-1])

	newTodoResponse := &TodoResponse{}
	newTodoResponse.Id = idIndex[len(idIndex)-1] + 1
	newTodoResponse.Title = u.Title
	newTodoResponse.Completed = u.Completed
	newTodoResponse.Url = "http://localhost:8080/todos/" + strconv.Itoa(newTodoResponse.Id)

	// ใส่
	mockTodos.TodoResponses = append(mockTodos.TodoResponses, *newTodoResponse)

	return c.JSON(http.StatusCreated, *newTodoResponse)
}

func GetTodos(c echo.Context) (err error) {

	// 	[
	//   {
	//     "title": "blah",
	//     "id": 1279,
	//     "completed": false,
	//     "url": "http://todo-backend-node-koa.herokuapp.com/todos/1279"
	//   }
	// 	]

	// todoResponses := []TodoResponse{}

	// ##### start.Mock data แบบอ่อนๆ ######
	// todoResponse := TodoResponse{}
	// todoResponse.Title = "blah"
	// todoResponse.Id = 8888
	// todoResponse.Completed = false
	// todoResponse.Url = "http://localhost:9000/todos/8888"
	// fmt.Println(todoResponse)
	// ##### end.Mock data แบบอ่อนๆ ######

	// ##### start.Mock data แบบเทพ (เพราะ mock แบบ In memory) ######
	// อ่านเพิ่มได้เรื่อง Pointer ได้ครัช
	// https://stackoverflow.com/questions/38172661/what-is-the-meaning-of-and-in-golang/45993495
	// todoResponses = append(todoResponses, *mockTodo)
	// fmt.Println("todoResponses ", todoResponses)
	// fmt.Println("mockTodos ", mockTodos.TodoResponses)

	return c.JSON(http.StatusOK, mockTodos.TodoResponses)
}

func postgresConnection() *sql.DB {

	conn, err := sql.Open("postgres", "host=yourhostname port=26257 user=yourusername password=yourpassword dbname=yourdatabase_name sslmode=require")

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SUCCESS CONNECTED DB")
	return conn

}
