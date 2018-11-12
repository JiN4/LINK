package controller

import (
	"net/http"

	"github.com/LINK/model"
	"github.com/gin-gonic/gin"
)

// // 自分のユーザー情報
// type User struct { //
// 	Id       uint      `json:"id"`
// 	Uid      string    `json:"-"`
// 	Name     string    `json:"name"`
// 	Email    string    `json:"email"`
// 	Articles []Article `json:"articles"`
// 	CreateAt uint      `create_at`
// 	//        val githubStatus : GitHubStatus,// ログインで得られるGitHubの情報によってデータクラスを作る
// }


func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	model.CreateUser(model.User{
		UID:   user.UID, //ToDo 本来は認証アリ
		Name:  user.Name,
		Email: user.Email,
	})

	c.JSON(http.StatusOK, user)
}
