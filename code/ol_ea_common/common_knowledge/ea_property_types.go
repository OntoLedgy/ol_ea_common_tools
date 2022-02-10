package common_knowledge

type EaPropertyTypes int //Enum):

const (
	CONNECTOR_PROPERTY EaPropertyTypes = iota
	ELEMENT_PROPERTY
	ATTRIBUTE_PROPERTY
)

func (eaPropertyTypes EaPropertyTypes) get_type_name() string {

	switch eaPropertyTypes {
	case CONNECTOR_PROPERTY:
		return "connector property"
	case ELEMENT_PROPERTY:
		return "element property"
	case ATTRIBUTE_PROPERTY:
		return "attribute property"
	}
	return "unknown"
}