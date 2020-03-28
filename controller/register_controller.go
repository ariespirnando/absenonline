package controller

import(
    "github.com/ariespirnando/absenonline/config" 
    "github.com/ariespirnando/absenonline/model"
	"log"   
    "net/http" 
    "github.com/gin-gonic/gin"   
)

func Register(c *gin.Context){
	var json model.Register //ini ngambile struct login 
	var status int
	DB := config.Connect() //config ke db
    defer DB.Close() //pastikan sedang non aktif
    if err := c.ShouldBindJSON(&json);err != nil{ //Cek requestnya sesuai atau tidak
        c.JSON(http.StatusBadRequest, gin.H{ 
            "message" : "Bad Request",  
            "error_code" : "000001", 
        }) 
    } else{

		err := DB.QueryRow("SELECT `register_pengguna`(?, ?, ?, ?, ?) as StatusRegister",
			json.Namapengguna,
			json.Nomorindentitas,
			json.Email,
			json.Handphone,
			json.Password, 
			).Scan(&status)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{ 
                "message" : "Database sedang Maintenance",
                "error_code" : "000003", 
            }) 
		}else{
			if status==1 {
				c.JSON(http.StatusOK, gin.H{ 
					"message" : "Username atau Email sudah terdaftar",  
					"error_code" : "000004", 
				})
			}else{
				//Send Email disini
				c.JSON(http.StatusOK, gin.H{ 
					"message" : "Berhasil register",  
					"error_code" : "000000", 
				})
			} 
		} 
	}
}


 