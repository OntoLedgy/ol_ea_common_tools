package temporary

//from nf_ea_common_tools_source.b_code.services.session.orchestrators.ea_tools_session_managers import \
//EaToolsSessionManagers
//from nf_ea_common_tools_source.b_code.services.session.processes.creators.empty_nf_ea_com_universe_creator import create_empty_nf_ea_universe
//from nf_common_source.code.constants.standard_constants import DEFAULT_NULL_VALUE
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.collection_types.nf_ea_com_collection_types import \
//NfEaComCollectionTypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.nf_ea_com_universes import NfEaComUniverses
//from pandas import concat
//from pandas import DataFrame

import (
	"github.com/OntoLedgy/ol_common_services/code/constants"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/common_knowledge/collection_types"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/processes/creators"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
)

//def convert_nf_ea_com_dictionary_to_nf_ea_com_universe(
func ConvertOlEaComMapToOlEaComUniverse(
	//ea_tools_session_manager: EaToolsSessionManagers,
	eaToolsSessionManagers orchestrators.EaToolsSessionManagers,
	//nf_ea_com_dictionary: dict,
	eaComMap map[collection_types.OlEaComCollectionTypes]dataframes.DataFrames,
	//short_name: str) \
	//-> NfEaComUniverses:
	shortName string) *com.OlEaComUniverses {

	//nf_ea_com_universe = \
	olEaComUniverse :=
		//create_empty_nf_ea_universe(
		creators.CreateEmptyOlEAUniverse(
			//ea_tools_session_manager=ea_tools_session_manager,
			eaToolsSessionManagers,
			//short_name=short_name)
			shortName)

	//for collection_key, nf_com_ea_dictionary_item in nf_ea_com_dictionary.items():
	for collectionKey, olComEaDictionaryItem := range eaComMap {
		//__convert_nf_com_ea_dictionary_item_to_nf_ea_com_universe_collection(
		convertOlComEaDictionaryItemToOlEaComUniverseCollection(
			//collection_key=collection_key,
			collectionKey,
			//nf_com_ea_dictionary_item=nf_com_ea_dictionary_item,
			olComEaDictionaryItem,
			//nf_ea_com_universe=nf_ea_com_universe)
			olEaComUniverse)

	}

	//return \
	//nf_ea_com_universe
	return olEaComUniverse
}

//def __convert_nf_com_ea_dictionary_item_to_nf_ea_com_universe_collection(
func convertOlComEaDictionaryItemToOlEaComUniverseCollection(
	//collection_key: NfEaComCollectionTypes,
	collectionKey collection_types.OlEaComCollectionTypes,
	//nf_com_ea_dictionary_item: DataFrame,
	olComEaDictionaryItem dataframes.DataFrames,
	//nf_ea_com_universe: NfEaComUniverses):
	olEaComUniverse *com.OlEaComUniverses) {

	//if collection_key == 'ea_attributes_order':
	//return
	if collectionKey.CollectionName() == "ea_attributes_order" {
		return
	}

	//if collection_key == NfEaComCollectionTypes.EA_CONNECTORS_PC:
	if collectionKey == collection_types.EA_CONNECTORS_PC {
		//collection_key = \
		collectionKey =
			//NfEaComCollectionTypes.EA_CONNECTORS
			collection_types.EA_CONNECTORS
	}

	//collection = \
	collection :=
		//nf_ea_com_universe.nf_ea_com_registry.dictionary_of_collections[collection_key]
		olEaComUniverse.OlEaComRegistries.MapOfCollections[collectionKey]

	//new_collection = \
	newCollection :=
		//concat(
		//[
		&dataframes.DataFrames{
			//collection,
			DataFrame: collection.DataFrame.Concat(
				//nf_com_ea_dictionary_item
				//])
				olComEaDictionaryItem.DataFrame)}

	//new_collection = \
	newCollection.FillNa(
		//new_collection.fillna(
		//DEFAULT_NULL_VALUE)
		constants.DEFAULT_NULL_VALUE)

	//nf_ea_com_universe.nf_ea_com_registry.dictionary_of_collections[collection_key] = \
	olEaComUniverse.OlEaComRegistries.MapOfCollections[collectionKey] =
		//new_collection
		newCollection

}
