You have been asked to create a web server where users can track how many games players have won. 

GET /players/{name} - should return a number indicating the total number of wins

POST /players/{name} - should record a win for that name, incrementing for every subsequent POST 

To create a web server in Go you will typically call ListenAndServe. 

    func ListenAndServe(addr string, handler Handler) error

This will start a web server listening on a port, creating a goroutine for every request & running it against a Handler.

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

