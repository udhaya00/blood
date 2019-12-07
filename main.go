package main
import(
_"github.com/go-sql-driver/mysql"
"html/template"
"database/sql"
"net/http"
"log"
)
type UserDetails struct {
	Name string
	Password string
	Contactno string
	City string
	Gender string
	Age string
	BloodGroup string
	Type string
}
func main() {
db,err :=sql.Open("mysql","root:admin123@(localhost:3306)/UV")
if err!=nil {
log.Fatal(err)
}
if err :=db.Ping();err!=nil {
log.Fatal(err)
}
tmpl:=template.Must(template.ParseFiles("signup.html"))
http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request) {
	if r.Method !=http.MethodPost {
	tmpl.Execute(w,nil)
	return
	} 
	details :=UserDetails {
	Name : r.FormValue("username"),
	Password : r.FormValue("password"),
	Contactno : r.FormValue("no"),
	City : r.FormValue("city"),
	Gender : r.FormValue("gender"),
	Age : r.FormValue("age"),
	BloodGroup : r.FormValue("Blood Group"),
	Type : r.FormValue("type"),
}
insert,_ :=db.Query("INSERT INTO bloodgrp VALUES (?,?,?,?,?,?,?,?)",details.Name,details.Password,details.Contactno,details.City,details.Gender,details.Age,details.BloodGroup,details.Type)
defer insert.Close()
tmpl.Execute(w,struct{Success bool}{true})
})
http.ListenAndServe(":8080",nil)
}