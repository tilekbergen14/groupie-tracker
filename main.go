package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tilekbergen14/groupie-tracker/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("Server running on port", *listenAddr)
	log.Fatal(server.Start())
}

// r := gin.Default()

// r.GET("/", handleHome)

// r.Run()
// type Artist struct {
// 	Id    int    `json:"id"`
// 	Image string `json:"image"`
// 	Name  string `json:"name"`
// }

// func handleHome(c *gin.Context) {
// 	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists/3")
// 	// if err != nil {
// 	// 	fmt.Println("Cannot get info")
// 	// 	return
// 	// }

// 	var artist Artist

// 	err := json.NewDecoder(res.Body).Decode(&artist)
// 	if err != nil {
// 		http.Error(c.Writer, err.Error(), 400)
// 		return
// 	}
// 	fmt.Println(artist)
// 	http.ServeFile(w, r, "./web/templates/index.html")
// }
