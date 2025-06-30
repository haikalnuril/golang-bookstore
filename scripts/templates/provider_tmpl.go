package templates

var ProviderTmpl = `package {{.ModuleName}}

import (
	"{{.Package}}/internal/modules/{{.ModuleName}}/entity"
	"{{.Package}}/internal/modules/{{.ModuleName}}/router"
	"{{.Package}}/internal/modules/{{.ModuleName}}/usecase"
	"{{.Package}}/internal/modules/{{.ModuleName}}/repository/postgres"
	"{{.Package}}/internal/modules/{{.ModuleName}}/controller"
	"{{.Package}}/internal/app/config""

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	)

type {{.Name}}Module struct {
	App *fiber.App
	DB  *gorm.DB
}

func (m *{{.Name}}Module) Register() {
	{{.ModuleName}}Repo := postgres.New{{.Name}}Postgres(m.DB)
	{{.ModuleName}}Usecase := usecase.New{{.Name}}Usecase({{.ModuleName}}Repo)
	validator := config.NewValidator()
	{{.ModuleName}}Controller := controller.New{{.Name}}Controller({{.ModuleName}}Usecase, validator)
	api := m.App.Group("/api/v1")
	router.New{{.Name}}Router(api, {{.ModuleName}}Controller)
}

func (m *{{.Name}}Module) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.{{.Name}}{})
	`
