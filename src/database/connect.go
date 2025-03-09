package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, context.Context, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")

	// Crear un contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Cerrar el contexto si hay un error o al finalizar
	defer cancel()

	// Configurar opciones del cliente
	clientOptions := options.Client().ApplyURI(MONGO_URI)

	// Conectar con MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	// Verificar la conexión con un ping
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("✅ Conexión exitosa a MongoDB")
	return client, ctx, nil
}
