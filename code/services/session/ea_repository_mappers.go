package session

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/objects"
)

//class EaRepositoryMappers:
type EaRepositoryMappers struct {
	//__map = \
	//NfBimappings({})
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
//def store_map(
func (EaRepositoryMappers) store_map(
	//ea_repository: EaRepositories,
	ea_repository *objects.EaRepositories,
	//i_dual_repository: IDualRepository):
	i_dual_repository i_dual_objects.IDualRepository) {

	//EaRepositoryMappers.__map.add_mapping(
	//domain_value=ea_repository,
	//range_value=i_dual_repository)

}

//@staticmethod
//def get_i_dual_repository(
func get_i_dual_repository(
	//ea_repository: EaRepositories
	ea_repository objects.EaRepositories) *i_dual_objects.IDualRepository {

	//i_dual_repository = \
	i_dual_repository :=
		&i_dual_objects.IDualRepository{}
	//TODO - implement map
	//EaRepositoryMappers.__map.try_get_range_using_domain(
	//domain_key=ea_repository).value

	//return \
	//i_dual_repository
	return i_dual_repository
}

//@staticmethod
//def close_all_ea_repositories():
func close_all_ea_repositories() {

	//for i_dual_repository in EaRepositoryMappers.__map.get_range():

	//i_dual_repository.exit()

	//EaRepositoryMappers.__map = \
	//NfBimappings({})
}
