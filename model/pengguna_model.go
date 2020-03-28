package model
type Register struct{
    Namapengguna string `form:"nama_pengguna" json:"nama_pengguna" binding:"required"`
    Nomorindentitas string `form:"nomor_indentitas" json:"nomor_indentitas" binding:"required"`  
	Email string `form:"email" json:"email" binding:"required"`
	Handphone string `form:"nomor_handphone" json:"nomor_handphone"` 
    Password string `form:"password" json:"password,omitempty" binding:"required"` 
}
type Login struct{
    Nomorindentitas string `form:"nomor_indentitas" json:"nomor_indentitas" binding:"required"` 
    Password string `form:"password" json:"password" binding:"required"` 
}