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
	placeService    service.PlaceService       = service.NewPlaceService(placeRepository)
	placeController controller.PlaceController = controller.NewPlaceController(placeService)

	ispaceRepository repository.ISpaceRepository = repository.NewISpaceRepository()
	ispaceService    service.ISpaceService       = service.ListSpacesService(ispaceRepository)
	ispaceController controller.ISpaceController = controller.NewISpaceController(ispaceService)

	httpRouter router.Router = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/places", placeController.GetPlaces)
	httpRouter.POST("/places", placeController.AddPlace)

	httpRouter.GET("/spaces", ispaceController.ListSpaces)

	httpRouter.SERVE(port)

}
