package skillcontroller

import (
	"context"
	"fmt"
	"strconv"

	skillservice "github.com/AlexSilverson/Short_Hack_Back/src/service/skillService"
	"github.com/gofiber/fiber/v2"
)

// GetSkillById Getting Skill by ID
//
//	@Summary		Getting Skill by Id
//	@Description	Getting Skill by Id in detail
//	@Tags			Skills
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of Skill"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/skill/{id} [get]
func GetSkillById(app *fiber.App, skillService skillservice.SkillService, ctx context.Context) fiber.Router {
	return app.Get("/skill/:id", func(c *fiber.Ctx) error {
		fmt.Println("here")
		skillId := c.Params("id")
		id, err := strconv.ParseInt(skillId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}

		skill, er := skillService.GetById(uint(id), ctx)
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(er)
		}

		return c.Status(fiber.StatusOK).JSON(skill)
	})
}

// GetSkillById Getting all Skill
//
//	@Summary		Getting all Skill
//	@Tags			Skills
//	@Accept			json
//	@Produce		json
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/skill/all/ [get]
func GetAllSkills(app *fiber.App, skillService skillservice.SkillService, ctx context.Context) fiber.Router {
	return app.Get("/skill/all/", func(c *fiber.Ctx) error {

		skills, er := skillService.GetAll(ctx)
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(er)
		}

		return c.Status(fiber.StatusOK).JSON(skills)
	})
}
