package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/hanafi407/bookings/pkg/config"
	"github.com/hanafi407/bookings/pkg/handlers"
	"github.com/hanafi407/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.Repo.NewRepository(&app)
	handlers.Repo.NewHandler(repo)
	render.NewTemplate(&app)
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	serv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serv.ListenAndServe()
	log.Fatal(err)
}
