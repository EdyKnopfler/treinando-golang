package main

import (
	"context"
	"log"
	"time"

	pbA "com.derso/testify/grpc/proto/assincrona"
	pbS "com.derso/testify/grpc/proto/sincrona"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	sincronaClient   pbS.SincronaClient
	assincronaClient pbA.AssincronaClient
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("não foi possível conectar: %v", err)
	}
	sincronaClient = pbS.NewSincronaClient(conn)
	assincronaClient = pbA.NewAssincronaClient(conn)

	go tarefaPeriodica()

	r := gin.Default()
	r.POST("/solicita", func(c *gin.Context) {
		var req struct {
			Conteudo string `json:"conteudo"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"erro": "JSON inválido"})
			return
		}
		resp, err := sincronaClient.Solicitar(context.Background(), &pbS.Solicitacao{
			Conteudo: req.Conteudo,
		})
		if err != nil {
			c.JSON(500, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(200, gin.H{"resposta": resp.Conteudo})
	})

	r.Run(":8080")
}

func tarefaPeriodica() {
	for {
		time.Sleep(30 * time.Second) // intervalo desejado
		_, err := assincronaClient.Notificar(context.Background(), &pbA.Notificacao{
			Conteudo: "Executar limpeza",
		})
		if err != nil {
			log.Println("Erro ao enviar tarefa assíncrona:", err)
		} else {
			log.Println("Tarefa assíncrona enviada com sucesso")
		}
	}
}
