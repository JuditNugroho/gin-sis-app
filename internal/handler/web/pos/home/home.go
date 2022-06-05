package home

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebPOSHomeHandler(ctx *fiber.Ctx) error {

	return ctx.Render("pos/home", constantsEntity.WebData{
		Title:        constantsEntity.WebPosHomeTitle,
		BaseURL:      constantsEntity.BaseURL,
		CurrentURL:   constantsEntity.WebPOSHomeURL,
		TemplateURL:  constantsEntity.TemplateUrl,
		LinkPageList: constantsEntity.LinkPageList,
	})
}