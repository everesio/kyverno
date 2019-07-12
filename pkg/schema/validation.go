package schema

import (
	"github.com/amzuko/gojsonschema"
	client "github.com/nirmata/kyverno/pkg/dclient"
	"github.com/nirmata/kyverno/pkg/schema/utils"
)

type Validator struct {
	cache Interface
}

func NewValidator(client *client.Client) Interface {
	initializerLoader()
	return &Validator{
		cache: NewClusterCacheFactory(client)}
}

func (v *Validator) validate(document []byte) (bool, error) {
	return v.cache.validate(document)
}
func initializerLoader() {
	// Without forcing these types the schema fails to load
	// Need to Work out proper handling for these types
	gojsonschema.FormatCheckers.Add("int64", utils.ValidFormat{})
	gojsonschema.FormatCheckers.Add("byte", utils.ValidFormat{})
	gojsonschema.FormatCheckers.Add("int32", utils.ValidFormat{})
	gojsonschema.FormatCheckers.Add("int-or-string", utils.ValidFormat{})
}
