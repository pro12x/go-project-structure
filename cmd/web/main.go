package main

import (
	"os"
)

func main() {
	// args := os.Args[1:] // Contrôle des arguments
	Run(os.Args[1:])
	// if len(args) == 0 {
	// 	var appConfig config.Config // Appel des éléments de configuration

	// 	templateCache, err := handlers.CreateTemplateCache() // Création de cache pour switcher entre les pages

	// 	if err != nil { // Retour en cas d'erreur
	// 		panic(err)
	// 	}

	// 	appConfig.TemplateCache = templateCache // Création du cache de la vue
	// 	appConfig.Port = ":1112"                // Configuration du port
	// 	appConfig.StaticDir = "assets"          // Configuration des fichiers statiques
	// 	handlers.CreateTemplates(&appConfig)    // Création de la vue

	// 	// Passage des fichiers statiques
	// 	statics := http.FileServer(http.Dir(appConfig.StaticDir))
	// 	http.Handle("/static/", http.StripPrefix("/static/", statics))

	// 	// Passages du gestionnaires des templates
	// 	http.HandleFunc("/", handlers.Home)

	// 	// Lancement du serveur d'application
	// 	fmt.Println("Server started on port http://localhost" + appConfig.Port)
	// 	http.ListenAndServe(appConfig.Port, nil)
	// 	os.Exit(0) // Sortie du programme
	// }
}
