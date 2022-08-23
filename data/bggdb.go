package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync/atomic"

	_ "github.com/mattn/go-sqlite3"
)

type BggDB struct {
	dbFile   string
	db       *sql.DB
	newIndex int32
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

func (rec *Comment) Populate(row *sql.Rows) error {
	return row.Scan(&rec.Id, &rec.User, &rec.Rating, &rec.Text, &rec.GameId)
}

const (
	FIND_GAME_BY_NAME         = "select * from game where name like ? limit ? offset ?"
	COUNT_GAME_BY_NAME        = "select count(*) as game_count from game where name like ?"
	FIND_GAME_BY_ID           = "select * from game where gid = ?"
	FIND_LARGEST_GAME_ID      = "select gid from game order by gid desc limit 1"
	INSERT_NEW_GAME           = "insert into game(gid, name, year, ranking, users_rated, url, image) values (?, ?, ?, ?, ?, ?, ?)"
	FIND_COMMENTS_BY_GAME_ID  = "select * from comment where gid = ?"
	COUNT_COMMENTS_BY_GAME_ID = "select count(*) from comment where gid = ?"
	INSERT_NEW_COMMENT        = "insert into comment(c_id, user, rating, c_text, gid) values (?, ?, ?, ?, ?)"
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

func (bggdb *BggDB) FindBoardgameById(ctx context.Context, gameId int32) (*Game, error) {

	rows, err := bggdb.db.QueryContext(ctx, FIND_GAME_BY_ID, gameId)
	if nil != err {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	result := Game{}
	if err := result.Populate(rows); nil != err {
		return nil, fmt.Errorf("read error: %v", err)
	}

	return &result, nil
}

func (bggdb *BggDB) FindLargestGameId(ctx context.Context) (*int32, error) {

	rows, err := bggdb.db.QueryContext(ctx, FIND_LARGEST_GAME_ID)
	if nil != err {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("no game records in the database?")
	}

	var lastIndex int32
	if err := rows.Scan(&lastIndex); nil != err {
		return nil, fmt.Errorf("cannot read last index from game table: %v", err)
	}

	return &lastIndex, nil
}

func (bggdb *BggDB) InsertNewBoardGame(ctx context.Context, game Game) (*int32, error) {

	newGameId := atomic.AddInt32(&bggdb.newIndex, 1)

	result, err := bggdb.db.ExecContext(ctx, INSERT_NEW_GAME, newGameId, game.Name,
		game.Year, game.Ranking, game.UsersRated, game.Url, game.Image)

	if nil != err {
		return nil, fmt.Errorf("insert error: %v", err)
	}

	rowCount, err := result.RowsAffected()
	if nil != err {
		return nil, fmt.Errorf("insert error: %v", err)
	}
	if rowCount < 1 {
		return nil, fmt.Errorf("inserted but row count is not 1: %d", rowCount)
	}

	return &newGameId, nil
}

func (bggdb *BggDB) FindCommentsByGameId(ctx context.Context, gameId int32) (*[]Comment, error) {

	rows, err := bggdb.db.QueryContext(ctx, FIND_COMMENTS_BY_GAME_ID, gameId)
	if nil != err {
		return nil, fmt.Errorf("query error: %v", err)
	}
	var count int32
	if err := rows.Scan(&count); nil != err {
		rows.Close()
		return nil, fmt.Errorf("read result error: %v", err)
	}
	rows.Close()

	rows, err = bggdb.db.QueryContext(ctx, FIND_COMMENTS_BY_GAME_ID, gameId)
	if nil != err {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	results := make([]Comment, count, count)
	for i := int32(0); i < count; i++ {
		if !rows.Next() {
			break
		}
		c := Comment{}
		if err := c.Populate(rows); nil != err {
			log.Printf("Comment.Populate() error: %v", err)
			continue
		}
		results[i] = c
	}

	return &results, nil
}

func (bggdb *BggDB) Open() error {
	if db, err := sql.Open("sqlite3", bggdb.dbFile); nil != err {
		return err
	} else {
		bggdb.db = db
		// Ignore error
		idx, _ := bggdb.FindLargestGameId(context.Background())
		bggdb.newIndex = *idx + 10
	}

	return nil
}

func (bggdb *BggDB) StartOfNewGameId() int32 {
	return bggdb.newIndex
}

func (bggdb *BggDB) Close() error {
	return bggdb.db.Close()
}
