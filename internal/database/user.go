package database

type UserStore struct {
	service
}


// does this works ?
func (s *UserStore) Create(name string) error {
	query := `insert into users (username) values $1`

	if _, err := s.db.Exec(query, name); err != nil {
		return err
	}
	return nil
}
