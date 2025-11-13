package domain

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	TeamID   int    `db:"team_id"`
	IsActive bool   `db:"is_active"`
}

type Team struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type TeamWithMembers struct {
	ID      int
	Name    string
	Members []User
}
