package orchestrators

import (
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/model_loader/native_xml_loader"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"path"
)

//https://github.com/boro-alpha/nf_ea_common_tools/blob/master/nf_ea_common_tools_source/b_code/services/general/nf_ea/domain_migration/domain_to_nf_ea_com_migration/orchestrators/nf_ea_com_universe_to_eapx_migration_orchestator.py
//
//import os
//from nf_common_source.code.services.file_system_service.objects.files import Files
//from nf_common_source.code.services.file_system_service.objects.folders import Folders
//from nf_common_source.code.services.reporting_service.wrappers.run_and_log_function_wrapper import run_and_log_function
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.nf_ea_com_universes import NfEaComUniverses
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.model_loader.native_xml_loader.nf_ea_xml_load_to_empty_ea_model_orchestrator
//import orchestrate_nf_ea_xml_load_to_empty_ea_model
//from nf_ea_common_tools_source.b_code.services.session.orchestrators.ea_tools_session_managers import EaToolsSessionManagers

//@run_and_log_function
//def orchestrate_nf_ea_com_universe_to_eapx_migration(

func OrchestrateOlEaComUniverseToEapxMigration(
	//ea_tools_session_manager: EaToolsSessionManagers,
	ea_tools_session_manager orchestrators.EaToolsSessionManagers,
	//nf_ea_com_universe: NfEaComUniverses,
	nf_ea_com_universe *com.OlEaComUniverses,
	//short_name: str,
	short_name string,
	//output_folder: Folders):
	output_folder *object_model.Folders) {

	//ea_export_folder_path = \
	ea_export_folder_path :=
		//os.path.join(
		path.Join(
			//output_folder.absolute_path_string,
			output_folder.AbsolutePathString(),
			//short_name,
			short_name,
			//short_name + '_ea_export')
			short_name+"_ea_export")

	//ea_repository_file_path = \
	ea_repository_file_path :=
		//os.path.join(
		path.Join(
			//ea_export_folder_path,
			ea_export_folder_path,
			//short_name + '_ea_export.eapx')
			short_name+"_ea_export.eapx")

	//ea_repository_file = \
	ea_repository_file := &object_model.Files{}
	//Files(
	ea_repository_file.Initialise(
		//absolute_path_string=ea_repository_file_path)
		ea_repository_file_path)

	//output_xml_file_full_path = \
	output_xml_file_full_path :=
		//os.path.join(
		path.Join(
			//ea_export_folder_path,
			ea_export_folder_path,
			//short_name + '_ea_export.xml'
			short_name+"_ea_export.xml")

	//orchestrate_nf_ea_xml_load_to_empty_ea_model(
	native_xml_loader.OrchestrateNfEaXmlLoadToEmptyEaModel(
		//ea_tools_session_manager=ea_tools_session_manager,
		ea_tools_session_manager,
		//nf_ea_com_universe=nf_ea_com_universe,
		nf_ea_com_universe,
		//output_xml_file_full_path=output_xml_file_full_path,
		output_xml_file_full_path,
		//ea_repository_file=ea_repository_file,
		ea_repository_file,
		//short_name=short_name + '_ea_export')
		short_name+"_ea_export")

}
