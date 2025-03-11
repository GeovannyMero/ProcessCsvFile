package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"example.com/csv/src/core"
	"example.com/csv/src/database"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 🔹 Definir la estructura con etiquetas BSON
//
//	type Person struct {
//		ID                        string `bson:"_id,omitempty"`
//		Ruc                       string `bson:"ruc"`
//		RazonSocial               string `bson:"razon_social"`
//		CodigoJuridiccion         string `bson:"codigo_juridiccion"`
//		EstadoContribuyente       string `bson:"estado_contribuyente"`
//		ClaseContribuyente        string `bson:"clase_contribuyente"`
//		FechaInicioActividades    string `bson:"fecha_inicio_actividades"`
//		FechaActualización        string `bson:"fecha_actualizacion"`
//		FechaSuspencionDefinitiva string `bson:"fecha_suspencion_definitiva"`
//		FechaReinicioActividades  string `bson:"fecha_reinicio_actividades"`
//		Obligado                  string `bson:"obligado"`
//		TipoContribuyente         string `bson:"tipo_contribuyente"`
//		NumeroEstablecimiento     int    `bson:"numero_establecimiento"`
//		NombreFantasiaComercial   string `bson:"nombre_fantasia_comercial"`
//		EstadoEstablecimiento     string `bson:"estado_establecimiento"`
//		DescripcionProvinciaEst   string `bson:"descripcion_provincia_est"`
//		DescripcionCantonEst      string `bson:"descripcion_canton_est"`
//		DescripcionParroquiaEst   string `bson:"descripcion_parroquia_est"`
//		CodigoCIIU                string `bson:"codigo_ciiu"`
//		ActividadEconomica        string `bson:"actividad_economica"`
//		AgenteRetencion           string `bson:"agente_retencion"`
//		Especial                  string `bson:"especial"`
//	}

func main() {
	people := core.ReadCsv()
	println("Cantidad de registros: ", len(people))

	var interfacePeople []interface{}

	for _, value := range people {
		interfacePeople = append(interfacePeople, value)
	}

	// ConnectToDb(people[50])
	//ConnectToDb(interfacePeople)

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

func ConnectToDb(person []interface{}) {
	err_env := godotenv.Load()

	if err_env != nil {
		log.Fatalf("Error loading .env file: %s", err_env)
	}

	// Get value from .env
	MONGO_URI := os.Getenv("MONGO_URI")
	fmt.Println(MONGO_URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the database.
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection.
	err_con := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err_con)
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}
	defer client.Disconnect(ctx)

	fmt.Println("Conexión exitosa a MongoDB")
	fmt.Println("✅ Conectado a MongoDB")

	collection := client.Database("testdb").Collection("test")

	if err != nil {
		log.Fatal(err)
	}

	//***//

	// var data Person
	// for _, person := range people {
	// 	if person.ruc == "0190301850001" {
	// 		// bsonData, err := bson.Marshal(person)
	// 		fmt.Println(person)

	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		// fmt.Println(bsonData)
	// 		fmt.Println(person.razon_social)
	// 		data = Person{
	// 			ruc:          "0190301850001",
	// 			razon_social: "FLORIPAMBA",
	// 		}

	// 		fmt.Printf("data: %v", data)
	// 		fmt.Println(data)
	// 		results, err_insert := collection.InsertOne(ctx, data)

	// 		if err_insert != nil {
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Println("📌 Documentos insertados con IDs:", results.InsertedID)

	// 	}

	// }

	// fmt.Println(collection)

	// // 🔹 3. Crear un objeto `User`
	// user := User{
	// 	Name:  "Carlos",
	// 	Age:   30,
	// 	Email: "carlos@example.com",
	// }

	// // 🔹 4. Insertar el documento en MongoDB
	// dat, err_ctv := bson.Marshal(data)

	// if err_ctv != nil {
	// 	log.Fatal(err_ctv)
	// }
	// fmt.Println(data.Ruc)
	// fmt.Println(string(dat))
	result_insert, err_insert := collection.InsertMany(ctx, person)
	if err_insert != nil {
		log.Fatal("❌  ERROR AL INSERTAR DATOS: ", err)
	}

	// // // 🔹 5. Mostrar el ID del documento insertado
	fmt.Println("📌 Documento insertado con ID:", len(result_insert.InsertedIDs))
}

// func ReadCsv() []models.Person {
// 	//abre el documento
// 	file, err := os.Open("C:\\Users\\gmero\\Documents\\GO\\04-csv\\resources\\SRI_RUC_Azuay(1).csv")

// 	if err != nil {
// 		fmt.Println("Error al abrir archivo: ", err)
// 		return nil
// 	}

// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.Comma = '|'

// 	//leer un lector
// 	records, err := reader.ReadAll()

// 	if err != nil {
// 		fmt.Println("ERROR al leer archivo CSV:", err)
// 		return nil
// 	}

// 	//crear una variable de una array de tipo person
// 	var people []models.Person

// 	for i, row := range records {
// 		if i == 0 {
// 			continue
// 		}

// 		numero_est, _ := strconv.ParseInt(row[11], 10, 64)

// 		person := models.Person{
// 			Ruc:                       row[0],
// 			RazonSocial:               row[1],
// 			CodigoJuridiccion:         row[2],
// 			EstadoContribuyente:       row[3],
// 			ClaseContribuyente:        row[4],
// 			FechaInicioActividades:    row[5],
// 			FechaActualización:        row[6],
// 			FechaSuspencionDefinitiva: row[7],
// 			FechaReinicioActividades:  row[8],
// 			Obligado:                  row[9],
// 			TipoContribuyente:         row[10],
// 			NumeroEstablecimiento:     int(numero_est),
// 			NombreFantasiaComercial:   row[12],
// 			EstadoEstablecimiento:     row[13],
// 			DescripcionProvinciaEst:   row[14],
// 			DescripcionCantonEst:      row[15],
// 			DescripcionParroquiaEst:   row[16],
// 			CodigoCIIU:                row[17],
// 			ActividadEconomica:        row[18],
// 			AgenteRetencion:           row[19],
// 			Especial:                  row[20],
// 		}

// 		people = append(people, person)
// 	}

// 	return people

// }

// func Connect() (*mongo.Client, context.Context, error) {
// 	err := godotenv.Load()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	MONGO_URI := os.Getenv("MONGO_URI")

// 	// Crear un contexto con timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

// 	// Cerrar el contexto si hay un error o al finalizar
// 	defer cancel()

// 	// Configurar opciones del cliente
// 	clientOptions := options.Client().ApplyURI(MONGO_URI)

// 	// Conectar con MongoDB
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// Verificar la conexión con un ping
// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	fmt.Println("✅ Conexión exitosa a MongoDB")
// 	return client, ctx, nil
// }
