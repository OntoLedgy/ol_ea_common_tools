package creators

import (
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
)

//from nf_ea_common_tools_source.b_code.services.session.orchestrators.ea_tools_session_managers import \
//EaToolsSessionManagers

//def create_empty_nf_ea_universe(
func CreateEmptyOlEAUniverse(
	//ea_tools_session_manager: EaToolsSessionManagers,
	eaToolsSessionManager orchestrators.EaToolsSessionManagers,
	//short_name: str):
	shortName string) *com.OlEaComUniverses {

	//empty_universe_repository = \
	emptyUniverseRepository :=
		//ea_tools_session_manager.create_empty_ea_repository_with_short_name(
		eaToolsSessionManager.CreateEmptyEaRepositoryWithShortName(
			//short_name=short_name)
			shortName)

	//nf_ea_com_universe_manager = \
	OlEaComUniverseManager :=
		//ea_tools_session_manager.nf_ea_com_endpoint_manager.nf_ea_com_universe_manager
		eaToolsSessionManager.OlEaComEndpointManagers.OlEaComUniverseManagers

	//empty_nf_ea_com_universe = \
	emptyOlEaComUniverse :=
		//nf_ea_com_universe_manager.nf_ea_com_universe_dictionary[empty_universe_repository]
		OlEaComUniverseManager.OlEaComUniverseMap[emptyUniverseRepository]

	//return \
	//empty_nf_ea_com_universe
	return emptyOlEaComUniverse

}
