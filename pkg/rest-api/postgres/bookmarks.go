package postgres

import (
	"log"
	"github.com/demas/cowl-go/pkg/rest-api/quzx"
)

// represent a PostgreSQL implementation of quzx.BookmarkRepository
type BookmarkRepository struct {
}

func (r *BookmarkRepository) insertBookmarkTags(tags []string, bookmarkId int) {

	insertTagQuery := `INSERT INTO BookmarkTag(Name, BookmarkCount)
			   VALUES($1, 0) ON CONFLICT(Name) DO UPDATE SET Name = EXCLUDED.Name RETURNING Id`

	insertTagConnectionQuery := `INSERT INTO BookmarkTagConnnection(BookmarkId, TagId)
				     VALUES($1, $2)`

	tx := db.MustBegin()

	for _, tag := range tags {

		var tagId int
		tagRows := db.QueryRow(insertTagQuery, tag)
		tagRows.Scan(&tagId)

		_, err := tx.Exec(insertTagConnectionQuery, bookmarkId, tagId)
		if err != nil {
			log.Println(err)
		}
	}

	tx.Commit()
}

func (r *BookmarkRepository) InsertBookmark(bookmark *quzx.BookmarkPOST) {

	insertBookmarkQuery := `INSERT INTO Bookmark(Url, Title, Description, ReadItLater)
				VALUES(:url, :title, :description, :readitlater)
				RETURNING Id`

	tx := db.MustBegin()

	var bookmarkId int
	rows, err := db.NamedQuery(insertBookmarkQuery, bookmark)
	if err != nil {
		log.Println(err)
	}

	if rows.Next() {
		rows.Scan(&bookmarkId)
	}

	r.insertBookmarkTags(bookmark.Tags, bookmarkId)

	tx.Commit()
}

