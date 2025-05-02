# Treinando Golang

Referência rápida para a realização de tarefas do dia a dia em Go.

* Parsing de JSON arbitrário e manipulação dos mapas resultantes
* Suíte de testes com _testify_
* Migrações de bases de dados com _goose_
* Genéricos
* Goroutines
* Unicode
* gORM
* gRPC

```bash
# Gerar os arquivos Go dos Protobufs
protoc --proto_path=. --go_out=. --go-grpc_out=. grpc/proto/sincrona.proto
protoc --proto_path=. --go_out=. --go-grpc_out=. grpc/proto/asssincrona.proto

# Rodar cliente e servidor
go run grpc/servicoA/main.go
go run grpc/servicoB/main.go

# Enviar requisição (envia para A, deve chegar no terminal do B)
curl -X POST -H 'Content-Type: application/json' -d '{"conteudo": "blelelê"}' localhost:8080/solicita
```

## Prompt no Chat-GPT para elaborar o esqueleto do gRPC

Resultado no repositório após devidos ajustes.

```
Preciso elaborar um conjunto com dois microsserviços golang trocando mensagens Protobuf via gRPC. O serviço A possui uma interface web onde recebe uma requisição de outra aplicação/front-end e faz uma requisição síncrona para o serviço B, que repassa algum dado para ser devolvido para o cliente.

O serviço A também de tempos em tempos envia mensagens assíncronas para o serviço B realizar alguma tarefa de manutenção.

O conteúdo dessas mensagens pode ser abstraído, por ex.: string conteudo = 1; e nada mais. Os nomes podem ser Sincrona e Assincrona.
```