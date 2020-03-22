package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ManuStoessel/wirvsvirus/backend/entity"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.User{}
		data = data.Read(id)
		if data != nil {
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"user": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal user data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("User found.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find user.")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.User{}
		data = data.Read(id)
		if data != nil {
			userToBeUpdated := User{}
			err := json.NewDecoder(r.Body).Decode(&userToBeUpdated)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"error": "could not parse body as user"}`))
				log.WithFields(log.Fields{
					"body": fmt.Sprintf("%+v", r.Body),
				}).Error("Unable to unmarshal body as user.")
				return
			}

			userUpdated := &entity.User{
				ID:       userToBeUpdated.ID,
				Username: userToBeUpdated.Username,
				Email:    userToBeUpdated.Email,
			}

			userUpdated.Update()

			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"user": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal user data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("User updated.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find user.")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.User{}
		data = data.Read(id)
		if data != nil {
			data.Delete()
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"user": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal user data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Debug("User deleted.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find user.")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		log.WithFields(log.Fields{
			"body": fmt.Sprintf("%+v", r.Body),
		}).Error("Unable to parse body.")
		return
	}

	userToBeCreated := User{}

	err = json.NewDecoder(r.Body).Decode(&userToBeCreated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "could not parse body as user"}`))
		log.WithFields(log.Fields{
			"body": fmt.Sprintf("%+v", r.Body),
		}).Error("Unable to unmarshal body as user.")
		return
	}

	userCreated := &entity.User{
		Username: userToBeCreated.Username,
		Email:    userToBeCreated.Email,
	}

	userCreated.Create()

	response, err := json.Marshal(User{
		ID:       userCreated.ID,
		Email:    userCreated.Email,
		Username: userCreated.Username,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "could not marshal user"}`))
		log.WithFields(log.Fields{
			"user": fmt.Sprintf("%+v", userCreated),
		}).Error("Unable to marshal user as body.")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	log.WithFields(log.Fields{
		"user": fmt.Sprintf("%+v", userCreated),
	}).Debug("User created.")
	return

}
