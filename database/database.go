package database

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres" // Use the PostgreSQL driver
	"gorm.io/gorm"
)

var DB *gorm.DB

// Open the database and establish the connection
func Init() *gorm.DB {
	// Specify the connection properties for the PostgreSQL container
	dsn := "user=docker password=docker dbname=hospital sslmode=disable  host=localhost port=5432"
	// The 'host' should match the service name defined in your Docker Compose file (in this case, 'db')
	// 'user', 'password', 'dbname', and 'sslmode' should match your PostgreSQL configuration

	// Open the database connection.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	return DB
}

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
