package links

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"url-shortener/src/infrastructure/cache"
	"url-shortener/src/infrastructure/db"
)

func errorHandler(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Error: %s", message)
}
func GetShortened(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	req := &GetShortenedUrlRequest{}
	err := json.NewDecoder(request.Body).Decode(req)

	if err != nil {
		writer.WriteHeader(422)
		resp := GetShortenedUrlResponse{nil, nil, " Wrong Input."}

		err = json.NewEncoder(writer).Encode(&resp)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}

		return
	}

	log.Info("Shortening request for: ", req.OriginLink)
	if !isValidURL(req.OriginLink) {
		writer.WriteHeader(422)
		resp := GetShortenedUrlResponse{nil, nil, "This is not a URL!."}

		err = json.NewEncoder(writer).Encode(&resp)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}

		return
	}
	rep := NewLinkRepository(db.PostgresDBProvider, cache.RedisCacheProvider)
	qhandler := NewLinkQueryHandler(rep)
	queryReq := GetShortenedLinkQuery{OriginLink: req.OriginLink}
	result, err := qhandler.GetShortenedLink(request.Context(), queryReq)

	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		resp := GetShortenedUrlResponse{&result.ShortenedLink, &result.ExpiresAt, "OK!"}
		writer.WriteHeader(http.StatusOK)
		err = json.NewEncoder(writer).Encode(&resp)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
		return
	}
	chandler := NewLinkCommandHandler(rep)
	cmd := CreateShortenedUrlCommand{req.OriginLink}
	cmdResult, err := chandler.CreateShortenedUrl(request.Context(), cmd)
	if err != nil {
		log.Fatalln("Error occured while creating SHort form, ", err)
	}

	resp := &GetShortenedUrlResponse{&cmdResult.ShortenedUrl, &cmdResult.ExpiresAt, "OK!"}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(&resp)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
func GetOrigin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	req := &GetOriginUrlRequest{}
	err := json.NewDecoder(request.Body).Decode(req)

	if err != nil || req.ShortenedLink == "" {
		writer.WriteHeader(422)
		resp := GetShortenedUrlResponse{nil, nil, " Wrong Input."}

		err = json.NewEncoder(writer).Encode(&resp)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}

		return
	}

	rep := NewLinkRepository(db.PostgresDBProvider, cache.RedisCacheProvider)
	qhandler := NewLinkQueryHandler(rep)
	queryReq := GetOriginLinkQuery{ShortenedLink: req.ShortenedLink}
	result, err := qhandler.GetOriginLink(request.Context(), queryReq)

	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		resp := GetOriginUrlResponse{&result.OriginLink, &result.ExpiresAt, "OK!"}
		writer.WriteHeader(http.StatusOK)
		err = json.NewEncoder(writer).Encode(&resp)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
		return
	} else {
		writer.WriteHeader(404)
		resp := GetOriginUrlResponse{nil, nil, "Shortened URL does not exist!."}

		err = json.NewEncoder(writer).Encode(&resp)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}

		return

	}

}
func UrlRedirect(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	shortURL := mux.Vars(request)["shortURL"]
	if shortURL == "" {
		errorHandler(writer, http.StatusNotFound, "Short URL not found")
		return
	}
	rep := NewLinkRepository(db.PostgresDBProvider, cache.RedisCacheProvider)
	qhandler := NewLinkQueryHandler(rep)
	queryReq := GetOriginLinkQuery{ShortenedLink: shortURL}
	result, err := qhandler.GetOriginLink(request.Context(), queryReq)
	log.Info(result.OriginLink)
	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		//http.Redirect(writer, request, result.OriginLink, http.StatusMovedPermanently)
		return

	} else {
		errorHandler(writer, http.StatusNotFound, "Short URL not found")

	}
}
