package creators

import (
	"github.com/OntoLedgy/ol_ea_common_tools/code/ol_ea_common/objects"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
)

//def register_universe_copy(
func RegisterUniverseCopy(
	//ea_tools_session_manager,
	eaToolsSessionManager orchestrators.EaToolsSessionManagers,
	//nf_ea_com_universe,
	OlEaComUniverse *com.OlEaComUniverses,
	//ea_repository_copy):
	EaRepositoryCopy *objects.EaRepositories) {

	//nf_ea_com_universe_manager = \
	OlEaComUniverseManager :=
		//ea_tools_session_manager.nf_ea_com_endpoint_manager.nf_ea_com_universe_manager
		eaToolsSessionManager.OlEaComEndpointManagers.OlEaComUniverseManagers

	//nf_ea_com_universe_manager.nf_ea_com_universe_dictionary[ea_repository_copy] = \
	OlEaComUniverseManager.OlEaComUniverseMap[EaRepositoryCopy] =
		//nf_ea_com_universe
		OlEaComUniverse

}
