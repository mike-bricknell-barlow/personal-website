package controllers

import (
    "os"
    "fmt"
)

func AssetTimeStamp(filepath string) int64 {
    f, err := os.Open("./build/scripts/" + filepath)
    if err != nil {
        fmt.Printf("Error reading file %s: %v\n", filepath, err)
        return 1
    }

    statinfo, err := f.Stat()
    if err != nil {
        fmt.Printf("Error getting file stats %s: %v\n", filepath, err)
        return 1
    }

    return statinfo.ModTime().Unix()
}
