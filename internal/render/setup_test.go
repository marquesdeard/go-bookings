package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/marquesdeard/go-bookings/internal/config"
	"github.com/marquesdeard/go-bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// what am i goint to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// app.InfoLog = infoLog

	// errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// app.ErrorLog = errorLog

	// Session setup
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}
func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)

	return length, nil
}
