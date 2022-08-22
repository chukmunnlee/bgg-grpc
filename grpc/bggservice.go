package grpc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/chukmunnlee/bgg-grpc/data"
)

type BGGService struct {
	UnimplementedBGGServiceServer

	bggDB data.BggDB
}

func New(bggDB data.BggDB) BGGService {
	svc := BGGService{bggDB: bggDB}
	log.Printf("Start of gameId index: %d", bggDB.StartOfNewGameId())
	return svc
}

func (svc *BGGService) FindBoardgameByName(ctx context.Context, req *FindBoardgameByNameRequest) (*FindBoardgameByNameResponse, error) {

	limit := req.GetLimit()
	offset := req.GetOffset()
	query := req.GetQuery()

	log.Printf("FindBoardgameByName: query=%s, offset=%d, limit=%d", query, offset, limit)

	rows, err := svc.bggDB.FindBoardgameByName(ctx, query, limit, offset)
	if nil != err {
		return nil, fmt.Errorf("Error: FindBoardgameByName: %v", err)
	}

	games := make([]*Game, limit)
	for _, v := range *rows {
		g := populateGame(v)
		games = append(games, &g)
	}

	total := limit
	if int32(len(games)) >= total {
		if t, err := svc.bggDB.CountBoardgamesByName(ctx, query); nil != err {
			return nil, fmt.Errorf("Error: CountBoardgamesByName: %v", err)
		} else {
			total = *t
		}
	}

	resp := FindBoardgameByNameResponse{Games: games, Limit: limit, Offset: offset, Total: total}

	return &resp, nil
}

func (svc *BGGService) FindBoardgameById(ctx context.Context, req *FindBoardgameByIdRequest) (*FindBoardgameByIdResponse, error) {

	gameId := req.GetGameId()

	log.Printf("FindBoardgameById: gameId=%d", gameId)

	game, err := svc.bggDB.FindBoardgameById(ctx, gameId)
	if nil != err {
		if errors.Is(err, sql.ErrNoRows) {
			resp := FindBoardgameByIdResponse{
				Found:   false,
				Message: fmt.Sprintf("Cannot find game %d", gameId),
			}
			return &resp, nil
		}
		return nil, fmt.Errorf("Error: FindBoardgameById: %v", err)
	}

	result := populateGame(*game)
	resp := FindBoardgameByIdResponse{Game: &result, Found: true}

	return &resp, nil
}

func (svc *BGGService) InsertNewBoardGame(ctx context.Context, req *InsertNewBoardgameRequest) (*InsertNewBoardgameResponse, error) {

	log.Printf("InsertNewBoardGame: game=%v", req.GetGame())

	newGame := data.Game{
		Name:       req.GetGame().Name,
		Year:       req.GetGame().Year,
		Ranking:    req.GetGame().Ranking,
		UsersRated: req.GetGame().UsersRated,
		Url:        req.GetGame().Url,
		Image:      req.GetGame().Image,
	}

	newGameId, err := svc.bggDB.InsertNewBoardGame(ctx, newGame)
	if nil != err {
		return nil, fmt.Errorf("Cannot insert new boardgame: %v", err)
	}

	resp := InsertNewBoardgameResponse{GameId: *newGameId}

	return &resp, nil
}

func populateGame(g data.Game) Game {
	return Game{
		GameId:     g.GameId,
		Name:       g.Name,
		Year:       g.Year,
		Ranking:    g.Ranking,
		UsersRated: g.UsersRated,
		Url:        g.Url,
		Image:      g.Image,
	}
}
