package testing

import (
	"fmt"
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/objects"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ea/com"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/processes"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/processes/populators"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"github.com/go-ole/go-ole"
	"testing"
)

func TestEASessionManager(t *testing.T) {

	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	eaRepositoryFile := &object_model.Files{}
	eaRepositoryFile.Initialise("D:\\S\\python\\bCLEARer\\nf_ea_common_tools\\testing\\_ea_export.eapx")

	eaRepository := com.CreateEaRepository(
		eaRepositoryFile,
		"")

	fmt.Println(eaRepository.EaRepositoryFile.Path.String())

}

func TestInitialiseEmptyComRepository(t *testing.T) {

	emptyEaComCollection :=
		populators.CreateEmptyOlEaComMapOfCollections()

	fmt.Println(emptyEaComCollection)

}

func TestInitialiseEaCom(t *testing.T) {

	emptyEaComCollection :=
		processes.InitialiseOlEaComMap()

	fmt.Println(emptyEaComCollection)

}

func TestEARepositoryMapper(t *testing.T) {

	eaRepositoryMapper := session.EaRepositoryMappers{}

	eaRepository := &objects.EaRepositories{}

	iDualRepository := i_dual_objects.IDualRepository{}

	eaRepositoryMapper.StoreMap(
		eaRepository,
		iDualRepository)

	fmt.Println(eaRepositoryMapper)
}
