package funcResource

import (
	"crud/modelResource"
)

func CheckResource(ID int) (exists bool) {
	_, exists = modelResource.MMyResources[ID]
	return exists
}
