package app

import (
	export_template "capp-api/cmd/panda/panda-fulfillment-api/app/handlers/export-template"
	panda_services "capp-api/exmsg/panda/services"
	"capp-api/pkg/codename"
	"net/http"

	"gitlab.shopbase.dev/brodev/beekit/core/bkmicro"
	"gitlab.shopbase.dev/brodev/beekit/core/transport/transhttp"
)

// Server --
type server struct {
	ExportTemplateClient panda_services.ExportTemplateClient
}

// NewServer return a new server
func NewServer(bkMicroApp *bkmicro.App) {
	m := bkMicroApp.InitGRPCClients(
		codename.DMSExportTemplate.String(),
	)

	s := &server{
		ExportTemplateClient: panda_services.NewExportTemplateClient(
			m[codename.DMSExportTemplate.String()]),
	}

	bkMicroApp.AddRoutes(s.InitRoutes(bkMicroApp.HTTPBasePath))
}

func (s *server) initS3Storage() {

}

// OnClose -- on close
func OnClose() {
	//rabbitMQProducer.Close()
}

// InitRoutes -- Initialize our routes
func (s *server) InitRoutes(basePath string) transhttp.Routes {
	// Initialize our routes

	return transhttp.Routes{

		//GET /v1/panda-fulfillment/templates
		transhttp.Route{
			Name:     "export list template",
			Method:   http.MethodGet,
			BasePath: basePath,
			Pattern:  "/templates",
			Handler: &export_template.ExportTemplateHanlder{
				ExportTemplateClient: s.ExportTemplateClient,
			},
			Timeout: 30000,
		},
	}
}
