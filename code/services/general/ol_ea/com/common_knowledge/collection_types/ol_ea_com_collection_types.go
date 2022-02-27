package collection_types

//from enum import auto
//from enum import unique
//from nf_common_source.code.nf.types.collection_types import CollectionTypes

//@unique
//class NfEaComCollectionTypes(
//CollectionTypes):
type OlEaComCollectionTypes int

const (
	EA_PACKAGES OlEaComCollectionTypes = iota
	EA_CLASSIFIERS
	EA_ATTRIBUTES
	EA_CONNECTORS
	EA_CONNECTORS_PC
	EA_STEREOTYPES
	STEREOTYPE_USAGE
	THIN_EA_EXPLICIT_OBJECTS
	THIN_EA_REPOSITORIED_OBJECTS
	THIN_EA_STEREOTYPEABLE_OBJECTS
	THIN_EA_PACKAGEABLE_OBJECTS
	THIN_EA_ELEMENTS
	THIN_EA_CLASSIFIERS
	THIN_EA_PACKAGES
	THIN_EA_DIAGRAMS
	THIN_EA_CONNECTORS
	THIN_EA_STEREOTYPES
	THIN_EA_ATTRIBUTES
	THIN_EA_ELEMENT_COMPONENTS
	THIN_EA_OPERATIONS
	EA_DIAGRAMS
	EA_STEREOTYPE_GROUPS
	EA_OPERATIONS
	EA_OBJECT_STEREOTYPES
	EA_CONNECTOR_TYPES
	EA_ELEMENT_TYPES
	EA_DIAGRAM_TYPES
	EA_CARDINALITIES
	EA_FULL_GENERALISATIONS
	EA_FULL_DEPENDENCIES
	EA_NEAREST_PACKAGES
	EA_FULL_PACKAGES
	EA_PACKAGE_CONTENTS_SUMMARY
	SUMMARY_TABLE_BY_TYPE
	DEPENDENCY_DEPTHS_TABLE
	ANALYSIS_METRICS
	//# From model stats code
	MODEL_STATS_GENERAL_CONNECTORS_BY_TYPE
	MODEL_STATS_GENERAL_EDGES
	MODEL_STATS_GENERAL_LEAVES
	MODEL_STATS_GENERAL_NODES
	MODEL_STATS_GENERAL_OBJECTS_BY_TYPE
	MODEL_STATS_GENERAL_PATHS
	MODEL_STATS_GENERAL_ROOTS
	MODEL_STATS_GENERAL_SUMMARY_TABLE

	MODEL_STATS_PROXY_PROCESS_CONNECTORS_BY_TYPE
	MODEL_STATS_PROXY_PROCESS_EDGES
	MODEL_STATS_PROXY_PROCESS_LEAVES
	MODEL_STATS_PROXY_PROCESS_NODES
	MODEL_STATS_PROXY_PROCESS_OBJECTS_BY_TYPE
	MODEL_STATS_PROXY_PROCESS_PATHS
	MODEL_STATS_PROXY_PROCESS_ROOTS
	MODEL_STATS_PROXY_PROCESS_SUMMARY_TABLE

	MODEL_STATS_FULL_DEPENDENCIES_EDGES
	MODEL_STATS_FULL_DEPENDENCIES_LEAVES
	MODEL_STATS_FULL_DEPENDENCIES_NODES
	MODEL_STATS_FULL_DEPENDENCIES_PATHS
	MODEL_STATS_FULL_DEPENDENCIES_ROOTS
	MODEL_STATS_FULL_DEPENDENCIES_SUMMARY_TABLE

	MODEL_STATS_FIRST_CLASS_RELATION_EDGES
	MODEL_STATS_FIRST_CLASS_RELATION_LEAVES
	MODEL_STATS_FIRST_CLASS_RELATION_NODES
	MODEL_STATS_FIRST_CLASS_RELATION_ROOTS
	MODEL_STATS_FIRST_CLASS_RELATION_IMPLICIT_EDGES

	MODEL_STATS_HIGH_ORDER_TYPES_EDGES
	MODEL_STATS_HIGH_ORDER_TYPES_LEAVES
	MODEL_STATS_HIGH_ORDER_TYPES_NODES
	MODEL_STATS_HIGH_ORDER_TYPES_ROOTS
	MODEL_STATS_HIGH_ORDER_TYPES_IMPLICIT_EDGES
)

