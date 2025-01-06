# simple web server

### structure

    ```mermaid
        ---
        title: server overview
        ---
            flowchart LR
            A[server]--->B[/]
            A--->C[/hello]
            A--->D[/form]
            B--->E[index.html]
            C--->F[hello func]
            D--->F[form func]
            F--->G[form.html]
    ```