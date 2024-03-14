package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "os"
    "fmt"
    "log"
    "io/ioutil"
    "time"
    "html/template"
    "encoding/json"
)

type Mini struct {
	Data []struct {
		ID         int `json:"id"`
		Attributes struct {
			Content []struct {
				Type     string `json:"type"`
				Children []struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"children"`
			} `json:"Content"`
			Title         string    `json:"Title"`
			CreatedAt     time.Time `json:"createdAt"`
			UpdatedAt     time.Time `json:"updatedAt"`
			PublishedAt   time.Time `json:"publishedAt"`
			Slug          string    `json:"Slug"`
			Completed     time.Time `json:"Completed"`
			FeaturedImage struct {
				Data struct {
					ID         int `json:"id"`
					Attributes struct {
						Name            string `json:"name"`
						AlternativeText any    `json:"alternativeText"`
						Caption         any    `json:"caption"`
						Formats         struct {
							Medium struct {
								Name   string  `json:"name"`
								Hash   string  `json:"hash"`
								Ext    string  `json:"ext"`
								Mime   string  `json:"mime"`
								Path   any     `json:"path"`
								Width  int     `json:"width"`
								Height int     `json:"height"`
								Size   float64 `json:"size"`
								URL    string  `json:"url"`
							} `json:"medium"`
							Small struct {
								Name   string  `json:"name"`
								Hash   string  `json:"hash"`
								Ext    string  `json:"ext"`
								Mime   string  `json:"mime"`
								Path   any     `json:"path"`
								Width  int     `json:"width"`
								Height int     `json:"height"`
								Size   float64 `json:"size"`
								URL    string  `json:"url"`
							} `json:"small"`
							Large struct {
								Name   string  `json:"name"`
								Hash   string  `json:"hash"`
								Ext    string  `json:"ext"`
								Mime   string  `json:"mime"`
								Path   any     `json:"path"`
								Width  int     `json:"width"`
								Height int     `json:"height"`
								Size   float64 `json:"size"`
								URL    string  `json:"url"`
							} `json:"large"`
						} `json:"formats"`
					} `json:"attributes"`
				} `json:"data"`
			} `json:"Featured_image"`
			Gallery struct {
				Data []struct {
					ID         int `json:"id"`
					Attributes struct {
						Name            string `json:"name"`
						AlternativeText any    `json:"alternativeText"`
						Caption         any    `json:"caption"`
						Width           int    `json:"width"`
						Height          int    `json:"height"`
						URL             string  `json:"url"`
						Formats         struct {
							Small struct {
                                Name   string  `json:"name"`
                                Hash   string  `json:"hash"`
                                Ext    string  `json:"ext"`
                                Mime   string  `json:"mime"`
                                Path   any     `json:"path"`
                                Width  int     `json:"width"`
                                Height int     `json:"height"`
                                Size   float64 `json:"size"`
                                URL    string  `json:"url"`
                            } `json:"small"`
                            Medium struct {
                                Name   string  `json:"name"`
                                Hash   string  `json:"hash"`
                                Ext    string  `json:"ext"`
                                Mime   string  `json:"mime"`
                                Path   any     `json:"path"`
                                Width  int     `json:"width"`
                                Height int     `json:"height"`
                                Size   float64 `json:"size"`
                                URL    string  `json:"url"`
                            } `json:"medium"`
							Large struct {
								Name   string  `json:"name"`
								Hash   string  `json:"hash"`
								Ext    string  `json:"ext"`
								Mime   string  `json:"mime"`
								Path   any     `json:"path"`
								Width  int     `json:"width"`
								Height int     `json:"height"`
								Size   float64 `json:"size"`
								URL    string  `json:"url"`
							} `json:"large"`
						} `json:"formats"`
					} `json:"attributes"`
				} `json:"data"`
			} `json:"Gallery"`
		} `json:"attributes"`
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

func MiniHandler (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    godotenv.Load()
    apiBase := os.Getenv("API_BASE_URL")
    url := apiBase + "miniatures/?populate=*&filters[Slug][$eq]=" + vars["mini"]

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

    var miniData Mini
    json.Unmarshal(bodyBytes, &miniData)

    filepaths := []string{
        "./build/css/global.css",
        "./build/css/mini__header.css",
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
        PageTitle: miniData.Data[0].Attributes.Title + " - Mini - Mike Bricknell-Barlow - Web Dev, Techie, & Geek",
        CriticalStyles: template.CSS(critStyles),
    }

    tmpl, err := template.New("main.html").Funcs(template.FuncMap{
        "AssetTimeStamp": AssetTimeStamp,
    }).ParseFiles(
        "templates/layout/main.html",
        "templates/mini.html",
        "templates/partials/mini/header.html",
        "templates/partials/mini/body.html",
        "templates/partials/mini/scripts.html",
        "templates/partials/mini/styles.html",
    )

    tmpl.Execute(w, map[string]interface{}{"PageTitle":data.PageTitle, "CriticalStyles":data.CriticalStyles, "Mini":[]Mini{miniData}})
}