package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/util"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uqq, err := r.Cookie("uqq")
	if err != nil || uqq == nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}

}

func UserIsLogin(uqq int, token string) int {
	t := util.CreateToken(uqq)
	res := util.ResolveToken(t)
	if token == res {
		return 1
	} else {
		return 0
	}
}
