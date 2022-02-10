package common_knowledge

type EaConnectorTypes int //(EaObjectTypes):

const (
	ABSTRACTION EaConnectorTypes = iota
	AGGREGATION
	ASSEMBLY
	ASSOCIATION
	COLLABORATION
	COMMUNICATION_PATH
	CONNECTOR
	CONTROL_FLOW
	DELEGATE
	DEPENDENCY
	DEPLOYMENT
	ER_LINK
	EXTENSION
	GENERALIZATION
	INFORMATION_FLOW
	INSTANTIATION
	INTERRUPT_FLOW
	MANIFEST
	NESTING
	NOTE_LINK
	OBJECT_FLOW
	PACKAGE
	PROTOCOL_CONFORMANCE
	PROTOCOL_TRANSITION
	REALISATION
	SEQUENCE
	STATE_FLOW
	SUBSTITUTION
	USAGE
	USE_CASE
)

func (eaConnectorTypes EaConnectorTypes) get_type_from_name() string {

	switch eaConnectorTypes {

	case ABSTRACTION:
		return "Abstraction"
	case AGGREGATION:
		return "Aggregation"
	case ASSEMBLY:
		return "Assembly"
	case ASSOCIATION:
		return "Association"
	case COLLABORATION:
		return "Collaboration"
	case COMMUNICATION_PATH:
		return "CommunicationPath"
	case CONNECTOR:
		return "Connector"
	case CONTROL_FLOW:
		return "ControlFlow"
	case DELEGATE:
		return "Delegate"
	case DEPENDENCY:
		return "Dependency"
	case DEPLOYMENT:
		return "Deployment"
	case ER_LINK:
		return "ERLink"
	case EXTENSION:
		return "Extension"
	case GENERALIZATION:
		return "Generalization"
	case INFORMATION_FLOW:
		return "InformationFlow"
	case INSTANTIATION:
		return "Instantiation"
	case INTERRUPT_FLOW:
		return "InterruptFlow"
	case MANIFEST:
		return "Manifest"
	case NESTING:
		return "Nesting"
	case NOTE_LINK:
		return "NoteLink"
	case OBJECT_FLOW:
		return "ObjectFlow"
	case PACKAGE:
		return "Package"
	case PROTOCOL_CONFORMANCE:
		return "ProtocolConformance"
	case PROTOCOL_TRANSITION:
		return "ProtocolTransition"
	case REALISATION:
		return "Realisation"
	case SEQUENCE:
		return "Sequence"
	case STATE_FLOW:
		return "StateFlow"
	case SUBSTITUTION:
		return "Substitution"
	case USAGE:
		return "Usage"
	case USE_CASE:
		return "UseCase"

	}
	return "unknown"
}



