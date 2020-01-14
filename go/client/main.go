package main

import (
	"github.com/csandiego/poc-grpc/go/data"
	grpcService "github.com/csandiego/poc-grpc/go/grpc"
	"github.com/csandiego/poc-grpc/go/protobuf"
	"github.com/csandiego/poc-grpc/go/service"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Insufficient arguments.")
	}
	conn, err := grpc.Dial("localhost:80", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := protobuf.NewBookServiceClient(conn)
	var service service.BookService = grpcService.NewBookService(client)
	switch os.Args[1] {
	case "list":
		books, err := service.List()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(books)
	case "create":
		if len(os.Args) != 4 {
			log.Fatal("Wrong number of arguments for create.")
		}
		if err := service.Create(data.Book{Title: os.Args[2], Author: os.Args[3]}); err != nil {
			log.Fatal(err)
		}
	case "get":
		if len(os.Args) != 3 {
			log.Fatal("Wrong number of argument for get.")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		book, err := service.Get(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(book)
	case "update":
		if len(os.Args) != 5 {
			log.Fatal("Wrong number of arguments for update.")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		if err := service.Update(id, data.Book{Title: os.Args[3], Author: os.Args[4]}); err != nil {
			log.Fatal(err)
		}
	case "delete":
		if len(os.Args) != 3 {
			log.Fatal("Wrong number of arguments for delete.")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		if err := service.Delete(id); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unsupported operation.")
	}
}
