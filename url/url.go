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
	page.Get("/", controller.Homepage)                                    //ujicoba panggil package musik
	page.Get("/presensi", controller.GetAllPresensi)                      //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID)                   //menampilkan data presensi berdasarkan id
	page.Get("/monitoring/:id", controller.GetMonitoringID)
	page.Get("/mahasiswa/:id", controller.GetMahasiswaID)
	page.Get("/orangtua/:id", controller.GetOrangTuaID)
	page.Get("/dosenwali/:id", controller.GetDosenWaliID)
	page.Get("/tema/:id", controller.GetTemaID)
	page.Get("/all-monitoring", controller.GetAllMonitoring)
	page.Get("/all-mahasiswa", controller.GetAllMahasiswa)
	page.Get("/all-orangtua", controller.GetAllOrangTua)
	page.Get("/all-dosenwali", controller.GetAllDosenWali)
	page.Get("/all-tema", controller.GetAllTema)
	page.Get("/all-dosenwali", controller.GetAllDosenwali)
	page.Post("/ins", controller.InsertData)
	page.Post("/ins/monitoring", controller.InsertDataMonitoring)
	page.Get("/docs/*", swagger.HandlerDefault)
	page.Put("/upd/:id", controller.UpdateData)
	page.Put("/monitoring/:id", controller.UpdateDataMonitoring)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
	page.Delete("/monitoring/:id", controller.DeleteMonitoringByID)
}
