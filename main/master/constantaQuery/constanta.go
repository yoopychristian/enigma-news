package constantaQuery

const (
	//karyawan permanen
	GETALLREGISTER     = "SELECT * FROM register"
	GETREGISTERBYUNAME = "SELECT * FROM register WHERE username= ?"
	GETADDREGISTER     = "INSERT INTO register VALUES (?,?,?,?,?,?)"
	GETUPDATEREGISTER  = "UPDATE register SET  email=?, full_name=?, phone_number=?, password=? WHERE username=?"
	GETDELETEREGISTER  = "DELETE FROM Register WHERE username=?"

	GETALLARTICLE    = "SELECT * FROM article"
	GETARTICLEBYID   = "SELECT * FROM article WHERE id=?"
	GETADDARTICLE    = "INSERT INTO register VALUES (?,?,?)"
	GETUPDATEARTICLE = "UPDATE register SET  title=?, article=? WHERE id=?"
	GETDELETEARTICLE = "DELETE FROM Register WHERE Register_id=?"
)
