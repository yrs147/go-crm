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

func main(){
	
}