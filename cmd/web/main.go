package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/marquesdeard/go-bookings/internal/config"
	"github.com/marquesdeard/go-bookings/internal/handlers"
	"github.com/marquesdeard/go-bookings/internal/helpers"
	"github.com/marquesdeard/go-bookings/internal/models"
	"github.com/marquesdeard/go-bookings/internal/render"
)

var port = ":8080"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting Application on port %s", port))

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() error {
	// what am i goint to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Session setup
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	// in production make it true
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(("cannot create template cache"))
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
