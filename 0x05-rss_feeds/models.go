package main

type User struct {
	Id uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
}
func databaseUserToUser(dbUser database.User) User {
	return User {
		Id: dbUser.Id
		CreatedAt: dbUser.CreatedAt
		UpdatedAt: dbUser.UpdatedAt
		Name: dbUser.Name
	}
}