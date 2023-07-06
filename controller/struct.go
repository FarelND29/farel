package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Karyawan struct {
	//ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	Jabatan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty" example:"Anonymous"`
	Jam_kerja   []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty" example:"8"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty" example:"08:00"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty" example:"16:00"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty" example:"7"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty" example:"2"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty" example:"Piket Z"`
}

type Presensi struct {
	//ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.11"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	//Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin string   `bson:"checkin,omitempty" json:"checkin,omitempty" example:"MASUK"`
	Biodata Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
// struct monitoring orang tua

type Mahasiswa struct {
	/* ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"2343242"`  */
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Farel"`
	NPM          int                `bson:"npm,omitempty" json:"npm,omitempty" example:"1214070"`
	Jekel        string             `bson:"jekel,omitempty" json:"jekel,omitempty" example:"laki-laki"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"083821157026"`
}

type OrangTua struct {
	/* ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"example:"78678654`  */
	Nama_OT      string             `bson:"nama_ot,omitempty" json:"nama_ot,omitempty" example:"Rini"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"083821157029"`
	Anak         Mahasiswa          `bson:"anak,omitempty" json:"anak,omitempty" example:"Farel"`
}

type DosenWali struct {
	/* ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"example:"98786576"` */
	Nama_Dosen   string             `bson:"nama_dosen,omitempty" json:"nama_dosen,omitempty" example:"Roni"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty" example:"jl.delima"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08917271575"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty" example:"roni12@gmail.com"`
}

type Tema struct {
/* 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"example:"2675642` */
	Nama_Tema string             `bson:"nama_tema,omitempty" json:"nama_tema,omitempty" example:"web"`
}

type Monitoring struct {
/* 	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"example:"289790` */
	OrangTua OrangTua           `bson:"ortu,omitempty" json:"ortu,omitempty" example:"Rini"`
	Tema     Tema               `bson:"tema,omitempty" json:"tema,omitempty" example:"web"`
	Dosen    DosenWali          `bson:"dosen,omitempty" json:"dosen,omitempty" example:"Roni"`
	Tanggal  string             `bson:"tanggal,omitempty" json:"tanggal,omitempty" example:"12-2-2023"`
	Hari     string             `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin"`
}