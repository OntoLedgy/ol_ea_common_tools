package native_xml_loader

import "github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com"

//https://github.com/boro-alpha/nf_ea_common_tools/blob/master/nf_ea_common_tools_source/b_code/services/general/nf_ea/model_loader/native_xml_loader/dictionary_of_dataframes_getter.py

//from nf_common_source.code.constants.standard_constants import DEFAULT_NULL_VALUE
//from pandas import DataFrame
//
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.collection_types.nf_ea_com_collection_types import \
//NfEaComCollectionTypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.column_types.nf_ea_com_additional_column_types import \
//NfEaComAdditionalColumnTypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.column_types.nf_ea_com_column_types import \
//NfEaComColumnTypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.nf_ea_com_universes import NfEaComUniverses

//def get_dictionary_of_dataframes(
//nf_ea_com_universe: NfEaComUniverses) \
//-> dict:
func getDictionaryOfDataframes(universe *com.OlEaComUniverses) interface{} {

	return nil
}

//output_dictionary = \
//{}

//collections_dictionary = \
//nf_ea_com_universe.nf_ea_com_registry.dictionary_of_collections
//
//for collection_key in collections_dictionary.keys():
//output_dictionary[collection_key] = \
//__clear_nf_ea_guids_from_dataframe(
//collection_key=collection_key,
//nf_ea_com_dataframe=collections_dictionary[collection_key])
//
//return \
//output_dictionary
//
//
//def __clear_nf_ea_guids_from_dataframe(
//collection_key: NfEaComCollectionTypes,
//nf_ea_com_dataframe: DataFrame) \
//-> DataFrame:
//nf_ea_com_dataframe_copy = \
//nf_ea_com_dataframe.copy()
//
//if collection_key == NfEaComCollectionTypes.EA_STEREOTYPES:
//nf_ea_com_dataframe_copy[NfEaComAdditionalColumnTypes.ORIGINAL_EA_GUIDS.column_name] = \
//nf_ea_com_dataframe_copy[NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name]
//
//if collection_key == NfEaComCollectionTypes.EA_CLASSIFIERS:
//nf_ea_com_dataframe_copy[NfEaComAdditionalColumnTypes.ORIGINAL_EA_GUIDS.column_name] = \
//nf_ea_com_dataframe_copy[NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name]
//
//if collection_key == NfEaComCollectionTypes.EA_CONNECTORS:
//nf_ea_com_dataframe_copy[NfEaComAdditionalColumnTypes.ORIGINAL_EA_GUIDS.column_name] = \
//nf_ea_com_dataframe_copy[NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name]
//
//if NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name in nf_ea_com_dataframe_copy.columns:
//nf_ea_com_dataframe_copy[NfEaComColumnTypes.EXPLICIT_OBJECTS_EA_GUID.column_name] = \
//DEFAULT_NULL_VALUE
//
//return \
//nf_ea_com_dataframe_copy
