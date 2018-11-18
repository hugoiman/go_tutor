//  https://dasarpemrogramangolang.novalagung.com/48-web.html

package main

import (
  "fmt"
  "net/http"   //  package web
  "html/template"
)


func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "apa kabar!")
}

func main() {
  //  http.HandleFunc() digunakan untuk routing aplikasi web. param2 adalah CALLBACK
  http.HandleFunc("/index", index)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintln(w, "halo!")
    var data = map[string] string {
      "Name" :" Ethan Hunt",
      "Message" : "Mission Succes",
    }

    var t, err = template.ParseFiles("template.html")   //  template.ParseFiles() u/ parsing template, mengembalikan 2 data yaitu instance template-nya dan error (jika ada).
    if err != nil {
      fmt.Println(err.Error())
      return
    }
    t.Execute(w, data)    //  Execute() akan membuat hasil parsing template ditampilkan ke layar web browser.
  })

  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)   //  http.listenAndServe() digunakan untuk menghidupkan server sekaligus menjalankan aplikasi menggunakan server tersebut.
                                      //  Di golang, 1 web aplikasi adalah 1 buah server berbeda.
                                      //  contoh berikut server dijalankan pada port 8080.
                                      //  Untuk menghentikan web server, tekan CTRL+C pada terminal
}
