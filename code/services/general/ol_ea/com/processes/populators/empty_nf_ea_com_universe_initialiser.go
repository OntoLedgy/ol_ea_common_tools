package populators

//from pandas import DataFrame
//from pandas import read_sql_query
//import importlib
//import os
//import pyodbc
//from nf_ea_common_tools_source.b_code.services.general.nf_ea.com.common_knowledge.collection_types.nf_ea_com_collection_types import NfEaComCollectionTypes

import (
	"container/list"
	"github.com/OntoLedgy/ol_ea_common_tools/code/services/general/ol_ea/com/common_knowledge/collection_types"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/lists"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

//def create_empty_nf_ea_com_dictionary_of_collections() \
//-> dict:
func CreateEmptyOlEaComMapOfCollections() *map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames {

	//database_connection = \
	databaseConnection :=
		getDatabaseConnection()
	//__get_database_connection()

	//dictionary_of_collections = \
	//__import_tables_from_database(
	//database_connection=database_connection)

	mapOfCollections :=
		importTablesFromDatabase(
			databaseConnection)

	//return \
	//dictionary_of_collections
	defer databaseConnection.Close()
	return mapOfCollections
}

//def __get_database_connection():
func getDatabaseConnection() contract.IDatabases {

	//module = \
	//importlib.import_module(
	//name='nf_ea_common_tools_source.resources.templates')

	//module_path_string = \
	//module.__path__._path[0]

	//ResourceFullFileName = \
	//os.path.join(
	//module_path_string,
	//TODO - Should be configuration
	ResourceFullFileName :=
		//'empty_nf_ea_com.accdb')
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\empty_ol_ea_com.accdb"

	//database_connection = \
	databaseConnection :=
		connectWithAccessDatabase(
			ResourceFullFileName)
	//__connect_with_access_database(
	//access_database_full_path=ResourceFullFileName)

	//return \
	//database_connection
	return databaseConnection

}

//def __connect_with_access_database(

func connectWithAccessDatabase(
	//access_database_full_path: str) \
	//-> pyodbc.Connection:
	resource_full_file_name string) contract.IDatabases {

	//access_driver = \
	//r'{Microsoft Access Driver (*.mdb, *.accdb)}'
	//
	//connection_string = \
	//r'Driver={};DBQ={}'.format(access_driver, access_database_full_path)
	//
	//database_connection = \
	//pyodbc.connect(
	//connection_string)
	//
	//return \
	//database_connection
	//TODO - wrap this in facade.

	databaseFactory :=
		&database_i_o_service.DatabaseFactory{
			DatabaseName:   resource_full_file_name,
			DatabaseType:   configurations.DbTypeAccess,
			SystemDatabase: "C:\\Users\\khanm\\AppData\\Roaming\\Microsoft\\Access\\System.mdw"}

	database :=
		databaseFactory.New()

	database.Connect()

	return database

}

//def __import_tables_from_database(
func importTablesFromDatabase(
	//database_connection: pyodbc.Connection) \
	databaseConnection contract.IDatabases) *map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames { //-> dict:

	//table_names_list = \
	tableNamesList := importTableNamesListFromDatabase(
		databaseConnection)
	//__import_table_names_list_from_database(
	//database_connection=database_connection)

	//dictionary_of_input_dataframes = \
	mapOfInputDataFrames :=
		getMapOfTablesFromDatabase(
			tableNamesList,
			databaseConnection)

	//__get_dictionary_of_tables_from_database(
	//table_names_list=table_names_list,
	//database_connection=database_connection)

	//return \
	//dictionary_of_input_dataframes

	return mapOfInputDataFrames

}

//def __get_dictionary_of_tables_from_database(
func getMapOfTablesFromDatabase(
	//table_names_list: list,
	tableNamesList lists.Lists,
	//database_connection: pyodbc.Connection) \
	connection contract.IDatabases) *map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames { //-> dict:

	//dictionary_of_dataframes = \
	mapOfDataframes := &map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames{}
	//{}

	//for table_name in table_names_list:
	for table := tableNamesList.Front(); table != nil; table = table.Next() {
		addTableToMap(
			table,
			connection,
			*mapOfDataframes)
	}

	//__add_table_to_dictionary(
	//table_name=table_name,
	//database_connection=database_connection,
	//dictionary_of_dataframes=dictionary_of_dataframes)

	//return \
	//dictionary_of_dataframes
	return mapOfDataframes

}

//def __add_table_to_dictionary(
func addTableToMap(
	//table_name: str,
	table *list.Element,
	//database_connection: pyodbc.Connection,
	connection contract.IDatabases,
	//dictionary_of_dataframes: dict):
	MapOfDataFrames map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames) {

	//dataframe = \
	dataframe :=
		//__import_dataframe_from_database(
		importDataFrameFromDatabase(
			//table_name=table_name,
			table.Value.(string),
			//database_connection=database_connection)
			connection)

	//collection_type = \
	//NfEaComCollectionTypes.get_collection_type_from_name(
	//name=table_name)

	collectionType := collection_types.GetCollectionTypeFromName(table.Value.(string))

	if collectionType == nil {

		//if not collection_type:
		//raise \
		//NotImplementedError()
		panic("not implemented collection type")

	}

	//dictionary_of_dataframes[collection_type] = \
	//dataframe
	MapOfDataFrames[*collectionType] = dataframe
}

//def __import_table_names_list_from_database(
//database_connection: pyodbc.Connection) \
func importTableNamesListFromDatabase(
	connection contract.IDatabases) lists.Lists {
	//-> list:

	//table_names_list_name = \
	//tableNamesListName :=
	////'table_name_list'
	//	"table_name_list"
	//select_query = \
	//"SELECT * FROM {}".format(table_names_list_name)
	//
	//table_names_dataframe = \
	tables_slice, _ :=
		connection.GetTables()

	var tables_struct_slice []string

	for _, item := range tables_slice {
		tables_struct_slice = append(tables_struct_slice, item.Name)
	}

	table_names_dataframe :=
		&dataframes.DataFrames{
			DataFrame: dataframe.New(
				series.New(tables_struct_slice, series.String, "Name"))}

	//read_sql_query(
	//select_query,
	//database_connection)

	//list_of_table_names = \
	listOfTableNames :=
		lists.Lists{
			*list.New()}

	//table_names_dataframe['name'].tolist()
	for _, table := range table_names_dataframe.Col("Name").Records() {
		listOfTableNames.PushFront(table)
	}

	//return \
	//list_of_table_names
	return listOfTableNames
}

//def __import_dataframe_from_database(
func importDataFrameFromDatabase(
	//table_name: str,
	table string,
	//database_connection: pyodbc.Connection) \
	connection contract.IDatabases) *dataframes.DataFrames {
	//-> DataFrame:

	transaction := connection.BeginDatabaseTransaction()

	var tableData [][]string

	//select_query = \
	//"SELECT * FROM {}".format(table_name)

	//dataframe = \
	//read_sql_query(
	//select_query,
	//database_connection)
	transaction.Select(
		&tableData,
		`SELECT * FROM $1`, table)

	tablesDataFrame := &dataframes.DataFrames{
		DataFrame: dataframe.LoadRecords(tableData),
	}

	//return \
	//dataframe
	return tablesDataFrame
}
