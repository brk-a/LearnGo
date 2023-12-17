# notes

* api is served at port 5000 on localhost
* no db; books are stored in mem
* to see books in collection

    ```bash
        curl localhost:5000/books
    ```

* to add the book in [body.json][def] to collection

    ```bash
        curl localhost:5000/books --include --header "Content-Type: application/json" -d @body.json
    ```

* if the `POST` request does not work, use this

    ```bash
        curl localhost:5000/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
    ```

* to borrow a book whose id is 2

    ```bash
        curl localhost:5000/borrow?id=2 --request "PATCH"
    ```

* to return a book whose id is 1

    ```bash
        curl localhost:5000/return?id=1 --request "PATCH"
    ```
[def]: ./body.json