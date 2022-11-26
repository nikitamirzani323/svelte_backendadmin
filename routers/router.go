package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/svelte_backendadmin/controller"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/dist", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	app.Get("api/healthz", controller.HealthCheck)
	app.Post("/api/login", controller.Login)
	app.Post("/api/home", controller.Home)
	app.Post("/api/dashboardlistpasaran", controller.Dashboardlistpasaran)
	app.Post("/api/dashboardwinlose", controller.Dashboardwinlose)
	app.Post("/api/dashboardwinlosepasaran", controller.Dashboardwinlosepasaran)
	app.Post("/api/reportwinlose", controller.Reportwinlose)
	app.Post("/api/generatorsave", controller.Generatorsave)

	app.Post("/api/allpasaran", controller.Pasaran)
	app.Post("/api/editpasaran", controller.Pasarandetail)
	app.Post("/api/savepasaran", controller.Pasaransave)
	app.Post("/api/savepasaranonline", controller.Pasaransaveonline)
	app.Post("/api/deletepasaranonline", controller.Pasarandeleteonline)
	app.Post("/api/savepasaranlimitline", controller.Pasaransavelimit)
	app.Post("/api/savepasaranconf432", controller.Pasaransaveconf432d)
	app.Post("/api/savepasaranconfcolokbebas", controller.Pasaransaveconfcbebas)
	app.Post("/api/savepasaranconfcolokmacau", controller.Pasaransaveconfcmacau)
	app.Post("/api/savepasaranconfcoloknaga", controller.Pasaransaveconfcnaga)
	app.Post("/api/savepasaranconfcolokjitu", controller.Pasaransaveconfcjitu)
	app.Post("/api/savepasaranconf5050umum", controller.Pasaransaveconf5050umum)
	app.Post("/api/savepasaranconf5050special", controller.Pasaransaveconf5050special)
	app.Post("/api/savepasaranconf5050kombinasi", controller.Pasaransaveconf5050kombinasi)
	app.Post("/api/savepasaranconfmacaukombinasi", controller.Pasaransaveconfmacaukombinasi)
	app.Post("/api/savepasaranconfdasar", controller.Pasaransaveconfdasar)
	app.Post("/api/savepasaranconfshio", controller.Pasaransaveconfshio)

	app.Post("/api/allperiode", controller.Periode)
	app.Post("/api/editperiode", controller.Periodedetail)
	app.Post("/api/saveperiode", controller.Periodesave)
	app.Post("/api/savepasarannew", controller.Periodesavenew)
	app.Post("/api/saveperioderevisi", controller.Periodesaverevisi)
	app.Post("/api/cancelbet", controller.Periodecancelbet)
	app.Post("/api/periodelistmember", controller.Periodelistmember)
	app.Post("/api/periodelistmemberbynomor", controller.Periodelistmemberbynomor)
	app.Post("/api/periodelistbet", controller.Periodelistbet)
	app.Post("/api/periodelistbetstatus", controller.Periodelistbetstatus)
	app.Post("/api/periodelistbetusername", controller.Periodelistbetbyusername)
	app.Post("/api/periodelistbettable", controller.Periodelistbettable)
	app.Post("/api/periodebettable", controller.Periodebettable)
	app.Post("/api/listpasaran", controller.Periodelistpasaran)
	app.Post("/api/listprediksi", controller.Periodeprediksi)

	app.Post("/api/alladmin", controller.Admin)
	app.Post("/api/editadmin", controller.Admindetail)
	app.Post("/api/saveadmin", controller.Adminsave)
	app.Post("/api/saveadminiplist", controller.Adminsaveiplist)
	app.Post("/api/deleteadminiplist", controller.Deleteiplist)

	app.Post("/api/alladminrule", controller.Adminrule)
	app.Post("/api/editadminrule", controller.Adminruledetail)
	app.Post("/api/saveadminrule", controller.Adminrulesave)
	app.Post("/api/saveadminruleconf", controller.Adminruleconfsave)

	app.Post("/api/log", controller.Log)
	return app
}
