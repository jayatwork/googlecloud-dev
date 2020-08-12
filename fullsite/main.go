gcloud package main
//Necessar imports for stdlib and external
import (
   "net/http"
   "fmt"
   "log"
   "github.com/gorilla/mux"
   "encoding/json"
   "net"
   "strings"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
   Name string
   Time string
}

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!!!! All the business is at landing /web/ ")
	fmt.Println("Endpoint Hit: homePage")
}

func greetingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	fmt.Println("Hello, World!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Servername: "+key +" STATUS:")
	
	//Strip the unnecessary squileys
	topen := strings.Replace(key, "{", "", -1)
	tclose := strings.Replace(topen, "}", "", -1)
	
	//Do the tcp dialer for realtime stats 
    conn, err := net.Dial("tcp", tclose + ":3389")

        if err != nil {
            fmt.Fprintf(w, "SERVER UNREACHABLE \n")
        return
        }

    //Print discovery output remote < -- > local OCP orgin
    fmt.Fprintf(w, "SERVER ALIVE \n\n")
    fmt.Fprintf(w, "Remote Address : \n" + conn.RemoteAddr().String())
    //Called from Openshift Paas platform address
    fmt.Fprintf(w, "OCP caller Address :" + conn.LocalAddr().String())
	
	
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	staticFileDirectory := http.Dir("./web/")
	staticFileHandler := http.StripPrefix("/web/", http.FileServer(staticFileDirectory))
	myRouter.PathPrefix("/web/").Handler(staticFileHandler).Methods("GET")
	
	//myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/greeting", greetingPage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
   //Report status of webserver and enable handler
   fmt.Println("Listening");
   handleRequests()
   
}
