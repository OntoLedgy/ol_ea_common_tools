package ea_tools_session_service

import (
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/objects"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ea/com"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
)

//import tkinter as tk
//from tkinter import simpledialog
//from nf_common_source.code.services.file_system_service.objects.files import Files
//from nf_ea_common_tools_source.b_code.services.general.ea.com.ea_com_managers import EaComManagers
//from nf_ea_common_tools_source.b_code.services.general.ea.com.ea_repository_factory import create_ea_repository
//from nf_ea_common_tools_source.b_code.services.general.ea.com.ea_repository_factory import create_empty_ea_repository
//from nf_ea_common_tools_source.b_code.nf_ea_common.objects.ea_repositories import EaRepositories

//def get_repository() \
//-> EaRepositories:
func GetRepository() *objects.EaRepositories {

	//with \
	//EaComManagers() \

	//as ea_com_manager:

	//ea_repository_file = \
	//ea_com_manager.get_ea_repository_file()
	eaRepositoryFile := com.GetEaRepositoryFile()

	//short_name = \
	//__get_short_name() //TODO
	shortName := "to be done"

	//ea_repository = \
	eaRepository :=
		//get_repository_using_file_and_short_name(
		getRepositoryUsingFileAndShortName(
			eaRepositoryFile,
			shortName)

	//ea_repository_file=ea_repository_file,
	//short_name=short_name)

	//return \
	//ea_repository

	return eaRepository

}

//def get_repository_using_file_and_short_name(
func getRepositoryUsingFileAndShortName(
	//ea_repository_file: Files,
	eaRepositoryFile *object_model.Files,
	//short_name: str) \
	shortName string) *objects.EaRepositories { //-> EaRepositories:

	//ea_repository = \
	eaRepository :=
		//create_ea_repository(
		com.CreateEaRepository(
			//ea_repository_file=ea_repository_file,
			eaRepositoryFile,
			//short_name=short_name)
			shortName)

	//return \
	//ea_repository
	return eaRepository

}

//def get_empty_ea_repository_with_short_name(
//short_name: str) \
//-> EaRepositories:
//ea_repository = \
//create_empty_ea_repository(
//short_name=short_name)
//
//return \
//ea_repository

//def __get_short_name():
//root = \
//tk.Tk()
//
//root.withdraw()
//
//short_name = \
//simpledialog.askstring(
//title="Input",
//prompt="Enter The Repository's Short Name")
//
//return \
//short_name
