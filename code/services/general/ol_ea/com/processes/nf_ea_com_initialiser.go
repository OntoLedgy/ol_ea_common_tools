package processes

//
//from nf_common_source.code.nf.types.nf_column_types import NfColumnTypes
//from pandas import DataFrame
//from nf_ea_common_tools_source.b_code.nf_ea_common.common_knowledge.column_types.ea_t.ea_t_xref_column_types import EaTXrefColumnTypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.collection_types.nf_ea_com_collection_types import NfEaComCollectionTypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.column_types.nf_ea_com_column_types import NfEaComColumnTypes

import (
	"github.com/OntoLedgy/ol_common_services/code/ol/types"
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/common_knowledge/column_types/ea_t"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/common_knowledge/collection_types"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/common_knowledge/column_types"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

//def initialise_nf_ea_com_dictionary() \
//-> dict:
func InitialiseOlEaComMap() *map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames {

	//nf_ea_com_dictionary = \
	OlEaComDictionary :=
		//dict()
		map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames{}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.EA_PACKAGES] = \
	OlEaComDictionary[collection_types.EA_PACKAGES] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//NfColumnTypes.NF_UUIDS.column_name,
				series.New([]string{}, series.String, types.OL_UUIDS.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name,
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_GUID.ColumnName()),
				//NfEaComColumnTypes.PACKAGEABLE_OBJECTS_PARENT_EA_ELEMENT.column_name,
				series.New([]string{}, series.String, column_types.PACKAGEABLE_OBJECTS_PARENT_EA_ELEMENT.ColumnName()),
				//NfEaComColumnTypes.PACKAGES_VIEW_TYPE.column_name])
				series.New([]string{}, series.String, column_types.PACKAGES_VIEW_TYPE.ColumnName()),
			)}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.EA_CLASSIFIERS] = \
	OlEaComDictionary[collection_types.EA_CLASSIFIERS] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//NfColumnTypes.NF_UUIDS.column_name,
				series.New([]string{}, series.String, types.OL_UUIDS.ColumnName()),
				//NfEaComColumnTypes.ELEMENTS_EA_OBJECT_TYPE.column_name,
				series.New([]string{}, series.String, column_types.ELEMENTS_EA_OBJECT_TYPE.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_OBJECT_NAME.column_name,
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_OBJECT_NAME.ColumnName()),
				//NfEaComColumnTypes.PACKAGEABLE_OBJECTS_PARENT_EA_ELEMENT.column_name,
				series.New([]string{}, series.String, column_types.PACKAGEABLE_OBJECTS_PARENT_EA_ELEMENT.ColumnName()),
				//NfEaComColumnTypes.ELEMENTS_CLASSIFIER.column_name,
				series.New([]string{}, series.String, column_types.ELEMENTS_CLASSIFIER.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name,
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_GUID.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_OBJECT_NOTES.column_name])
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_OBJECT_NOTES.ColumnName()),
			)}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.EA_ATTRIBUTES] = \
	OlEaComDictionary[collection_types.EA_ATTRIBUTES] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//NfColumnTypes.NF_UUIDS.column_name,
				series.New([]string{}, series.String, types.OL_UUIDS.ColumnName()),
				//NfEaComColumnTypes.ELEMENT_COMPONENTS_CONTAINING_EA_CLASSIFIER.column_name,
				series.New([]string{}, series.String, column_types.ELEMENT_COMPONENTS_CONTAINING_EA_CLASSIFIER.ColumnName()),
				//NfEaComColumnTypes.ELEMENT_COMPONENTS_CLASSIFYING_EA_CLASSIFIER.column_name,
				series.New([]string{}, series.String, column_types.ELEMENT_COMPONENTS_CLASSIFYING_EA_CLASSIFIER.ColumnName()),
				//NfEaComColumnTypes.ELEMENT_COMPONENTS_UML_VISIBILITY_KIND.column_name,
				series.New([]string{}, series.String, column_types.ELEMENT_COMPONENTS_UML_VISIBILITY_KIND.ColumnName()),
				//NfEaComColumnTypes.ELEMENT_COMPONENTS_TYPE.column_name,
				series.New([]string{}, series.String, column_types.ELEMENT_COMPONENTS_TYPE.ColumnName()),
				//NfEaComColumnTypes.ELEMENT_COMPONENTS_DEFAULT.column_name,
				series.New([]string{}, series.String, column_types.ELEMENT_COMPONENTS_DEFAULT.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_OBJECT_NAME.column_name,
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_OBJECT_NAME.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name])
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_GUID.ColumnName()),
			)}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.EA_CONNECTORS] = \
	OlEaComDictionary[collection_types.EA_CONNECTORS] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//NfColumnTypes.NF_UUIDS.column_name,
				series.New([]string{}, series.String, types.OL_UUIDS.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name])
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_GUID.ColumnName()),
			)}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.EA_CONNECTORS_PC] = \
	OlEaComDictionary[collection_types.EA_CONNECTORS_PC] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//NfColumnTypes.NF_UUIDS.column_name,
				series.New([]string{}, series.String, types.OL_UUIDS.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name])
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_GUID.ColumnName()),
			)}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.EA_STEREOTYPE_GROUPS] = \
	OlEaComDictionary[collection_types.EA_STEREOTYPE_GROUPS] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//NfColumnTypes.NF_UUIDS.column_name,
				series.New([]string{}, series.String, types.OL_UUIDS.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_OBJECT_NAME.column_name])
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_OBJECT_NAME.ColumnName()),
			)}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.EA_STEREOTYPES] = \
	OlEaComDictionary[collection_types.EA_STEREOTYPE_GROUPS] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//NfColumnTypes.NF_UUIDS.column_name,
				series.New([]string{}, series.String, types.OL_UUIDS.ColumnName()),
				//NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name])
				series.New([]string{}, series.String, column_types.EXPLICIT_OBJECTS_EA_GUID.ColumnName()),
			)}

	//nf_ea_com_dictionary[NfEaComCollectionTypes.STEREOTYPE_USAGE] = \
	OlEaComDictionary[collection_types.STEREOTYPE_USAGE] =
		//DataFrame(
		&dataframes.DataFrames{
			//columns=[
			DataFrame: dataframe.New(
				//EaTXrefColumnTypes.T_XREF_CLIENT_EA_GUIDS.nf_column_name,
				series.New([]string{}, series.String, ea_t.T_XREF_CLIENT_EA_GUIDS.OlColumnName()),
				//'stereotype_guids',
				series.New([]string{}, series.String, "stereotype_guids"),
				//NfEaComColumnTypes.STEREOTYPE_CLIENT_NF_UUIDS.column_name,
				series.New([]string{}, series.String, column_types.STEREOTYPE_CLIENT_NF_UUIDS.ColumnName()),
				//'stereotype_nf_uuids'])
				series.New([]string{}, series.String, "stereotype_nf_uuids"),
			)}

	//TODO : discuss with Andy

	////nf_ea_com_dictionary['ea_attributes_order'] = \
	//OlEaComDictionary["ea_attributes_order"] =
	////DataFrame(
	//	&dataframes.DataFrames{
	//		//columns=[
	//		DataFrame: dataframe.New(
	//			//EaTXrefColumnTypes.T_XREF_CLIENT_EA_GUIDS.nf_column_name,
	//			series.New([]string{}, series.String, ea_t.T_XREF_CLIENT_EA_GUIDS.OlColumnName()),
	//			//'naming_space_names',
	//			series.New([]string{}, series.String, "naming_space_names"),
	//			//'attribute_order'])
	//			series.New([]string{}, series.String, "attribute_order"),
	//		)}

	//return \
	//nf_ea_com_dictionary
	return &OlEaComDictionary
}
