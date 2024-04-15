package database

type Migrator interface {
	Up() error
	Down(int) error
	Force(int) error
}
