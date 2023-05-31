package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/FarelND29/farel/config"
	inimodul "github.com/FarelND29/monitoring_orang_tua/module"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	inimodel "github.com/indrariksa/be_presensi/model"
	inimodul1 "github.com/indrariksa/be_presensi/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodul1.GetAllPresensi(config.Ulbimongoconn2, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodul1.GetPresensiFromID(objID, config.Ulbimongoconn2, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetMonitoring(c *fiber.Ctx) error {
	ps := inimodul.GetMonitoringFromNamaMahasiswa(config.Ulbimongoconn, "monitoring", "Farel")
	return c.JSON(ps)
}

func GetMahasiswa(c *fiber.Ctx) error {
	ps := inimodul.GetMahasiswaFromNpm(config.Ulbimongoconn, "mahasiswa", 1214070)
	return c.JSON(ps)
}

func GetOrangTua(c *fiber.Ctx) error {
	ps := inimodul.GetOrangTuaFromNamaMahasiswa(config.Ulbimongoconn, "orangtua", "toni")
	return c.JSON(ps)
}

func GetTema(c *fiber.Ctx) error {
	ps := inimodul.GetTemaFromNamaTema(config.Ulbimongoconn, "tema", "Kewirausahaan")
	return c.JSON(ps)
}

func GetAllMonitoring(c *fiber.Ctx) error {
	ps := inimodul.GetAllMonitoring(config.Ulbimongoconn, "monitoring")
	return c.JSON(ps)
}

func GetAllMahasiswa(c *fiber.Ctx) error {
	ps := inimodul.GetAllMahasiswa(config.Ulbimongoconn, "mahasiswa")
	return c.JSON(ps)
}

func GetAllOrangTua(c *fiber.Ctx) error {
	ps := inimodul.GetAllOrangTua(config.Ulbimongoconn, "orangtua")
	return c.JSON(ps)
}

func GetAllTema(c *fiber.Ctx) error {
	ps := inimodul.GetAllTema(config.Ulbimongoconn, "tema")
	return c.JSON(ps)
}

func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn2
	var presensi inimodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodul1.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
