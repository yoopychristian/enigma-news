package master

import (
	"database/sql"
	"enigma-news/main/master/controllers"
	"enigma-news/main/master/repositories"
	"enigma-news/main/master/usecases"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	registerRepo := repositories.InitRegisterRepoImpl(db)
	registerUseCase := usecases.InitRegisterUseCase(registerRepo)
	controllers.RegisterController(r, registerUseCase)

	articleRepo := repositories.InitArticleRepoImpl(db)
	articleUseCase := usecases.InitArticleUseCase(articleRepo)
	controllers.ArticleController(r, articleUseCase)

	userRepo := repositories.InitUserRepoImpl(db)
	userUseCase := usecases.InitUserUsecase(userRepo)
	controllers.UserController(r, userUseCase)
}
