package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "pencethren/go-messageboard/data/postgres/model"
	"pencethren/go-messageboard/entity"
	"pencethren/go-messageboard/repository"
)

type postgresBoardRepository struct {
	*sql.DB
}

func NewPostgresBoardRepository(db *sql.DB) repository.IBoardRepository {
	return &postgresBoardRepository{DB: db}
}

func (r *postgresBoardRepository) Add(board *entity.Board) (uuid.UUID, error) {
	newBoard := models.Board{
		Name:        board.Name,
		Description: null.StringFrom(board.Description),
		Status:      models.BoardStatus(board.State.String()),
	}
	err := newBoard.Insert(context.TODO(), r.DB, boil.Infer())
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.MustParse(newBoard.ID), nil
}

func (r *postgresBoardRepository) List(pageSize int, filter *entity.BoardSearch) ([]*entity.BoardSummary, error) {
	clauses := []qm.QueryMod{
		models.BoardWhere.Status.EQ(models.BoardStatus(filter.State.String())),
		qm.Limit(pageSize + 1),
		qm.OrderBy(models.BoardColumns.Name),
	}

	if filter.SearchTerm != "" {
		clauses = append(clauses, qm.Where("name ilike ?", fmt.Sprintf("%%%s%%", filter.SearchTerm)))
	}

	if filter.Bookmark != "" {
		clauses = append(clauses, models.BoardWhere.Name.GTE(filter.Bookmark))
	}

	boardModels, err := models.Boards(
		clauses...,
	).All(context.TODO(), r.DB)

	if err != nil {
		return nil, err
	}

	boards := []*entity.BoardSummary{}
	for _, bm := range boardModels {
		state, err := entity.ParseBoardState(bm.Status.String())
		if err != nil {
			return nil, err
		}

		boards = append(boards, &entity.BoardSummary{
			Id:    uuid.MustParse(bm.ID),
			Name:  bm.Name,
			State: state,
		})
	}

	return boards, nil
}

func (r *postgresBoardRepository) ExistsWithName(name string) (bool, error) {
	count, err := models.Boards(
		models.BoardWhere.Name.EQ(name),
	).Count(context.TODO(), r.DB)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
