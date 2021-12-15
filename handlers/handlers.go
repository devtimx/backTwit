package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/devtimx/backTwit/middlew"
	"github.com/devtimx/backTwit/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routes.GraboTweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlew.ChequeoBD(middlew.ValidoJWT(routes.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middlew.ChequeoBD(middlew.ValidoJWT(routes.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subiravatar", middlew.ChequeoBD(middlew.ValidoJWT(routes.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirbanner", middlew.ChequeoBD(middlew.ValidoJWT(routes.SubirBanner))).Methods("POST")
	router.HandleFunc("/obteneravatar", middlew.ChequeoBD(middlew.ValidoJWT(routes.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/obtenerbanner", middlew.ChequeoBD(middlew.ValidoJWT(routes.ObtenerBanber))).Methods("GET")
	router.HandleFunc("/altarelacion", middlew.ChequeoBD(middlew.ValidoJWT(routes.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajarelacion", middlew.ChequeoBD(middlew.ValidoJWT(routes.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultarelacion", middlew.ChequeoBD(middlew.ValidoJWT(routes.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listarusuarios", middlew.ChequeoBD(middlew.ValidoJWT(routes.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leotweetsseguidores", middlew.ChequeoBD(middlew.ValidoJWT(routes.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
