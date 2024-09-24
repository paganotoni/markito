package health

import (
	"net/http"

	"github.com/leapkit/leapkit/core/db"
	"github.com/leapkit/leapkit/core/server"
)

func Check(db db.ConnFn) func(w http.ResponseWriter, r *http.Request) {
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
