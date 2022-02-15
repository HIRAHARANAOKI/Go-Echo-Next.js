package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

var e = createMux()

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(mysql)/aadhp")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getRows(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT id,name FROM users")
	if err != nil {
		panic(err.Error())
	}
	return rows
}

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getAllUsers(c echo.Context) error {
	db := dbConnect()
	defer db.Close()
	fmt.Println(db)

	rows := getRows(db)
	defer rows.Close()

	users := Users{}
	var results []Users
	for rows.Next() {
		err := rows.Scan(&users.ID, &users.Name)
		if err != nil {
			panic(err.Error())
		} else {
			results = append(results, users)
		}
	}
	fmt.Println(users, "users")
	fmt.Println(results, "results")

	return c.JSON(http.StatusOK, results)
}

func main() {
	e.GET("/", articleIndex)

	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
