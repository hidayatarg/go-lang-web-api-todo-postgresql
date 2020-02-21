package postgres

// create DB
import (
	_ "github.com/lib/pq"
	"github.com/go-pg/pg/v9"
)

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}