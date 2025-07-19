package models

import (
	"time"

	"github.com/google/uuid"
)

//

type Dusun struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Nama      string    `json:"nama" db:"nama"`
	Warna     string    `json:"warna" db:"warna"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type StatistikDusun struct {
	ID              uuid.UUID `json:"id" db:"id"`
	DusunID         uuid.UUID `json:"dusun_id" db:"dusun_id"`
	JumlahKK        int       `json:"jumlah_kk" db:"jumlah_kk"`
	JumlahLaki      int       `json:"jumlah_laki" db:"jumlah_laki"`
	JumlahPerempuan int       `json:"jumlah_perempuan" db:"jumlah_perempuan"`
	JumlahPenduduk  int       `json:"jumlah_penduduk" db:"jumlah_penduduk"`
	JumlahRumah     int       `json:"jumlah_rumah" db:"jumlah_rumah"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

//
