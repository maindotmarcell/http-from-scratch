package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/maindotmarcell/http-from-scratch/internal/constants"
	"github.com/maindotmarcell/http-from-scratch/internal/http"
	"github.com/maindotmarcell/http-from-scratch/internal/router"
)

type HTTPServer struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	Router     router.Router
}

// Returns an initialized Server struct, with it's listener address set as listenAddr
func NewHTTPServer(listenAddr string) *HTTPServer {
	return &HTTPServer{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		Router:     *router.New(),
	}
}

// Starts the server and keeps receiving connections until the quitch is closed
func (s *HTTPServer) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to bind to %s", s.listenAddr)
	}

	s.ln = ln

	fmt.Printf("Server is now listening to connections on %s\n", s.listenAddr)

	go s.acceptLoop()

	go s.handleSignals()

	<-s.quitch

	return nil
}

// Creates a channel and handels interrupt signals to facilitate graceful shutdown
func (s *HTTPServer) handleSignals() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigCh
	fmt.Printf("\nreceived signal %v, shutting down...\n", sig)
	close(s.quitch)
}

// Accepts TCP connections concurrently and calls handleConn for each of them
func (s *HTTPServer) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Accepted connection. Reading from %s\n", conn.RemoteAddr().String())
		go s.handleConn(conn)
	}
}

// Reads from the connection, parses the http request, calls the appropriate handler and writes the response to the connection.
// If the method/path combination is not found it will write a 404 NOT FOUND response to the connection.
func (s *HTTPServer) handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error reading from connection: ", err.Error())
		}
		return
	}

	fmt.Printf("Message received from %s: %s\n", conn.RemoteAddr().String(), string(buf[:n]))

	req := http.ParseRequest(buf[:n])
	handler := s.Router.Route(req)
	res := ""
	if handler != nil {
		res = handler(req)
	} else {
		res = http.FormatResponse(http.Response{Status: constants.StatusNotFound})
	}

	conn.Write([]byte(res))
	fmt.Println("Reply has been sent to the client.")
}
