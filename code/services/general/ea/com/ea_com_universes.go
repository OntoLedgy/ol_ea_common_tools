package com

import (
	"fmt"
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"
	"github.com/OntoLedgy/ea_interop_service/code/processes"
	uuid "github.com/satori/go.uuid"
)

//class EaComUniverses:
type EaComUniverses struct {
	iDualRepository  i_dual_objects.IRepository
	InstanceGuid     uuid.UUID
	ConnectionString string
}

//def __init__(
//self,
func (eaComUniverses *EaComUniverses) Initialise(
	//ea_project_filename: str):
	eaProjectFileName string) {

	//i_dual_repository_creation_result = \
	iDualRepositoryCreationResult :=
		//get_i_dual_repository_creation_result(
		processes.GetIDualRepositoryCreationResult(
			//ea_project_filename=ea_project_filename)
			eaProjectFileName)

	//report_i_dual_repository_creation_result(
	ReportIDualRepositoryCreationResult(
		//i_dual_repository_creation_result=i_dual_repository_creation_result,
		iDualRepositoryCreationResult,
		//ea_project_filename=ea_project_filename)
		eaProjectFileName)

	//TODO - Generalise this check to IRepository level.
	//if not isinstance(i_dual_repository_creation_result.i_dual_repository, IDualRepository):

	//log_message(
	//message='Cannot create repository from ' + ea_project_filename)
	//
	//sys.exit(
	//-1)

	//self.i_dual_repository = \
	eaComUniverses.iDualRepository =
		//i_dual_repository_creation_result.i_dual_repository
		iDualRepositoryCreationResult.IDualRepository

	//self.__open_project_report()
	eaComUniverses.open_project_report()
}

//def __enter__(
//self):
//return \
//self

//def __exit__(
//self,
//exception_type,
//exception_value,
//traceback):
//self.close_project()

//def __open_project_report(
//self):
func (eaComUniverses EaComUniverses) open_project_report() {
	//log_message(
	fmt.Println(
		//"Instance GUID: " +
		"Instance GUID: " +
			//self.i_dual_repository.instance_guid)
			eaComUniverses.InstanceGuid.String())

	//log_message(
	fmt.Println(
		//"Connection string: " +
		"Connection string: " +
			//self.i_dual_repository.connection_string)
			eaComUniverses.ConnectionString)

	//iDualRepository := eaComUniverses.iDualRepository.(i_dual_objects.IDualRepository)

	//log_message(
	fmt.Println(
		//"Library version: " +
		"Library version: " + "XX") //TODO- add reference
	//str(self.i_dual_repository.library_version))
	//	iDualRepository.LiberaryVersion)
}

//def close_project(
//self):
func (EaComUniverses) close_project() {

	//log_message(
	//"Closing EA Project: " +
	//self.i_dual_repository.instance_guid)
	//
	//self.i_dual_repository.exit()
	//
	//log_message(
	//"EA Project Closed")
}
