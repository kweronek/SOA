package funcResource

import "crud/modelResource"

func DeleteResource(ID int) {
	delete(modelResource.MMyResources, ID)
}