//@staticmethod
//def get_collection_type_from_name(
//name: str):
func GetCollectionTypeFromName(
	name string) *OlEaComCollectionTypes {

	//for collection_type, collection_name in collection_name_mapping.items():
	for collection_type, collection_name := range collection_name_mapping {
		//if collection_name == name:
		if collection_name == name {

			//return \
			//collection_type
			return &collection_type
		}
	}
	return nil
}

//def __collection_name(
//self) \
//-> str:
//collection_name = \
//collection_name_mapping[self]
//
//return \
//collection_name

//collection_name = \
//property(
//fget=__collection_name)

//collection_name_mapping = \
var collection_name_mapping = map[OlEaComCollectionTypes]string{
	EA_PACKAGES:                    "ea_packages",
	EA_CLASSIFIERS:                 "ea_classifiers",
	EA_ATTRIBUTES:                  "ea_attributes",
	EA_CONNECTORS:                  "ea_connectors",
	EA_CONNECTORS_PC:               "ea_connectors_pc",
	EA_STEREOTYPES:                 "ea_stereotypes",
	STEREOTYPE_USAGE:               "stereotype_usage",
	THIN_EA_EXPLICIT_OBJECTS:       "thin_ea_explicit_objects",
	THIN_EA_REPOSITORIED_OBJECTS:   "thin_ea_repositoried_objects",
	THIN_EA_STEREOTYPEABLE_OBJECTS: "thin_ea_stereotypeable_objects",
	THIN_EA_PACKAGEABLE_OBJECTS:    "thin_ea_packageable_objects",
	THIN_EA_ELEMENTS:               "thin_ea_elements",
	THIN_EA_CLASSIFIERS:            "thin_ea_classifiers",
	THIN_EA_PACKAGES:               "thin_ea_packages",
	THIN_EA_DIAGRAMS:               "thin_ea_diagrams",
	THIN_EA_CONNECTORS:             "thin_ea_connectors",
	THIN_EA_STEREOTYPES:            "thin_ea_stereotypes",
	THIN_EA_ATTRIBUTES:             "thin_ea_attributes",
	THIN_EA_ELEMENT_COMPONENTS:     "thin_ea_element_components",
	THIN_EA_OPERATIONS:             "thin_ea_operations",
	EA_DIAGRAMS:                    "ea_diagrams",
	EA_STEREOTYPE_GROUPS:           "ea_stereotype_groups",
	EA_OPERATIONS:                  "ea_operations",
	EA_OBJECT_STEREOTYPES:          "ea_object_stereotypes",
	EA_CONNECTOR_TYPES:             "ea_connector_types",
	EA_ELEMENT_TYPES:               "ea_element_types",
	EA_DIAGRAM_TYPES:               "ea_diagram_types",
	EA_CARDINALITIES:               "ea_cardinalities",
	EA_FULL_GENERALISATIONS:        "ea_full_generalisations",
	EA_FULL_DEPENDENCIES:           "ea_full_dependencies",
	EA_NEAREST_PACKAGES:            "ea_nearest_packages",
	EA_FULL_PACKAGES:               "ea_full_packages",
	EA_PACKAGE_CONTENTS_SUMMARY:    "ea_package_contents_summary",
	SUMMARY_TABLE_BY_TYPE:          "summary_table_by_type",
	DEPENDENCY_DEPTHS_TABLE:        "dependency_depths_table",
	ANALYSIS_METRICS:               "analysis_metrics",
	//# From		model		stats		code
	MODEL_STATS_GENERAL_CONNECTORS_BY_TYPE: "general_connector_by_type",
	MODEL_STATS_GENERAL_EDGES:              "general_edges",
	MODEL_STATS_GENERAL_LEAVES:             "general_leaves",
	MODEL_STATS_GENERAL_NODES:              "general_nodes",
	MODEL_STATS_GENERAL_OBJECTS_BY_TYPE:    "general_objects_by_type",
	MODEL_STATS_GENERAL_PATHS:              "general_paths",
	MODEL_STATS_GENERAL_ROOTS:              "general_roots",
	MODEL_STATS_GENERAL_SUMMARY_TABLE:      "general_summary_table",

	MODEL_STATS_PROXY_PROCESS_CONNECTORS_BY_TYPE: "proxies_connector_by_type",
	MODEL_STATS_PROXY_PROCESS_EDGES:              "proxies_edges",
	MODEL_STATS_PROXY_PROCESS_LEAVES:             "proxies_leaves",
	MODEL_STATS_PROXY_PROCESS_NODES:              "proxies_nodes",
	MODEL_STATS_PROXY_PROCESS_OBJECTS_BY_TYPE:    "proxies_objects_by_type",
	MODEL_STATS_PROXY_PROCESS_PATHS:              "proxies_paths",
	MODEL_STATS_PROXY_PROCESS_ROOTS:              "proxies_roots",
	MODEL_STATS_PROXY_PROCESS_SUMMARY_TABLE:      "proxies_summary_table",

	MODEL_STATS_FULL_DEPENDENCIES_EDGES:         "full_dependencies_edges",
	MODEL_STATS_FULL_DEPENDENCIES_LEAVES:        "full_dependencies_leaves",
	MODEL_STATS_FULL_DEPENDENCIES_NODES:         "full_dependencies_nodes",
	MODEL_STATS_FULL_DEPENDENCIES_PATHS:         "full_dependencies_paths",
	MODEL_STATS_FULL_DEPENDENCIES_ROOTS:         "full_dependencies_roots",
	MODEL_STATS_FULL_DEPENDENCIES_SUMMARY_TABLE: "full_dependencies_summary_table",

	MODEL_STATS_FIRST_CLASS_RELATION_EDGES:          "FCR_edges",
	MODEL_STATS_FIRST_CLASS_RELATION_LEAVES:         "FCR_leaves",
	MODEL_STATS_FIRST_CLASS_RELATION_NODES:          "FCR_nodes",
	MODEL_STATS_FIRST_CLASS_RELATION_ROOTS:          "FCR_roots",
	MODEL_STATS_FIRST_CLASS_RELATION_IMPLICIT_EDGES: "FCR_implicit_edges",

	MODEL_STATS_HIGH_ORDER_TYPES_EDGES:          "HOT_edges",
	MODEL_STATS_HIGH_ORDER_TYPES_LEAVES:         "HOT_leaves",
	MODEL_STATS_HIGH_ORDER_TYPES_NODES:          "HOT_nodes",
	MODEL_STATS_HIGH_ORDER_TYPES_ROOTS:          "HOT_roots",
	MODEL_STATS_HIGH_ORDER_TYPES_IMPLICIT_EDGES: "HOT_implicit_edges"}

