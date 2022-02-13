package com

import (
	"fmt"
	"github.com/OntoLedgy/ea_interop_service/code/return_results"
)

//def ReportIDualRepositoryCreationResult(
func ReportIDualRepositoryCreationResult(

	//i_dual_repository_creation_result: IDualRepositoryCreationResults,
	i_dual_repository_creation_result *return_results.IDualRepositoryCreationResults,

	//ea_project_filename: str):
	ea_project_filename string) {

	//i_dual_repository_creation_result_type = \
	i_dual_repository_creation_result_type :=
		//i_dual_repository_creation_result.i_dual_repository_creation_result_type
		i_dual_repository_creation_result.IDualRepositoryCreationResultType

	switch i_dual_repository_creation_result_type {

	//if i_dual_repository_creation_result_type == IDualRepositoryCreationResultTypes.SUCCEEDED:
	case return_results.SUCCEEDED:

		//log_message(
		//'Opened EA project: ' + ea_project_filename)
		fmt.Println(
			"Opened EA project: " + ea_project_filename)

		return

	//if i_dual_repository_creation_result_type == IDualRepositoryCreationResultTypes.FAILED_TO_OPEN_EA:
	case return_results.FAILED_TO_OPEN_EA:
		//log_message(
		fmt.Println(

			//'Failed to open EA')
			"Failed to open EA")

		return

	//if i_dual_repository_creation_result_type == IDualRepositoryCreationResultTypes.FAILED_TO_OPEN_EA_PROJECT:
	case return_results.FAILED_TO_OPEN_EA_PROJECT:

		//log_message(
		fmt.Println(

			//'Failed to open EA project: ' + ea_project_filename)
			"Failed to open EA project: " + ea_project_filename)

		//else:
	default:
		//log_message(
		fmt.Printf(

			//'Unhandled type: ' + str(i_dual_repository_creation_result_type))
			"Unhandled type: %v\n ", i_dual_repository_creation_result_type)

		//sys.exit(
		//-1)
		panic("unknown type")
	}

}
