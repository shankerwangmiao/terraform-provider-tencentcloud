package ratelimit

import (
	"os"
	"strconv"
)

//default cgi limit

const (
	DefaultLimit                int64 = 15
	PROVIDER_NEED_LIMIT               = "TENCENTCLOUD_NEED_LIMIT"
	PROVIDER_CVM_CREATE_LIMIT         = "TENCENTCLOUD_CVM_CREATE_LIMIT"
	PROVIDER_CVM_DESCRIBE_LIMIT       = "TENCENTCLOUD_CVM_DESCRIBE_LIMIT"
	PROVIDER_CBS_DESCRIBE_LIMIT       = "TENCENTCLOUD_CBS_DESCRIBE_LIMIT"
	PROVIDER_CVM_DELETE_LIMIT         = "TENCENTCLOUD_CVM_DELETE_LIMIT"
	PROVIDER_CBS_DELETE_LIMIT         = "TENCENTCLOUD_CBS_DELETE_LIMIT"
)

func getEnvDefault(key string, defVal int) int {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	int, err := strconv.Atoi(val)
	if err != nil {
		panic("TENCENTCLOUD_READ_RETRY_TIMEOUT or TENCENTCLOUD_WRITE_RETRY_TIMEOUT must be int.")
	}
	return int
}

func init() {
	// cvm
	// create
	var cvmCreateLimit = getEnvDefault(PROVIDER_CVM_CREATE_LIMIT, 10)
	limitConfig["resource_tc_instance.create"] = int64(cvmCreateLimit)
	// describe
	var cvmDescribeLimit = getEnvDefault(PROVIDER_CVM_DESCRIBE_LIMIT, 20)
	var cbsDescribeLimit = getEnvDefault(PROVIDER_CBS_DESCRIBE_LIMIT, 20)
	limitConfig["service_tencentcloud_cvm.DescribeInstances"] = int64(cvmDescribeLimit)
	limitConfig["service_tencentcloud_cbs.DescribeDisks"] = int64(cbsDescribeLimit)
	// delete
	var cvmDeleteLimit = getEnvDefault(PROVIDER_CVM_DELETE_LIMIT, 20)
	var cbsDeleteLimit = getEnvDefault(PROVIDER_CBS_DELETE_LIMIT, 20)
	limitConfig["service_tencentcloud_cvm.TerminateInstances"] = int64(cvmDeleteLimit)
	limitConfig["service_tencentcloud_cbs.TerminateDisks"] = int64(cbsDeleteLimit)

	//old  (filename . key)
	limitConfig["resource_tc_instance"] = 50
	limitConfig["resource_tc_instance.update"] = 10
	limitConfig["resource_tc_instance.delete"] = 10

	//new(filename . action)
	limitConfig["service_tencentcloud_mysql"] = 50
	limitConfig["service_tencentcloud_mysql.CreateDBInstanceHour"] = 20
	limitConfig["service_tencentcloud_mysql.OfflineIsolatedInstances"] = 20
	limitConfig["service_tencentcloud_mysql.CreateBackup"] = 5
	limitConfig["service_tencentcloud_mysql.ModifyInstanceParam"] = 20

	//new(filename)
	limitConfig["service_tencentcloud_cos"] = DefaultLimit
	limitConfig["service_tencentcloud_vpc"] = DefaultLimit
	limitConfig["service_tencentcloud_redis"] = DefaultLimit
	limitConfig["service_tencentcloud_mongodb"] = DefaultLimit
	limitConfig["service_tencentcloud_dcg"] = DefaultLimit
	limitConfig["service_tencentcloud_dc"] = 5
	limitConfig["service_tencentcloud_ccn"] = DefaultLimit
	limitConfig["service_tencentcloud_cbs"] = DefaultLimit
	limitConfig["service_tencentcloud_as"] = DefaultLimit
}
