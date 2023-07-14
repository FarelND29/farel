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
	page.Delete("/del/monitoring/:id", controller.DeleteMonitoringByID)
	// END POINT MAHASISWA
	page.Get("/mahasiswa/:id", controller.GetMahasiswaID)
	page.Get("/all-mahasiswa", controller.GetAllMahasiswa)
	page.Post("/ins/mahasiswa", controller.InsertDataMahasiswa)
	page.Put("/upd/mahasiswa/:id", controller.UpdateDataMahasiswa)
	page.Delete("/del/mahasiswa/:id", controller.DeleteMahasiswaByID)
	//END POINT ORANGTUA
	page.Get("/orangtua/:id", controller.GetOrangTuaID)
	page.Get("/all-orangtua", controller.GetAllOrangTua)
	page.Post("/ins/orangtua", controller.InsertDataOrangTua)
	page.Put("/upd/orangtua/:id", controller.UpdateDataOrangTua)
	page.Delete("/del/orangtua/:id", controller.DeleteOrangTuaByID)
	//END POINT DOSEN
	page.Get("/dosenwali/:id", controller.GetDosenWaliID)
	page.Get("/all-dosenwali", controller.GetAllDosenWali)
	page.Post("/ins/dosenwali", controller.InsertDataDosenWali)
	page.Put("/upd/dosenwali/:id", controller.UpdateDataDosenWali)
	page.Delete("/del/dosenwali/:id", controller.DeleteDosenWaliByID)
	//END POINT TEMA
	page.Get("/tema/:id", controller.GetTemaID)
	page.Get("/all-tema", controller.GetAllTema)
	page.Post("/ins/tema", controller.InsertDataTema)
	page.Put("/upd/tema/:id", controller.UpdateDataTema)
	page.Delete("/del/tema/:id", controller.DeleteTemaByID)
	
}
