package employeeHandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	employeeEntity "github.com/MaulanaAhmadSulami/juke_test.git/internal/entities/employees"
	repository "github.com/MaulanaAhmadSulami/juke_test.git/internal/repository/postgres"
	"github.com/MaulanaAhmadSulami/juke_test.git/internal/server/http/protocol"
	"github.com/MaulanaAhmadSulami/juke_test.git/internal/service"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type HttpHandler struct {
	employeeService service.EmployeesService
	logger *zap.SugaredLogger
}

func newHttpHandler(employeeService service.EmployeesService, logger *zap.SugaredLogger) *HttpHandler {
	return &HttpHandler{
		employeeService: employeeService,
		logger: logger,
	}
}


// Get Employees godoc
//
// @Summary Get All Employees
// @Description Get a list of all employees available
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {object} employeeEntity.Employee
// @Failure 500 {object} map[string]string	"Internal server error"
// @Failure 404 {object} map[string]string	"not found"
// @Router /employees [get]
func (h *HttpHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	employees, err := h.employeeService.GetAll(ctx)
	if err != nil {
		h.logger.Errorw("failed to get all employees", "error", err)
		protocol.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	protocol.WriteJSON(w, http.StatusOK, employees)
}

// GetEmployeeById godoc
//
// @Summary Get Employee By ID
// @Description Get employee By ID
// @Tags employees
// @Accept json
// @Produce json
// @Param employeeId path int true "Employee ID"
// @Success 200 {object}  employeeEntity.Employee
// @Failure 500 {object} map[string]string	"Internal server error"
// @Failure 404 {object} map[string]string	"not found"
// @Router /employees/{employeeId} [get]
func(h *HttpHandler) GetById(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	idStr := chi.URLParam(r, "employeeId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		protocol.WriteJSONError(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	employee, err := h.employeeService.GetById(ctx, id)
	if err != nil{
		switch {
		case errors.Is(err, repository.ErrNotFound):
			protocol.WriteJSONError(w, http.StatusNotFound, "employee not found")
		default:
			h.logger.Errorw("failed to get employee", "error", err, "id", id)
			protocol.WriteJSONError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	protocol.WriteJSON(w, http.StatusOK, employee)
}

// CreateEmployee godoc
// @Summary Create new employee
// @Description Create a new employeee
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body employeeEntity.Employee true "Employee data"
// @Success 201 {object} employeeEntity.Employee
// @Failure 500 {object} map[string]string	"Internal server error"
// @Failure 404 {object} map[string]string	"not found"
// @Router /employees [post]
func (h *HttpHandler) Create(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	var emp employeeEntity.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		protocol.WriteJSONError(w, http.StatusBadRequest, "invalid request bdoy")
		return
	}

	if err := h.employeeService.Create(ctx, &emp); err != nil {
		switch{
		case errors.Is(err, repository.ErrUniqueViolation):
			protocol.WriteJSONError(w, http.StatusConflict, "email already exists")
		case errors.Is(err, repository.ErrNullEmail):
			protocol.WriteJSONError(w, http.StatusConflict, "email cannot be null")
		case errors.Is(err, repository.ErrNullOrNegSalary):
			protocol.WriteJSONError(w, http.StatusConflict, "salary cannot be null or negative")
		default:
			h.logger.Errorw("failed to createa", "error", err)
			protocol.WriteJSONError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	protocol.WriteJSON(w, http.StatusOK, emp)
}

// UpdateEmployee godoc
// @Summary Update employee
// @Description Update an employee
// @Tags employees
// @Accept json
// @Produce json
// @Param employeeId path int true "Employee ID"
// @Param employee body employeeEntity.Employee true "Employee data"
// @Success 200 {object} employeeEntity.Employee
// @Failure 404 {object} map[string]string	"not found"
// @Failure 500 {object} map[string]string	"Internal server error"
// @Router /employees/{employeeId} [put]
func (h *HttpHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "employeeId")
	id, err := strconv.ParseInt(idStr,10,64)
	if err != nil {
		protocol.WriteJSONError(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	var emp employeeEntity.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		protocol.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	emp.ID = id

	if err := h.employeeService.Update(ctx, &emp); err != nil {
		switch{
		case errors.Is(err, repository.ErrNotFound):
			protocol.WriteJSONError(w, http.StatusNotFound, "employee not found")
		case errors.Is(err, repository.ErrUniqueViolation):
			protocol.WriteJSONError(w, http.StatusConflict, "email already exists")
		case errors.Is(err, repository.ErrNullOrNegSalary):
			protocol.WriteJSONError(w, http.StatusConflict, "salary cannot be null or negative")
		default:
			h.logger.Errorw("failed to update", "error", err, "id", id)
			protocol.WriteJSONError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	protocol.WriteJSON(w, http.StatusOK, emp)
}

// DeleteEmnployee godoc
// @Summary delete employee
// @Description delete an empkloye
// @Tags employees
// @Accept json
// @Produce json
// @Param employeeId path int true "Employee ID"
// @Success 200 {object} employeeEntity.Employee
// @Failure 404 {object} map[string]string	"not found"
// @Failure 500 {object} map[string]string	"Internal server error"
// @Router /employees/{employeeId} [delete]
func (h *HttpHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "employeeId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		protocol.WriteJSONError(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	if err := h.employeeService.Delete(ctx, id); err != nil {
		switch{
		case errors.Is(err, repository.ErrNotFound):
			protocol.WriteJSONError(w, http.StatusNotFound, "employee not found")
		default:
			h.logger.Errorw("failed to delete", "error", err, "id", id)
			protocol.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	protocol.WriteJSON(w, http.StatusOK, "deleted successfully")
}
