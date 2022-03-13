package ea_t

//from enum import auto
//from enum import unique
//from nf_ea_common_tools_source.b_code.nf_ea_common.common_knowledge.column_types.ea_t.ea_t_column_types import EaTColumnTypes
//
//
//@unique
//class EaTXrefColumnTypes(
//EaTColumnTypes):
type EaTXrefColumnTypes int

const (

	//T_XREF_EA_GUIDS = auto()
	T_XREF_EA_GUIDS EaTXrefColumnTypes = iota
	//T_XREF_NAMES = auto()
	T_XREF_NAMES
	//T_XREF_TYPES = auto()
	T_XREF_TYPES
	//T_XREF_DESCRIPTIONS = auto()
	T_XREF_DESCRIPTIONS
	//T_XREF_CLIENT_EA_GUIDS = auto()
	T_XREF_CLIENT_EA_GUIDS
)

//def __column_name(
//self) \
//-> str:
func (eaTXrefColumnType EaTXrefColumnTypes) ColumnName() string {

	//column_name = \
	//column_name_mapping[self]
	//
	//return \
	//column_name

	//column_name_mapping = \
	switch eaTXrefColumnType {
	//{

	//EaTXrefColumnTypes.T_XREF_EA_GUIDS: 'XrefID',
	case T_XREF_EA_GUIDS:
		return "XrefID"
	//EaTXrefColumnTypes.T_XREF_NAMES: 'Name',
	case T_XREF_NAMES:
		return "Name"
	//EaTXrefColumnTypes.T_XREF_TYPES: 'Type',
	case T_XREF_TYPES:
		return "Type"
	//EaTXrefColumnTypes.T_XREF_DESCRIPTIONS: 'Description',
	case T_XREF_DESCRIPTIONS:
		return "Description"
	//EaTXrefColumnTypes.T_XREF_CLIENT_EA_GUIDS: 'Client'
	case T_XREF_CLIENT_EA_GUIDS:
		return "Client"

		//}
	}

	return "not implemented"
}

//column_name = \
//property(
//fget=__column_name)

//def __nf_column_name(
//self) \
//-> str:
func (eaTXrefColumnType EaTXrefColumnTypes) OlColumnName() string {

	//nf_column_name = \
	//nf_column_name_mapping[self]
	//
	//return \
	//nf_column_name

	//nf_column_name = \
	//property(
	//fget=__nf_column_name)

	//nf_column_name_mapping = \
	switch eaTXrefColumnType {
	//{
	//EaTXrefColumnTypes.T_XREF_EA_GUIDS: 't_xref_ea_guids',
	case T_XREF_EA_GUIDS:
		return "t_xref_ea_guids"
	//EaTXrefColumnTypes.T_XREF_NAMES: 't_xref_names',
	case T_XREF_NAMES:
		return "t_xref_names"
	//EaTXrefColumnTypes.T_XREF_TYPES: 't_xref_types',
	case T_XREF_TYPES:
		return "t_xref_types"
	//EaTXrefColumnTypes.T_XREF_DESCRIPTIONS: 't_xref_descriptions',
	case T_XREF_DESCRIPTIONS:
		return "t_xref_descriptions"
	//EaTXrefColumnTypes.T_XREF_CLIENT_EA_GUIDS: 't_xref_client_ea_guids'
	case T_XREF_CLIENT_EA_GUIDS:
		return "t_xref_client_ea_guids"
		//}
	}

	return ""
}
