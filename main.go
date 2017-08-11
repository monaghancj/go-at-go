package.main

type Person struct {
	ID 				string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string
}

type Address struct {
	
}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
  router := mux.NewRouter()
  router.HandleFunc('/people'), GetPeopleEndPoint).Methods('GET')

  log.Fatal(http.ListenAndServe(':12345', router))
}