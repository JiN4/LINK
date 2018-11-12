package createSeeds

import (
	"fmt"

	"github.com/LINK/model"
)

// // 自分のユーザー情報
// type User struct {
// 	//Model
// 	gorm.Model
// 	Uid   string
// 	Name  string
// 	Email string
// 	//        val githubStatus : GitHubStatus,// ログインで得られるGitHubの情報によってデータクラスを作る
// }

func CreateSeedUsers() {

	users_infos := []map[string]string{
		map[string]string{
			"Uid":   "SuOEVHYgdjJGD9fffdsarR27NX8h",
			"Name":  "野口英世",
			"Email": "noguti_hideyo@gmail.com",
		},
		map[string]string{
			"Uid":   "UFhjdr86rNC764VNY7vE$u8VHHGD",
			"Name":  "樋口一葉",
			"Email": "higuti_itiyo@yahoo.co.jp",
		},
		map[string]string{
			"Uid":   "dFzSuOE4F1Vrmb7aHTmNmarR2ES2",
			"Name":  "福沢諭吉",
			"Email": "hukuzawa_yukiti@ezweb.ne.jp",
		},
	}

	for _, info := range users_infos {
		createUser(model.User{
			UID:   info["Uid"],
			Name:  info["Name"],
			Email: info["Email"],
		})
	}
}

func createUser(user model.User) {
	user, err := model.CreateUser(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user created\n")
}
