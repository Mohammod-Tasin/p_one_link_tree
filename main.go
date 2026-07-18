package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Link struct{
	Platform string `json:"platform"`
	URL string `json:"url"`
}

func setHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
}

func githubHandler(w http.ResponseWriter, r *http.Request){
	setHeaders(w)
	data:= Link{
		Platform: "Github",
		URL: "https://github.com/Mohammod-Tasin",
	}
	json.NewEncoder(w).Encode(data)
}

func linkedinHandler(w http.ResponseWriter, r *http.Request){
	setHeaders(w)
	data:= Link{
		Platform: "LinkedIn",
		URL: "https://www.linkedin.com/in/ibnay-sufian-mohammod-tasin-29a82a3b8?utm_source=share_via&utm_content=profile&utm_medium=member_android",
	}
	json.NewEncoder(w).Encode(data)
}

func facebookHandler(w http.ResponseWriter, r *http.Request){
	setHeaders(w)
	data:= Link{
		Platform: "Facebook",
		URL: "https://www.facebook.com/share/1BnjQiMkSQ/",
	}
	json.NewEncoder(w).Encode(data)
}

func main(){
	http.HandleFunc("/api/github", githubHandler)
	http.HandleFunc("/api/linkedin", linkedinHandler)
	http.HandleFunc("/api/facebook", facebookHandler)

	fmt.Println("Server is running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}