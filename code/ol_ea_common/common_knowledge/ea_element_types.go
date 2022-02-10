package common_knowledge

type EaElementTypes int //EaObjectTypes)

const (
	ACTION EaElementTypes = iota
	ACTION_PIN
	ACTIVITY
	ACTIVITY_PARAMETER
	ACTIVITY_PARTITION
	ACTIVITY_REGION
	ACTOR
	ARTIFACT
	ASSOCIATION_ELEMENT
	BOUNDARY
	CENTRAL_BUFFER_NODE
	CHANGE
	CLASS
	COLLABORATION_ELEMENT
	COLLABORATION_OCCURRENCE
	COMMENT
	COMPONENT
	CONDITIONAL_NODE
	CONSTRAINT
	DATA_STORE
	DATA_TYPE
	DECISION
	DEFECT
	DEPLOYMENT_SPECIFICATION
	DEVICE
	DIAGRAM_FRAME
	ENTITY
	ENTRY_POINT
	ENUMERATION
	EVENT
	EXCEPTION_HANDLER
	EXECUTION_ENVIRONMENT
	EXIT_POINT
	EXPANSION_NODE
	EXPANSION_REGION
	FEATURE
	GUI_ELEMENT
	INFORMATION_ITEM
	INTERACTION
	INTERACTION_FRAGMENT
	INTERACTION_OCCURRENCE
	INTERACTION_STATE
	INTERFACE
	INTERRUPTIBLE_ACTIVITY_REGION
	ISSUE
	LABEL
	LOOP_NODE
	MERGE_NODE
	MESSAGE_ENDPOINT
	NODE
	NOTE
	OBJECT
	OBJECT_NODE
	PACKAGE_ELEMENT
	PARAMETER
	PART
	PORT
	PRIMITIVE_TYPE
	PROTOCOL_STATE_MACHINE
	PROVIDED_INTERFACE
	REGION
	REPORT
	REQUIRED_INTERFACE
	REQUIREMENT
	RISK
	SCREEN
	SEQUENCE_ELEMENT
	SIGNAL
	STATE
	STATE_MACHINE
	STATE_NODE
	SYNCHRONIZATION
	TASK
	TEST
	TEXT
	TIMELINE
	TRIGGER
	UML_DIAGRAM
	USE_CASE_ELEMENT
	USER

	PROXY_CONNECTOR
)

func (eaElementTypes EaElementTypes) get_type_from_name() string {
	switch eaElementTypes {

	case ACTION:
		return "Action"
	case ACTION_PIN:
		return "ActionPin"
	case ACTIVITY:
		return "Activity"
	case ACTIVITY_PARAMETER:
		return "ActivityParameter"
	case ACTIVITY_PARTITION:
		return "ActivityPartition"
	case ACTIVITY_REGION:
		return "ActivityRegion"
	case ACTOR:
		return "Actor"
	case ARTIFACT:
		return "Artifact"
	case ASSOCIATION_ELEMENT:
		return "Association"
	case BOUNDARY:
		return "Boundary"
	case CENTRAL_BUFFER_NODE:
		return "CentralBufferNode"
	case CHANGE:
		return "Change"
	case CLASS:
		return "Class"
	case COLLABORATION_ELEMENT:
		return "Collaboration"
	case COLLABORATION_OCCURRENCE:
		return "CollaborationOccurrence"
	case COMMENT:
		return "Comment"
	case COMPONENT:
		return "Component"
	case CONDITIONAL_NODE:
		return "ConditionalNode"
	case CONSTRAINT:
		return "Constraint"
	case DATA_STORE:
		return "DataStore"
	case DATA_TYPE:
		return "DataType"
	case DECISION:
		return "Decision"
	case DEFECT:
		return "Defect"
	case DEPLOYMENT_SPECIFICATION:
		return "DeploymentSpecification"
	case DEVICE:
		return "Device"
	case DIAGRAM_FRAME:
		return "DiagramFrame"
	case ENTITY:
		return "Entity"
	case ENTRY_POINT:
		return "EntryPoint"
	case ENUMERATION:
		return "Enumeration"
	case EVENT:
		return "Event"
	case EXCEPTION_HANDLER:
		return "ExceptionHandler"
	case EXECUTION_ENVIRONMENT:
		return "ExecutionEnvironment"
	case EXIT_POINT:
		return "ExitPoint"
	case EXPANSION_NODE:
		return "ExpansionNode"
	case EXPANSION_REGION:
		return "ExpansionRegion"
	case FEATURE:
		return "Feature"
	case GUI_ELEMENT:
		return "GUIElement"
	case INFORMATION_ITEM:
		return "InformationItem"
	case INTERACTION:
		return "Interaction"
	case INTERACTION_FRAGMENT:
		return "InteractionFragment"
	case INTERACTION_OCCURRENCE:
		return "InteractionOccurrence"
	case INTERACTION_STATE:
		return "InteractionState"
	case INTERFACE:
		return "Interface"
	case INTERRUPTIBLE_ACTIVITY_REGION:
		return "InterruptibleActivityRegion"
	case ISSUE:
		return "Issue"
	case LABEL:
		return "Label"
	case LOOP_NODE:
		return "LoopNode"
	case MERGE_NODE:
		return "MergeNode"
	case MESSAGE_ENDPOINT:
		return "MessageEndpoint"
	case NODE:
		return "Node"
	case NOTE:
		return "Note"
	case OBJECT:
		return "Object"
	case OBJECT_NODE:
		return "ObjectNode"
	case PACKAGE_ELEMENT:
		return "Package"
	case PARAMETER:
		return "Parameter"
	case PART:
		return "Part"
	case PORT:
		return "Port"
	case PRIMITIVE_TYPE:
		return "PrimitiveType"
	case PROTOCOL_STATE_MACHINE:
		return "ProtocolStateMachine"
	case PROVIDED_INTERFACE:
		return "ProvidedInterface"
	case REGION:
		return "Region"
	case REPORT:
		return "Report"
	case REQUIRED_INTERFACE:
		return "RequiredInterface"
	case REQUIREMENT:
		return "Requirement"
	case RISK:
		return "Risk"
	case SCREEN:
		return "Screen"
	case SEQUENCE_ELEMENT:
		return "Sequence"
	case SIGNAL:
		return "Signal"
	case STATE:
		return "State"
	case STATE_MACHINE:
		return "StateMachine"
	case STATE_NODE:
		return "StateNode"
	case SYNCHRONIZATION:
		return "Synchronization"
	case TASK:
		return "Task"
	case TEST:
		return "Test"
	case TEXT:
		return "Text"
	case TIMELINE:
		return "TimeLine"
	case TRIGGER:
		return "Trigger"
	case UML_DIAGRAM:
		return "UMLDiagram"
	case USE_CASE_ELEMENT:
		return "UseCase"
	case USER:
		return "User"
	case PROXY_CONNECTOR:
		return "ProxyConnector"

	}
	return "unknown"
}
