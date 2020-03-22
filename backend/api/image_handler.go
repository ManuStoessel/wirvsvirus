package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ManuStoessel/wirvsvirus/backend/entity"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func getImage(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Image{}
		data = data.Read(id)
		if data != nil {
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"image": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal image data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("Image found.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find image.")
}

func updateImage(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Image{}
		data = data.Read(id)
		if data != nil {
			imageToBeUpdated := entity.Image{}
			err := json.NewDecoder(r.Body).Decode(&imageToBeUpdated)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"error": "could not parse body as image"}`))
				log.WithFields(log.Fields{
					"body": fmt.Sprintf("%+v", r.Body),
				}).Error("Unable to unmarshal body as image.")
				return
			}

			imageToBeUpdated.ID = id
			imageToBeUpdated.Update()

			responseBody, err := json.Marshal(imageToBeUpdated)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"image": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal image data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Trace("image updated.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find Image.")
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if id, ok := queries["id"]; ok {
		data := &entity.Image{}
		data = data.Read(id)
		if data != nil {
			data.Delete()
			responseBody, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				log.WithFields(log.Fields{
					"image": fmt.Sprintf("%+v", data),
				}).Error("Unable to marshal image data.")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
			log.WithFields(log.Fields{
				"id": id,
			}).Debug("Image deleted.")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
	log.WithFields(log.Fields{
		"queries": fmt.Sprintf("%+v", queries),
	}).Error("Unable to find image.")
}

func createImage(w http.ResponseWriter, r *http.Request) {
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

	imageToBeCreated := entity.Image{}

	err = json.NewDecoder(r.Body).Decode(&imageToBeCreated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "could not parse body as image"}`))
		log.WithFields(log.Fields{
			"body": fmt.Sprintf("%+v", r.Body),
		}).Error("Unable to unmarshal body as image.")
		return
	}

	imageToBeCreated.Create()

	response, err := json.Marshal(imageToBeCreated)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "could not marshal image"}`))
		log.WithFields(log.Fields{
			"image": fmt.Sprintf("%+v", imageToBeCreated),
		}).Error("Unable to marshal image as body.")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	log.WithFields(log.Fields{
		"image": fmt.Sprintf("%+v", imageToBeCreated),
	}).Debug("Image created.")
	return

}

func listImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	imageList := ImageList{}
	image := entity.Image{}

	imageList.Images = image.ListAll()
	imageList.Count = len(imageList.Images)

	responseBody, err := json.Marshal(imageList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		log.WithFields(log.Fields{
			"imagelist": fmt.Sprintf("%+v", imageList),
		}).Error("Unable to marshal imagelist data.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	log.WithFields(log.Fields{
		"listlength": fmt.Sprintf("%+v", imageList.Count),
	}).Trace("Imagelist returned.")
	return
}
