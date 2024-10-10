package types

type MstrDatabaseType string

const (
	MstrDatabaseTypeReserved                       MstrDatabaseType = "reserved"
	MstrDatabaseTypeAccess                         MstrDatabaseType = "access"
	MstrDatabaseTypeAltibase                       MstrDatabaseType = "altibase"
	MstrDatabaseTypeAmazonAthena                   MstrDatabaseType = "amazon_athena"
	MstrDatabaseTypeAmazonAurora                   MstrDatabaseType = "amazon_aurora"
	MstrDatabaseTypeAmazonDocumentDb               MstrDatabaseType = "amazon_document_db"
	MstrDatabaseTypeAmazonDynamoDb                 MstrDatabaseType = "amazon_dynamo_db"
	MstrDatabaseTypeAmazonRedshift                 MstrDatabaseType = "amazon_redshift"
	MstrDatabaseTypeArcadiaPlatform                MstrDatabaseType = "arcadia_platform"
	MstrDatabaseTypeAster                          MstrDatabaseType = "aster"
	MstrDatabaseTypeAzureCosmos                    MstrDatabaseType = "azure_cosmos"
	MstrDatabaseTypeBigDataEngine                  MstrDatabaseType = "big_data_engine"
	MstrDatabaseTypeCassandra                      MstrDatabaseType = "cassandra"
	MstrDatabaseTypeCirro                          MstrDatabaseType = "cirro"
	MstrDatabaseTypeCloudElement                   MstrDatabaseType = "cloud_element"
	MstrDatabaseTypeCloudGateway                   MstrDatabaseType = "cloud_gateway"
	MstrDatabaseTypeCloudGatewayAwsS3              MstrDatabaseType = "cloud_gateway_aws_s3"
	MstrDatabaseTypeCloudGatewayAzureAdls2         MstrDatabaseType = "cloud_gateway_azure_adls_2"
	MstrDatabaseTypeCloudGatewayGoogleCloudStorage MstrDatabaseType = "cloud_gateway_google_cloud_storage"
	MstrDatabaseTypeComposite                      MstrDatabaseType = "composite"
	MstrDatabaseTypeConcur                         MstrDatabaseType = "concur"
	MstrDatabaseTypeConnectionCloud                MstrDatabaseType = "connection_cloud"
	MstrDatabaseTypeDatabricks                     MstrDatabaseType = "databricks"
	MstrDatabaseTypeDataDirectCloud                MstrDatabaseType = "data_direct_cloud"
	MstrDatabaseTypeDatallegro                     MstrDatabaseType = "datallegro"
	MstrDatabaseTypeDb2                            MstrDatabaseType = "db2"
	MstrDatabaseTypeDenodo                         MstrDatabaseType = "denodo"
	MstrDatabaseTypeDremio                         MstrDatabaseType = "dremio"
	MstrDatabaseTypeDrill                          MstrDatabaseType = "drill"
	MstrDatabaseTypeDropbox                        MstrDatabaseType = "dropbox"
	MstrDatabaseTypeDruid                          MstrDatabaseType = "druid"
	MstrDatabaseTypeElasticsearch                  MstrDatabaseType = "elasticsearch"
	MstrDatabaseTypeEloqua                         MstrDatabaseType = "eloqua"
	MstrDatabaseTypeEnterpriseDb                   MstrDatabaseType = "enterprise_db"
	MstrDatabaseTypeEssBase                        MstrDatabaseType = "ess_base"
	MstrDatabaseTypeExaSolution                    MstrDatabaseType = "exa_solution"
	MstrDatabaseTypeExcel                          MstrDatabaseType = "excel"
	MstrDatabaseTypeFacebook                       MstrDatabaseType = "facebook"
	MstrDatabaseTypeFinancialforce                 MstrDatabaseType = "financialforce"
	MstrDatabaseTypeGbase8a                        MstrDatabaseType = "gbase_8a"
	MstrDatabaseTypeGeneric                        MstrDatabaseType = "generic"
	MstrDatabaseTypeGenericDataConnector           MstrDatabaseType = "generic_data_connector"
	MstrDatabaseTypeGitHub                         MstrDatabaseType = "github"
	MstrDatabaseTypeGoogleAnalytics                MstrDatabaseType = "google_analytics"
	MstrDatabaseTypeGoogleBigQuery                 MstrDatabaseType = "google_big_query"
	MstrDatabaseTypeGoogleBigQueryFFSql            MstrDatabaseType = "google_big_query_ff_sql"
	MstrDatabaseTypeGoogleCloudSpanner             MstrDatabaseType = "google_cloud_spanner"
	MstrDatabaseTypeGoogleDrive                    MstrDatabaseType = "google_drive"
	MstrDatabaseTypeHive                           MstrDatabaseType = "hive"
	MstrDatabaseTypeHiveThrift                     MstrDatabaseType = "hive_thrift"
	MstrDatabaseTypeHubspot                        MstrDatabaseType = "hubspot"
	MstrDatabaseTypeIbmIps                         MstrDatabaseType = "ibm_ips"
	MstrDatabaseTypeImpala                         MstrDatabaseType = "impala"
	MstrDatabaseTypeInformatica                    MstrDatabaseType = "informatica"
	MstrDatabaseTypeInformix                       MstrDatabaseType = "informix"
	MstrDatabaseTypeJira                           MstrDatabaseType = "jira"
	MstrDatabaseTypeKafka                          MstrDatabaseType = "kafka"
	MstrDatabaseTypeKognitiowx2                    MstrDatabaseType = "kognitiowx2"
	MstrDatabaseTypeKyvos_mdx                      MstrDatabaseType = "kyvos_mdx"
	MstrDatabaseTypeMapd                           MstrDatabaseType = "mapd"
	MstrDatabaseTypeMarketo                        MstrDatabaseType = "marketo"
	MstrDatabaseTypeMarkLogic                      MstrDatabaseType = "mark_logic"
	MstrDatabaseTypeMemSql                         MstrDatabaseType = "mem_sql"
	MstrDatabaseTypeMetamatrix                     MstrDatabaseType = "metamatrix"
	MstrDatabaseTypeMicrosoftAS                    MstrDatabaseType = "microsoft_as"
	MstrDatabaseTypeMicrosoftDynamicsCRM           MstrDatabaseType = "microsoft_dynamics_crm"
	MstrDatabaseTypeMicrosoftDynamicsERP           MstrDatabaseType = "microsoft_dynamics_erp"
	MstrDatabaseTypeMicrosoftDynamics365           MstrDatabaseType = "microsoft_dynamics_365"
	MstrDatabaseTypeMicrosoftFabric                MstrDatabaseType = "microsoft_fabric"
	MstrDatabaseTypeMongoBI                        MstrDatabaseType = "mongo_bi"
	MstrDatabaseTypeMongoDB                        MstrDatabaseType = "mongo_db"
	MstrDatabaseTypeMySQL                          MstrDatabaseType = "my_sql"
	MstrDatabaseTypeNeo4j                          MstrDatabaseType = "neo4j"
	MstrDatabaseTypeNeoview                        MstrDatabaseType = "neoview"
	MstrDatabaseTypeNetezza                        MstrDatabaseType = "netezza"
	MstrDatabaseTypeOdata                          MstrDatabaseType = "odata"
	MstrDatabaseTypeOpenAccess                     MstrDatabaseType = "open_access"
	MstrDatabaseTypeOracle                         MstrDatabaseType = "oracle"
	MstrDatabaseTypeOracleCxSales                  MstrDatabaseType = "oracle_cx_sales"
	MstrDatabaseTypeOracleCxService                MstrDatabaseType = "oracle_cx_service"
	MstrDatabaseTypeOracleEloqua                   MstrDatabaseType = "oracle_eloqua"
	MstrDatabaseTypePalantirFoundry                MstrDatabaseType = "palantir_foundry"
	MstrDatabaseTypeParAccel                       MstrDatabaseType = "par_accel"
	MstrDatabaseTypeParStream                      MstrDatabaseType = "par_stream"
	MstrDatabaseTypePaypal                         MstrDatabaseType = "paypal"
	MstrDatabaseTypePhoenix                        MstrDatabaseType = "phoenix"
	MstrDatabaseTypePig                            MstrDatabaseType = "pig"
	MstrDatabaseTypePivotalHawq                    MstrDatabaseType = "pivotal_hawq"
	MstrDatabaseTypePostgreSQL                     MstrDatabaseType = "postgre_sql"
	MstrDatabaseTypePresto                         MstrDatabaseType = "presto"
	MstrDatabaseTypePython                         MstrDatabaseType = "python"
	MstrDatabaseTypeRedBrick                       MstrDatabaseType = "red_brick"
	MstrDatabaseTypeSalesforce                     MstrDatabaseType = "salesforce"
	MstrDatabaseTypeSand                           MstrDatabaseType = "sand"
	MstrDatabaseTypeSAP                            MstrDatabaseType = "sap"
	MstrDatabaseTypeSAPBW4Hana                     MstrDatabaseType = "sap_bw4_hana"
	MstrDatabaseTypeSAPBWOdata                     MstrDatabaseType = "sap_bw_odata"
	MstrDatabaseTypeSAPEccOdata                    MstrDatabaseType = "sap_ecc_odata"
	MstrDatabaseTypeSAPHana                        MstrDatabaseType = "sap_hana"
	MstrDatabaseTypeSAPHanaMDX                     MstrDatabaseType = "sap_hana_mdx"
	MstrDatabaseTypeSAPS4Hana                      MstrDatabaseType = "sap_s4_hana"
	MstrDatabaseTypeSearchEngine                   MstrDatabaseType = "search_engine"
	MstrDatabaseTypeServicemax                     MstrDatabaseType = "servicemax"
	MstrDatabaseTypeServiceNow                     MstrDatabaseType = "servicenow"
	MstrDatabaseTypeShopify                        MstrDatabaseType = "shopify"
	MstrDatabaseTypeSnowflake                      MstrDatabaseType = "snow_flake"
	MstrDatabaseTypeSparkConfig                    MstrDatabaseType = "spark_config"
	MstrDatabaseTypeSparkSQL                       MstrDatabaseType = "spark_sql"
	MstrDatabaseTypeSplunk                         MstrDatabaseType = "splunk"
	MstrDatabaseTypeSqlServer                      MstrDatabaseType = "sql_server"
	MstrDatabaseTypeSquare                         MstrDatabaseType = "square"
	MstrDatabaseTypeStarburst                      MstrDatabaseType = "starburst"
	MstrDatabaseTypeSugarCRM                       MstrDatabaseType = "sugar_crm"
	MstrDatabaseTypeSybase                         MstrDatabaseType = "sybase"
	MstrDatabaseTypeSybaseIQ                       MstrDatabaseType = "sybase_iq"
	MstrDatabaseTypeSybaseSQLAny                   MstrDatabaseType = "sybase_sql_any"
	MstrDatabaseTypeTandem                         MstrDatabaseType = "tandem"
	MstrDatabaseTypeTeamCity                       MstrDatabaseType = "teamcity"
	MstrDatabaseTypeTeradata                       MstrDatabaseType = "teradata"
	MstrDatabaseTypeTM1                            MstrDatabaseType = "tm1"
	MstrDatabaseTypeTrino                          MstrDatabaseType = "trino"
	MstrDatabaseTypeTwitter                        MstrDatabaseType = "twitter"
	MstrDatabaseTypeUnknown                        MstrDatabaseType = "unknown"
	MstrDatabaseTypeUrlAuth                        MstrDatabaseType = "url_auth"
	MstrDatabaseTypeVectorwise                     MstrDatabaseType = "vectorwise"
	MstrDatabaseTypeVertica                        MstrDatabaseType = "vertica"
	MstrDatabaseTypeXQuery                         MstrDatabaseType = "xquery"
	MstrDatabaseTypeYellowbrick                    MstrDatabaseType = "yellowbrick"
	MstrDatabaseTypeCloudGatewaySharepoint         MstrDatabaseType = "cloud_gateway_sharepoint"
)

