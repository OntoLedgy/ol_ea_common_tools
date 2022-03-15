package com

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
)

//class EaComManagers:
type EaComManagers struct {
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

//@staticmethod
//def get_ea_repository_file() \
func GetEaRepositoryFile() *object_model.Files { //-> Files:

	//file = \
	file :=
		//select_file()
		file_system_service.SelectFile()

	//return \
	//file
	return file
}
