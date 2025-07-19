package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"

	dbsqlc "example.com/virtual-card-x/db/sqlc"
)

type Server struct {
	queries *dbsqlc.Queries
	router  *gin.Engine
}

func New(dsn string) (*Server, error) {
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	s := &Server{
		queries: dbsqlc.New(db),
	}
	s.setupRouter()
	return s, nil
}

func (s *Server) setupRouter() {
	r := gin.Default()
	r.POST("/users", s.createUser)
	r.GET("/users", s.listUsers)
	s.router = r
}

func (s *Server) Run(addr string) error {
	if addr == "" {
		addr = ":8080"
	}
	return s.router.Run(addr)
}

func (s *Server) createUser(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := s.queries.CreateUser(c, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) listUsers(c *gin.Context) {
	users, err := s.queries.ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
