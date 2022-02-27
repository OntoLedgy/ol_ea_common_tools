package endpoint_managers

//from nf_ea_common_tools_source.b_code.services.session.nf_ea_com_endpoint.orchestrators.nf_managers import NfManagers
//from nf_ea_common_tools_source.b_code.services.session.nf_ea_com_endpoint.orchestrators.universe_managers.nf_ea_com_universe_managers import \
//NfEaComUniverseManagers
import (
	ol_ea_com_endpoint_orchestrators "github.com/OntoLedgy/ol_ea_common_tools/code/services/session/ol_ea_com_endpoint/orchestrators"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/session/ol_ea_com_endpoint/orchestrators/universe_managers"
	session_orchestrators "github.com/OntoLedgy/ol_ea_common_tools/code/services/session/orchestrators"
)

//class NfEaComEndpointManagers(
type OlEaComEndpointManagers struct {
	//NfManagers):
	*ol_ea_com_endpoint_orchestrators.OlManagers
	*universe_managers.OlEaComUniverseManagers
}

//def __init__(
//self,
func (olEaComEndpointManagers *OlEaComEndpointManagers) Initialise(
	//ea_tools_session_manager):
	EaToolsSessionManager session_orchestrators.EaToolsSessionManagers) {

	//NfManagers.__init__(
	//self)

	//self.nf_ea_com_universe_manager = \
	olEaComEndpointManagers.OlEaComUniverseManagers =
		//NfEaComUniverseManagers(
		&universe_managers.OlEaComUniverseManagers{}

	//EaToolsSessionManager=EaToolsSessionManager)
	//TODO - OlManagers should be an Interface

}

//def close(
//self):

func (olEaComEndpointManagers OlEaComEndpointManagers) Close() {

	//self.nf_ea_com_universe_manager.close()

}
