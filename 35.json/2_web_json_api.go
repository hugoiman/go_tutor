//  Web API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML.

package main

import "encoding/json"
import "net/http"
import "fmt"

func main() {
  http.HandleFunc("/users", users)
  http.HandleFunc("/user", user)

  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)
}

type student struct {
  ID    string
  Name  string
  Grade int
}

var data = []student{
  student{"E001", "ethan", 21},
  student{"W001", "wick", 22},
  student{"B001", "bourne", 23},
  student{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")  //  w.Header().Set("Content-Type", "application/json") u/ menentukan tipe response, yaitu sebagai JSON.

  // Jika request adalah POST, maka data yang di-encode ke JSON dijadikan sebagai response.

  if r.Method == "POST" {                           //  r.Method() u/ deteksi jenis request  (Post/Get/Lainnya)
    var result, err = json.Marshal(data)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Write(result)   //  r.Write() digunakan untuk mendaftarkan data sebagai response.
    return
  }

  http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST" {
    var id = r.FormValue("id")
    var result []byte
    var err error

    for _, each := range data {
      if each.ID == id {
        result, err = json.Marshal(each)

        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }

        w.Write(result)
        return
      }
    }

    http.Error(w, "User not found", http.StatusBadRequest)
    return
  }

  http.Error(w, "", http.StatusBadRequest)
}
