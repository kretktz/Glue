# Glue

Back end to the Glue project, an app for individuals who wish to use co-working spaces.

## Requirements

1. Go ver 1.15.2\
\
[Golang Website](https://golang.org/)
2. Firebase Tools:
    ```bash
    curl -sL https://firebase.tools | bash
    ```
    
    Login to your Firebase account:
    ```bash
    firebase login
    ```
    
    [Firebase Tools Reference](https://firebase.google.com/docs/cli)
2. Firebase Admin SDK for Go\
\
    [Firebase Admin Pre Requisites](https://firebase.google.com/docs/admin/setup)
    ```bash
    go get firebase.google.com/go
    ```
3. Dependencies:
    * Mux Router
    ```bash
    go get github.com/gorilla/mux
    ```
    * Chi Router
     ```bash
    go get github.com/go-chi/chi
    ```
    _Either one of the routers can be used, make sure to select the desired router in **var** section in the server.go file_\
        ```
        httpRouter router.Router = router.NewMuxRouter()
        ```
