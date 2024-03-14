package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    sassDir := "./assets/sass/"
    cssDir := "./build/css/"

    // Execute the Sass compiler
    cmd := exec.Command("sass", "--update", "--style", "compressed", sassDir+":"+cssDir)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error compiling Sass:", err)
    } else {
        fmt.Println("Sass compilation successful!")
    }

    jsDir := "./assets/scripts/"
    buildDir := "./build/scripts/"

    files := []string{
        "mini",
    }

    cmd = exec.Command("mkdir", "-p", buildDir)
    cmd.Run()

    for _, file := range files {
        cmd = exec.Command("./node_modules/.bin/uglifyjs", jsDir+file+".js", "--compress", "--mangle", "--output", buildDir+file+".min.js")

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
            fmt.Println("Error compiling script file:", err)
        } else {
            fmt.Println("Script compilation successful!")
        }
    }
}