package column_types

//from enum import auto
//from enum import unique
//from nf_common_source.code.nf.types.column_types import ColumnTypes

//@unique
//class NfEaComColumnTypes(
//ColumnTypes):
type OlEaComColumnTypes int

const (
	CLASSIFIERS_ALL_COMPONENT_EA_ATTRIBUTES OlEaComColumnTypes = iota
	CLASSIFIERS_ALL_COMPONENT_EA_OPERATIONS
	CLASSIFIERS_CONTAINING_EA_ELEMENT
	ELEMENTS_EA_OBJECT_TYPE
	ELEMENTS_SUPPLIER_PLACE1_END_CONNECTORS
	ELEMENTS_CLIENT_PLACE2_END_CONNECTORS
	ELEMENTS_CONTAINED_EA_DIAGRAMS
	ELEMENTS_CONTAINED_EA_CLASSIFIERS
	ELEMENTS_CLASSIFIER
	EXPLICIT_OBJECTS_EA_OBJECT_NAME
	EXPLICIT_OBJECTS_EA_OBJECT_NOTES
	EXPLICIT_OBJECTS_EA_GUID
	PACKAGEABLE_OBJECTS_PARENT_EA_ELEMENT
	PACKAGES_CONTAINED_EA_PACKAGES
	PACKAGES_VIEW_TYPE
	REPOSITORIED_OBJECTS_EA_REPOSITORY
	STEREOTYPE_EA_STEREOTYPE_GROUP
	STEREOTYPEABLE_OBJECTS_EA_OBJECT_STEREOTYPES
	CONNECTORS_DIRECTION_TYPE_NAME
	CONNECTORS_ELEMENT_TYPE_NAME
	CONNECTORS_SOURCE_CARDINALITY
	CONNECTORS_DEST_CARDINALITY
	ELEMENT_COMPONENTS_CONTAINING_EA_CLASSIFIER
	ELEMENT_COMPONENTS_CLASSIFYING_EA_CLASSIFIER
	ELEMENT_COMPONENTS_UML_VISIBILITY_KIND
	ELEMENT_COMPONENTS_DEFAULT
	ELEMENT_COMPONENTS_TYPE
	ATTRIBUTES_LOWER_BOUNDS
	ATTRIBUTES_UPPER_BOUNDS
	ANALYSIS_METRICS_METRICS
	ANALYSIS_METRICS_VALUES
	STEREOTYPE_CLIENT_NF_UUIDS
	STEREOTYPE_NAMES
	STEREOTYPE_APPLIES_TOS
	STEREOTYPE_STYLE
	STEREOTYPE_PROPERTY_TYPE
)

//def __column_name(
//self) \
//-> str:
//column_name = \
//column_name_mapping[self]
//
//return \
//column_name
//
//column_name = \
//property(
//fget=__column_name)

//column_name_mapping = \
func (olEaComColumnTypes OlEaComColumnTypes) ColumnName() string {

	switch olEaComColumnTypes {

	case CLASSIFIERS_ALL_COMPONENT_EA_ATTRIBUTES:
		return "all_component_ea_attributes"
	case CLASSIFIERS_ALL_COMPONENT_EA_OPERATIONS:
		return "all_component_ea_operations"
	case CLASSIFIERS_CONTAINING_EA_ELEMENT:
		return "containing_ea_element"
	case ELEMENTS_EA_OBJECT_TYPE:
		return "ea_object_type"
	case ELEMENTS_SUPPLIER_PLACE1_END_CONNECTORS:
		return "supplier_place1_end_connectors"
	case ELEMENTS_CLIENT_PLACE2_END_CONNECTORS:
		return "client_place2_end_connectors"
	case ELEMENTS_CONTAINED_EA_DIAGRAMS:
		return "contained_ea_diagrams"
	case ELEMENTS_CONTAINED_EA_CLASSIFIERS:
		return "contained_ea_classifiers"
	case ELEMENTS_CLASSIFIER:
		return "classifier"
	case EXPLICIT_OBJECTS_EA_OBJECT_NAME:
		return "ea_object_name"
	case EXPLICIT_OBJECTS_EA_OBJECT_NOTES:
		return "ea_object_notes"
	case EXPLICIT_OBJECTS_EA_GUID:
		return "ea_guid"
	case PACKAGEABLE_OBJECTS_PARENT_EA_ELEMENT:
		return "parent_ea_element"
	case PACKAGES_CONTAINED_EA_PACKAGES:
		return "contained_ea_packages"
	case PACKAGES_VIEW_TYPE:
		return "view_type"
	case REPOSITORIED_OBJECTS_EA_REPOSITORY:
		return "ea_repository"
	case STEREOTYPE_EA_STEREOTYPE_GROUP:
		return "ea_stereotype_group"
	case STEREOTYPEABLE_OBJECTS_EA_OBJECT_STEREOTYPES:
		return "ea_object_stereotypes"
	case CONNECTORS_DIRECTION_TYPE_NAME:
		return "ea_connector_direction_type_name"
	case CONNECTORS_ELEMENT_TYPE_NAME:
		return "ea_connector_element_type_name"
	case CONNECTORS_SOURCE_CARDINALITY:
		return "ea_connector_supplier_cardinality"
	case CONNECTORS_DEST_CARDINALITY:
		return "ea_connector_client_cardinality"
	case ELEMENT_COMPONENTS_CONTAINING_EA_CLASSIFIER:
		return "containing_ea_classifier"
	case ELEMENT_COMPONENTS_CLASSIFYING_EA_CLASSIFIER:
		return "classifying_ea_classifier"
	case ELEMENT_COMPONENTS_UML_VISIBILITY_KIND:
		return "uml_visibility_kind"
	case ELEMENT_COMPONENTS_DEFAULT:
		return "default"
	case ELEMENT_COMPONENTS_TYPE:
		return "type"
	case ATTRIBUTES_LOWER_BOUNDS:
		return "lower_bounds"
	case ATTRIBUTES_UPPER_BOUNDS:
		return "upper_bounds"
	case ANALYSIS_METRICS_METRICS:
		return "metrics"
	case ANALYSIS_METRICS_VALUES:
		return "values"
	case STEREOTYPE_CLIENT_NF_UUIDS:
		return "client_nf_uuids"
	case STEREOTYPE_NAMES:
		return "stereotype_names"
	case STEREOTYPE_APPLIES_TOS:
		return "stereotype_applies_tos"
	case STEREOTYPE_STYLE:
		return "stereotype_style"
	case STEREOTYPE_PROPERTY_TYPE:
		return "property_type"
	}
	return "unknown"
}
