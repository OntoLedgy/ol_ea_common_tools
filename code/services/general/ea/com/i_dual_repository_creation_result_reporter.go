package com

func create_ea_repository(
	ea_repository_file string, //Files,
	short_name string) EaRepositories {

	eaRepository := new
	EaRepositories()

	return eaRpository
}

i_dual_repository_creation_result = \
get_i_dual_repository_creation_result(
ea_project_filename = ea_repository_file.absolute_path_string)

report_i_dual_repository_creation_result(
i_dual_repository_creation_result = i_dual_repository_creation_result,
ea_project_filename = ea_repository_file.absolute_path_string)

i_dual_repository = \
i_dual_repository_creation_result.i_dual_repository

if not isinstance(i_dual_repository, IDualRepository):
log_message(
message = 'Cannot create repository from ' + ea_repository_file.absolute_path_string)

sys.exit(
-1)

ea_repository = \
EaRepositories(
ea_repository_file = ea_repository_file,
short_name =short_name)

EaRepositoryMappers.store_map(
ea_repository = ea_repository,
i_dual_repository = i_dual_repository)

return \
ea_repository


def create_empty_ea_repository(
short_name: str) \
-> EaRepositories:
ea_repository = \
EaRepositories(
short_name = short_name)

return \
ea_repository
