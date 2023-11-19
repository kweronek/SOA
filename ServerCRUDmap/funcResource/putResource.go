package funcResource

import (
	"crud/modelResource"
	"crud/serviceGlobals"
	"encoding/json"
	"log"
	"net/http"
)

func PutResource(ID int, putJSON string) (httpStatus int) {
	var putResource modelResource.Resource
	if json.Unmarshal([]byte(putJSON), &putResource) != nil {
		log.Println("** Data not valid")
		return http.StatusPartialContent
	}
	modelResource.MMyResources[ID] = putResource
	if ID > serviceGlobals.RecCnt.Count {
		serviceGlobals.RecCnt.SetValue(ID)
	}
	return http.StatusCreated
}
