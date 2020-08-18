package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jialijelly/sample_blog_server/db"
)

type Server struct {
	handler *handler
	router  *mux.Router
}

type responseWriterLogger struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriterLogger(w http.ResponseWriter) *responseWriterLogger {
	return &responseWriterLogger{ResponseWriter: w}
}

func (r *responseWriterLogger) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func requestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Received request %v %v", req.Method, req.URL)
		start := time.Now()
		res := newResponseWriterLogger(w)
		next.ServeHTTP(res, req)
		log.Printf("Completed request %v %s in %v", res.statusCode, http.StatusText(res.statusCode), time.Since(start))
	})
}

func (s *Server) registerRoutes() {
	s.router.Methods(http.MethodPost).Name("CreateArticle").Path("/articles").HandlerFunc(s.handler.CreateArticle)
	s.router.Methods(http.MethodGet).Name("ListArticles").Path("/articles").HandlerFunc(s.handler.ListArticles)
	s.router.Methods(http.MethodGet).Name("GetArticle").Path("/articles/{article_id}").HandlerFunc(s.handler.GetArticle)
	s.router.Use(requestLoggingMiddleware)
	log.Printf("Serving the following APIs:")
	s.router.Walk(
		func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			path, _ := route.GetPathTemplate()
			name := route.GetName()
			log.Printf("%s: %s", name, path)
			return nil
		})
}

func (s *Server) Run() {
	log.Print("Server started...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", DefaultConfiguration.Server.Host, DefaultConfiguration.Server.Port), s.router))
}

func NewServer() *Server {
	var database db.DB
	var err error
	switch DefaultConfiguration.DB.Type {
	case "mysql":
		database, err = db.NewSQLServer(DefaultConfiguration.DB)
	default:
		log.Fatalf("Unsupported database type = %v", DefaultConfiguration.DB.Type)
		return nil
	}
	if err != nil {
		log.Fatalf("Failed to setup database = %v", err)
		return nil
	}
	server := &Server{
		handler: NewHandler(database),
		router:  mux.NewRouter(),
	}
	server.registerRoutes()
	return server
}
