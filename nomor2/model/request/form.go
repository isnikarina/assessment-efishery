package request

type FormRequest struct {
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	Alamat       string `json:"alamat"`
	Kota         string `json:"kota"`
	Provinsi     string `json:"provinsi"`
	KodePos      string `json:"kode_pos"`
	NoTelp       string `json:"no_telp"`
	NoHp         string `json:"no_hp"`
	TanggalLahir string `json:"tanggal_lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Pendidikan   string `json:"pendidikan"`
	Pengalaman   string `json:"pengalaman"`
	Pengajuan    int32  `json:"pengajuan"`
}
