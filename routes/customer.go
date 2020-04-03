package routes

import (
	"fmt"
	"net/http"

	"../controllers"

	"github.com/gorilla/mux"
)

var SetupServer = func(appPort string) {
	r := mux.NewRouter()
	/* Users Routes */
	r.HandleFunc("/customers", controllers.InsertCustomer).Methods("POST")
	r.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", controllers.GetCustomerById).Methods("GET")
	r.HandleFunc("/customers/{name}/list", controllers.GetCustomersByName).Methods("GET")
	r.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")

	err := http.ListenAndServe(":"+appPort, r)
	if err != nil {
		fmt.Print(err)
	}
}
