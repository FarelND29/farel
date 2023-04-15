package controller

import (
	"github.com/FarelND29/farel/config"
	inimodul "github.com/FarelND29/monitoring_orang_tua/module"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
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