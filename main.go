/* 
http://www.cplusplus.com/articles/2LywvCM9/
 http://www.alternatestack.com/development/app-development/simple-http-server-with-templated-response-in-go/

 */

package main
import (
 "fmt"
 "net/http"
 "os"
)

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "index.html")
}

func main() {
 var PORT string

 if PORT = os.Getenv("PORT"); PORT == "" {
  PORT = "8001"
 }
 
 fmt.Println("Server Starting Listening on port "+ PORT)
 
 http.HandleFunc("/", HttpFileHandler)

 http.ListenAndServe(":" + PORT, nil)

}