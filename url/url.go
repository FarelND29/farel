package url

import (
	"github.com/FarelND29/farel/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage) //ujicoba panggil package musik
	page.Get("/presensi", controller.GetAllPresensi) //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Get("/monitoring", controller.GetMonitoring)
	page.Get("/mahasiswa", controller.GetMahasiswa)
	page.Get("/orangtua", controller.GetOrangTua)
	page.Get("/tema", controller.GetTema)
	page.Get("/all-monitoring", controller.GetAllMonitoring)
	page.Get("/all-mahasiswa", controller.GetAllMahasiswa)
	page.Get("/all-orangtua", controller.GetAllOrangTua)
	page.Get("/all-tema", controller.GetAllTema)
	page.Post("/ins", controller.InsertData)
	page.Get("/docs/*", swagger.HandlerDefault)
}
