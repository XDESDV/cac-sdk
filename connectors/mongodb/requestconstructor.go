package mongodb

import (
	"strconv"
	"strings"

	"github.com/xdesdv/cac-sdk/queries"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SelectConstructeur : Construit la requete du select
func SelectConstructeur(params queries.QueryParams) bson.M {

	var findQuery bson.M
	findQuery = make(map[string]interface{})
	if params.TestDeleted {
		findQuery = Alive(findQuery)
	}
	findQuery = FilterConstructeur(params, findQuery)
	findQuery = FilterLikeConstructeur(params, findQuery)
	return findQuery

}

// FilterConstructeur : Préparation du where
func FilterConstructeur(params queries.QueryParams, fq bson.M) bson.M {
	var (
		key      string
		value    interface{}
		operator string
	)

	query := make(bson.M)

	if len(params.FilterClause) > 0 {
		for _, param := range params.FilterClause {

			// Récupération du filtre demandé
			filter := strings.Split(param, ",")

			// Operator retrieve
			if len(filter) == 2 {
				operator = "="
			} else {
				operator = filter[2]
			}

			key = filter[0]

			// conversion of the value into its data type
			if vBool, sErr := strconv.ParseBool(filter[1]); sErr != nil {
				if vInt, sErr := strconv.ParseInt(filter[1], 10, 64); sErr != nil {
					if vFloat, sErr := strconv.ParseFloat(filter[1], 64); sErr != nil {
						value = filter[1]
					} else {
						value = vFloat
					}
				} else {
					value = vInt
				}
			} else {
				value = vBool
			}

			// operator management
			switch operator {
			case "=":
				fq[key] = value
			case ">":
				query["$gt"] = filter[1]
				fq[key] = query
			case ">=":
				query["$gte"] = filter[1]
				fq[key] = query
			case "<":
				query["$lt"] = filter[1]
				fq[key] = query
			case "<=":
				query["$lte"] = filter[1]
				fq[key] = query
			case "!=":
				query["$ne"] = filter[1]
				fq[key] = query
			default:
				fq[key] = filter[1]
			}

		}
	}
	return fq
}

// FilterLikeConstructeur : préparation du where pour un like
func FilterLikeConstructeur(params queries.QueryParams, fq bson.M) bson.M {

	if len(params.FilterLikeClause) > 0 {
		for _, param := range params.FilterLikeClause {
			// Récupération du filtre like demandé
			filterLike := strings.Split(param, ",")
			fq[filterLike[0]] = primitive.Regex{Pattern: "^.*" + filterLike[1] + ".*$", Options: "i"}
			//fq = append(fq, bson.M{filterLike[0]: primitive.Regex{Pattern: "^.*" + filterLike[1] + ".*$", Options: "i"}})
		}
	}
	return fq
}

// Alive allows to set the element to deleted (Archive)
func Alive(fq bson.M) bson.M {
	fq["deleted"] = false
	return fq
}
