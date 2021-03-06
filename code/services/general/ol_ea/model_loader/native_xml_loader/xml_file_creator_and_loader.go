package native_xml_loader

import "github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"

///https://github.com/boro-alpha/nf_ea_common_tools/blob/master/nf_ea_common_tools_source/b_code/services/general/nf_ea/model_loader/native_xml_loader/xml_file_creator_and_loader.py

//from nf_common_source.code.services.reporting_service.reporters.log_with_datetime import log_message
//from nf_ea_common_tools_source.b_code.nf_ea_common.objects.ea_repositories import EaRepositories
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.collection_types.nf_ea_com_collection_types import \
//NfEaComCollectionTypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.model_loader.ea_interop_loader.object_loaders.ea_stereotypes_loader import map_and_load_ea_stereotypes
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.model_loader.native_xml_loader.ea_xml_file_loader import load_ea_xml_file_to_ea_repository
//from nf_ea_common_tools_source.b_code.services.general.ea.xml.ea_model_xml_orchestrator import orchestrate_creation_xml_native_file_for_import
//from nf_ea_common_tools_source.b_code.services.session.orchestrators.ea_tools_session_managers import \
//EaToolsSessionManagers

//def create_and_load_xml_file(

func createAndLoadXmlFiles(
	//ea_tools_session_manager: EaToolsSessionManagers,
	manager orchestrators.EaToolsSessionManagers,
	//output_xml_file_full_path: str,
	path string,
	//ea_repository: EaRepositories,
	repository interface{},
	//nf_ea_com_dataframes_dictionary: dict,
	dictionary interface{},
	//default_model_package_ea_guid: str = None,
	default_model_package_ea_guid string,
	//just_create_xml=False):)
	just_create_xml bool) {

	//log_message(
	//message='Creating xml file')
	//
	//ea_project_common_nf_ea_sql_universe = \
	//ea_tools_session_manager.nf_ea_sql_stage_manager.nf_ea_sql_universe_manager.get_nf_ea_sql_universe(
	//ea_repository=ea_repository)
	//
	//ea_project_common_nf_ea_sql_universe.add_ea_guids_to_ea_identifiers_for_objects()
	//
	//ea_project_common_nf_ea_sql_universe.add_ea_guids_to_ea_identifiers_for_packages()
	//
	//last_ea_identifier_for_objects = \
	//ea_project_common_nf_ea_sql_universe.get_last_ea_identifier_for_objects()
	//
	//last_ea_identifier_for_packages = \
	//ea_project_common_nf_ea_sql_universe.get_last_ea_identifier_for_packages()
	//
	//last_ea_identifier_for_attributes = \
	//ea_project_common_nf_ea_sql_universe.get_last_ea_identifier_for_attributes()
	//
	//map_and_load_ea_stereotypes(
	//ea_stereotypes=nf_ea_com_dataframes_dictionary[NfEaComCollectionTypes.EA_STEREOTYPES],
	//ea_tools_session_manager=ea_tools_session_manager,
	//ea_repository=ea_repository)
	//
	//orchestrate_creation_xml_native_file_for_import(
	//output_xml_file_full_path=output_xml_file_full_path,
	//nf_ea_com_dataframes_dictionary=nf_ea_com_dataframes_dictionary,
	//start_ea_identifier_for_new_objects=last_ea_identifier_for_objects + 1,
	//start_ea_identifier_for_new_packages=last_ea_identifier_for_packages + 1,
	//start_ea_identifier_for_new_attributes=last_ea_identifier_for_attributes + 1,
	//default_model_package_ea_guid=default_model_package_ea_guid)
	//
	//if just_create_xml:
	//return
	//
	//load_ea_xml_file_to_ea_repository(
	//ea_repository=ea_repository,
	//load_xml_file_full_path=output_xml_file_full_path)

}
