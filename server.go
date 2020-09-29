package main

import (
	"fmt"
	"net/http"

	controller "glue/glue-backend-golang/controller"
	router "glue/glue-backend-golang/http"
	repository "glue/glue-backend-golang/repository"
	service "glue/glue-backend-golang/service"
)

var (
	placeRepository repository.PlaceRepository = repository.NewFirestoreRepository()
	placeService    service.PlaceService       = service.NewPlacesService(placeRepository)
	placeController controller.PlaceController = controller.NewPlaceController(placeService)
	httpRouter      router.Router              = router.NewChiRouter()
)

func main() {

	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/places", placeController.GetPlaces)
	httpRouter.POST("/places", placeController.AddPlace)

	httpRouter.SERVE(port)

}
