package main

import "github.com/gin-gonic/gin"

func main() {

	// Creo router
	router := gin.Default()

	/*
	   EJERCICIO 1
	   Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al
	   pegarle responda un texto que diga “pong”
	   El endpoint deberá ser de método GET
	   La respuesta de “pong” deberá ser enviada como texto, NO como JSON
	*/

	// Capturo la solicitud GET /ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	/*
	   EJERCICIO 2
	   Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al
	   pegarle deberá responder en texto “Hola + nombre + apellido”

	   El endpoint deberá ser de método POST
	   Se deberá usar el package JSON para resolver el ejercicio
	   La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
	   La estructura deberá ser como esta:
	   {
	   		“nombre”: “Andrea”,
	   		“apellido”: “Rivas”
	   }
	*/

	// Creo la estructura
	type persona struct {
		Nombre   string
		Apellido string
	}

	// Capturo la solicitud GET /ping
	router.POST("/saludo", func(ctx *gin.Context) {
		body := persona{}

		if err := ctx.ShouldBind(&body); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(400)
		}

		ctx.String(200, "Hola "+body.Nombre+" "+body.Apellido)

	})

	router.Run(":8080")

}
