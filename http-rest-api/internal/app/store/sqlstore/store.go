package sqlstore

import (
	"database/sql"

	"http-rest-api/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// package sqlstore

// import (
// 	"database/sql"

// 	_ "github.com/lib/pq" // ..
// )

// type Store struct {
// 	db             *sql.DB
// 	userRepository *UserRepository
// }

// func New(db *sql.DB) *Store {
// 	return &Store{
// 		db: db,
// 	}
// }

// //
// func (s *Store) User() *UserRepository {
// 	if s.userRepository != nil {
// 		return s.userRepository
// 	}
// 	s.userRepository = &UserRepository{
// 		Store: s,
// 	}
// 	return s.userRepository
// }