type MstrDataSourceType string

const (
	MstrDataSourceTypeReserved         MstrDataSourceType = "reserved"
	MstrDataSourceTypeNormal           MstrDataSourceType = "normal"
	MstrDataSourceTypeDataImport       MstrDataSourceType = "data_import"
	MstrDataSourceTypeDatImportPrimary MstrDataSourceType = "data_import_primary"
	MstrDataSourceTypeDataScript       MstrDataSourceType = "script"
)

type OdbcVersion string

const (
	OdbcVersion2x OdbcVersion = "version2x"
	OdbcVersion3x OdbcVersion = "version3x"
)

type MstrRestDataSource struct {
	MstrObject
	AccessRightsFlags               int          `json:"acg"`
	Database                        DatabaseInfo `json:"database"`
	TablePrefix                     string       `json:"tablePrefix"`
	OdbcVerion                      OdbcVersion  `json:"odbcVersion"`
	IntermediateStoreDatabaseName   string       `json:"intermediateStoreDbName"`
	IntermediateStoreTablespaceName string       `json:"intermediateStoreTableSpaceName"`
	DBMS                            MstrObject   `json:"dbms"`
	Owner                           MstrObject   `json:"owner"`
	Hidden                          bool         `json:"hidden"`
	Script                          struct {
		ID string `json:"id"`
	} `json:"script"`
}

