package internal

import (
	"net/http"

	"go.leapkit.dev/core/db"
	"go.leapkit.dev/core/server"
)

func check(db db.ConnFn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := db()
		if err != nil {
			server.Error(w, err, http.StatusInternalServerError)
			return
		}

		err = db.Ping()
		if err != nil {
			server.Error(w, err, http.StatusInternalServerError)
			return
		}

		w.Write([]byte("OK"))
	}
}