func (olEaCollectionTypes OlEaComCollectionTypes) String() string {

	switch olEaCollectionTypes {
	case EA_PACKAGES:
		return "ea_packages"
	case EA_CLASSIFIERS:
		return "ea_classifiers"
	case EA_ATTRIBUTES:
		return "ea_attributes"
	case EA_CONNECTORS:
		return "ea_connectors"
	case EA_CONNECTORS_PC:
		return "ea_connectors_pc"
	case EA_STEREOTYPES:
		return "ea_stereotypes"
	case STEREOTYPE_USAGE:
		return "stereotype_usage"
	case THIN_EA_EXPLICIT_OBJECTS:
		return "thin_ea_explicit_objects"
	case THIN_EA_REPOSITORIED_OBJECTS:
		return "thin_ea_repositoried_objects"
	case THIN_EA_STEREOTYPEABLE_OBJECTS:
		return "thin_ea_stereotypeable_objects"
	case THIN_EA_PACKAGEABLE_OBJECTS:
		return "thin_ea_packageable_objects"
	case THIN_EA_ELEMENTS:
		return "thin_ea_elements"
	case THIN_EA_CLASSIFIERS:
		return "thin_ea_classifiers"
	case THIN_EA_PACKAGES:
		return "thin_ea_packages"
	case THIN_EA_DIAGRAMS:
		return "thin_ea_diagrams"
	case THIN_EA_CONNECTORS:
		return "thin_ea_connectors"
	case THIN_EA_STEREOTYPES:
		return "thin_ea_stereotypes"
	case THIN_EA_ATTRIBUTES:
		return "thin_ea_attributes"
	case THIN_EA_ELEMENT_COMPONENTS:
		return "thin_ea_element_components"
	case THIN_EA_OPERATIONS:
		return "thin_ea_operations"
	case EA_DIAGRAMS:
		return "ea_diagrams"
	case EA_STEREOTYPE_GROUPS:
		return "ea_stereotype_groups"
	case EA_OPERATIONS:
		return "ea_operations"
	case EA_OBJECT_STEREOTYPES:
		return "ea_object_stereotypes"
	case EA_CONNECTOR_TYPES:
		return "ea_connector_types"
	case EA_ELEMENT_TYPES:
		return "ea_element_types"
	case EA_DIAGRAM_TYPES:
		return "ea_diagram_types"
	case EA_CARDINALITIES:
		return "ea_cardinalities"
	case EA_FULL_GENERALISATIONS:
		return "ea_full_generalisations"
	case EA_FULL_DEPENDENCIES:
		return "ea_full_dependencies"
	case EA_NEAREST_PACKAGES:
		return "ea_nearest_packages"
	case EA_FULL_PACKAGES:
		return "ea_full_packages"
	case EA_PACKAGE_CONTENTS_SUMMARY:
		return "ea_package_contents_summary"
	case SUMMARY_TABLE_BY_TYPE:
		return "summary_table_by_type"
	case DEPENDENCY_DEPTHS_TABLE:
		return "dependency_depths_table"
	case ANALYSIS_METRICS:
		return "analysis_metrics"
	//# From		model		stats		code
	case MODEL_STATS_GENERAL_CONNECTORS_BY_TYPE:
		return "general_connector_by_type"
	case MODEL_STATS_GENERAL_EDGES:
		return "general_edges"
	case MODEL_STATS_GENERAL_LEAVES:
		return "general_leaves"
	case MODEL_STATS_GENERAL_NODES:
		return "general_nodes"
	case MODEL_STATS_GENERAL_OBJECTS_BY_TYPE:
		return "general_objects_by_type"
	case MODEL_STATS_GENERAL_PATHS:
		return "general_paths"
	case MODEL_STATS_GENERAL_ROOTS:
		return "general_roots"
	case MODEL_STATS_GENERAL_SUMMARY_TABLE:
		return "general_summary_table"

	case MODEL_STATS_PROXY_PROCESS_CONNECTORS_BY_TYPE:
		return "proxies_connector_by_type"
	case MODEL_STATS_PROXY_PROCESS_EDGES:
		return "proxies_edges"
	case MODEL_STATS_PROXY_PROCESS_LEAVES:
		return "proxies_leaves"
	case MODEL_STATS_PROXY_PROCESS_NODES:
		return "proxies_nodes"
	case MODEL_STATS_PROXY_PROCESS_OBJECTS_BY_TYPE:
		return "proxies_objects_by_type"
	case MODEL_STATS_PROXY_PROCESS_PATHS:
		return "proxies_paths"
	case MODEL_STATS_PROXY_PROCESS_ROOTS:
		return "proxies_roots"
	case MODEL_STATS_PROXY_PROCESS_SUMMARY_TABLE:
		return "proxies_summary_table"

	case MODEL_STATS_FULL_DEPENDENCIES_EDGES:
		return "full_dependencies_edges"
	case MODEL_STATS_FULL_DEPENDENCIES_LEAVES:
		return "full_dependencies_leaves"
	case MODEL_STATS_FULL_DEPENDENCIES_NODES:
		return "full_dependencies_nodes"
	case MODEL_STATS_FULL_DEPENDENCIES_PATHS:
		return "full_dependencies_paths"
	case MODEL_STATS_FULL_DEPENDENCIES_ROOTS:
		return "full_dependencies_roots"
	case MODEL_STATS_FULL_DEPENDENCIES_SUMMARY_TABLE:
		return "full_dependencies_summary_table"

	case MODEL_STATS_FIRST_CLASS_RELATION_EDGES:
		return "FCR_edges"
	case MODEL_STATS_FIRST_CLASS_RELATION_LEAVES:
		return "FCR_leaves"
	case MODEL_STATS_FIRST_CLASS_RELATION_NODES:
		return "FCR_nodes"
	case MODEL_STATS_FIRST_CLASS_RELATION_ROOTS:
		return "FCR_roots"
	case MODEL_STATS_FIRST_CLASS_RELATION_IMPLICIT_EDGES:
		return "FCR_implicit_edges"

	case MODEL_STATS_HIGH_ORDER_TYPES_EDGES:
		return "HOT_edges"
	case MODEL_STATS_HIGH_ORDER_TYPES_LEAVES:
		return "HOT_leaves"
	case MODEL_STATS_HIGH_ORDER_TYPES_NODES:
		return "HOT_nodes"
	case MODEL_STATS_HIGH_ORDER_TYPES_ROOTS:
		return "HOT_roots"
	case MODEL_STATS_HIGH_ORDER_TYPES_IMPLICIT_EDGES:
		return "HOT_implicit_edges"
	}
	return "unknown"
}
