// Code generated by glean from glean.go:9. DO NOT EDIT.
package rounds

import (
	"context"
	"database/sql"

	"github.com/corverroos/unsure/engine/internal"
)

const cols = " `id`, `match_id`, `index`, `team`, `status`, `state`, `error`, `created_at`, `updated_at` "
const selectPrefix = "select " + cols + " from engine_rounds where "

func Lookup(ctx context.Context, dbc dbc, id int64) (*internal.Round, error) {
	return lookupWhere(ctx, dbc, "id=?", id)
}

// lookupWhere queries the engine_rounds table with the provided where clause, then scans
// and returns a single row.
func lookupWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) (*internal.Round, error) {
	return scan(dbc.QueryRowContext(ctx, selectPrefix+where, args...))
}

// listWhere queries the engine_rounds table with the provided where clause, then scans
// and returns all the rows.
func listWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) ([]internal.Round, error) {

	rows, err := dbc.QueryContext(ctx, selectPrefix+where, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []internal.Round
	for rows.Next() {
		r, err := scan(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, *r)
	}

	return res, rows.Err()
}

func scan(row row) (*internal.Round, error) {
	var g glean

	err := row.Scan(&g.ID, &g.MatchID, &g.Index, &g.Team, &g.Status, &g.State, &g.Error, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &internal.Round{
		ID:        g.ID,
		MatchID:   g.MatchID,
		Index:     g.Index,
		Team:      g.Team,
		Status:    g.Status,
		State:     g.State,
		Error:     g.Error.String,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

// dbc is a common interface for *sql.DB and *sql.Tx.
type dbc interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// row is a common interface for *sql.Rows and *sql.Row.
type row interface {
	Scan(dest ...interface{}) error
}
