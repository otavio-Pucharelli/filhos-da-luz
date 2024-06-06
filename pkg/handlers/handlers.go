package handlers

import (
	"fmt"
	"net/http"

	"github.com/otavio-Pucharelli/filhos-da-luz/pkg/config"
	"github.com/otavio-Pucharelli/filhos-da-luz/pkg/models"
	"github.com/otavio-Pucharelli/filhos-da-luz/pkg/render"
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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tpl.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tpl.html", &models.TemplateData{})
}

// Resident is the handler for the about page
func (m *Repository) Resident(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "resident.page.tpl.html", &models.TemplateData{})
}

// PostResident is the handler for the about page
func (m *Repository) PostResident(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	address := r.Form.Get("address")
	city := r.Form.Get("city")
	state := r.Form.Get("state")
	zip := r.Form.Get("zip")

	w.Write([]byte(fmt.Sprintf("Name: %s\nEmail: %s\nPhone: %s\nAddress: %s\nCity: %s\nState: %s\nZip: %s", name, email, phone, address, city, state, zip)))
}
