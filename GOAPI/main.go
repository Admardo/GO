package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

//Struct dos Dados
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
}

//Lista dos Albuns
var albums = []album{
    {ID: "1", Title: "Heaven and Hell", Artist: "Black Sabbath"},
    {ID: "2", Title: "Holy Diver", Artist: "Dio"},
    {ID: "3", Title: "Black Album", Artist: "Metallica"},
}

//Função GET que retorna a Lista dos Albuns no formato JSON
func getAlbums(c *gin.Context) {
    c.JSON(http.StatusOK, albums)
}

// Função POST que adiciona Albuns na Lista
func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.JSON(http.StatusCreated, newAlbum)
}

// Função GET que retorna um Item da Lista conforme ID
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.ID == id {
            c.JSON(http.StatusOK, a)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Album Não Encontrado"})
}

//Função Principal com os Endpoints
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)	   
    router.Run("localhost:8080")
}
