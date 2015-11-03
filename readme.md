# phpcgi

usage:

`

	import (
	    "net/http"
	    "github.com/kiancyc/gophp"
	)

	func main() {    
	    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	    	gophp.Cgi(w, r, "C:/php/php-cgi", "E:/goproject/src/bamboo/api/articles/get.php")
	    })
	    http.ListenAndServe(":8080", nil)
	}
	
`