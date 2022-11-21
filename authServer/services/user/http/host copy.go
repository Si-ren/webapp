package http

// import (
// 	"cmdb/services/user"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	query := user.NewQueryHostRequestFromHTTP(r)
// 	set, err := h.service.QueryHost(r.Context(), query)
// 	if err != nil {
// 		response.Failed(w, err)
// 		return
// 	}
// 	response.Success(w, set)
// }

// func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	ins := user.NewDefaultHost()
// 	if err := request.GetDataFromRequest(r, ins); err != nil {
// 		response.Failed(w, err)
// 		return
// 	}

// 	ins, err := h.service.SaveHost(r.Context(), ins)
// 	if err != nil {
// 		response.Failed(w, err)
// 		return
// 	}

// 	response.Success(w, ins)
// }

// func (h *handler) DescribeHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	req := user.NewDescribeHostRequestWithID(ps.ByName("id"))
// 	set, err := h.service.DescribeHost(r.Context(), req)
// 	if err != nil {
// 		response.Failed(w, err)
// 		return
// 	}
// 	response.Success(w, set)
// }

// func (h *handler) DeleteHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	req := user.NewDeleteHostRequestWithID(ps.ByName("id"))
// 	set, err := h.service.DeleteHost(r.Context(), req)
// 	if err != nil {
// 		response.Failed(w, err)
// 		return
// 	}
// 	response.Success(w, set)
// }

// func (h *handler) PutHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	req := user.NewUpdateHostRequest(ps.ByName("id"))

// 	if err := request.GetDataFromRequest(r, req.UpdateHostData); err != nil {
// 		response.Failed(w, err)
// 		return
// 	}

// 	ins, err := h.service.UpdateHost(r.Context(), req)
// 	if err != nil {
// 		response.Failed(w, err)
// 		return
// 	}

// 	response.Success(w, ins)
// }

// func (h *handler) PatchHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	req := user.NewUpdateHostRequest(ps.ByName("id"))
// 	req.UpdateMode = user.PATCH

// 	if err := request.GetDataFromRequest(r, req.UpdateHostData); err != nil {
// 		response.Failed(w, err)
// 		return
// 	}

// 	ins, err := h.service.UpdateHost(r.Context(), req)
// 	if err != nil {
// 		response.Failed(w, err)
// 		return
// 	}

// 	response.Success(w, ins)
// }
