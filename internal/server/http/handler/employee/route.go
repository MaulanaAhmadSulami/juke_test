package employeeHandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/MaulanaAhmadSulami/juke_test.git/internal/service"
	"go.uber.org/zap"
)

func RegisterRoute(
	employeService service.EmployeesService,
	logger *zap.SugaredLogger,
) func(chi.Router){
	return func(r chi.Router){
		handler := newHttpHandler(employeService, logger)
		r.Get("/", handler.GetAll)
		r.Get("/{employeeId}", handler.GetById)
		r.Post("/", handler.Create)
		r.Put("/{employeeId}", handler.Update)
		r.Delete("/{employeeId}", handler.Delete)
	}
}