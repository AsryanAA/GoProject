package main

import ("fmt" //библиотека для вывода (println)
       "net/http")

type Employee struct {
  login string
  surname string
  name string
  age uint
  position string
  phone int
  vacation_days float32
}

func (e Employee) getAllInfo() string {
  return fmt.Sprintf("Это : %s %s! Он является : %s. Ему : %d лет", e.surname, e.name, e.position,
  e.age)
}

func home_page(w http.ResponseWriter, r *http.Request)  {
  aram := Employee{"AAsryan", "Асрян", "Арам", 28, "dev", 9155915671, 13.5}
  fmt.Fprintf(w, "Good job\n" + aram.getAllInfo())
}

func about_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "About company")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/about", about_page)
  http.ListenAndServe(":8080", nil)
}

func main()  {
  handleRequest()
}
