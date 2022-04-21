package repository

// import (
// 	"testing"
// )

// // func initAuthRepo(req *require.Assertions) (*Repository, *sqlx.DB) {
// // 	viper.AddConfigPath("../../configs")
// // 	viper.SetConfigName("config")

// // 	db, err := NewPostgresDB(Config{
// // 		Host:     viper.GetString("db.host"),
// // 		Port:     viper.GetString("db.port"),
// // 		Username: viper.GetString("db.username"),
// // 		DBName:   viper.GetString("db.dbname"),
// // 		SSLMode:  viper.GetString("db.sslmode"),
// // 		Password: os.Getenv("DB_PASSWORD"),
// // 	})
// // 	req.NoErrorf(err, "failed to initialize db: %s", err.Error())
// // 	return NewRepository(db), db
// // }

// func TestAuthPostgres_CreateUser(t *testing.T) {
// 	// if testing.Short() {
// 	// 	t.Skip("skip repo tests")
// 	// }
// 	/* req := require.New(t)
// 	// repo, db := initAuthRepo(req)
// 	// defer db.Close()

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	repoMock := NewMockAuthorization(ctrl)
// 	user := todoapp.User{
// 		Name:     "Make",
// 		Username: "userName",
// 		Password: "password",
// 	}
// 	// repo.CreateUser(user)
// 	repoMock.EXPECT().CreateUser(user).Return(1, nil).Times(1)
// 	id, err := repoMock.CreateUser(user)
// 	req.NoError(err, "User Create err")
// 	req.Equal(1, id, "crea te user exp != act")*/
// }
