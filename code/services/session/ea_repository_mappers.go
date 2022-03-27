//https://github.com/boro-alpha/nf_ea_common_tools/blob/master/nf_ea_common_tools_source/b_code/services/session/ea_repository_mappers.py
package session

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"
	"github.com/OntoLedgy/ol_common_services/code/ol/golang_extensions/collections"
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/objects"
)

//class EaRepositoryMappers:
type EaRepositoryMappers struct {
	//__map = \
	//NfBimappings({})

	olBimapping collections.OlBimappings
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
func (eaRepositoryMapper EaRepositoryMappers) StoreMap(
	//ea_repository: EaRepositories,
	eaRepository *objects.EaRepositories,
	//i_dual_repository: IDualRepository):
	iDualRepository i_dual_objects.IDualRepository) {

	//EaRepositoryMappers.__map.add_mapping(
	eaRepositoryMapper.olBimapping.AddMapping(
		//domain_value=ea_repository,
		eaRepository,
		//range_value=i_dual_repository)
		iDualRepository)
}

//@staticmethod
//def get_i_dual_repository(
func (eaRepositoryMapper EaRepositoryMappers) GetIDualRepository(
	//ea_repository: EaRepositories
	ea_repository objects.EaRepositories) *i_dual_objects.IDualRepository {

	//i_dual_repository = \
	iDualRepository :=
		//EaRepositoryMappers.__map.try_get_range_using_domain(
		eaRepositoryMapper.olBimapping.TryGetRangeUsingDomain(
			//domain_key=ea_repository).value
			ea_repository).(i_dual_objects.IDualRepository)

	//return \
	//i_dual_repository
	return &iDualRepository
}

//@staticmethod
//def close_all_ea_repositories():
func close_all_ea_repositories() {

	//for i_dual_repository in EaRepositoryMappers.__map.get_range():
	//i_dual_repository.exit()

	//EaRepositoryMappers.__map = \
	//NfBimappings({})

}
