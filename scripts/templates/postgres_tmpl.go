package templates

var PostgresTmpl = `package postgres

import (
	"{{.Package}}/internal/modules/{{.Module}}/entity"
	"{{.Package}}/internal/modules/{{.Module}}/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{.ModuleName}}Postgres struct {
	db *gorm.DB
}

func New{{.Name}}Postgres(db *gorm.DB) *{{.ModuleName}}Postgres {
	return &{{.ModuleName}}Postgres{db}
}

func (r *{{.ModuleName}}Postgres) Create({{.ModuleName}} *entity.{{.Name}}) error {
	return r.db.Create({{.ModuleName}}).Error
}

func (r *{{.ModuleName}}Postgres) Find() ([]entity.{{.Name}}, error) {
	{{.ModuleName}}s := []entity.{{.Name}}{}
	err := r.db.Find(&{{.ModuleName}}s).Error
	if err != nil {
		return nil, err
	}
	return {{.ModuleName}}s, nil
}

func (r *{{.ModuleName}}Postgres) FindById(id uuid.UUID) (entity.{{.Name}}, error) {
	{{.ModuleName}} := &entity.{{.Name}}{}
	result := r.db.Where("id = ?", id).First({{.ModuleName}})

	if result.Error != nil {
		return nil, result.Error
	}

	return {{.ModuleName}}, nil
}

func (r *{{.ModuleName}}Postgres) Update(id string, {{.ModuleName}} *entity.{{.Name}}) (*entity.{{.Name}}, error) {
	err := r.db.Model(&entity.{{.Name}}{}).Where("id = ?", id).Updates({{.ModuleName}}).Error
	if err != nil {
		return nil, err
	}
	return {{.ModuleName}}, nil
}

func (r *{{.ModuleName}}Postgres) Delete({{.ModuleName}} *entity.{{.Name}}) error {
	err := r.db.Delete(&entity.{{.Name}}{}, {{.ModuleName}}.Id).Error
	if err != nil {
		return err
	}
	return nil
}
`
