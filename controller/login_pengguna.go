package controller

import(
    "github.com/ariespirnando/absenonline/config" 
    "github.com/ariespirnando/absenonline/model"
	"log"  
	"database/sql" 
    "net/http" 
    "github.com/gin-gonic/gin"   
)

func Login(c *gin.Context){
    var json model.Login //ini ngambile struct login 
    if err := c.ShouldBindJSON(&json);err != nil{ //Cek requestnya sesuai atau tidak
        c.JSON(http.StatusBadRequest, gin.H{ 
            "message" : "Bad Request",  
            "error_code" : "000001", 
        }) 
    }else{ //jika sesuai
        var user model.Register //ambil struct pengguna 
        DB := config.Connect() //config ke db
        defer DB.Close() //pastikan sedang non aktif
        
        err := DB.QueryRow("CALL `login_pengguna`(?,?)", //panggil query pake Query row untuk validasi jumlahnya
                json.Nomorindentitas, //parsing data
                json.Password,  
        ).Scan(
			&user.Namapengguna, //scan data yang didapatkan
            &user.Nomorindentitas,
            &user.Email,
            &user.Handphone) 
    
        if err != nil {
            log.Print(err) 
            c.JSON(http.StatusBadRequest, gin.H{ 
                "message" : "Database sedang Maintenance",
                "error_code" : "000003", 
            }) 
        }else{
            switch {
            case err == sql.ErrNoRows:
                c.JSON(http.StatusUnauthorized, gin.H{ 
                    "message" : "Pengguna tidak terautorisasi",
                    "error_code" : "000002", 
                })           
			default: 
				Token := createToken(&user)
				c.JSON(http.StatusOK, gin.H{ 
					"message" : "Berhasil Login",
					"error_code" : "000000",
					"detail":user,
					"token":Token,
				})   
            } 
        }
        
    }
}


func createToken(pengguna *model.Register) string{
    var tokenstring string
        DB := config.Connect() //config ke db
        defer DB.Close() //pastikan sedang non aktif
        err := DB.QueryRow("SELECT `create_token`(?)",
            pengguna.Nomorindentitas, 
            ).Scan(&tokenstring)
        if err != nil {
            log.Print(err)
        }
    return tokenstring
}