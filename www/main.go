package main

import ("fmt" //библиотека для вывода (println)
       "net/http"; "html/template")

type Employee struct {
  Login string
  Surname string
  Name string
  Age uint
  Position string
  Phone int
  Vacation_days float32
  Hobbies [] string
}

func (e Employee) getAllInfo() string {
  return fmt.Sprintf("Это : %s %s! Он является : %s. Ему : %d лет", e.Surname, e.Name, e.Position,
  e.Age)
}

func (e Employee) setNewPosition(newPosition string)  {
  e.Position = newPosition
}

func home_page(w http.ResponseWriter, r *http.Request)  {
  aram := Employee{"AAsryan", "Асрян", "Арам", 28, "dev", 9155915671, 13.5, [] string{}}
  fmt.Fprintf(w, "Good job\n" + aram.getAllInfo())
}

func about_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "About company")
}

func main_page(w http.ResponseWriter, r *http.Request) {
  aram := Employee{"AAsryan", "Асрян", "А", 27, "dev", 9155915671, 13.8, [] string{"Football", "Guitar", "Paint"}}
  tmpl, _ := template.ParseFiles("templates/main.html")
  tmpl.Execute(w, aram)
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/about", about_page)
  http.HandleFunc("/main", main_page)
  http.ListenAndServe(":8080", nil)
}

//главная функция
func main()  {
  handleRequest()
}
