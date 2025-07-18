package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
    "github.com/gorilla/mux"
)

type WebHandler struct {
    templates *template.Template
}

func NewWebHandler() *WebHandler {
    // Load all HTML templates
    templates := template.Must(template.ParseGlob("templates/*.html"))
    
    return &WebHandler{
        templates: templates,
    }
}

// Serve static files (CSS, JS, images)
func (h *WebHandler) ServeStatic(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := vars["filename"]
    
    // Security: prevent directory traversal
    if filepath.Clean(filename) != filename {
        http.Error(w, "Invalid file path", http.StatusBadRequest)
        return
    }
    
    http.ServeFile(w, r, "static/"+filename)
}

// Home page handler
func (h *WebHandler) HomePage(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Title   string
        Message string
    }{
        Title:   "Welcome",
        Message: "Hello from Go Backend!",
    }
    
    err := h.templates.ExecuteTemplate(w, "index.html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Dashboard handler
func (h *WebHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
    // Check authentication here
    
    err := h.templates.ExecuteTemplate(w, "dashboard.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Register web routes
func (h *WebHandler) RegisterRoutes(r *mux.Router) {
    // Static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", 
        http.FileServer(http.Dir("static/"))))
    
    // Web pages
    r.HandleFunc("/", h.HomePage).Methods("GET")
    r.HandleFunc("/dashboard", h.Dashboard).Methods("GET")
}