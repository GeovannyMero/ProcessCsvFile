package main

import (
	"context"
	"fmt"
	"log"

	"example.com/csv/src/core"
	"example.com/csv/src/database"
)

func main() {
	people := core.ReadCsv()
	println("Cantidad de registros: ", len(people))

	var interfacePeople []interface{}

	for _, value := range people {
		interfacePeople = append(interfacePeople, value)
	}

	client, ctx, err := database.Connect()

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	// Acceder a una base de datos y colección (ejemplo)
	db := client.Database("testdb")
	collection := db.Collection("test")

	fmt.Println("📂 Base de datos seleccionada:", db.Name())
	fmt.Println("📌 Colección seleccionada:", collection.Name())

	errPing := client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("❌ Error de conexión a MongoDB:", errPing)
	}

	resultInsert, err := collection.InsertMany(context.Background(), interfacePeople)

	if err != nil {
		log.Fatal("❌ Error al insertar: ", err)
	}

	fmt.Println("✅  Insertados:", len(resultInsert.InsertedIDs))

	defer client.Disconnect(ctx) // Cerrar conexión al final
}
