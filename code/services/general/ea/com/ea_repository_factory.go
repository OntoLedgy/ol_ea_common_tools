package com

import (
	"github.com/OntoLedgy/ea_interop_service/code/processes"
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/objects"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
)

func CreateEaRepository(
	eaRepositoryFile *object_model.Files,
	eaRepositoryShortName string) *objects.EaRepositories {

	//i_dual_repository_creation_result :=
	i_dual_repository_creation_result :=
		//get_i_dual_repository_creation_result(
		processes.GetIDualRepositoryCreationResult(
			//eaRepositoryFile.absolute_path_string)
			eaRepositoryFile.AbsolutePathString())

	//ReportIDualRepositoryCreationResult(
	ReportIDualRepositoryCreationResult(
		//	i_dual_repository_creation_result = i_dual_repository_creation_result,
		i_dual_repository_creation_result,
		//	ea_project_filename = eaRepositoryFile.absolute_path_string)
		eaRepositoryFile.AbsolutePathString())

	//i_dual_repository = \
	i_dual_repository :=
		//i_dual_repository_creation_result.i_dual_repository
		i_dual_repository_creation_result.IDualRepository

	//if not isinstance(i_dual_repository, IDualRepository):
	//log_message(
	//	message = 'Cannot create repository from ' + eaRepositoryFile.absolute_path_string)
	//
	//sys.exit(
	//	-1)

	i_dual_repository.InitialiseRepository()

	//ea_repository = \
	eaRepository :=
		//EaRepositories(
		&objects.EaRepositories{
			//	ea_repository_file=ea_repository_file,
			eaRepositoryShortName,
			//	short_name=short_name)
			eaRepositoryFile}

	//EaRepositoryMappers.store_map(
	//	ea_repository = ea_repository,
	//	i_dual_repository = i_dual_repository)

	//return \
	//ea_repository
	return eaRepository

}

func CreateEmptyEaRepository(short_name string) *objects.EaRepositories {

	eaRepository := new(objects.EaRepositories)

	eaRepository.ShortName = short_name

	return eaRepository

}
