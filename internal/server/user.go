package server

import (
	"fmt"
	"kreol/cmd/web"
	"kreol/internal/database"
	"net/http"
)

type UserHandler struct {
	user_db database.UserStore
}

func (h *UserHandler) GetCreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := web.HelloForm().Render(r.Context(), w); err != nil {
		fmt.Println(err)
	}
}
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Get it")

	if err := h.user_db.Create("tiffany"); err != nil {
		fmt.Println("this error occurred")
		return
	}
}
