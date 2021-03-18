package e

const (
	SUCCESS                 = 200
	UPDATE_PASSWORD_SUCCESS = 201
	NOT_EXIST_INENTIFIER    = 202
	ERROR                   = 500
	INVALID_PARAMS          = 400

	//成员错误
	ERROR_EXIST_NICK           = 10001
	ERROR_EXIST_USER           = 10002
	ERROR_NOT_EXIST_USER       = 10003
	ERROR_NOT_COMPARE          = 1004
	ERROR_NOT_COMPARE_PASSWORD = 10005
	ERROR_FAIL_ENCRYPTION      = 10006
	ERROR_NOT_EXIST_PRODUCT    = 10007
	ERROR_NOT_EXIST_ADDRESS    = 10008
	ERROR_EXIST_FAVORITE       = 10009

	//店家错误
	ERROR_BOSS_CHECK_TOKEN_FAIL       = 20001
	ERROR_BOSS_CHECK_TOKEN_TIMEOUT    = 20002
	ERROR_BOSS_TOKEN                  = 20003
	ERROR_BOSS                        = 20004
	ERROR_BOSS_INSUFFICIENT_AUTHORITY = 20005
	ERROR_BOSS_PRODUCT                = 20006

	//管理员错误
	ERROR_AUTH_CHECK_TOKEN_FAIL       = 30001 //token 错误
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT    = 30002 //token 过期
	ERROR_AUTH_TOKEN                  = 30003
	ERROR_AUTH                        = 30004
	ERROR_AUTH_INSUFFICIENT_AUTHORITY = 30005
	ERROR_READ_FILE                   = 30006
	ERROR_SEND_EMAIL                  = 30007
	ERROR_CALL_API                    = 30008
	ERROR_UNMARSHAL_JSON              = 30009
	ERROR_ADMIN_FIND_USER             = 30010

	//数据库错误
	ERROR_DATABASE = 40001

	//对象存储错误
	ERROR_OSS = 50001
)
