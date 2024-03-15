package controllers

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "encoding/json"
    "github.com/joho/godotenv"
    "os"
    "time"
    "log"
)

type Attributes struct {
    Content      []struct {
        Type     string `json:"type"`
        Children []struct {
            Type string `json:"type"`
            Text string `json:"text"`
        } `json:"children"`
    } `json:"Content"`
    Title        string `json:"Title"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
    PublishedAt  time.Time `json:"publishedAt"`
    Slug         string `json:"Slug"`
    FeaturedImage struct {
        Data struct {
            ID          int         `json:"id"`
            Attributes  interface{} `json:"attributes"`
        } `json:"data"`
    } `json:"Featured_image"`
    Gallery struct {
        Data []struct {
            ID         int         `json:"id"`
            Attributes interface{} `json:"attributes"`
        } `json:"data"`
    } `json:"Gallery"`
}

type Minis struct {
    Data []struct {
        ID         int `json:"id"`
        Attributes Attributes `json:"attributes"`
    } `json:"data"`
    Meta struct {
        Pagination struct {
            Page      int `json:"page"`
            PageSize  int `json:"pageSize"`
            PageCount int `json:"pageCount"`
            Total     int `json:"total"`
        } `json:"pagination"`
    } `json:"meta"`
}

func MinisHandler (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    var pageNum = vars["pageNum"]
    if (pageNum == "") {
        pageNum = "1"
    }

    var pagination = "&pagination[page]=" + pageNum + "&pagination[pageSize]=12"

    godotenv.Load()
    apiBase := os.Getenv("API_BASE_URL")
    url := apiBase + "miniatures/?populate=*&sort=createdAt:DESC" + pagination

    var bearer = "Bearer " + os.Getenv("API_TOKEN")
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Authorization", bearer)
    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Add JSON data to Struct
    var miniData Minis
    json.Unmarshal(bodyBytes, &miniData)

    filepaths := []string{
        "./build/css/global.css",
        "./build/css/homepage__intro.css",
        "./build/css/miniatures__intro.css",
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
        PageTitle: "Miniatures - Mike Bricknell-Barlow - Web Dev, Techie, & Geek",
        CriticalStyles: template.CSS(critStyles),
    }

    tmpl, err := template.New("main.html").Funcs(template.FuncMap{
        "inc": func (num int) int {
            return num + 1;
        },
        "dec": func (num int) int {
            return num - 1;
        },
    }).ParseFiles(
        "templates/layout/main.html",
        "templates/miniatures.html",
        "templates/partials/home/styles.html",
        "templates/partials/miniatures/archive.html",
        "templates/partials/miniatures/list.html",
        "templates/partials/miniatures/styles.html",
        "templates/partials/pagination.html",
    )

    if err := tmpl.Execute(os.Stdout, data); err != nil {
        //fmt.Println(err)
    }

    tmpl.Execute(w, map[string]interface{}{"PageTitle":data.PageTitle, "CriticalStyles":data.CriticalStyles, "Minis":[]Minis{miniData}})
}