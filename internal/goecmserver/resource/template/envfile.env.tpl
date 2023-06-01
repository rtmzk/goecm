STORAGEPLATFORM={{ .ECMDeploySpec.ECMStorageSpec.StorageType }}
HTTP_PROTOCOL_TYPE={{ .ECMDeploySpec.Scheme }}
S3_Url={{ .ECMDeploySpec.ECMStorageSpec.StorageUrl }}
S3_SecretKey={{ .ECMDeploySpec.ECMStorageSpec.StorageSK }}
S3_AccessKey={{ .ECMDeploySpec.ECMStorageSpec.StorageAK }}
S3_Bucket={{ .ECMDeploySpec.ECMStorageSpec.StorageBucket }}
S3_SignatureVersion=4
Ceph_HealthUrl={{ .ECMDeploySpec.ECMStorageSpec.StorageHealthCheck }}
OSS_ISCNAME=false
ESURL=http://es:9200
CACHEMODULE=redis
ISUDC=0
PRODUCTION=edoc2
INBIZ=true
INBIZURL=http://inbiz
INBIZGRPCURL=http://inbiz:5003
ECMURL=http://edoc2
ECMGRPCURL=http://edoc2:5002
Region_TSAP={Region_TSAP}
{{ if .MiddlewareDeploySpec.Database.IsExternal }}
{{ if eq .MiddlewareDeploySpec.Database.DBType "mysql" }}
DatabaseType=1
{{ else if eq .MiddlewareDeploySpec.Database.DBType "mssql" }}
DatabaseType=4
{{ end }}
DatabaseServerName={{ .MiddlewareDeploySpec.Database.DBHosts }}
DatabaseServerPort={{ .MiddlewareDeploySpec.Database.DBPort }}
DatabaseUserName={{ .MiddlewareDeploySpec.Database.DBUser }}
DatabasePassword={{ .MiddlewareDeploySpec.Database.DBPass }}
{{ end }}
#IsSendNotNeedIndexToMQ=true
#InWiseIndexQueueName=index_inwise_queue|true
#OCR_MAXFILESIZE=10,80
#NUMS=25
#ThreadCheck=true
#IsUpgrade=5.16.0.2
#ENCRYPTENGINE=SmartSec
#SharedStorageId=0
#DecryptDebug=1
#DECRYPTGRPCURL=http://decryption:9090#

