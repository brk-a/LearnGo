package model

type Todo struct {
	Id uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool `json:"done"`
}

func CreateTodo(todo string) err {
	statement := `insert into todos(todo, done) values($1, $2);`
	_, err := db.Query(statement, todo, false)

	return err
}

func GetAllTodos()([]Todo, err){
	todos := []Todo{}
	statement := `select * from todos;`

	rows, err := db.Query(statement)

	if err!=nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next(){
		var title string
		var done bool
		var id uint64

		err := db.Scan(&id, &title, &done)
		todo := Todo{
			Id: id,
			Todo: todo,
			Done: done,
		}

		todos = append(todos, todo)
	}
}