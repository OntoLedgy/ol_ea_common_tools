package com

//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.nf_ea_com_registries import NfEaComRegistries
//from nf_ea_common_tools_source.b_code.nf_ea_common.objects.ea_repositories import EaRepositories
//from pandas import DataFrame
//
//from nf_ea_common_tools_source.b_code.services.session.processes.creators.nf_ea_com_universe_copier import \
//register_universe_copy

import (
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/objects"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/processes/creators"
)

//class NfEaComUniverses:
type OlEaComUniverses struct {
	*orchestrators.EaToolsSessionManagers
	*objects.EaRepositories
	*OlEaComRegistries
}

//def __init__(
//self,
func (olEaComUniverses *OlEaComUniverses) Initialise(
	//ea_tools_session_manager,
	eaToolsSessionManager *orchestrators.EaToolsSessionManagers,
	//ea_repository: EaRepositories):
	eaRepository *objects.EaRepositories) {

	//self.ea_tools_session_manager = \
	olEaComUniverses.EaToolsSessionManagers =
		//ea_tools_session_manager
		eaToolsSessionManager

	//self.ea_repository = \
	olEaComUniverses.EaRepositories =
		//ea_repository
		eaRepository

	//self.nf_ea_com_registry = \
	olEaComUniverses.OlEaComRegistries =
		//NfEaComRegistries(
		//self)
		&OlEaComRegistries{}

}

//def __enter__(
//self):
//return \
//self
//
//def __exit__(
//self,
//exception_type,
//exception_value,
//traceback):
//pass

//def copy(
//self,
func (olEaComUniverses *OlEaComUniverses) Copy(
	//short_name=str()):
	shortName string) *OlEaComUniverses {

	//ea_repository_copy = \
	eaRepositoryCopy :=
		//EaRepositories(
		&objects.EaRepositories{
			//short_name=short_name,
			ShortName: shortName,
			//ea_repository_file=self.ea_repository.ea_repository_file)
			EaRepositoryFile: olEaComUniverses.EaRepositories.EaRepositoryFile}

	//self_copy = \
	selfCopy :=
		//NfEaComUniverses(
		&OlEaComUniverses{
			//ea_tools_session_manager=self.ea_tools_session_manager,
			olEaComUniverses.EaToolsSessionManagers,
			//ea_repository=ea_repository_copy)
			eaRepositoryCopy,
			//self_copy.nf_ea_com_registry = \
			//self.nf_ea_com_registry.copy()
			olEaComUniverses.OlEaComRegistries.Copy()}

	//register_universe_copy(
	creators.RegisterUniverseCopy(
		//ea_tools_session_manager=self.ea_tools_session_manager,
		olEaComUniverses.EaToolsSessionManager,
		//nf_ea_com_universe=self_copy,
		selfCopy,
		//ea_repository_copy=ea_repository_copy)
		eaRepositoryCopy)

	//return \
	//self_copy
	return selfCopy
}

//def export_dataframes(
//self,
//project_short_name: str,
//output_folder_name: str):
//self.nf_ea_com_registry.export_dataframes(
//short_name=project_short_name,
//output_folder_name=output_folder_name)

//def get_ea_stereotypes(
//self) \
//-> DataFrame:
//ea_stereotypes = \
//self.nf_ea_com_registry.get_ea_stereotypes()
//
//return \
//ea_stereotypes

//def get_ea_stereotype_groups(
//self) \
//-> DataFrame:
//ea_stereotype_groups = \
//self.nf_ea_com_registry.get_ea_stereotype_groups()
//
//return \
//ea_stereotype_groups
