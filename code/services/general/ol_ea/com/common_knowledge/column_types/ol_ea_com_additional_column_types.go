package column_types

//from enum import auto
//from enum import unique
//from nf_common_source.code.nf.types.column_types import ColumnTypes
//
//
//@unique
//class NfEaComAdditionalColumnTypes(
//ColumnTypes):
type OlEaComAdditionalColumnTypes int

const (
	//ORIGINAL_EA_GUIDS = auto()
	ORIGINAL_EA_GUIDS OlEaComAdditionalColumnTypes = iota
)

func (olEaComAdditionalColumnTypes OlEaComAdditionalColumnTypes) ColumnName() string {
	//def __column_name(
	//self) \
	//-> str:
	//column_name = \
	//column_name_mapping[self]
	//
	//return \
	//column_name
	//
	//column_name = \
	//property(
	//fget=__column_name)
	//

	//column_name_mapping = \
	switch olEaComAdditionalColumnTypes {
	//{
	//NfEaComAdditionalColumnTypes.ORIGINAL_EA_GUIDS: 'original_ea_guids',
	case ORIGINAL_EA_GUIDS:
		return "original_ea_guids"
		//}
	}

	return "unknown"
}
