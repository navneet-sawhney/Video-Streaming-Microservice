package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/video", streamVideo)

    log.Println("Starting server on :3000...")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}

func streamVideo(w http.ResponseWriter, r *http.Request) {
    // Open the video file
    file, err := os.Open("Vectors.mp4")  
    if err != nil {
        http.Error(w, "Could not open video file", http.StatusInternalServerError)
        log.Printf("Error opening video file: %s", err.Error())
        return
    }
    defer file.Close()

    // Get the file's metadata for the modification time
    fileInfo, err := file.Stat()
    if err != nil {
        http.Error(w, "Could not retrieve file info", http.StatusInternalServerError)
        log.Printf("Error getting file info: %s", err.Error())
        return
    }

    // Set headers
    w.Header().Set("Content-Type", "video/mp4")
    w.Header().Set("Accept-Ranges", "bytes")

    // Serve the file with its modification time
    http.ServeContent(w, r, "Vectors.mp4", fileInfo.ModTime(), file) 
}
