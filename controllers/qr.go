package controllers

import (
	"fmt"

	"github.com/Kchanit/qr-generator/models"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func HandleGenerateQRCode(c *fiber.Ctx) error {
	qr := new(models.QRCode)
	if err := c.BodyParser(qr); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(err)
	}
	if qr.Content == "" {
		fmt.Println("Could not determine the desired QR code content.")
		return c.Status(400).JSON("Could not determine the desired QR code content.")
	}
	if qr.Size == 0 {
		fmt.Println("Could not determine the desired QR code size.")
		return c.Status(400).JSON("Could not determine the desired QR code size.")
	}

	codeData, err := qrcode.Encode(qr.Content, qrcode.Medium, qr.Size)
	if err != nil {
		fmt.Println("could not generate a QR code: ", err)
		return c.Status(400).JSON(err)
	}

	c.Set(fiber.HeaderContentType, "image/png")

	return c.Status(200).Send(codeData)
}
