package main
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"html/template"
	"log"
)
type progres struct{
	id int
	vendor_code string
	name_ru string
	name_kz string
	description string
	price int
}
var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("select * from postgresdb.Products")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []progres{}

	for rows.Next(){
		p := progres{}
		err := rows.Scan(&p.id, &p.vendor_code, &p.name_ru, &p.name_kz, &p.description, &p.price)
		if err != nil{
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, products)
}

func main() {

	db, err := sql.Open("postgresql", "postgres:1234@/postgresdb")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}