type DatabaseInfo struct {
	Type               MstrDatabaseType              `json:"type"`
	Version            string                        `json:"version"`
	Connection         MstrRestObjectEmbedded        `json:"connection"`
	EmbeddedConnection *EmbeddedDataSourceConnection `json:"embeddedConnection,omitempty"`
	PrimaryDataSource  *MstrObject                   `json:"primaryDatasource,omitempty"`
	DataMartDataSource *MstrObject                   `json:"dataMartDatasource,omitempty"`
}

type MstrRestObjectEmbedded struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Embedded bool   `json:"isEmbedded"`
}

type DataSourceLogin struct {
	MstrObject
	Hidden bool `json:"hidden"`
}

type ExecutionMode string

const (
	ExecutionModeReserved        ExecutionMode = "reserved"
	ExecutionModeAsyncConnection ExecutionMode = "async_connection"
	ExecutionModeAsyncStatement  ExecutionMode = "async_statement"
	ExecutionModeSynchronous     ExecutionMode = "synchronous"
)

type CharEncoding string

const (
	CharEncodingMultiByte CharEncoding = "multibyte"
	CharEncodingUTF8      CharEncoding = "utf8"
)

type DriverType string

const (
	DriverTypeReserved DriverType = "reserved"
	DriverTypeODBC     DriverType = "odbc"
	DriverTypeNative   DriverType = "native"
)

