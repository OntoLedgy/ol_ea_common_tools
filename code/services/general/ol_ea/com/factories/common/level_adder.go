package common

//from nf_common_source.code.services.dataframe_service.dataframe_mergers import left_merge_dataframes
//from nf_co t "github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"

import "github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"

//def add_level(
func AddLevel(
	//dataframe: DataFrame,
	dataframe,
	//next_level_dataframe: DataFrame):
	nextLevelDataFrame string) *dataframes.DataFrames {

	//thin_ea_elements_columns = \
	//list(next_level_dataframe.columns)

	thinEaElementsColumns := &dataframes.DataFrames{}

	//foreign_key_dataframe_other_column_rename_dictionary = \
	//dict()

	//nf_uuids_column_name = \
	//NfColumnTypes.NF_UUIDS.column_name

	//for thin_ea_elements_column in thin_ea_elements_columns:
	//if thin_ea_elements_column != nf_uuids_column_name:
	//foreign_key_dataframe_other_column_rename_dictionary[thin_ea_elements_column] = \
	//thin_ea_elements_column

	//dataframe = \
	//left_merge_dataframes(
	//master_dataframe=dataframe,
	//master_dataframe_key_columns=[nf_uuids_column_name],
	//merge_suffixes=['_this_level', '_next_level'],
	//foreign_key_dataframe=next_level_dataframe,
	//foreign_key_dataframe_fk_columns=[nf_uuids_column_name],
	//foreign_key_dataframe_other_column_rename_dictionary=foreign_key_dataframe_other_column_rename_dictionary)

	//return \
	//dataframe
	return thinEaElementsColumns
}
