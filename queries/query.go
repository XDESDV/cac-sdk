package queries

import (
	"cac-sdk/functions"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Headers : Headers spécifique
type Headers struct {
	TimeZone string
}

// QueryParams : Parametres requetes HTTP
type QueryParams struct {
	ID               string
	Headers          Headers
	Path             string
	Body             []byte
	View             string
	FilterClause     []string
	FilterLikeClause []string
	SortClause       []string
	Offset           int
	Count            int
	Export           string
	GroupBy          string
	Columns          []string
	SearchClause     []string
	Collection       string
	TestDeleted      bool
}

// HTTPParser : Parser pour QueryParams
type HTTPParser interface {
	Parse(*http.Request)
}

// Parse : QueryParams parser
func (q QueryParams) Parse(httpReq *http.Request) QueryParams {
	query := httpReq.URL.Query()
	vars := mux.Vars(httpReq)
	count, _ := strconv.Atoi(query.Get("count"))
	offset, _ := strconv.Atoi(query.Get("offset"))
	view := query.Get("view")
	groupby := query.Get("col")

	var headers Headers
	// Récupération des entêtes
	// ...
	// ...
	// ...

	path := httpReq.URL.Path
	// Récupération des éléments à rechercher
	search := []string{}
	if len(query.Get("search")) > 0 {
		search = strings.Split(query.Get("search"), "+")
		value := search[0]
		tabkeyword := strings.Split(value, " ")
		search = search[:0]
		for i := 0; i < len(tabkeyword); i++ {
			// on double les aspostrophes
			keyword := strings.Replace(tabkeyword[i], "'", "''", -1)
			search = append(search, keyword)
		}
	}

	// Récupération du tri
	tri := []string{}
	if len(query.Get("sort")) > 0 {
		tri = strings.Split(query.Get("sort"), ",")
	}

	body, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		log.Fatal(err)
	}

	filter := query["filter"]
	likeFilter := query["filter_like"]
	// Supression de tous les doublons possibles des requetes envoyées
	functions.RemoveDuplicate(&filter)
	functions.RemoveDuplicate(&likeFilter)
	functions.RemoveDuplicate(&tri)
	functions.RemoveDuplicate(&search)

	var params = QueryParams{
		ID:               vars["id"],
		FilterClause:     filter,
		FilterLikeClause: likeFilter,
		SortClause:       tri,
		Headers:          headers,
		Body:             body,
		Count:            count,
		Offset:           offset,
		GroupBy:          groupby,
		View:             view,
		SearchClause:     search,
		Path:             path,
		TestDeleted:      true,
	}

	return params
}
