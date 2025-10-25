package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type Paper struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Authors      string `json:"authors"`
	Journal      string `json:"journal"`
	PublishDate  string `json:"publish_date"`
	Keywords     string `json:"keywords"`
	Abstract     string `json:"abstract"`
}


func getPapers(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM papers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var papers []Paper
	for rows.Next() {
		var p Paper
		if err := rows.Scan(&p.ID, &p.Title, &p.Authors, &p.Journal, &p.PublishDate, &p.Keywords, &p.Abstract); err == nil {
			papers = append(papers, p)
		}
	}
	c.JSON(http.StatusOK, papers)
}

func addPaper(c *gin.Context) {
	var p Paper
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := DB.Exec(`INSERT INTO papers (title, authors, journal, publish_date, keywords, abstract) VALUES (?, ?, ?, ?, ?, ?)`,
		p.Title, p.Authors, p.Journal, p.PublishDate, p.Keywords, p.Abstract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Paper added successfully!"})
}

func deletePaper(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM papers WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Paper deleted successfully!"})
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/papers", getPapers)
	r.POST("/api/papers", addPaper)
	r.DELETE("/api/papers/:id", deletePaper)
}
