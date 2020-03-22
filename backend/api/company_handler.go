package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ManuStoessel/wirvsvirus/backend/entity"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func getCompany(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Company{}
		data = data.Read(id)
		if data != nil {
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"company": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal company data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("Company found.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find company.")
}

func updateCompany(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Company{}
		data = data.Read(id)
		if data != nil {
			companyToBeUpdated := entity.Company{}
			err := json.NewDecoder(r.Body).Decode(&companyToBeUpdated)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"error": "could not parse body as company"}`))
				log.WithFields(log.Fields{
					"body": fmt.Sprintf("%+v", r.Body),
				}).Error("Unable to unmarshal body as company.")
				return
			}

			companyToBeUpdated.ID = id
			companyToBeUpdated.Update()

			responseBody, err := json.Marshal(companyToBeUpdated)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"company": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal company data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("Company updated.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find company.")
}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Company{}
		data = data.Read(id)
		if data != nil {
			data.Delete()
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"company": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal company data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Debug("Company deleted.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find company.")
}

func createCompany(w http.ResponseWriter, r *http.Request) {
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

	companyToBeCreated := entity.Company{}

	err = json.NewDecoder(r.Body).Decode(&companyToBeCreated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "could not parse body as company"}`))
		log.WithFields(log.Fields{
			"body": fmt.Sprintf("%+v", r.Body),
		}).Error("Unable to unmarshal body as company.")
		return
	}

	companyToBeCreated.Create()

	response, err := json.Marshal(companyToBeCreated)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "could not marshal company"}`))
		log.WithFields(log.Fields{
			"company": fmt.Sprintf("%+v", companyToBeCreated),
		}).Error("Unable to marshal company as body.")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	log.WithFields(log.Fields{
		"company": fmt.Sprintf("%+v", companyToBeCreated),
	}).Debug("Company created.")
	return

}

func listCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	companyList := CompanyList{}
	company := entity.Company{}

	companyList.Companies = company.ListAll()
	companyList.Count = len(companyList.Companies)

	responseBody, err := json.Marshal(companyList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		log.WithFields(log.Fields{
			"companylist": fmt.Sprintf("%+v", companyList),
		}).Error("Unable to marshal companylist data.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	log.WithFields(log.Fields{
		"listlength": fmt.Sprintf("%+v", companyList.Count),
	}).Trace("Companylist returned.")
	return
}

func listCompaniesByTown(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if name, ok := queries["name"]; ok {

		name = strings.ToLower(strings.TrimSpace(name))

		companyList := CompanyList{}
		company := entity.Company{}

		companies := company.ListAll()

		for _, c := range companies {
			if name == strings.ToLower(strings.TrimSpace(c.Town)) {
				companyList.Companies = append(companyList.Companies, c)
				companyList.Count++
			}
		}

		responseBody, err := json.Marshal(companyList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			log.WithFields(log.Fields{
				"companylist": fmt.Sprintf("%+v", companyList),
			}).Error("Unable to marshal companylist data.")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
		log.WithFields(log.Fields{
			"listlength": fmt.Sprintf("%+v", companyList.Count),
		}).Trace("Companylist returned.")
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find town.")
}
