package lead

import (
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yrs147/go-crm/database"

)
type Lead struct{
	gorm.Model
	Name string
	Company string
	Email string 
	Phone int
}

func GetLeads(c *fiber.Ctx){
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx){
	id := c.Params
	db := database.DBConn
	car lead lead 
	db.Find(&lead,id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx){
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err !=nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
	
}

func DeleteLead(c *fiber.Ctx){

}
