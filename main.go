package main
import( 
    "github.com/ariespirnando/absenonline/controller"   
    "github.com/gin-gonic/gin" 
    "github.com/subosito/gotenv" 
    "log"
    "os" 
    "fmt"
)
func init()  {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
    } 
    if(os.Getenv("GIN_MODE")=="release"){
        gin.SetMode(gin.ReleaseMode)
    }
}
func main(){   
    fmt.Printf("Welcome API \n")
    router := gin.Default()  
    v1 := router.Group("/api/pengguna/")
    {  
        v1.POST("/login", controller.Login) 
        v1.POST("/register", controller.Register)  
    }  
    router.Run(os.Getenv("PORT"))
}