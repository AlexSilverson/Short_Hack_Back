package skillservice

import (
	"context"
	"errors"
	"fmt"

	entity "github.com/AlexSilverson/Short_Hack_Back/src/entity/skill"
	"github.com/jackc/pgx/v5/pgxpool"
)

type skillService struct {
	p *pgxpool.Pool
}

type SkillService interface {
	AddSkill(skill entity.Skill, ctx context.Context) error
	GetAll(ctx context.Context) ([]entity.Skill, error)
	GetById(id uint, ctx context.Context) (entity.Skill, error)
}

func (r skillService) AddSkill(skill entity.Skill, ctx context.Context) error {
	conn, err := r.p.Acquire(ctx)
	if err != nil {
		fmt.Printf("Unable to acquire a database connection: %v\n", err)
		return errors.New("unable to acquire a database connection")
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO skill (name) VALUES ($1) RETURNING id",
		skill.Name)
	var id uint
	err = row.Scan(&id)
	if err != nil {
		fmt.Printf("Unable to INSERT: %v\n", err)

		return errors.New("unable to insert")
	}
	return nil
}

func (r skillService) GetAll(ctx context.Context) ([]entity.Skill, error) {
	var skills []entity.Skill
	rows, err := r.p.Query(ctx, "SELECT * FROM skill")
	if err != nil {
		fmt.Printf("Unable to make query a database connection: %v\n", err)
		return skills, errors.New("unable to acquire a database connection")
	}
	for rows.Next() {
		var id uint
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println(err)
			return skills, err
		}
		skills = append(skills, entity.Skill{ID: id, Name: name})
	}

	return skills, nil
}

func (r skillService) GetById(id uint, ctx context.Context) (entity.Skill, error) {
	var skill entity.Skill

	rows, err := r.p.Query(ctx, "SELECT (name) FROM skill WHERE (id) = $1 ", id)
	if err != nil {
		fmt.Printf("Unable to make query a database connection: %v\n", err)
		return skill, errors.New("unable to acquire a database connection")
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println(err)
			return skill, err
		}
		skill = entity.Skill{ID: id, Name: name}
	}

	return skill, nil
}

func NewSkillSevice(p *pgxpool.Pool) SkillService {
	return &skillService{p: p}
}
