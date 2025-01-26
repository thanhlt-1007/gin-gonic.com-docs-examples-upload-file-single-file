package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func postFileHandler(context *gin.Context) {
    file, _ := context.FormFile("file")
    dst := "tmp/" + file.Filename
    context.SaveUploadedFile(file, dst)
    context.JSON(
        http.StatusOK,
        gin.H {
            "file": file.Filename,
        },
    )
}

func main() {
    engine := gin.Default()
    engine.MaxMultipartMemory = 8 << 20
    engine.POST("/file", postFileHandler)
    engine.Run()
}
