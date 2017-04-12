package tst

import (
	"os"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
)

var feedId int

func TestMain(m *testing.M) {

	prepare()
	retCode := m.Run()
	os.Exit(retCode)
}

func prepare() {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"))

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// clean up all tables
	db.Exec(`DELETE FROM Settings`)
	db.Exec(`DELETE FROM HackerNews`)

	db.Exec(`DELETE FROM StackQuestions`)
	db.Exec(`INSERT INTO Settings(Name, Value) VALUES('one', 'one_value')`)
	db.Exec(`INSERT INTO HackerNews(Id, By, Score, Time, Title, Type, Url, Readed, Favorite)
			VALUES(1, 'by', 10, 123456, 'title', 'type', 'http://google.com', 0, 0)`)

	insertFeedQuery := `INSERT INTO RssFeed(Title, Description, Link, UpdateUrl, ImageTitle, ImageUrl, ImageHeight,
			     		        ImageWidth, LastSyncTime, Total, Unreaded, SyncInterval, AlternativeName,
					        RssType, ShowContent, ShowOrder, Folder, LimitFull, LimitHeadersOnly, Broken)
	      		     VALUES('feed_title', 'description', 'www.ya.ru', '', '', '', 0, 0, 0, 0, 0, 1000, 'alternative',
	      		            2, 0, 0, 'devel', 0, 0, 0)
	      		     RETURNING Id`

	rows, err := db.Query(insertFeedQuery)
	if err != nil {
		log.Print(err)
	}
	if rows.Next() {
		rows.Scan(&feedId)
	}

	db.Exec(`INSERT INTO RssItem(FeedId, Title, Summary, Content, Link, Date, ItemId, Readed, Favorite)
			VALUES($1, 'item_title', 'item_summary', 'content', 'www.google.com', 0, '0', 0, 0)`, feedId)

	db.Exec(`INSERT INTO StackQuestions(Title, Link, QuestionId, Tags, Score, AnswerCount, ViewCount,
				                    UserId, UserReputation, UserDisplayName, UserProfileImage,
				                    Classification, Details, Readed, CreationDate, Favorite, Classified)
			VALUES('title', 'www.google.com', 10, 'linux,docker', 5, 10, 100, 1, 12, 'user', 'profile',
			        'docker', '', 0,  0, 0, 1)`)
}



