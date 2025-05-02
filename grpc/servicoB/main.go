package main

import (
	"context"
	"log"
	"net"

	pbA "com.derso/testify/grpc/proto/assincrona"
	pbS "com.derso/testify/grpc/proto/sincrona"
	"google.golang.org/grpc"
)

type server struct {
	pbS.UnimplementedSincronaServer
	pbA.UnimplementedAssincronaServer
}

func (s *server) Solicitar(ctx context.Context, req *pbS.Solicitacao) (*pbS.Resposta, error) {
	log.Println("Solicitação recebida:", req.Conteudo)
	return &pbS.Resposta{Conteudo: "Resposta para: " + req.Conteudo}, nil
}

func (s *server) Notificar(ctx context.Context, req *pbA.Notificacao) (*pbA.Resultado, error) {
	log.Println("Tarefa de manutenção recebida:", req.Conteudo)
	// simulação de execução de tarefa
	return &pbA.Resultado{Status: "Executado"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("erro ao escutar: %v", err)
	}
	grpcServer := grpc.NewServer()
	pbS.RegisterSincronaServer(grpcServer, &server{})
	pbA.RegisterAssincronaServer(grpcServer, &server{})
	log.Println("Serviço B escutando na porta 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("falha ao servir: %v", err)
	}
}
