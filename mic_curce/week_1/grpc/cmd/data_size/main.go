package main

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit"
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/proto"

	desc "github.com/vivaldi7/golang_code/mic_curce/grpc/pkg/note_v1"
)

func main() {
	session := &desc.NoteInfo{
		Title:    gofakeit.BeerName(),
		Content:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}

	dataJson, _ := json.Marshal(session)
	fmt.Printf("\n\n dataJson len %d byte \n%v\n", len(dataJson), dataJson)

	dataPb, _ := proto.Marshal(session)
	fmt.Printf("dataPb len %d byte \n%v\n", len(dataPb), dataPb)

}
