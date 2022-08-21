package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type BggDB struct {
	dbFile string
	db     *sql.DB
}

/*
CREATE TABLE IF NOT EXISTS "game" (
	"gid" INTEGER NOT NULL  ,
	"name" VARCHAR(256) NOT NULL  ,
	"year" INTEGER NOT NULL  ,
	"ranking" INTEGER NOT NULL  ,
	"users_rated" INTEGER NOT NULL  ,
	"url" VARCHAR(512) NULL  ,
	"image" VARCHAR(512) NULL  ,
	PRIMARY KEY ("gid")
);

CREATE TABLE IF NOT EXISTS "comment" (
	"c_id" CHARACTER(8) NOT NULL  ,
	"user" VARCHAR(64) NOT NULL  ,
	"rating" INTEGER NULL  ,
	"c_text" TEXT NULL  ,
	"gid" INTEGER NOT NULL  ,
	PRIMARY KEY ("c_id")
);
*/

type Game struct {
	GameId     int32  `db:"gid"`
	Name       string `db:"name"`
	Year       int32  `db:"year"`
	Ranking    int32  `db:"ranking"`
	UsersRated int32  `db:"users"`
	Url        string `db:"url"`
	Image      string `db:"image"`
}

func (rec *Game) Populate(row *sql.Rows) error {
	return row.Scan(&rec.GameId, &rec.Name, &rec.Year, &rec.Ranking, &rec.UsersRated, &rec.Url, &rec.Image)
}

type Comment struct {
	Id     string
	User   string
	Rating int32
	Text   string
	GameId int32
}

func (rec *Comment) Polulate(row *sql.Rows) error {
	return row.Scan(&rec.Id, &rec.User, &rec.Rating, &rec.Text, &rec.GameId)
}

const (
	FIND_GAME_BY_NAME  = "select * from game where name like ? limit ? offset ?"
	COUNT_GAME_BY_NAME = "select count(*) as game_count from game where name like ?"
)

func New(dbFile string) BggDB {
	return BggDB{dbFile: dbFile}
}

func (bggdb *BggDB) FindBoardgameByName(ctx context.Context, query string, limit int32, offset int32) (*[]Game, error) {

	rows, err := bggdb.db.QueryContext(ctx, FIND_GAME_BY_NAME, fmt.Sprintf("%%%s%%", query), limit, offset)
	if nil != err {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	results := make([]Game, 0, limit)
	for rows.Next() {
		g := Game{}
		if err := g.Populate(rows); nil != err {
			log.Printf("Game.Populate() error: %v", err)
			continue
		}
		results = append(results, g)
	}

	return &results, nil
}

func (bggdb *BggDB) CountBoardgamesByName(ctx context.Context, query string) (*int32, error) {
	rows, err := bggdb.db.QueryContext(ctx, COUNT_GAME_BY_NAME, fmt.Sprintf("%%%s%%", query))
	if nil != err {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var count int32 = 0
	if rows.Next() {
		if err := rows.Scan(&count); nil != err {
			return nil, fmt.Errorf("rows.Scan: %v", err)
		}
	}
	return &count, nil
}

func (bggdb *BggDB) Open() error {
	if db, err := sql.Open("sqlite3", bggdb.dbFile); nil != err {
		return err
	} else {
		bggdb.db = db
	}
	return nil
}

func (bggdb *BggDB) Close() error {
	return bggdb.db.Close()
}
