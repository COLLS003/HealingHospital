package database

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Open the database and establish the connection
func Init() *gorm.DB {

	// Specify connection properties.
	// Bulk@SMS2088
	//http://bdigismat.com/index.php?route=/table/structure&db=bulksms&table=users
	// dsn := "colls:1234@tcp(127.0.0.1:3306)/skin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "bulksms:Bulk@SMS2088@tcp(dbs.bdigismat.com:3306)/bulksms?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "colls:12345Ft(0!6@tcp(173.212.241.254:3306)/cradlevoices?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the database connection.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	return DB
}

// Function to get the database connection

// Function to get the database connection
func GetConnection() *gorm.DB {
	return DB
}

// bidner
func Bind(c *gin.Context, object interface{}) error {
	binder := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(object, binder)
}

// return customised error info
// helps return a custmised error info
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

//validators

func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		res1 := fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		fmt.Println(res1)
		if v.Param() != "" {

			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}
	return res
}

// wraping error into an object
// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

//
