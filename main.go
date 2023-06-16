package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"syahrul/connection"
	"syahrul/middleware"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id           int
	ProjectName  string
	StartDate    time.Time
	EndDate      time.Time
	Duration     string
	Description  string
	postingTime  string
	Html         bool
	Css          bool
	Javascript   bool
	Java         bool
	Technologies []string
	Image        string
	Author       string
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

func main() {
	connection.DatabaseConnect()
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.Static("/public", "public")
	e.Static("/uploads", "uploads")
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/detailproject/:id", detailproject)
	e.GET("/myproject", myProject)
	e.GET("/testimoni", testimoni)
	e.POST("/add-Project", middleware.UploadFile(addProject))
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/edit-project/:id", ressEditProject)
	e.GET("/edit-project/:id", editProject)
	// Register
	e.GET("/form-register", formRegister)
	e.POST("/register", register)

	// Login
	e.GET("/form-login", formLogin)
	e.POST("/login", login)

	e.POST("/logout", logout)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_projects.id, tb_projects.name, start_date, end_date,duration, description, html, css, javascript, java, image, tb_user.name AS author FROM tb_projects JOIN tb_user ON tb_projects.author_id = tb_user.id")

	var ress []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Duration, &each.Description, &each.Html, &each.Css, &each.Javascript, &each.Java, &each.Image, &each.Author)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}
		ress = append(ress, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	projects := map[string]interface{}{
		"Projects":     ress,
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimoni(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	var tmpl, err = template.ParseFiles("views/testimoni.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/myproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func detailproject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var DetailProject = Project{}

	err := connection.Conn.QueryRow(context.Background(),
		"SELECT id, name, start_date, end_date,duration, description, technologies, html, css, javascript, java FROM tb_projects WHERE id=$1", id).Scan(
		&DetailProject.Id, &DetailProject.ProjectName, &DetailProject.StartDate, &DetailProject.EndDate, &DetailProject.Duration, &DetailProject.Description, &DetailProject.Technologies, &DetailProject.Html, &DetailProject.Css, &DetailProject.Javascript, &DetailProject.Java)

	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
	}

	data := map[string]interface{}{
		"Project": DetailProject,
	}

	var tmpl, errTemplate = template.ParseFiles("views/detailproject.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	projectName := c.FormValue("input-name")
	startDate := c.FormValue("input-start")
	endDate := c.FormValue("input-end")
	description := c.FormValue("input-description")
	html := c.FormValue("input-check-html")
	css := c.FormValue("input-check-css")
	javascript := c.FormValue("input-check-javascript")
	java := c.FormValue("input-check-java")
	// konversi value cekbox, string to boolean
	htmlValue := html != ""
	cssValue := css != ""
	javascriptValue := javascript != ""
	javaValue := java != ""
	// parsing string to time.Time
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	sess, _ := session.Get("session", c)
	author := sess.Values["id"].(int)

	image := c.Get("dataFile").(string)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (name, start_date, end_date, description, duration, html, css, javascript, java, image, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		projectName, start, end, description, getDuration(startDate, endDate), htmlValue, cssValue, javascriptValue, javaValue, image, author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, duration, html, css, javascript, java FROM tb_projects WHERE id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Description, &ProjectDetail.Duration, &ProjectDetail.Html, &ProjectDetail.Css, &ProjectDetail.Javascript, &ProjectDetail.Java)

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/edit-project.html")
	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func ressEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Id :", id)

	projectName := c.FormValue("input-name")
	startDate := c.FormValue("input-start")
	endDate := c.FormValue("input-end")
	description := c.FormValue("input-description")
	html := c.FormValue("input-check-html")
	css := c.FormValue("input-check-css")
	javascript := c.FormValue("input-check-javascript")
	java := c.FormValue("input-check-java")
	// postingTime := time.Now()

	// konversi cekbox string to boolean
	htmlValue := html != ""
	cssValue := css != ""
	javascriptValue := javascript != ""
	javaValue := java != ""
	// parsing string to time.Time
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	_, err := connection.Conn.Exec(
		context.Background(), "UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, duration=$5, html=$6, css=$7, javascript=$8, java=$9 WHERE id=$10",
		projectName, start, end, description, getDuration(startDate, endDate), htmlValue, cssValue, javascriptValue, javaValue, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	fmt.Println("edit :", projectName)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func getDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears == 1 {
		duration = strconv.Itoa(durationYears) + " year"
	} else if durationMonths > 1 {
		duration = strconv.Itoa(durationMonths) + " months"
	} else if durationMonths == 1 {
		duration = strconv.Itoa(durationMonths) + " month"
	} else if durationWeeks > 1 {
		duration = strconv.Itoa(durationWeeks) + " weeks"
	} else if durationWeeks == 1 {
		duration = strconv.Itoa(durationWeeks) + " week"
	} else if durationDays > 1 {
		duration = strconv.Itoa(durationDays) + " days"
	} else {
		duration = strconv.Itoa(durationDays) + " day"
	}

	return duration
}

func formRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func register(c echo.Context) error {
	// to make sure request body is form data format, not JSON, XML, etc.
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success!", true, "/form-login")
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}
