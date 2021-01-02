-------------------------- REDIS MONITOR -------------------------

This package contains the backend of Redis Monitor written in GoLang
Redis Monitor is a web app that publishes Redis real-time statistics
via web socket to the client, which would be consumed by the web client.


1. Basic Structure:

    [ REDIS DATABASE ---------> SERVER -------(WEBSOCKET)------> CLIENT ]

2. Prerequisites:
    a) GoLang (go)
    b) Redis (redis-cli)

3. To run this Redis_Monitor GoLang package, developer first has to install
    the following dependencies:
    http router : "go get github.com/gorilla/mux"
    Redis       : "go get github.com/go-redis/redis"
    WebSocket   : "go get github.com/gorilla/websocket"

4. To run this application:
    a) Pull this package from git - "git pull Redis_Monitor"
    b) Start Redis service (on mac - "brew services start redis")
    c) go run main.go (on terminal) - this will start the server
    d) Open browser and point it to "http://localhost:8080"
    e) From inspect element select console in the browser
    d) Run these commands in the browser console:
        i) This command will connect to websocket connection of the server:
            var ws = new WebSocket("ws://localhost:8080/ws");
        ii) This command will print messages on browser console passed by server:
            ws.addEventListener("message", function(e) {console.log(e.data);})
    e) To stop the connection press CTRL+C on terminal
    f) And then finally stop the redis-cli kernel (on mac - "brew services stop redis")


                                Thank You.
