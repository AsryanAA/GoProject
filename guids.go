package main

import "fmt"

func main() {
  fmt.Println("Типы данных Go")

  var i int = 0; //целочисленная переменная
  var f float64 = 0.0 //число с плавающей точкой
  var str string = " string";
  var isTrue bool = true;

  //условный оператор
  if i < 0 {
    fmt.Println("i < 0")
  } else if i == 0 {
    fmt.Println("i == 0")
  } else {
    fmt.Println("i > 0")
  }

  //конструкция switch case
  switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    default: fmt.Println("default")
  }

  //for
  for i := 0; i < 10; i++ {
    fmt.Print(i)
    fmt.Println("")
  }

  //arrays
  var arr[3] int
  arr[0] = 45;
  fmt.Println(i, f, str, isTrue)

  //defer - откладывание
  //type (struct) - аналог классов
}
