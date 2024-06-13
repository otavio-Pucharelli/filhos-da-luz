package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/otavio-Pucharelli/filhos-da-luz/internal/config"
	"github.com/otavio-Pucharelli/filhos-da-luz/internal/forms"
	"github.com/otavio-Pucharelli/filhos-da-luz/internal/helpers"
	"github.com/otavio-Pucharelli/filhos-da-luz/internal/models"
	"github.com/otavio-Pucharelli/filhos-da-luz/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tpl.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tpl.html", &models.TemplateData{})
}

// Resident is the handler for the about page
func (m *Repository) Resident(w http.ResponseWriter, r *http.Request) {
	var emptyResident models.Resident
	data := make(map[string]interface{})
	data["resident"] = emptyResident
	render.RenderTemplate(w, r, "resident.page.tpl.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// jsonResponse is a generic JSON response used by the API
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// PostResident is the handler for the about page
func (m *Repository) PostResident(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	resident := models.Resident{
		Name:    r.Form.Get("name"),
		Email:   r.Form.Get("email"),
		Phone:   r.Form.Get("phone"),
		Address: r.Form.Get("address"),
		City:    r.Form.Get("city"),
		State:   r.Form.Get("state"),
		Zip:     r.Form.Get("zip"),
	}

	// Validate the form
	form := forms.New(r.PostForm)
	form.Required("name", "email", "phone", "address", "city", "state", "zip")
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["resident"] = resident

		render.RenderTemplate(w, r, "resident.page.tpl.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	resp := jsonResponse{
		OK:      true,
		Message: "Resident saved",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Name: %s\nEmail: %s\nPhone: %s\nAddress: %s\nCity: %s\nState: %s\nZip: %s", resident.Name, resident.Email, resident.Email, resident.Address, resident.City, resident.State, resident.Zip)))
	w.Write([]byte(out))
}
