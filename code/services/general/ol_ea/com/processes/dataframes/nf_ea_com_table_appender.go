package dataframes

//
//import pandas
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.collection_types.nf_ea_com_collection_types import \
//NfEaComCollectionTypes
//
//
//def append_nf_ea_com_table(
//nf_ea_com_dictionary: dict,
//new_nf_ea_com_collection: pandas.DataFrame,
//nf_ea_com_collection_type: NfEaComCollectionTypes) \
//-> dict:
//nf_ea_com_collection = \
//nf_ea_com_dictionary[nf_ea_com_collection_type]
//
//nf_ea_com_collection = \
//pandas.concat(
//[
//nf_ea_com_collection,
//new_nf_ea_com_collection
//])
//
//nf_ea_com_collection = \
//nf_ea_com_collection.reset_index(
//drop=True)
//
//nf_ea_com_dictionary[nf_ea_com_collection_type] = \
//nf_ea_com_collection
//
//return \
//nf_ea_com_dictionary
