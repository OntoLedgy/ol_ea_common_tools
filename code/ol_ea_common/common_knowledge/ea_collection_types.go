package common_knowledge

type EaCollectionTypes int //CollectionTypes

const (
	COLLECTION_NOT_SET EaCollectionTypes = iota
	T_OBJECT
	T_CONNECTOR
	T_PACKAGE
	T_ATTRIBUTE
	T_XREF
	T_STEREOTYPES
	T_DIAGRAM
	T_DIAGRAMLINKS
	T_DIAGRAMOBJECTS
	T_OPERATION
	T_CONNECTORTYPES
	T_OBJECTTYPES
	T_DIAGRAMTYPES
	T_CARDINALITY
	EXTENDED_T_OBJECT
	EXTENDED_T_CONNECTOR
	EXTENDED_T_PACKAGE
	EXTENDED_T_ATTRIBUTE
	EXTENDED_T_OPERATION
	EXTENDED_T_XREF
	EXTENDED_T_STEREOTYPES
	EXTENDED_T_DIAGRAM
	EXTENDED_T_DIAGRAMLINKS
	EXTENDED_T_DIAGRAMOBJECTS
	EXTENDED_T_CONNECTORTYPES
	EXTENDED_T_OBJECTTYPES
	EXTENDED_T_DIAGRAMTYPES
	EXTENDED_T_CARDINALITY
	OBJECT_STEREOTYPES
	PACKAGE_HIERARCHY
)

func (eaCollectionTypes EaCollectionTypes) GetCollectionName() string {

	switch eaCollectionTypes {

	case T_OBJECT:
		return "t_object"
	case T_CONNECTOR:
		return "t_connector"
	case T_PACKAGE:
		return "t_package"
	case T_ATTRIBUTE:
		return "t_attribute"
	case T_OPERATION:
		return "t_operation"
	case T_XREF:
		return "t_xref"
	case T_STEREOTYPES:
		return "t_stereotypes"
	case T_DIAGRAM:
		return "t_diagram"
	case T_DIAGRAMLINKS:
		return "t_diagramlinks"
	case T_DIAGRAMOBJECTS:
		return "t_diagramobjects"
	case T_CONNECTORTYPES:
		return "t_connectortypes"
	case T_OBJECTTYPES:
		return "t_objecttypes"
	case T_DIAGRAMTYPES:
		return "t_diagramtypes"
	case T_CARDINALITY:
		return "t_cardinality"
	case EXTENDED_T_OBJECT:
		return "extended_t_object"
	case EXTENDED_T_CONNECTOR:
		return "extended_t_connector"
	case EXTENDED_T_PACKAGE:
		return "extended_t_package"
	case EXTENDED_T_ATTRIBUTE:
		return "extended_t_attribute"
	case EXTENDED_T_OPERATION:
		return "extended_t_operation"
	case EXTENDED_T_XREF:
		return "extended_t_xref"
	case EXTENDED_T_STEREOTYPES:
		return "extended_t_stereotypes"
	case EXTENDED_T_DIAGRAM:
		return "extended_t_diagram"
	case EXTENDED_T_DIAGRAMLINKS:
		return "extended_t_diagramlinks"
	case EXTENDED_T_DIAGRAMOBJECTS:
		return "extended_t_diagramobjects"
	case EXTENDED_T_CONNECTORTYPES:
		return "extended_t_connectortypes"
	case EXTENDED_T_OBJECTTYPES:
		return "extended_t_objecttypes"
	case EXTENDED_T_DIAGRAMTYPES:
		return "extended_t_diagramtypes"
	case EXTENDED_T_CARDINALITY:
		return "extended_t_cardinality"
	case OBJECT_STEREOTYPES:
		return "object_stereotypes"
	case PACKAGE_HIERARCHY:
		return "package_hierarchy"

	}
	return "unknown"
}
