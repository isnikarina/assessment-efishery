package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-efishery/db"
	"test-efishery/model"
	"test-efishery/model/request"

	"github.com/gorilla/mux"
)

func CreateForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	form := new(request.FormRequest)
	json.NewDecoder(r.Body).Decode(&form)

	newForm := model.FormModel{
		Nama:         form.Nama,
		Email:        form.Email,
		Alamat:       form.Alamat,
		Kota:         form.Kota,
		Provinsi:     form.Provinsi,
		KodePos:      form.KodePos,
		NoTelp:       form.NoTelp,
		NoHp:         form.NoHp,
		TanggalLahir: form.TanggalLahir,
		Pekerjaan:    form.Pekerjaan,
		Pendidikan:   form.Pendidikan,
		Pengalaman:   form.Pengalaman,
		Pengajuan:    form.Pengajuan,
	}
	db.DB.Create(&newForm)
	json.NewEncoder(w).Encode(newForm)
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var form []model.FormModel

	db.DB.Find(&form)

	json.NewEncoder(w).Encode(form)
}

func UpdateForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	formId := mux.Vars(r)

	formID, _ := strconv.Atoi(formId["id"])

	formUpdate := new(request.FormRequest)

	json.NewDecoder(r.Body).Decode(&formUpdate)

	var form model.FormModel

	db.DB.First(&form, "id = ?", formID)

	form.Nama = formUpdate.Nama
	form.Email = formUpdate.Email
	form.Alamat = formUpdate.Alamat
	form.Kota = formUpdate.Kota
	form.Provinsi = formUpdate.Provinsi
	form.KodePos = formUpdate.KodePos
	form.NoTelp = formUpdate.NoTelp
	form.NoHp = formUpdate.NoHp
	form.TanggalLahir = formUpdate.TanggalLahir
	form.Pekerjaan = formUpdate.Pekerjaan
	form.Pendidikan = formUpdate.Pendidikan
	form.Pengalaman = formUpdate.Pengalaman
	form.Pengajuan = formUpdate.Pengajuan

	db.DB.Save(&form)

	json.NewEncoder(w).Encode(form)
}

func DeleteForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	formId := mux.Vars(r)

	formID, _ := strconv.Atoi(formId["id"])

	var form model.FormModel

	db.DB.First(&form, "id = ?", formID)

	db.DB.Delete(&form)

	json.NewEncoder(w).Encode(form)
}
