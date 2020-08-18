package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jialijelly/sample_blog_server/models"
	"github.com/jmoiron/sqlx"
)

var (
	createArticlesTable = `CREATE TABLE IF NOT EXISTS Articles (
		id CHAR(36),
		title TEXT,
		content TEXT,
		author TEXT,
		PRIMARY KEY (id)
	)`
)

type mySQL struct {
	*sqlx.DB
}

func connect(config models.DBConfig) (*mySQL, error) {
	retries := 10
	var mysql *sqlx.DB
	var err error
	for i := 0; i < retries && mysql == nil; i++ {
		mysql, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name))
		time.Sleep(1 * time.Second)
	}
	return &mySQL{DB: mysql}, err
}

func NewSQLServer(config models.DBConfig) (DB, error) {
	db, err := connect(config)
	if err != nil {
		return nil, err
	}
	err = db.init()
	return db, err
}

func (db *mySQL) init() error {
	createArticlesTableStatement, err := db.Prepare(createArticlesTable)
	if err != nil {
		return err
	}
	_, err = createArticlesTableStatement.Exec()
	if err != nil {
		return err
	}
	log.Println("Initialise database successfully")
	return nil
}

func (db *mySQL) CreateArticle(article models.ArticleInfo) models.Error {
	_, err := db.NamedExec("REPLACE into Articles values(:id, :title, :content, :author)", article)
	if err != nil {
		log.Printf("Failed to create article in database : %v", err)
		return models.DatabaseError()
	}
	return nil
}

func (db *mySQL) ListArticles() ([]models.ArticleInfo, models.Error) {
	res := []models.ArticleInfo{}
	err := db.Select(&res, "SELECT * FROM Articles")
	if err != nil {
		log.Printf("Failed to get all articles from database : %v", err)
		return nil, models.DatabaseError()
	}
	return res, nil
}

func (db *mySQL) GetArticle(id string) (models.ArticleInfo, models.Error) {
	res := models.ArticleInfo{}
	err := db.Get(&res, "SELECT * FROM Articles WHERE id=?", id)
	if err != nil {
		log.Printf("Failed to get article %v from database : %v", id, err)
		if err == sql.ErrNoRows {

			return res, models.NotFound(id)
		}
		return res, models.DatabaseError()
	}
	return res, nil
}