type DataSourceConnection struct {
	MstrObject
	Hidden                   bool              `json:"hidden"`
	ExecutionMode            ExecutionMode     `json:"executionMode"`
	MaxCancelAttemptTime     int               `json:"maxCancelAttemptTime"`
	MaxQueryExecTime         int               `json:"maxQueryExeTime"`
	MaxConnectionAttemptTime int               `json:"maxConnectionAttemptTime"`
	ConnectionLifetime       int               `json:"connectionLifetime"`
	ConnectionIdleTimeout    int               `json:"connectionIdleTimeout"`
	WindowsCharEncoding      CharEncoding      `json:"charEncodingWindows"`
	UnixCharEncoding         CharEncoding      `json:"charEncodingUnix"`
	TablePrefix              string            `json:"tablePrefix"`
	ConnectionString         string            `json:"connectionString"`
	ParameterisedQueries     bool              `json:"parameterizedQueries"`
	ExtendedFetch            bool              `json:"extendedFetch"`
	Database                 DatabaseInfo      `json:"database"`
	DriverType               DriverType        `json:"driverType"`
	OAuthParameter           string            `json:"oauthParameter"`
	WalletInfo               string            `json:"walletInfo"`
	IAM                      MstrObject        `json:"iam"`
	Resource                 string            `json:"resource"`
	Scope                    string            `json:"scope"`
	EnableSSO                bool              `json:"enableSso"`
	ExtraSensitiveFields     map[string]string `json:"extraSensitiveFields"`
}

type EmbeddedDataSourceConnection struct {
	DataSourceConnection
	Login         MstrRestObjectEmbedded `json:"login"`
	EmbeddedLogin *DataSourceLogin       `json:"embeddedLogin,omitempty"`
}

type DataSourceUpdateOperation string

const (
	DataSourceUpdateOperationAdd            DataSourceUpdateOperation = "add"
	DataSourceUpdateOperationReplace        DataSourceUpdateOperation = "replace"
	DataSourceUpdateOperationRemove         DataSourceUpdateOperation = "remove"
	DataSourceUpdateOperationIncr           DataSourceUpdateOperation = "incr"
	DataSourceUpdateOperationRemoveElement  DataSourceUpdateOperation = "removeElement"
	DataSourceUpdateOperationAddElement     DataSourceUpdateOperation = "addElement"
	DataSourceUpdateOperationRemoveElements DataSourceUpdateOperation = "removeElements"
	DataSourceUpdateOperationAddElements    DataSourceUpdateOperation = "addElements"
)

type DataSourceUpdate struct {
	Operation DataSourceUpdateOperation `json:"op"`
	Path      string                    `json:"path"`
	Value     interface{}               `json:"value"`
}
