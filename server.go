package main

import (
	"fmt"
	"net/http"

	"glue/glue-backend-golang/controller"
	router "glue/glue-backend-golang/http"
	"glue/glue-backend-golang/repository"
	"glue/glue-backend-golang/service"
)

var (
	placeRepository = repository.NewFirestoreRepository()
	placeService    = service.NewPlaceService(placeRepository)
	placeController = controller.NewPlaceController(placeService)

	ISpaceRepository = repository.NewISpaceRepository()
	ISpaceService    = service.SpacesService(ISpaceRepository)
	ISpaceController = controller.NewISpaceController(ISpaceService)

	ITicketRepository = repository.NewITicketRepository()
	ITicketService = service.TicketService(ITicketRepository)
	ITicketController = controller.NewITicketController(ITicketService)

	httpRouter = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/places", placeController.GetPlaces)
	httpRouter.POST("/places", placeController.AddPlace)

	httpRouter.GET("/spaces", ISpaceController.ListSpaces)
	httpRouter.GET("/spaceByID", ISpaceController.GetSpaceByID)

	httpRouter.GET("/availableTickets", ITicketController.ListAllAvailableTickets)

	httpRouter.SERVE(port)

}
