package entities

const (
	UserRole Role = iota + 1
	AuthorRole
	AdminRole
)

type Role int

func (r Role) String() string {
	return [...]string{"User", "Author", "Admin"}[r-1]
}
