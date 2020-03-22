package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ManuStoessel/wirvsvirus/backend/entity"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func getDonation(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Donation{}
		data = data.Read(id)
		if data != nil {
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"user": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal donation data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("Donation found.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find donation.")
}

func updateDonation(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Donation{}
		data = data.Read(id)
		if data != nil {
			donationToBeUpdated := entity.Donation{ID: id}
			err := json.NewDecoder(r.Body).Decode(&donationToBeUpdated)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"error": "could not parse body as donation"}`))
				log.WithFields(log.Fields{
					"body": fmt.Sprintf("%+v", r.Body),
				}).Error("Unable to unmarshal body as donation.")
				return
			}

			donationToBeUpdated.ID = id
			donationToBeUpdated.Update()

			responseBody, err := json.Marshal(donationToBeUpdated)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"donation": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal donation data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("Donation updated.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find user.")
}

func deleteDonation(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Donation{}
		data = data.Read(id)
		if data != nil {
			data.Delete()
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"donation": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal donation data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Debug("Donation deleted.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find donation.")
}

func createDonation(w http.ResponseWriter, r *http.Request) {
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

	donationToBeCreated := entity.Donation{}

	err = json.NewDecoder(r.Body).Decode(&donationToBeCreated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "could not parse body as donation"}`))
		log.WithFields(log.Fields{
			"body": fmt.Sprintf("%+v", r.Body),
		}).Error("Unable to unmarshal body as donation.")
		return
	}

	donationToBeCreated.Create()

	response, err := json.Marshal(donationToBeCreated)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "could not marshal donation"}`))
		log.WithFields(log.Fields{
			"donation": fmt.Sprintf("%+v", donationToBeCreated),
		}).Error("Unable to marshal donation as body.")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	log.WithFields(log.Fields{
		"donation": fmt.Sprintf("%+v", donationToBeCreated),
	}).Debug("Donation created.")
	return

}

func listDonations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	donationList := DonationList{}
	donation := entity.Donation{}

	donationList.Donations = donation.ListAll()
	donationList.Count = len(donationList.Donations)

	responseBody, err := json.Marshal(donationList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		log.WithFields(log.Fields{
			"donationlist": fmt.Sprintf("%+v", donationList),
		}).Error("Unable to marshal donationlist data.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	log.WithFields(log.Fields{
		"listlength": fmt.Sprintf("%+v", donationList.Count),
	}).Trace("Donationlist returned.")
	return
}
