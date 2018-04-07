package myhttp

import (
	"api/setting"
	"net/http"
)

// Start api les
func Start() error {
	fs := http.FileServer(http.Dir(setting.APIStaticDir))

	http.Handle("/", fs)

	http.HandleFunc("/user", switchMetod(
		// Выводить пользователя по id.
		accGetUser,
		// Добавлять пользователя.
		accSetUser,
		// Изменять имя пользователя.
		accUpdateUser,
		// Удалять пользователя.
		accDeleteUser,
	))

	return http.ListenAndServe(setting.APIListenAddr, nil)
}
