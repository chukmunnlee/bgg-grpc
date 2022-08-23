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

	limit := req.Limit
	offset := req.Offset
	query := req.Query

	log.Printf("FindBoardgameByName: query=%s, offset=%d, limit=%d", query, offset, limit)

	rows, err := svc.bggDB.FindBoardgameByName(ctx, query, limit, offset)
	if nil != err {
		return nil, fmt.Errorf("Error: FindBoardgameByName: %v", err)
	}

	games := make([]*Game, 0, limit)
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
	} else {
		total = int32(len(games))
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

	log.Printf("InsertNewBoardGame: game=%v", req.Game)

	newGame := data.Game{
		Name:       req.Game.Name,
		Year:       req.Game.Year,
		Ranking:    req.Game.Ranking,
		UsersRated: req.Game.UsersRated,
		Url:        req.Game.Url,
		Image:      req.Game.Image,
	}

	newGameId, err := svc.bggDB.InsertNewBoardGame(ctx, newGame)
	if nil != err {
		return nil, fmt.Errorf("Cannot insert new boardgame: %v", err)
	}

	resp := InsertNewBoardgameResponse{GameId: *newGameId}

	return &resp, nil
}

func (svc *BGGService) FindCommentsByGameId(req *FindCommentsByGameIdRequest,
	stream BGGService_FindCommentsByGameIdServer) error {

	log.Printf("FindCommentsByGameId: %d", req.GameId)

	ctx := stream.Context()

	count, err := svc.bggDB.CountCommentsByGameId(ctx, req.GameId)
	if nil != err {
		return err
	}
	iter, err := svc.bggDB.FindCommentsByGameId(ctx, req.GameId)
	if nil != err {
		return err
	}
	defer iter.Close()

	var i = int32(0)
	for iter.HasNext() {
		c, err := iter.Get()
		if nil != err {
			(*count)--
			continue
		}
		i++
		comment := polulateComment(*c)
		resp := FindCommentsByGameIdResponse{Comment: &comment, Cursor: i, Total: *count}
		stream.Send(&resp)
	}

	return nil
}

func (svc *BGGService) InsertNewComment(ctx context.Context, req *InsertNewCommentRequest) (*InsertNewCommentResponse, error) {

	log.Printf("InsertNewComment: comment=%v", req.Comment)

	newComment := data.Comment{
		User:   req.Comment.User,
		Rating: req.Comment.Rating,
		Text:   req.Comment.Text,
		GameId: req.Comment.GameId,
	}

	commentId, err := svc.bggDB.InsertNewComment(ctx, newComment)
	if nil != err {
		return nil, fmt.Errorf("Cannot insert new comment: %v", err)
	}

	resp := InsertNewCommentResponse{CommentId: *commentId}
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
func polulateComment(g data.Comment) Comment {
	return Comment{
		Id:     g.Id,
		User:   g.User,
		Rating: g.Rating,
		Text:   g.Text,
		GameId: g.GameId,
	}
}
