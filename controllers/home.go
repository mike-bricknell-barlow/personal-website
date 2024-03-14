package controllers

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
)

type PageData struct {
    PageTitle string
    CriticalStyles template.CSS
}

func HomeHandler (w http.ResponseWriter, r *http.Request) {
    filepaths := []string{
        "./build/css/global.css",
        "./build/css/homepage__hero.css",
    }

    var critStyles string

    for _, filepath := range filepaths {
        contentBytes, err := ioutil.ReadFile(filepath)
        if err != nil {
            fmt.Printf("Error reading file %s: %v\n", filepath, err)
            continue // Move to the next file if there's an error
        }
        critStyles += string(contentBytes)
    }

    data := PageData{
        PageTitle: "Mike Bricknell-Barlow - Web Dev, Techie, & Geek",
        CriticalStyles: template.CSS(critStyles),
    }

    tmpl := template.Must(template.ParseFiles(
        "templates/layout/main.html",
        "templates/home.html",
        "templates/partials/home/hero.html",
        "templates/partials/home/intro.html",
        "templates/partials/home/styles.html",
        "templates/partials/external-link-icon.html",
    ))
    tmpl.Execute(w, data)
}