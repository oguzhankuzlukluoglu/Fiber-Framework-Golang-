package handlers

import (
	"firstAPI/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

var customers = []models.Customer{
	{ID: 1, Name: "Mert Sahin", Username: "mertshn", Phone: "05433434343"},
	{ID: 2, Name: "Tony Stark", Username: "ironMan", Phone: "3423423423"},
	{ID: 3, Name: "Blue", Username: "red", Phone: "443523423"},
} //mock data

func GetAllCustomer(c *fiber.Ctx) error {
	if len(customers) <= 0 {
		return c.Status(http.StatusNotFound).JSON("no data")
	}
	return c.Status(http.StatusOK).JSON(customers)
}

func GetByUsernameWithData(c *fiber.Ctx) error {
	username := c.Params("username")

	for _, element := range customers {
		if username == element.Username {
			return c.Status(http.StatusOK).JSON(element)
		}
	}

	return c.Status(http.StatusNotFound).JSON("Username you were looking for was not found :(")
}
func CreateCustomer(c *fiber.Ctx) error {
	var newCustomers models.Customer

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Message: Bad Request")
	}

	customers = append(customers, newCustomers)
	return c.Status(http.StatusOK).JSON("Success Created")
}
func UpdateCustomer(c *fiber.Ctx) error {
	var newCustomers models.Customer

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Message: Bad Request")
	}

	for i, element := range customers {
		if element.Username == newCustomers.Username {
			element.Name = newCustomers.Name
			element.Phone = newCustomers.Phone
			customers = append(customers[:i], customers[i+1:]...)
			customers = append(customers, element)
		}
	}

	return c.Status(http.StatusOK).JSON("Updated Data")
}
func DeleteCustomer(c *fiber.Ctx) error {
	username := c.Params("username")

	log.Println(username)

	for i, element := range customers {
		if element.Username == username {
			customers = append(customers[:i], customers[i+1:]...)
			fmt.Println(customers)
		}
	}
	return c.Status(http.StatusOK).JSON("Deleted Data")
}
