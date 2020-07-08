package constantaQuery

const (
	//karyawan permanen
	GETALLREGISTER     = "SELECT * FROM registrasi"
	GETREGISTERBYUNAME = "SELECT * FROM registrasi WHERE username= ?"
	GETADDREGISTER     = "INSERT INTO registrasi VALUES (?,?,?,?,?,?)"
	GETUPDATEREGISTER  = "UPDATE registrasi SET email=?, full_name=?, phone_number=?, username=?, password=? WHERE id=?"
	GETDELETEREGISTER  = "DELETE FROM registrasi WHERE username=?"

	GETALLARTICLE    = "SELECT * FROM article"
	GETARTICLEBYID   = "SELECT * FROM article WHERE id=?"
	GETADDARTICLE    = "INSERT INTO article VALUES (?,?,?)"
	GETUPDATEARTICLE = "UPDATE article SET  title=?, article=? WHERE id=?"
	GETDELETEARTICLE = "DELETE FROM article WHERE id=?"
)
