package gophp

import (
	"net/http/cgi"
	"net/http"
)

// serve http with php script(cgi mode)
func Cgi(w http.ResponseWriter, r *http.Request, phpBin string, scriptFileName string) {
    handler 	:= new(cgi.Handler)
	handler.Path = phpBin
	handler.Env  = append(handler.Env, "REDIRECT_STATUS=CGI") 
	handler.Env  = append(handler.Env, "SCRIPT_FILENAME="+scriptFileName) 

    handler.ServeHTTP(w, r)
}
