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
	// END POINT PRESENSI
	page.Get("/presensi", controller.GetAllPresensi)                      //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID)                   //menampilkan data presensi berdasarkan id
	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
	page.Get("/docs/*", swagger.HandlerDefault)
	// END PONT MONITORING
	page.Get("/monitoring/:id", controller.GetMonitoringID)
	page.Get("/all-monitoring", controller.GetAllMonitoring)
	page.Post("/ins/monitoring", controller.InsertDataMonitoring)
	page.Put("/upd/monitoring/:id", controller.UpdateDataMonitoring)
	page.Delete("/del-monitoring/:id", controller.DeleteMonitoringByID)
	// END POINT MAHASISWA
	page.Get("/mahasiswa/:id", controller.GetMahasiswaID)
	page.Get("/all-mahasiswa", controller.GetAllMahasiswa)
	//END POINT ORANGTUA
	page.Get("/orangtua/:id", controller.GetOrangTuaID)
	page.Get("/all-orangtua", controller.GetAllOrangTua)
	//END POINT DOSEN
	page.Get("/dosenwali/:id", controller.GetDosenWaliID)
	page.Get("/all-dosenwali", controller.GetAllDosenWali)
	//END POINT TEMA
	page.Get("/tema/:id", controller.GetTemaID)
	page.Get("/all-tema", controller.GetAllTema)
}
