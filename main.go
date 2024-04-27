package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/AlexSilverson/Short_Hack_Back/docs"
	skillcontroller "github.com/AlexSilverson/Short_Hack_Back/src/controller/skillController"
	service "github.com/AlexSilverson/Short_Hack_Back/src/service/skillService"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
)

// @title			Short Hack
// @version		1.0
// @description	This is a sample swagger for Short Hack
// @termsOfService	http://swagger.io/terms/

func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgresql://SuperPuperUser:SuperStrongPassword@larek.tech:50000/dev")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	/*t := entitySkill.Skill{
		ID:   1,
		Name: "C++",
	}


	er := s.AddSkill(t, ctx)*/
	s := service.NewSkillSevice(pool)
	ans, er := s.GetById(uint(3), ctx)
	if er != nil {
		fmt.Print(er)

	}
	skillcontroller.GetSkillById(app, s, ctx)
	skillcontroller.GetAllSkills(app, s, ctx)
	skillcontroller.AddSkill(app, s, ctx)
	fmt.Print(ans)
	app.Listen(":3000")
}
