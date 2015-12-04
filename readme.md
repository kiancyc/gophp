# gophp

usage(on windows):


	import (
	    "net/http"
	    "github.com/kiancyc/gophp"
	)

	func main() {    
	    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	    	gophp.Cgi(w, r, "C:/php/php-cgi", "E:/php/test.php")
	    })
	    http.ListenAndServe(":8080", nil)
	}
	
test.php:

	<?php
	echo "Hello, {$_GET['name']}";
	
visit: http://localhost:8080/?name=John

	Hello, John
