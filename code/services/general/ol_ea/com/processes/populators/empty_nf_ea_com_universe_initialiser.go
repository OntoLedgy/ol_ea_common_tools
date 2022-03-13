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
		//__get_database_connection()
		getDatabaseConnection()

	//dictionary_of_collections = \
	mapOfCollections :=
		//__import_tables_from_database(
		importTablesFromDatabase(
			//database_connection=database_connection)
			databaseConnection)

	//return \
	//dictionary_of_collections
	databaseConnection.Close()

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
	//TODO - Should be configuration
	ResourceFullFileName :=
		//os.path.join(
		//module_path_string,
		//'empty_nf_ea_com.accdb')
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\ol_ea_common_tools\\resources\\templates\\empty_ol_ea_com.accdb"

	//database_connection = \
	databaseConnection :=
		//__connect_with_access_database(
		connectWithAccessDatabase(
			//access_database_full_path=ResourceFullFileName)
			ResourceFullFileName)

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

	//TODO - wrap this in facade.
	databaseFactory :=
		&database_i_o_service.DatabaseFactory{
			DatabaseName:   resource_full_file_name,
			DatabaseType:   configurations.DbTypeAccess,
			SystemDatabase: "C:\\Users\\khanm\\AppData\\Roaming\\Microsoft\\Access\\System.mdw"}

	//database_connection = \
	database :=
		databaseFactory.
			New()

	//pyodbc.connect(
	//connection_string)
	database.
		Connect()

	//return \
	//database_connection

	return database

}

//def __import_tables_from_database(
func importTablesFromDatabase(
	//database_connection: pyodbc.Connection) \
	databaseConnection contract.IDatabases) *map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames { //-> dict:

	//table_names_list = \
	tableNamesList :=
		//__import_table_names_list_from_database(
		importTableNamesListFromDatabase(
			//database_connection=database_connection)
			databaseConnection)

	//dictionary_of_input_dataframes = \
	mapOfInputDataFrames :=
		//__get_dictionary_of_tables_from_database(
		getMapOfTablesFromDatabase(
			//table_names_list=table_names_list,
			tableNamesList,
			//database_connection=database_connection)
			databaseConnection)

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
	mapOfDataframes :=
		//{}
		&map[collection_types.OlEaComCollectionTypes]*dataframes.DataFrames{}

	//for table_name in table_names_list:
	for table := tableNamesList.Front(); table.Next() != nil; table = table.Next() {

		//__add_table_to_dictionary(
		addTableToMap(
			//table_name=table_name,
			table,
			//database_connection=database_connection,
			connection,
			//dictionary_of_dataframes=dictionary_of_dataframes)
			*mapOfDataframes)
	}

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
	collectionType :=
		//NfEaComCollectionTypes.get_collection_type_from_name(
		collection_types.GetCollectionTypeFromName(
			//name=table_name)
			table.Value.(string))

	//if not collection_type:
	if collectionType == nil {

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
	tableNamesListTableName := "table_name_list"

	//select_query = \
	selectQuery :=
		//"SELECT * FROM {}".format(table_names_list_name)
		"SELECT * FROM " + tableNamesListTableName

	var transaction = connection.
		BeginDatabaseTransaction()

	defer transaction.
		Commit()

	tables := new([]string)

	//read_sql_query(
	//select_query,
	//database_connection)
	transaction.Select(
		tables,
		selectQuery)

	tableNameSeries := series.New(
		*tables,
		series.String,
		"Name")

	//table_names_dataframe = \
	table_names_dataframe :=
		&dataframes.DataFrames{
			DataFrame: dataframe.New(
				tableNameSeries)}

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

	transaction :=
		connection.BeginDatabaseTransaction()

	defer transaction.
		Commit()

	var tableData [][]string

	//select_query = \
	selectQuery :=
		//"SELECT * FROM {}".format(table_name)
		`SELECT * FROM ` + table

	//dataframe = \
	//read_sql_query(
	//select_query,
	//database_connection)
	transaction.Select(
		&tableData,
		selectQuery)
	//		`SELECT * FROM $1`, table)

	tablesDataFrame := &dataframes.DataFrames{
		DataFrame: dataframe.
			LoadRecords(
				tableData),
	}

	//return \
	//dataframe
	return tablesDataFrame
}
