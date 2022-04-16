package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// "strings"

	_ "github.com/go-sql-driver/mysql"

	"net/http"
    
    "github.com/dgrijalva/jwt-go"
)

/*
 * Tag... - a very simple struct
 */
type Tag struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Merchant string `json:"merchant_name"`
	Outlet string `json:"outlet_name"`
	Omzet int `json:"omzet"`
}
// username, password := "admin1", "e00cf25ad42683b3df678c61f42c6bda"
func main() {
	username, password := "admin1", "e00cf25ad42683b3df678c61f42c6bda"
	login(username, password)
	// fmt.Println("Login success")
	fmt.Println("server live at port 8000")
	handleRequests()
}

func login(username string, password string) {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/majoo")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("select name, password from users where user_name = ? and password = ?", username, password)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.Name, &tag.Password)
		if err != nil {
			// panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println("Incorrect username or password")
		} else {
			// and then print out the tag's Name attribute
			log.Printf(tag.Name)
		}
	}
}

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Super Secret Information")
}
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Header["Token"] != nil {

      token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf(("Invalid Signing Method"))
        }
        aud := "billing.jwtgo.io"
        checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
        if !checkAudience {
          return nil, fmt.Errorf(("invalid aud"))
        }
        // verify iss claim
        iss := "jwtgo.io"
        checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
        if !checkIss {
          return nil, fmt.Errorf(("invalid iss"))
        }

        return MySigningKey, nil
      })
      if err != nil {
        fmt.Fprintf(w, err.Error())
      }

      if token.Valid {
        endpoint(w, r)
      }

    } else {
		//   fmt.Fprintf(w, "No Authorization Token provided")
		get_mercant_omzet()
    }
  })
}

func get_mercant_omzet() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/majoo")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT a.merchant_name, case when c.created_at BETWEEN '2022-11-01' and '2022-11-30' and c.bill_total = 0 then 0 else c.bill_total end as omzet FROM `merchants` a left join outlets b on a.id = b.merchant_id left join transactions c on a.id = c.merchant_id and b.id = c.outlet_id")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.Merchant, &tag.Omzet)
		// log.Printf(tag.Merchant, tag.Omzet)
		fmt.Println(tag.Merchant, tag.Omzet)
	}
}

func get_mercant_outlet_omzet() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/majoo")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT a.merchant_name, b.outlet_name, case when c.created_at BETWEEN '2022-11-01' and '2022-11-30' and c.bill_total = 0 then 0 else c.bill_total end as omzet FROM `merchants` a left join outlets b on a.id = b.merchant_id left join transactions c on a.id = c.merchant_id and b.id = c.outlet_id")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.Merchant, &tag.Omzet)
		// log.Printf(tag.Merchant, tag.Omzet)
		fmt.Println(tag.Merchant, tag.Outlet, tag.Omzet)
	}
}

func get_mercant_omzet_by_user() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/majoo")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var user = "Admin 1"

	// Execute the query
	results, err := db.Query("SELECT a.merchant_name, case when c.created_at BETWEEN '2022-11-01' and '2022-11-30' and c.bill_total = 0 then 0 else c.bill_total end as omzet FROM `merchants` a left join outlets b on a.id = b.merchant_id left join transactions c on a.id = c.merchant_id left join users d on a.user_id = d.id WHERE d.name = ?", user)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.Merchant, &tag.Omzet)
		// log.Printf(tag.Merchant, tag.Omzet)
		fmt.Println(tag.Merchant, tag.Omzet)
	}
}

func get_mercant_outlet_omzet_by_user() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/majoo")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var user = "Admin 1"

	// Execute the query
	results, err := db.Query("SELECT a.merchant_name, b.outlet_name, case when c.created_at BETWEEN '2022-11-01' and '2022-11-30' and c.bill_total = 0 then 0 else c.bill_total end as omzet FROM `merchants` a left join outlets b on a.id = b.merchant_id left join transactions c on a.id = c.merchant_id and b.id = c.outlet_id left join users d on a.user_id = d.id WHERE d.name = ?", user)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.Merchant, &tag.Omzet)
		// log.Printf(tag.Merchant, tag.Omzet)
		fmt.Println(tag.Merchant, tag.Outlet, tag.Omzet)
	}
}

func handleRequests() {
  http.Handle("/", isAuthorized(homePage))
  log.Fatal(http.ListenAndServe(":8000", nil))
}