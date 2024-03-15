package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "bricknellbarlow/website/controllers"
    "os"
    "github.com/joho/godotenv"
)


func main() {
    r := mux.NewRouter()

    godotenv.Load()
    domain := os.Getenv("DOMAIN")
    r.Host(domain)

    r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("assets/images/"))))
    r.PathPrefix("/icons/").Handler(http.StripPrefix("/icons/", http.FileServer(http.Dir("assets/icons/"))))
    r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("build/"))))
    r.PathPrefix("/cms/images/").Handler(http.StripPrefix("/cms/images/", http.FileServer(http.Dir("cms/public/"))))

    //r.HandleFunc("/", controllers.HomeHandler)
    r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/miniatures/", http.StatusSeeOther)
    })

    r.HandleFunc("/miniatures/", controllers.MinisHandler)
    r.HandleFunc("/miniatures/page/{pageNum}/", controllers.MinisHandler)
    r.HandleFunc("/miniatures/{mini}/", controllers.MiniHandler)

    httpPort := os.Getenv("HTTP_PORT")
    httpsPort := os.Getenv("HTTPS_PORT")

    certFile := os.Getenv("CERT_FILE")
    keyFile := os.Getenv("KEY_FILE")

    http.ListenAndServe(":" + httpPort, r)
    http.ListenAndServeTLS(":" + httpsPort, certFile, keyFile, r)
}

