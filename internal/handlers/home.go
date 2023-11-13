package handlers

import (
	"bim/internal/metier/models"
	"bim/internal/pkg"
	"fmt"
	"net/http"
)

// Gestionnaire de la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	// Vérification de la méthode
	if r.Method == http.MethodGet {
		// Contrôle des URLs
		if r.URL.Path != "/" {
			err := pkg.Error(http.StatusNotFound)
			renderTemplates(w, "error", &models.TemplateData{Error: models.ErrorData{Code: err["code"], Message: err["msg"]}})
			fmt.Println(http.StatusNotFound, "Not Found")
			return
		}
		names := make(map[string]string)
		names["owner"] = "Janel"
		names["name"] = "Proverbes"
		renderTemplates(w, "home", &models.TemplateData{StringData: names})
		fmt.Println(http.StatusOK, "OK")
	}
}
