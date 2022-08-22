package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/chukmunnlee/bgg-grpc/data"
)

type BGGService struct {
	UnimplementedBGGServiceServer

	bggDB data.BggDB
}

func New(bggDB data.BggDB) BGGService {
	return BGGService{bggDB: bggDB}
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
