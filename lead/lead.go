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

}

func GetLead(c *fiber.Ctx){
	id := c.Params
	db := database.DBConn
	car lead lead 
	db.Find(&lead,id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx){

}

func DeleteLead(c *fiber.Ctx){

}
