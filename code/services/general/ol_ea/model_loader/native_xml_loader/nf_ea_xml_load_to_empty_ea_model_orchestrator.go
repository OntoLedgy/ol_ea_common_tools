package native_xml_loader

import (
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/processes/nf_ea_com_loaders"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
)

//https://github.com/boro-alpha/nf_ea_common_tools/blob/master/nf_ea_common_tools_source/b_code/services/general/nf_ea/model_loader/native_xml_loader/nf_ea_xml_load_to_empty_ea_model_orchestrator.py

//from nf_common_source.code.services.file_system_service.objects.files import Files
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.nf_ea_com_universes import NfEaComUniverses
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.model_loader.native_xml_loader.dictionary_of_dataframes_getter import get_dictionary_of_dataframes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.model_loader.native_xml_loader.xml_file_creator_and_loader import create_and_load_xml_file
//from nf_ea_common_tools_source.b_code.services.session.orchestrators.ea_tools_session_managers import \
//EaToolsSessionManagers
//from nf_ea_common_tools_source.b_code.services.session.processes.nf_ea_com_loaders.ea_empty_repository_getter import \
//get_ea_repository_for_empty_ea_model

//def orchestrate_nf_ea_xml_load_to_empty_ea_model(
func OrchestrateNfEaXmlLoadToEmptyEaModel(
	//ea_tools_session_manager: EaToolsSessionManagers,
	ea_tools_session_manager orchestrators.EaToolsSessionManagers,
	//nf_ea_com_universe: NfEaComUniverses,
	nf_ea_com_universe *com.OlEaComUniverses,
	//output_xml_file_full_path: str,
	output_xml_file_full_path string,
	//ea_repository_file: Files,
	ea_repository_file *object_model.Files,
	//short_name: str):
	short_name string) {

	//ea_repository =\
	ea_repository :=
		//get_ea_repository_for_empty_ea_model(
		nf_ea_com_loaders.GetEaRepositoryForEmptyEaModel(
			//ea_tools_session_manager=ea_tools_session_manager,
			ea_tools_session_manager,
			//ea_repository_file=ea_repository_file,
			ea_repository_file,
			//short_name=short_name),
			short_name)

	//nf_ea_com_dataframes_dictionary = \
	nf_ea_com_dataframes_dictionary :=
		//get_dictionary_of_dataframes(
		getDictionaryOfDataframes(
			nf_ea_com_universe)
	//nf_ea_com_universe=nf_ea_com_universe)

	//create_and_load_xml_file(
	createAndLoadXmlFiles(
		//ea_tools_session_manager=ea_tools_session_manager,
		ea_tools_session_manager,
		//output_xml_file_full_path=output_xml_file_full_path,
		output_xml_file_full_path,
		//ea_repository=ea_repository,
		ea_repository,
		//nf_ea_com_dataframes_dictionary=nf_ea_com_dataframes_dictionary)
		nf_ea_com_dataframes_dictionary,
		"",
		false)

}
