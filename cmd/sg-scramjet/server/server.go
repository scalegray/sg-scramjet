package server


  import (
	log "code.google.com/p/log4go"
	//"encoding/json"
	"github.com/scalegray/sg-golib/amqp"

	"github.com/scalegray/sg-scramjet/api/http"

	"github.com/scalegray/sg-scramjet/cmd/sg-scramjet/server/queue"

	"os"
	"fmt"

)

type Server struct {
 httpConn        *http.HttpServer
 QueueServers    []*queue.QueueServer
 stopped          bool
}



func NewServer() (*Server, error) {
  log.Info("Starting New server")
  httpConn := http.NewHttpServer()

  return &Server{
    httpConn: httpConn,
  }, nil
}

func (self *Server) ListenAndServe() error {
  log.Info("Starting the admin interface on the port")
  var queues [1]string
  queues[0] = "morpheyesh-laptop"
  self.queueChecker()

  for i := range queues {
    singleQ := queues[i]
    Qserver := queue.NewServer(singleQ)
    fmt.Println(Qserver)
    go Qserver.ListenAndServe()
  }
  self.httpConn.ListenAndServe()

  return nil
 }


func (self *Server) queueChecker() {

  log.Info("Dialing RabbitMQ")
  rmq, err := amqp.Factory()

  log.Info("rmq-->  %v", rmq)


  if err != nil {
    log.Error("Something is terribly wrong..we failed to get the queue instance: %s", err)

  }
  log.Info("Entered here for dial")
  connection, cerr := rmq.Dial()
  log.Info("connection-->  %v", connection)

   if cerr != nil {
     fmt.Fprintf(os.Stderr, "Error: %v \nPlease start RabbitMQ service\n", cerr)
     os.Exit(1)
   }
   log.Info("Rabbitmq connected")

}

func (self *Server) Stop() {

  if self.stopped {
    return
  }
  log.Info("Stopping servers....")
  self.stopped = true

  log.Info("Stopping API Server")
  log.Info("API Server stopped")

}
