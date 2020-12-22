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
	ITicketService    = service.TicketService(ITicketRepository)
	ITicketController = controller.NewITicketController(ITicketService)

	httpRouter = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	//Places routes
	httpRouter.GET("/places", placeController.FireStoreGetPlaces)
	httpRouter.POST("/places", placeController.FireStoreAddPlace)

	//ISpace routes
	httpRouter.GET("/spaces", ISpaceController.PsqlListSpaces)
	httpRouter.GET("/spaceID", ISpaceController.PsqlGetSpaceByID)
	httpRouter.GET("/spaces-tickets", ISpaceController.PsqlListSpacesWithTickets)
	httpRouter.POST("/newSpace", ISpaceController.PsqlCreateNewSpace)

	//ITicket routes
	httpRouter.GET("/availableTickets", ITicketController.FireStoreListAllAvailableTickets)
	httpRouter.POST("/newTicket", ITicketController.PsqlCreateNewTicket)

	httpRouter.SERVE(port)

}
