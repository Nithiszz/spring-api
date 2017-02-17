package app

import (
	"net/http"

	"github.com/Nithiszz/sprint-api/pkg/api"
)

// RegisterBookService register book service to http server
func RegisterBookService(mux *http.ServeMux, s api.BookService) {
	sv := "/book"

	mux.HandleFunc(sv+".create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			renderString(w, http.StatusMethodNotAllowed, "Method not Allowed")
			return
		}

		var req api.BookRequest
		err := bindJSON(r.Body, &req)
		if err != nil {
			renderString(w, http.StatusBadRequest, "Bad Request")
			return
		}
		res, err := s.CreateBook(r.Context(), &req)
		if err != nil {
			renderError(w, err)
			return
		}

		renderJSON(w, http.StatusOK, res)
	})

	mux.HandleFunc(sv+".list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			renderString(w, http.StatusMethodNotAllowed, "Method not Allowed")
			return
		}

		res, err := s.ListBooks(r.Context())
		if err != nil {
			renderError(w, err)

			return
		}

		renderJSON(w, http.StatusOK, res)
	})

	mux.HandleFunc(sv+".get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			renderString(w, http.StatusMethodNotAllowed, "Method not Allowed")
			return
		}

		var req api.BookRequest
		err := bindJSON(r.Body, &req)
		if err != nil {
			renderString(w, http.StatusBadRequest, "Bad Request")
			return
		}
		res, err := s.GetBook(r.Context(), &req)
		if err != nil {
			renderError(w, err)
			return
		}

		renderJSON(w, http.StatusOK, res)
	})

	mux.HandleFunc(sv+".update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			renderString(w, http.StatusMethodNotAllowed, "Method not Allowed")
			return
		}

		var req api.BookRequest
		err := bindJSON(r.Body, &req)
		if err != nil {
			renderString(w, http.StatusBadRequest, "Bad Request")
			return
		}
		res, err := s.UpdateBook(r.Context(), &req)
		if err != nil {
			renderError(w, err)

			return
		}

		renderJSON(w, http.StatusOK, res)
	})

	mux.HandleFunc(sv+".delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			renderString(w, http.StatusMethodNotAllowed, "Method not Allowed")
			return
		}

		var req api.BookRequest
		err := bindJSON(r.Body, &req)
		if err != nil {
			renderString(w, http.StatusBadRequest, "Bad Request")
			return
		}
		err = s.DeleteBook(r.Context(), &req)
		if err != nil {
			renderError(w, err)
			return
		}

		renderNoContent(w)
	})
}
