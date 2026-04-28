package main
import(
	"fmt"
	"net/http"
)
func hello(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Привет, веб на GO!")
	}

func main(){
	http.HandleFunc("/", hello)

	http.ListenAndServe(":3000", nil)
	
}