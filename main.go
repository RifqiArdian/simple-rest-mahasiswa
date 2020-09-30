package main

import (
    "context"
    "fmt"
    "simple-rest-mahasiswa/mahasiswa"
    "simple-rest-mahasiswa/models"
    "simple-rest-mahasiswa/utils"
    "log"
    "net/http"
    "strconv"
    "encoding/json"
    jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)
 
func main() {
 
    http.HandleFunc("/mahasiswa", GetMahasiswa)
    http.HandleFunc("/login", Login)
    http.HandleFunc("/mahasiswa/create", PostMahasiswa)
    http.HandleFunc("/mahasiswa/update", UpdateMahasiswa)
    http.HandleFunc("/mahasiswa/delete", DeleteMahasiswa)
    err := http.ListenAndServe(":8000", nil)
 
    if err != nil {
        log.Fatal(err)
    }
}

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ctx, cancel := context.WithCancel(context.Background())
 
        defer cancel()
 
        mahasiswas, err := mahasiswa.GetAll(ctx)
 
        if err != nil {
            fmt.Println(err)
        }
 
        utils.ResponseJSON(w, mahasiswas, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
    return
}

func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
 
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var mhs models.Mahasiswa
 
        if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        if err := mahasiswa.Insert(ctx, mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
    if r.Method == "PUT" {
 
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var mhs models.Mahasiswa
 
        if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        fmt.Println(mhs)
 
        if err := mahasiswa.Update(ctx, mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
 
    if r.Method == "DELETE" {
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var mhs models.Mahasiswa
 
        id := r.URL.Query().Get("id")
 
        if id == "" {
            utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
            return
        }
        mhs.ID, _ = strconv.Atoi(id)
 
        if err := mahasiswa.Delete(ctx, mhs); err != nil {
 
            kesalahan := map[string]string{
                "error": fmt.Sprintf("%v", err),
            }
 
            utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}