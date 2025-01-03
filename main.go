package main

import (
	"blog-fiber/validaciones"
	"encoding/json"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

type StructValidator struct {
	validate *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	engine := html.New("./views", ".html")
	validate := validator.New()
	validaciones.RegistrarValiaciones(validate)

	app := fiber.New(fiber.Config{
		StructValidator: &StructValidator{validate: validate},
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		Views:           engine,
	})

	// Middleware, esto cargar todas lass paginas estaticas
	// app.Get("/*", static.New("./public"))

	// Ruta para renderizar la plantilla
	app.Get("/prueba", func(c fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title":       "Go Fiber Template Example 1",
			"Description": "An example template 3",
			"Greeting":    "Hello, World! 2",
		})
	})

	app.Use("/", static.New("./public/index.html"))
	app.Use("/about", static.New("./public/about.html"))
	app.Use("/post", static.New("./public/post.html"))
	app.Use("/contact", static.New("./public/contact.html"))

	// HTTP
	log.Fatal(app.Listen(":3000"))

	// para agregar TLS, HTTPS
	// log.Fatal(app.Listen(":3001", fiber.ListenConfig{
	// 	CertFile:    "certs/server.crt",
	// 	CertKeyFile: "certs/server.key",
	// }))
}
