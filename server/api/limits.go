package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerLimitsRoutes(r *mux.Router) {
	// limits
	r.HandleFunc("/limits", a.sessionRequired(a.handleCloudLimits)).Methods("GET")
}

func (a *API) handleCloudLimits(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /limits cloudLimits
	//
	// Fetches the cloud limits of the server.
	//
	// ---
	// produces:
	// - application/json
	// security:
	// - BearerAuth: []
	// responses:
	//   '200':
	//     description: success
	//     schema:
	//         "$ref": "#/definitions/BoardsCloudLimits"
	//   default:
	//     description: internal error
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"

	boardsCloudLimits, err := a.app.GetBoardsCloudLimits()
	if err != nil {
		a.errorResponse(w, r, err)
		return
	}

	data, err := json.Marshal(boardsCloudLimits)
	if err != nil {
		a.errorResponse(w, r, err)
		return
	}

	jsonBytesResponse(w, http.StatusOK, data)
}
