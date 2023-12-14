# notes
what to do at various steps

### `Dockerfile` and `initdb` (plus contents) created
run

```bash
podman build -t psql
```

### `_db_data` created
run 

```bash
podman run --name postgresql -p 5432:5432 -v ./_db_data:/var/lib/postgresql/data psql:latest
```

### `db.go` and `main.go` created
run the following

```go
    go mod init htmx
    go mod tidy
```