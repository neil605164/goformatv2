package errorcode

// newErrorCode 錯誤代碼格式
type newErrorCode struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

var errorCode = map[string]newErrorCode{
	/** 共同 **/
	"SUCCESS":                 {1, "SUCCESS"},                    // 呼叫API成功
	"PERMISSION_DENIED":       {403, "PERMISSION DENIED"},        // 權限不足
	"CREATE_DIR_ERROR":        {1001, "CREATE DIR ERROR"},        // 建立資料夾失敗
	"GET_UPLOAD_FILE_ERROR":   {1002, "GET UPLOAD FILE ERROR"},   // 取得上傳檔案失敗
	"CREATE_FILE_ERROR":       {1003, "CREATE FILE ERROR"},       // 建立檔案失敗
	"GET_UPLOAD_TYPE_ERROR":   {1004, "GET UPLOAD TYPE ERROR"},   // 取得上傳類型錯誤
	"JSON_MARSHAL_ERROR":      {1005, "JSON MARSHAL ERROR"},      // json encode 失敗
	"JSON_UNMARSHAL_ERROR":    {1006, "JSON UNMARSHAL ERROR"},    // json decode 失敗
	"CHANGE_PARAMS_TYPE_FAIL": {1007, "CHANGE PARAMS TYPE FAIL"}, // 資料轉型失敗
	"PARSE_TIME_ERROR":        {1008, "PARSE TIME ERROR"},        // 時間格式轉換錯誤
	"VALIDATE_PARAMS_FAIL":    {1009, "VALIDATE PARAMS FAIL"},    // 參數型態驗證失敗
	"IMAGES_TOO_LARGE":        {1010, "IMAGES TOO LARGE"},        // 檔案過大
	"BIND_PARAMS_FAIL":        {1011, "BIND PARAMS FAIL"},        // 組合參數失敗
	"CRYPTION_ERROR":          {1012, "CRYPTION ERROR"},          // 密碼加密錯誤
	"GET_TIME_ZONE_ERROR":     {1013, "GET TIME ZONE ERROR"},     // 取當前時區錯誤
	"LOG_ID_NOT_EXIST":        {1014, "LOG ID NOT EXIST"},        // Log 身份證
	"PARSE_BOOL_ERROR":        {1015, "PARSE BOOL ERROR"},        // 資料轉型失敗
	"SERVICE_MAINTAIN":        {1016, "SERVICE MAINTAIN"},        // 服務維護中

	/** DB 錯誤 **/
	"DB_CONNECT_ERROR":                    {2000, "DB CONNECT ERROR"},        // DB連線失敗
	"GET_USER_ACCOUNT_ERROR":              {2001, "GET USER ACCOUNT ERROR"},  // 查詢帳號是否存在失敗
	"UPDATE_LOGIN_TIME_ERROR":             {2002, "UPDATE LOGIN TIME ERROR"}, // 更新登入時間失敗
	"CREATE_SESSION_ERROR":                {2003, "CREATE SESSION ERROR"},    // 建立session資料失敗
	"DELETE_SESSION_ERROR":                {2003, "DELETE SESSION ERROR"},    // 登出時，清除DB session資料失敗
	"CREATE_ADMIN_USER_ERROR":             {2015, "CREATE MENU LIST ERROR"},  // 建立類別清單錯誤
	"QUERY_ADMIN_USER_ERROR":              {2016, "QUERY ADMIN USER ERROR"},  // 查詢使用者錯誤
	"UPDATE_ADMIN_DATA_ERROR":             {2005, "UPDATE ADMIN DATA ERROR"},
	"DELETE_ADMIN_ERROR":                  {2006, "DELETE ADMIN ERROR"},                  // 刪除管理者帳號失敗// 更新管理者帳號權限失敗
	"GET_TAGS_LIST_ERROR":                 {2007, "GET TAGS LIST ERROR"},                 // 取得標籤列表失敗
	"CREATE_OR_UPDATE_TAG_ERROR":          {2008, "CREATE OR UPDATE TAG ERROR"},          // 新增/更新標籤失敗
	"QUERY_TAG_ERROR":                     {2009, "QUERY TAG ERROR"},                     // 取得標籤失敗
	"DELETE_TAG_ERROR":                    {2010, "DELETE TAG ERROR"},                    // 刪除標籤錯誤
	"GET_SERIES_LIST_ERROR":               {2011, "GET SERIES LIST ERROR"},               // 取得系列列表失敗
	"CREATE_OR_UPDATE_SERIES_ERROR":       {2012, "CREATE OR UPDATE SERIES ERROR"},       // 新增/更新系列標籤失敗
	"QUERY_SERIES_ERROR":                  {2013, "QUERY SERIES ERROR"},                  // 取得系列標籤失敗
	"DELETE_SERIES_ERROR":                 {2014, "DELETE SERIES ERROR"},                 // 刪除系列標籤失敗
	"TAG_ID_NOT_EXISTS":                   {2015, "TAG ID NOT EXISTS"},                   // 標籤ID不存在
	"CREATE_OR_UPDATE_GAME_DOMAIN_ERROR":  {2016, "CREATE OR UPDATE GAME DOMAIN ERROR"},  // 新增/更新遊戲網域失敗
	"GET_GAME_DOMAIN_LIST_ERROR":          {2017, "GET GAME DOMAIN LIST ERROR"},          // 取得遊戲網域列表失敗
	"QUERY_GAME_DOMAIN_ERROR":             {2018, "QUERY GAME DOMAIN ERROR"},             // 查詢遊戲網域ID失敗
	"GAME_DOMAIN_ID_NOT_EXISTS":           {2019, "GAME DOMAIN ID NOT EXISTS"},           // 查詢遊戲網域ID不存在
	"DELETE_GAME_DOMAIN_ERROR":            {2020, "DELETE GAME DOMAIN ERROR"},            // 刪除遊戲網域失敗
	"GAME_ID_NOT_EXISTS":                  {2023, "GAME ID NOT EXISTS"},                  // game_id 不存在
	"GET_LANG_LIST_ERROR":                 {2024, "GET LANG LIST ERROR"},                 // 取得語系清單錯誤
	"UPDATE_SETTING_ERROR":                {2025, "UPDATE SETTING ERROR"},                // 更新 setting 設定值錯誤
	"SELECT_SETTING_ERROR":                {2026, "SELECT SETTING ERROR"},                // 搜尋 setting 清單錯誤
	"SELECT_MAINTAIN_STATUS_ERROR":        {2027, "SELECT MAINTAIN STATUS ERROR"},        // 搜尋維護開關狀態
	"GET_SIT_ERROR":                       {2028, "GET SIT ERROR"},                       // 取得站別錯誤
	"CORNER_TAG_ID_NOT_EXISTS":            {2029, "CORNER TAG ID NOT EXISTS"},            // 查詢角標ID不存在
	"GET_CORNER_TAGS_LIST_BY_SITE_ERROR":  {2030, "GET CORNER TAGS LIST BY SITE ERROR"},  // 查詢角標列表 By site 錯誤
	"CREATE_OR_UPDATE_CORNER_TAG_ERROR":   {2031, "CREATE OR UPDATE CORNER TAG ERROR"},   // 建立/編輯角標資料錯誤
	"CREATE_OR_UPDATE_THEME_ERROR":        {2032, "CREATE OR UPDATE THEME ERROR"},        // 建立/編輯主題資料錯誤
	"GET_THEME_BY_ID_ERROR":               {2033, "GET THEME BY ID ERROR"},               // 查詢主題資料 By ID 錯誤
	"UPDATE_THEME_STATUS_ERROR":           {2034, "UPDATE THEME STATUS ERROR"},           // 更新主題狀態錯誤
	"DELETE_THEME_ERROR":                  {2035, "DELETE THEME ERROR"},                  // 刪除主題錯誤
	"GET_THEME_LIST_ERROR":                {2036, "GET THEME LIST ERROR"},                // 取得主題列表錯誤
	"SELECT_ENABLE_THEME_ERROR":           {2037, "SELECT ENABLE THEME ERROR"},           // 查詢啟用主題錯誤
	"ENABLE_THEME_OVER_LIMIT_ERROR":       {2038, "ENABLE THEME OVER LIMIT ERROR"},       // 主題啟用總數超過上限
	"SELECT_SEARCH_RECORD_ERROR":          {2039, "SELECT SEARCH RECORD ERROR"},          // 尋找搜尋紀錄錯誤
	"CREATE_SEARCH_RECORD_ERROR":          {2040, "CREATE SEARCH RECORD ERROR"},          // 建立搜尋紀錄錯誤
	"GET_SEARCH_RECORD_LIST_ERROR":        {2041, "GET SEARCH RECORD LIST ERROR"},        // 取搜尋紀錄清單錯誤
	"UPDATE_NEW_CORNER_TAG_ERROR":         {2042, "UPDATE NEW CORNER TAG ERROR"},         // 更新最新角標錯誤
	"UPDATE_HOT_CORNER_TAG_ERROR":         {2043, "UPDATE HOT CORNER TAG ERROR"},         // 更新熱門角標錯誤
	"UPDATE_RECO_CORNER_TAG_ERROR":        {2044, "UPDATE RECO CORNER TAG ERROR"},        // 更新推薦角標錯誤
	"GET_CYPRESS_DB_DATA_ERROR":           {2046, "GET CYPRESS DB DATA ERROR"},           // 取 cypress DB sp_parentid_most_popular_game 錯誤
	"CREATE_OR_UPDATE_NEWSLETTER_ERROR":   {2048, "CREATE OR UPDATE NEWSLETTER ERROR"},   // 新增/編輯 電子報失敗
	"GET_SUBSCRIBE_LIST_ERROR":            {2049, "GET SUBSCRIBE LIST ERROR"},            // 取訂閱電子報失敗
	"GET_EMAILSEND_LIST_ERROR":            {2050, "GET EMAILSEND LIST ERROR"},            // 取待發送電子報清單失敗
	"CHANGE_SEND_STATUS_FAIL":             {2051, "CHANGE SEND STATUS FAIL"},             // 修改電子報發送狀態失敗
	"GET_EMAIL_LIST_ERROR":                {2052, "GET EMAIL LIST ERROR"},                // 取得電子報列表失敗
	"CHANGE_MAIL_STATUS_FAIL":             {2053, "CHANGE MAIL STATUS FAIL"},             // 修改電子報狀態失敗
	"CHECK_EMAIL_EXISTS_ERROR":            {2054, "CHECK EMAIL EXISTS ERROR"},            // 檢查電子報是否存在失敗
	"MAIL_ID_ERROR":                       {2055, "MAIL ID ERROR"},                       // 帶入的MailID錯誤
	"CREATE_SUBSCRIBE_ERROR":              {2056, "CREATE SUBSCRIBE ERROR"},              // 新增訂閱電子報失敗
	"UNSUBSCRIBE_ERROR":                   {2057, "UNSUBSCRIBE ERROR"},                   // 取消訂閱電子報失敗
	"GET_SEARCH_RECORD_TOTAL_COUNT_ERROR": {2058, "GET SEARCH RECORD TOTAL COUNT ERROR"}, // 取關鍵字搜尋清單總筆數錯誤
	"CREATE_OR_UPDATE_LINKS_ERROR":        {2059, "CREATE OR UPDATE LINKS ERROR"},        // 新增/編輯 連結資料錯誤
	"GET_LINKS_LIST_ERROR":                {2060, "GET LINKS LIST ERROR"},                // 取連結資料錯誤
	"SORT_LINKS_ERROR":                    {2061, "SORT LINKS ERROR"},                    // 排序連結順序錯誤
	"NO_DATA_BE_AFFECTED":                 {2062, "NO DATA BE AFFECTED"},                 // 沒有更動資料
	"MAIL_STATUS_ERROR":                   {2063, "MAIL STATUS ERROR"},                   // 電子報狀態非1
	"NO_RECEIVER":                         {2064, "NO RECEIVER"},                         // 預覽電子報無收件人
	"LANG_NOT_EXIST":                      {2065, "LANG NOT EXIST"},                      // 語系不存在
	"GET_AD_LIST_SQL_ERROR":               {2066, "GET AD LIST SQL ERROR"},               // 取廣告列表SQL錯誤
	"CREATE_NEWSLETTER_SEND_ERROR":        {2067, "CREATE NEWSLETTER SEND ERROR"},        // 建立發送事件錯誤

	/** Redis 錯誤 **/
	"REDIS_CONNECT_ERROR":              {3000, "REDIS CONNECT ERROR"},              // Redis連線失敗
	"REDIS_INSERT_ERROR":               {3001, "REDIS INSERT ERROR"},               // Redis寫入失敗
	"REDIS_DELETE_ERROR":               {3002, "REDIS DELETE ERROR"},               // Redis刪除失敗
	"REDIS_APPEND_ERROR":               {3003, "REDIS APPEND ERROR"},               // Redis增加值失敗
	"REDIS_SET_EXPIRE_ERROR":           {3004, "REDIS SET EXPIRE ERROR"},           // Redis設定過期時間失敗
	"REDIS_CHECK_EXIST_ERROR":          {3005, "REDIS CHECK EXIST ERROR"},          // 檢查Redis值是否存在時發生錯誤
	"REDIS_PING_ERROR":                 {3006, "REDIS PING ERROR"},                 // Redis Ping 錯誤
	"REDIS_GET_VALUE_ERROR":            {3007, "REDIS GET VALUE ERROR"},            // Redis 取值錯誤
	"REDIS_HASH_GET_VALUE_ERROR":       {3008, "REDIS HASH GET VALUE ERROR"},       // Redis hash get value 錯誤
	"REDIS_HASH_GET_ALL_VALUE_ERROR":   {3009, "REDIS HASH GET ALL VALUE ERROR"},   // Redis hash get all value 錯誤
	"REDIS_HASH_GET_SCAN_STRUCT_ERROR": {3010, "REDIS HASH GET SCAN STRUCT ERROR"}, // Redis hash get scan struct 錯誤
	"REDIS_HASH_MSET_VALUE_ERROR":      {3011, "REDIS HASH MSET VALUE ERROR"},      // Redis hash set mutli value struct 錯誤
	"REDIS_GET_TTL_ERROR":              {3012, "REDIS GET TTL ERROR"},              // Redis TTL get value 錯誤

	/** 呼叫 API 錯誤 **/
	"CURL_CREATE_FAIL":                              {4000, "CURL CREATE FAIL"},                              // CURL建立失敗
	"CURL_GET_FAIL":                                 {4001, "CURL GET FAIL"},                                 // CURL GET 失敗
	"CURL_POST_FAIL":                                {4002, "CURL POST FAIL"},                                // 取API失敗
	"API_CONNECT_ERROR":                             {4003, "API CONNECT ERROR"},                             // 對外連線失敗
	"API_STATUS_ERROR":                              {4004, "API STATUS ERROR"},                              // 對外連線回傳code異常
	"CREATE_CHERRY_GAME_INFO_ERROR":                 {4005, "CREATE CHERRY GAME INFO ERROR"},                 // 呼叫「遊戲中心」創立遊戲資訊失敗
	"UPDATE_CHERRY_GAME_INFO_ERROR":                 {4006, "UPDATE CHERRY GAME INFO ERROR"},                 // 呼叫「遊戲中心」更新遊戲資訊失敗
	"CURL_DELETE_IMAGE_API_ERROR":                   {4007, "CURL DELETE IMAGE API ERROR"},                   // 呼叫「圖片中心」刪除圖片API失敗
	"GET_CHERRY_GAME_INFO_ERROR":                    {4008, "GET CHERRY GAME INFO ERROR"},                    // 呼叫「遊戲中心」取遊戲資訊失敗
	"CURL_CYPRESS_GAME_LIST_API_ERROR":              {4009, "CURL CYPRESS GAME LIST API ERROR"},              // 呼叫「cypress game list api」失敗
	"SMTP_SEND_FAIL":                                {4010, "SMTP SEND FAIL"},                                // 使用SMTP發送Email失敗
	"CURL_CYPRESS_PLAYER_LOGIN_API_ERROR":           {4011, "CURL CYPRESS PLAYER LOGIN API ERROR"},           // 呼叫 cypresss player login api 失敗
	"CURL_CYPRESS_AUTH_API_ERROR":                   {4012, "CURL CYPRESS AUTH API ERROR"},                   // 呼叫 cypresss auth api 失敗
	"CURL_CYPRESS_USER_INFO_API_ERROR":              {4013, "CURL CYPRESS USER INFO API ERROR"},              // 呼叫 cypresss userinfo api 失敗
	"REDIS_CY_LOBBY_GAME_LIST_JSON_UNMARSHAL_ERROR": {4014, "REDIS CY LOBBY GAME LIST JSON UNMARSHAL ERROR"}, // Redis 的 Lobby 遊戲列表 JSON UNMARSHAL 錯誤
	"GET_CHERRY_GAME_GUIDE_ERROR":                   {4015, "GET CHERRY GAME GUIDE ERROR"},                   // 呼叫「遊戲中心」取遊戲攻略清單失敗
	"GET_CHERRY_GAME_GUIDE_GAMEID_ERROR":            {4016, "GET CHERRY GAME GUIDE GAMEID ERROR"},            // 呼叫「遊戲中心」取遊戲攻略 game_id 清單失敗
	"CURL_CYPRESS_LOBBY_BALANCE_API_ERROR":          {4017, "CURL CYPRESS LOBBY BALANCE API ERROR"},          // 取用戶餘額,幣別錯誤
	"CURL_CYPRESS_GAME_LINK_API_ERROR":              {4018, "CURL CYPRESS GAME LINK API ERROR"},              // 呼叫 cypress game link 錯誤

	/** AUTH **/
	"AUTH_VALIDATE_FAIL":      {5000, "AUTH VALIDATE FAIL"},      // 登入驗證失敗
	"USER_OR_PASSWORD_ERROR":  {5001, "USER OR PASSWORD ERROR"},  // 用戶不存在
	"USER_ACCOUNT_DISABLE":    {5002, "USER ACCOUNT DISABLE"},    // 用戶帳號遭凍結
	"ACCOUNT_RULE_ERROR":      {5003, "ACCOUNT RULE ERROR"},      // 帳號內容不符合正則規則
	"PASSWORD_RULE_ERROR":     {5004, "PASSWORD RULE ERROR"},     // 密碼內容不符合正則規則
	"USER_IS_EXIST":           {5005, "USER IS EXIST"},           // 用戶已經存在，不可重複註冊
	"TOKEN_PERMISSION_DENIED": {5006, "TOKEN PERMISSION DENIED"}, // Token錯誤

	/** GAME 遊戲資訊 **/
	"NO_PERMISSION_EDIT_LOBBY_SITE":                     {6000, "NO PERMISSION EDIT LOBBY SITE"},                     // 沒有權限編輯 Lobby 站台
	"NO_PERMISSION_EDIT_DEMO_SITE":                      {6001, "NO PERMISSION EDIT DEMO SITE"},                      // 沒有權限編輯 Demo 站台
	"GAME_ID_ALREADY_EXISTS":                            {6002, "GAME ID ALREADY EXISTS"},                            // 該款遊戲已經存在
	"GAME_CREATE_BIND_PARAMS_FAIL":                      {6003, "GAME CREATE BIND PARAMS FAIL"},                      // 建立遊戲 綁定參數錯誤
	"GAME_CREATE_VALIDATE_PARAMS_FAIL":                  {6004, "GAME CREATE VALIDATE PARAMS FAIL"},                  // 建立遊戲 驗證參數錯誤
	"GAME_CREATE_GAME_ID_ALREADY_EXISTS":                {6005, "GAME CREATE GAME ID ALREADY EXISTS"},                // 建立遊戲 game_id 已存在
	"GAME_CREATE_CHERRY_DATA_JSON_UNMARSHAL_ERROR":      {6006, "GAME CREATE CHERRY DATA JSON UNMARSHAL ERROR"},      // 建立遊戲取 Cherry data JSON UNMARSHAL 錯誤
	"GAME_CREATE_CHERRY_CREATE_GAME_API_ERROR":          {6007, "GAME CREATE CHERRY CREATE GAME API ERROR"},          // 建立遊戲 call Cherry api 回傳錯誤
	"CHECK_GAMEID_EXISTS_SQL_ERROR":                     {6008, "CHECK GAMEID EXISTS SQL ERROR"},                     // 檢查 GameID 是否存在 SQL 錯誤
	"GAME_EDIT_STATUS_JSON_MARSHAL_ERROR":               {6009, "GAME EDIT STATUS JSON MARSHAL ERROR"},               // 建立/更新遊戲 狀態 JSON MARSHAL錯誤
	"GAME_EDIT_DOMAIN_JSON_MARSHAL_ERROR":               {6010, "GAME EDIT DOMAIN JSON MARSHAL ERROR"},               // 建立/更新遊戲 網域 JSON MARSHAL錯誤
	"GAME_EDIT_PARAM_JSON_MARSHAL_ERROR":                {6011, "GAME EDIT PARAM JSON MARSHAL ERROR"},                // 建立/更新遊戲 特色參數 JSON MARSHAL錯誤
	"GAME_EDIT_GUIDE_SWITCH_JSON_MARSHAL_ERROR":         {6012, "GAME EDIT GUIDE SWITCH JSON MARSHAL ERROR"},         // 建立/更新遊戲 攻略開關 JSON MARSHAL錯誤
	"GAME_EDIT_FEATURE_SWITCH_JSON_MARSHAL_ERROR":       {6013, "GAME EDIT FEATURE SWITCH JSON MARSHAL ERROR"},       // 建立/更新遊戲 特色開關 JSON MARSHAL錯誤
	"GAME_EDIT_VALIDATE_PARENT_ID_STATUS_ERROR":         {6014, "GAME EDIT VALIDATE PARENT ID STATUS ERROR"},         // 建立/更新遊戲 驗證系統商 Parent ID 狀態錯誤
	"GAME_EDIT_VENDOR_JSON_MARSHAL_ERROR":               {6015, "GAME EDIT VENDOR JSON MARSHAL ERROR"},               // 建立/更新遊戲 系統商 JSON MARSHAL錯誤
	"CREATE_OR_UPDATE_GAME_SQL_ERROR":                   {6016, "CREATE OR UPDATE GAME SQL ERROR"},                   // 建立/更新遊戲 SQL 錯誤
	"DELETE_GAME_SQL_ERROR":                             {6017, "DELETE GAME SQL ERROR"},                             // 刪除遊戲 SQL 錯誤
	"GAME_UPDATE_BIND_PARAMS_FAIL":                      {6018, "GAME UPDATE BIND PARAMS FAIL"},                      // 更新遊戲 綁定參數錯誤
	"GAME_UPDATE_VALIDATE_PARAMS_FAIL":                  {6019, "GAME UPDATE VALIDATE PARAMS FAIL"},                  // 更新遊戲 驗證參數錯誤
	"GAME_UPDATE_ID_VALIDATE_FAIL":                      {6020, "GAME UPDATE ID VALIDATE FAIL"},                      // 更新遊戲 驗證 ID 錯誤
	"GET_GAME_BY_ID_NOT_EXISTS":                         {6021, "GET GAME BY ID NOT EXISTS"},                         // 遊戲ID 不存在
	"GET_GAME_BY_ID_SQL_ERROR":                          {6022, "GET GAME BY ID SQL ERROR"},                          // 取得遊戲 By ID SQL 錯誤
	"GAME_UPDATE_CHERRY_CONTENT_JSON_UNMARSHAL_ERROR":   {6023, "GAME UPDATE CHERRY CONTENT JSON UNMARSHAL ERROR"},   // 更新遊戲 Cherry 更新遊戲 API 回傳內容 JSON UNMARSHAL 錯誤
	"GAME_UPDATE_CHERRY_UPDATE_GAME_API_ERROR":          {6024, "GAME UPDATE CHERRY UPDATE GAME API ERROR"},          // 更新遊戲 Cherry call Cherry api 回傳錯誤
	"GAME_STATUS_JSON_UNMARSHAL_ERROR":                  {6025, "GAME STATUS JSON UNMARSHAL ERROR"},                  // 遊戲 狀態 JSON UNMARSHAL 錯誤
	"GAME_DOMAIN_JSON_UNMARSHAL_ERROR":                  {6026, "GAME DOMAIN JSON UNMARSHAL ERROR"},                  // 遊戲 網域 JSON UNMARSHAL 錯誤
	"GAME_PARAM_JSON_UNMARSHAL_ERROR":                   {6027, "GAME PARAM JSON UNMARSHAL ERROR"},                   // 遊戲 特色參數 JSON UNMARSHAL 錯誤
	"GAME_FEATURE_SWITCH_JSON_UNMARSHAL_ERROR":          {6028, "GAME FEATURE SWITCH JSON UNMARSHAL ERROR"},          // 遊戲 特色開關 JSON UNMARSHAL 錯誤
	"GAME_GUIDE_SWITCH_JSON_UNMARSHAL_ERROR":            {6029, "GAME GUIDE SWITCH JSON UNMARSHAL ERROR"},            // 遊戲 攻略開關 JSON UNMARSHAL 錯誤
	"GAME_VENDOR_JSON_UNMARSHAL_ERROR":                  {6030, "GAME VENDOR JSON UNMARSHAL ERROR"},                  // 遊戲 系統商 JSON UNMARSHAL 錯誤
	"GET_GAME_LIST_ORDER_FOLLOW_SQL_ERROR":              {6031, "GET GAME LIST ORDER FOLLOW SQL ERROR"},              // 查詢遊戲清單錯誤 ORDER BY follow SQL 錯誤
	"GET_GAME_TYPE_LIST_SQL_ERROR":                      {6032, "GET GAME TYPE LIST SQL ERROR"},                      // 取遊戲類型清單 SQL 錯誤
	"GAME_TAG_LIST_NAME_JSON_UNMARSHAL_ERROR":           {6033, "GAME TAG LIST NAME JSON UNMARSHAL ERROR"},           // 取遊戲標籤清單 語系名稱 JSON UNMARSHAL 錯誤
	"GAME_TAG_JSON_UNMARSHAL_ERROR":                     {6034, "GAME TAG JSON UNMARSHAL ERROR"},                     // 遊戲 標籤 JSON UNMARSHAL 錯誤
	"GET_GAME_TAG_LIST_SQL_ERROR":                       {6035, "GET GAME TAG LIST SQL ERROR"},                       // 取遊戲標籤清單 SQL 錯誤
	"GET_GAME_TAG_LIST_BY_SITE_SQL_ERROR":               {6036, "GET GAME TAG LIST BY SITE SQL ERROR"},               // 取遊戲標籤清單 By Site SQL 錯誤
	"FRONTEND_GAME_TAG_LIST_NAME_JSON_UNMARSHAL_ERROR":  {6037, "FRONTEND GAME TAG LIST NAME JSON UNMARSHAL ERROR"},  // 前台 取遊戲標籤清單 遊戲名稱 JSON UNMARSHAL 錯誤
	"GET_LANG_BY_LANG_SQL_ERROR":                        {6040, "GET LANG BY LANG SQL ERROR"},                        // 取得語系 By lang SQL 錯誤
	"LOBBY_GAME_LIST_REDIS_GAME_JSON_UNMARSHAL_ERROR":   {6041, "LOBBY GAME LIST REDIS GAME JSON UNMARSHAL ERROR"},   // 前台 Lobby Redis 取得 Game JSON UNMARSHAL 錯誤
	"LOBBY_GAME_LIST_LANG_NOT_EXIST":                    {6042, "LOBBY GAME LIST LANG NOT EXIST"},                    // 前台 Lobby 該語系遊戲列表不存在
	"LOBBY_GAME_LIST_LANG_GAME_ID_JSON_UNMARSHAL_ERROR": {6043, "LOBBY GAME LIST LANG GAME ID JSON UNMARSHAL ERROR"}, // 前台 Lobby 該語系遊戲ID黑名單 JSON UNMARSHAL 錯誤
	"GET_GAME_LIST_SQL_ERROR":                           {6044, "GET GAME LIST SQL ERROR"},                           // 取遊戲列表 SQL 錯誤
	"REDIS_GAME_TYPE_LIST_JSON_UNMARSHAL_ERROR":         {6045, "REDIS GAME TYPE LIST JSON UNMARSHAL ERROR"},         // 遊戲類型列表 From Redis JSON UNMARSHAL 錯誤
	"GAME_TYPE_LIST_JSON_UNMARSHAL_ERROR":               {6046, "GAME TYPE LIST JSON UNMARSHAL ERROR"},               // 遊戲類型列表 JSON UNMARSHAL 錯誤
	"GET_GAME_DOMAIN_LIST_SQL_ERROR":                    {6047, "GET GAME DOMAIN LIST SQL ERROR"},                    // 遊戲網域列表 SQL 錯誤                                               //
	"REDIS_LOBBY_GAME_LIST_INSERT_ERROR":                {6049, "REDIS LOBBY GAME LIST INSERT ERROR"},                // Redis Lobby 遊戲列表 insert 錯誤
	"GAME_LIST_STATUS_JSON_UNMARSHAL_ERROR":             {6051, "GAME LIST STATUS JSON UNMARSHAL ERROR"},             // 遊戲狀態列表 JSON UNMARSHAL 錯誤
	"GAME_ID_CREATE_VALIDATE_PARAMS_FAILED":             {6052, "GAME ID CREATE VALIDATE PARAMS FAILED"},             // 新增遊戲 ID驗證參數錯誤
	"CHECK_GAME_TYPE2_EXISTS_SQL_ERROR":                 {6051, "CHECK GAME TYPE2 EXISTS SQL ERROR"},                 // 檢查 GameType2 是否存在 SQL 錯誤
	"GUIDE_LANG_IS_EMPTY":                               {6052, "GUIDE LANG IS EMPTY"},                               // 未提供攻略語系
	"GET_GAME_DATA_BY_GID_ERROR":                        {6053, "GET GAME DATA BY GID ERROR"},                        // 透過 gameID 取資料錯誤
	"GAME_GUIDE_NOT_OPEN":                               {6054, "GAME GUIDE NOT OPEN"},                               // 攻略沒有開啟
	"PRESENT_TIME_NOT_FOUND":                            {6055, "PRESENT TIME NOT FOUND"},                            // 未帶入遊戲創立時間且CY無此遊戲資料

	/** 紀錄搜尋紀錄錯誤 **/
	"LOAD_TRANSLATION_FILE_FAIL":      {7000, "LOAD TRANSLATION FILE FAIL"},      // 載入翻譯字典檔錯誤
	"TRADITIONAL_TO_SIMPLIFIED_ERROR": {7001, "TRADITIONAL TO SIMPLIFIED ERROR"}, // 繁體轉簡體錯誤

	/*其他 錯誤*/
	"DOMAIN_NAME_RULE_ERROR":              {8001, "DOMAIN NAME RULE ERROR"},              // 網域名稱錯誤
	"MARSHAL_USER_INFO_ERROR":             {8002, "MARSHAL USER INFO ERROR"},             // middleware 處理 userinfo 資料錯誤
	"CAN_NOT_ENTER_PREVIOUS_DATE":         {8003, "CAN NOT ENTER PREVIOUS DATE"},         // 禁止輸入當天以前的日期
	"END_TIME_CANT_NOT_BEFORE_START_TIME": {8004, "END TIME CANT NOT BEFORE START TIME"}, // 結束時間不得小於開始時間
	"UPLOAD_TO_CENTER_ERROR":              {8005, "UPLOAD TO CENTER ERROR"},              // 上傳道圖片中心錯誤
	"CREATE_FOLDER_ERROR":                 {8006, "CREATE FOLDER ERROR"},                 // 建立資料夾錯誤

	/** Project 錯誤 **/
	"SESSION_NOT_EXIST": {9001, "SESSION NOT EXIST"},

	/** 角標錯誤 **/
	"CORNER_NAME_IS_NOT_RECO": {10001, "CORNER NAME IS NOT RECO"}, // 角標類型非推薦
	"CORNER_SITE_NOTE_EXIST":  {10002, "CORNER SITE NOTE EXIST"},  // 站台不存在

	/** 主題錯誤 **/
	"NOT_DELETE_ENABLE_THEME":               {11001, "NOT DELETE ENABLE THEME"},               // 啟用中主題不可刪除
	"FRONTEND_WRITE_THEME_REDIS_DATA_ERROR": {11002, "FRONTEND WRITE THEME REDIS DATA ERROR"}, // 寫入redis 主題資料錯誤
	"THEME_READ_JSON_MARSHAL_ERROR":         {11003, "THEME READ JSON MARSHAL ERROR"},         // 讀取主題JSON MARSHAL錯誤

	/** 連結錯誤 **/
	"CAN_NOT_DELETE_ENABLE_LINKS":  {12001, "CAN NOT DELETE ENABLE LINKS"},    // 啟用中連結不可刪除
	"DELETE_LINK_ERROR":            {12002, "DELETE LINK ERROR"},              // 刪除連結錯誤
	"SELECT_LINK_BY_TYPE_ID_ERROR": {12003, "SELECT LINK BY TYPE ID ERROR"},   // 刪除連結錯誤
	"LINK_IMG_JSON_MARSHAL_ERROR":  {12004, "LINK urlIMG JSON MARSHAL ERROR"}, // link image ummarshal 錯誤
	"LINK_NAME_JSON_MARSHAL_ERROR": {12005, "LINK NAME JSON MARSHAL ERROR"},   // link name ummarshal 錯誤
	"LINK_URL_JSON_MARSHAL_ERROR":  {12006, "LINK IMG JSON MARSHAL ERROR"},    // link url ummarshal 錯誤

	/** lobby 測試登入功能 **/
	"GET_LOGIN_DATA_ERROR":  {13001, "GET LOGIN DATA ERROR"},  // 取登入資訊錯誤
	"CY_PLAYER_TOKEN_EMPTY": {13002, "CY PLAYER TOKEN EMPTY"}, // cypress 回傳的 token 為空

	/** lobby 用戶 token 驗證 **/
	"LOBBY_TOKEN_NOT_EXIST":         {14001, "LOBBY TOKEN NOT EXIST"},         // 遊戲大廳 token 不存在
	"SELECT_USER_INFO_ERROR":        {14002, "SELECT USER INFO ERROR"},        // 取會員資料失敗
	"BALANCE_CHANGE_TYPE_ERROR":     {14003, "BALANCE CHANGE TYPE ERROR"},     // 額度轉換型態錯誤
	"GAME_IS_ALREADY_EXIST_IN_LIST": {14004, "GAME IS ALREADY EXIST IN LIST"}, // 遊戲已存在最愛清單
	"GAME_NOT_EXIST_IN_LIST":        {14005, "GAME NOT EXIST IN LIST"},        // 遊戲不存在最愛清單
	"UPDATE_USER_DETAIL_ERROR":      {14006, "UPDATE USER DETAIL ERROR"},      // 更新用戶資料錯誤

	/** 效能監控錯誤 **/
	"NOT_ALLOW_TO_USE_PPROF": {15001, "NOT ALLOW TO USE PPROF"}, // 權限不足以開啟 pprof 圖

	/**  公告 **/
	"CREATE_ANNOUNCEMENT_BIND_PARAMS_FAIL":                  {16001, "CREATE ANNOUNCEMENT BIND PARAMS FAIL"},                  // 建立公告綁定參數錯誤
	"CREATE_ANNOUNCEMENT_VALIDATE_PARAMS_FAIL":              {16002, "CREATE ANNOUNCEMENT VALIDATE PARAMS FAIL"},              // 建立公告驗證參數錯誤
	"CREATE_ANNOUNCEMENT_TITLE_JSON_MARSHAL_ERROR":          {16003, "CREATE ANNOUNCEMENT TITLE JSON MARSHAL ERROR"},          // 建立公告標題JSON MARSHAL錯誤
	"CREATE_ANNOUNCEMENT_CONTENT_JSON_MARSHAL_ERROR":        {16004, "CREATE ANNOUNCEMENT CONTENT JSON MARSHAL ERROR"},        // 建立公告內容JSON MARSHAL錯誤
	"CREATE_ANNOUNCEMENT_VENDOR_JSON_MARSHAL_ERROR":         {16005, "CREATE ANNOUNCEMENT VENDOR JSON MARSHAL ERROR"},         // 建立公告系統商JSON MARSHAL錯誤
	"CREATE_ANNOUNCEMENT_SQL_ERROR":                         {16006, "CREATE ANNOUNCEMENT SQL ERROR"},                         // 建立公告SQL錯誤
	"UPDATE_ANNOUNCEMENT_BIND_PARAMS_FAIL":                  {16007, "UPDATE ANNOUNCEMENT BIND PARAMS FAIL"},                  // 編輯公告綁定參數錯誤
	"UPDATE_ANNOUNCEMENT_VALIDATE_PARAMS_FAIL":              {16008, "UPDATE ANNOUNCEMENT VALIDATE PARAMS FAIL"},              // 編輯公告驗證參數錯誤
	"UPDATE_ANNOUNCEMENT_TITLE_JSON_MARSHAL_ERROR":          {16009, "UPDATE ANNOUNCEMENT TITLE JSON MARSHAL ERROR"},          // 編輯公告標題JSON MARSHAL錯誤
	"UPDATE_ANNOUNCEMENT_CONTENT_JSON_MARSHAL_ERROR":        {16010, "UPDATE ANNOUNCEMENT CONTENT JSON MARSHAL ERROR"},        // 編輯公告內容JSON MARSHAL錯誤
	"UPDATE_ANNOUNCEMENT_VENDOR_JSON_MARSHAL_ERROR":         {16011, "UPDATE ANNOUNCEMENT VENDOR JSON MARSHAL ERROR"},         // 編輯公告系統商JSON MARSHAL錯誤
	"UPDATE_ANNOUNCEMENT_SQL_ERROR":                         {16012, "UPDATE ANNOUNCEMENT SQL ERROR"},                         // 編輯公告SQL錯誤
	"DELETE_ANNOUNCEMENT_BIND_PARAMS_FAIL":                  {16013, "DELETE ANNOUNCEMENT BIND PARAMS FAIL"},                  // 刪除公告綁定參數錯誤
	"DELETE_ANNOUNCEMENT_VALIDATE_PARAMS_FAIL":              {16014, "DELETE ANNOUNCEMENT VALIDATE PARAMS FAIL"},              // 刪除公告驗證參數錯誤
	"DELETE_ANNOUNCEMENT_TO_UPDATE_EDITEDBY_SQL_ERROR":      {16015, "DELETE ANNOUNCEMENT TO UPDATE EDITEDBY SQL ERROR"},      // 刪除公告修改 EDITEDBY SQL錯誤
	"DELETE_ANNOUNCEMENT_SQL_ERROR":                         {16016, "DELETE ANNOUNCEMENT SQL ERROR"},                         // 刪除公告SQL錯誤
	"LIST_ANNOUNCEMENT_BIND_PARAMS_FAIL":                    {16017, "LIST ANNOUNCEMENT BIND PARAMS FAIL"},                    // 公告列表綁定參數錯誤
	"LIST_ANNOUNCEMENT_VALIDATE_PARAMS_FAIL":                {16018, "LIST ANNOUNCEMENT VALIDATE PARAMS FAIL"},                // 公告列表驗證參數錯誤
	"LIST_ANNOUNCEMENT_TITLE_JSON_UNMARSHAL_ERROR":          {16019, "LIST ANNOUNCEMENT TITLE JSON UNMARSHAL ERROR"},          // 公告列表標題JSON UNMARSHAL 錯誤
	"LIST_ANNOUNCEMENT_CONTENT_JSON_UNMARSHAL_ERROR":        {16020, "LIST ANNOUNCEMENT CONTENT JSON UNMARSHAL ERROR"},        // 公告列表內容JSON UNMARSHAL 錯誤
	"LIST_ANNOUNCEMENT_VENDOR_JSON_UNMARSHAL_ERROR":         {16021, "LIST ANNOUNCEMENT VENDOR JSON UNMARSHAL ERROR"},         // 公告列表系統商JSON UNMARSHAL 錯誤
	"LIST_ANNOUNCEMENT_SQL_ERROR":                           {16022, "LIST ANNOUNCEMENT SQL ERROR"},                           // 公告列表SQL錯誤
	"COUNT_ANNOUNCEMENT_SQL_ERROR":                          {16022, "COUNT ANNOUNCEMENT SQL ERROR"},                          // 公告總數SQL錯誤
	"LOBBY_LIST_ANNOUNCEMENT_USERINFO_JSON_UNMARSHAL_ERROR": {16023, "LOBBY LIST ANNOUNCEMENT USERINFO JSON UNMARSHAL ERROR"}, // Lobby 前台公告列表 user info JSON UNMARSHAL 錯誤
	"LOBBY_LIST_ANNOUNCEMENT_TITLE_JSON_UNMARSHAL_ERROR":    {16024, "LOBBY LIST ANNOUNCEMENT TITLE JSON UNMARSHAL ERROR"},    // Lobby 前台公告列表標題JSON UNMARSHAL 錯誤
	"LOBBY_LIST_ANNOUNCEMENT_CONTENT_JSON_UNMARSHAL_ERROR":  {16025, "LOBBY LIST ANNOUNCEMENT CONTENT JSON UNMARSHAL ERROR"},  // Lobby 前台公告列表內容JSON UNMARSHAL 錯誤
	"LOBBY_LIST_ANNOUNCEMENT_VENDOR_JSON_UNMARSHAL_ERROR":   {16026, "LOBBY LIST ANNOUNCEMENT VENDOR JSON UNMARSHAL ERROR"},   // Lobby 前台公告列表系統商JSON UNMARSHAL 錯誤
	"FRONT_LIST_ANNOUNCEMENT_SQL_ERROR":                     {16027, "FRONT LIST ANNOUNCEMENT SQL ERROR"},                     // 前台公告列表 SQL 錯誤

	/** 跑馬燈錯誤 **/
	"NEWS_TICKER_ID_NOT_EXISTS":        {17001, "NEWS TICKER ID NOT EXISTS"},        // 跑馬燈ID不存在
	"GET_NEWS_TICKER_LIST_ERROR":       {17002, "GET NEWS TICKER LIST ERROR"},       // 取跑馬燈清單錯誤
	"DELETE_NEWS_TICKER_ERROR":         {17003, "DELETE NEWS TICKER ERROR"},         // 刪除跑馬燈錯誤
	"UPDATE_NEWS_STICKER_EDITED_ERROR": {17004, "UPDATE NEWS STICKER EDITED ERROR"}, // 更新跑f馬燈編輯者錯誤
	"SWITCH_NEWS_TICKER_STATUS_ERROR":  {17005, "SWITCH NEWS TICKER STATUS ERROR"},  // 更改跑馬燈狀態錯誤

	/** cy team gateway錯誤 **/
	"VALIDATE_PARENT_ID_STATUS_ERROR":      {18001, "VALIDATE PARENT ID STATUS ERROR"},      // 取得parentID狀態錯誤
	"CURL_CYPRESS_TEAM_GATEWAY_API_ERROR":  {18002, "CURL CYPRESS TEAM GATEWAY API ERROR"},  // 取cypress team gateway api錯誤
	"CURL_CYPRESS_BAD_PARAMETERS_ERROR":    {18003, "CURL CYPRESS BAD PARAMETERS ERROR"},    // 取cy gateway api參數錯誤
	"CURL_CYPRESS_GET_PARENT_ID_API_ERROR": {18004, "CURL CYPRESS GET PARENT ID API ERROR"}, // 取cy gateway取parent ID 錯誤
	"CY_TEAM_GATEWAY_TOKEN_EMPTY":          {18005, "CY TEAM GATEWAY TOKEN EMPTY"},          // 取cy api token錯誤

	/** teamplus package error **/
	"SEND_MESSAGE_TO_TEAM_PLUS_ERROR": {19001, "SEND MESSAGE TO TEAM PLUS ERROR"}, // teamplus通知錯誤

	/** 娛樂城錯誤 **/
	"GET_RECREATION_LIST_ERROR":                 {20001, "GET RECREATION LIST ERROR"},                 // 取娛樂城清單錯誤
	"UNMARSHAL_RECREATION_NAME_ERROR":           {20002, "UNMARSHAL RECREATION NAME ERROR"},           // 處理娛樂城名稱錯誤
	"UNMARSHAL_RECREATION_FACUS_ERROR":          {20003, "UNMARSHAL RECREATION FACUS ERROR"},          // 處理娛樂城關注錯誤
	"UNMARSHAL_RECREATION_STATUS_ERROR":         {20004, "UNMARSHAL RECREATION STATUS ERROR"},         // 處理娛樂城狀態錯誤
	"UNMARSHAL_RECREATION_IMAGE_ERROR":          {20005, "UNMARSHAL RECREATION IMAGE ERROR"},          // 處理娛樂城圖片錯誤
	"CREATE_OR_UPDATE_RECREATION_ERROR":         {20006, "CREATE OR UPDATE RECREATION ERROR"},         // 新增修改娛樂城資料錯誤
	"BIND_CREATE_RECREATION_PARAMS_FAIL":        {20007, "BIND CREATE RECREATION PARAMS FAIL"},        // 取建立娛樂城參數錯誤
	"VALIDATE_CREATE_RECREATION_PARAMS_FAIL":    {20008, "VALIDATE CREATE RECREATION PARAMS FAIL"},    // 驗證建立娛樂城參數錯誤
	"MARSHAL_RECREATION_NAME_TO_BYTE_ERROR":     {20009, "MARSHAL RECREATION NAME TO BYTE ERROR"},     // 將娛樂城名稱轉型態為 byte 錯誤
	"MARSHAL_RECREATION_FOCUS_TO_BYTE_ERROR":    {20010, "MARSHAL RECREATION FOCUS TO BYTE ERROR"},    // 將娛樂城關注狀態轉型態為 byte 錯誤
	"MARSHAL_RECREATION_STATUS_TO_BYTE_ERROR":   {20011, "MARSHAL RECREATION STATUS TO BYTE ERROR"},   // 將娛樂城開關狀態轉型態為 byte 錯誤
	"MARSHAL_RECREATION_IMAGE_TO_BYTE_ERROR":    {20012, "MARSHAL RECREATION IMAGE TO BYTE ERROR"},    // 將娛樂城圖片名稱轉型態為 byte 錯誤
	"BIND_DELETE_RECREATION_PARAMS_FAIL":        {20013, "BIND DELETE RECREATION PARAMS FAIL"},        // 取刪除娛樂城參數錯誤
	"VALIDATE_DELETE_RECREATION_PARAMS_FAIL":    {20014, "VALIDATE DELETE RECREATION PARAMS FAIL"},    // 驗證刪除娛樂城參數錯誤
	"DELETE_RECREATION_ERROR":                   {20015, "DELETE RECREATION ERROR"},                   // 刪除娛樂城資料錯誤
	"GET_RECREATION_FOCUS_ERROR":                {20016, "GET RECREATION FOCUS ERROR"},                // 取娛樂城 focus 清單錯誤
	"FOCUS_SITE_IS_NOTE_EXIST":                  {20017, "FOCUS SITE IS NOTE EXIST"},                  // 狀態的站台不存在
	"STATUS_SITE_IS_NOTE_EXIST":                 {20018, "STATUS SITE IS NOTE EXIST"},                 // 狀態的站台不存在
	"REDIS_GET_RECREATION_LIST_UNMARSHAL_ERROR": {20019, "REDIS GET RECREATION LIST UNMARSHAL ERROR"}, // 處理 redis 資料錯誤
	"WRITE_RECREATION_LIST_INTO_REDIS_ERROR":    {20020, "WRITE RECREATION LIST INTO REDIS ERROR"},    // 寫娛樂城清單到 redis 錯誤
	"UNMARSHAL_RECREATION_URL_ERROR":            {20021, "UNMARSHAL RECREATION URL ERROR"},            // 處理娛樂城 url 錯誤
	"BIND_UPDATE_RECREATION_PARAMS_FAIL":        {20007, "BIND UPDATE RECREATION PARAMS FAIL"},        // 取建立娛樂城參數錯誤

	/** 影片類別錯誤 **/
	"JSON_UNMARSHAL_VIDEO_LIST_URL_ERROR":          {21001, "JSON UNMARSHAL VIDEO LIST URL ERROR"},          // 處理影片清單路徑錯誤
	"CREATE_VIDEO_BIND_PARAMS_FAIL":                {21002, "CREATE VIDEO BIND PARAMS FAIL"},                // 新增影片 綁定參數錯誤
	"CREATE_VIDEO_VALIDATE_PARAMS_FAIL":            {21003, "CREATE VIDEO VALIDATE PARAMS FAIL"},            // 新增影片 驗證參數錯誤
	"VIDEO_NAME_ALREADY_EXISTS":                    {21004, "VIDEO NAME ALREADY EXISTS"},                    // 影片名稱重複
	"CREATE_VIDEO_PERMISSION_DENIED":               {21005, "CREATE VIDEO PERMISSION DENIED"},               // 新增影片 權限不足
	"CREATE_VIDEO_JSON_MARSHAL_ERROR":              {21006, "CREATE VIDEO JSON MARSHAL ERROR"},              // 新增影片 處理JSON錯誤
	"UPDATE_VIDEO_STATUS_ERROR":                    {21007, "UPDATE VIDEO STATUS ERROR"},                    // 更新影片 DB更改狀態失敗
	"CREATE_OR_UPDATE_VIDEO_ERROR":                 {21008, "CREATE OR UPDATE VIDEO ERROR"},                 // 新增/更新 影片錯誤
	"UPDATE_VIDEO_BIND_PARAMS_FAIL":                {21009, "UPDATE VIDEO BIND PARAMS FAIL"},                // 更新影片 綁定參數錯誤
	"UPDATE_VIDEO_VALIDATE_PARAMS_FAIL":            {21010, "UPDATE VIDEO VALIDATE PARAMS FAIL"},            // 更新影片 驗證參數錯誤
	"UPDATE_VIDEO_VALIDATE_ID_FAIL":                {21011, "UPDATE VIDEO VALIDATE ID FAIL"},                // 更新影片 驗證參數ID錯誤
	"DELETE_VIDEO_BIND_PARAMS_FAIL":                {21012, "DELETE VIDEO BIND PARAMS FAIL"},                // 刪除影片 綁定參數錯誤
	"DELETE_VIDEO_VALIDATE_PARAMS_FAIL":            {21013, "DELETE VIDEO VALIDATE PARAMS FAIL"},            // 刪除影片 驗證參數錯誤
	"NOT_DELETE_ENABLE_VIDEO":                      {21014, "NOT DELETE ENABLE VIDEO"},                      // 刪除影片 啟用中影片不得刪除
	"DELETE_VIDEO_BY_ID_ERROR":                     {21015, "DELETE VIDEO BY ID ERROR"},                     // 刪除影片 刪除BY ID 失敗
	"DELETE_VIDEO_LIST_JSON_UNMARSHAL_ERROR":       {21016, "DELETE VIDEO LIST JSON UNMARSHAL ERROR"},       // 刪除影片 處理JSON清單錯誤
	"CURL_DELETE_VIDEO_API_ERROR":                  {21017, "CURL DELETE VIDEO API ERROR"},                  // 刪除影片 CURL 圖片中心 刪除影片失敗
	"WRITE_VIDEO_DATA_REDIS_ERROR":                 {21018, "WRITE VIDEO DATA REDIS ERROR"},                 // 前台 影片寫入Redis失敗
	"FRONTEND_REDIS_GET_DATA_JSON_UNMARSHAL_ERROR": {21019, "FRONTEND REDIS GET DATA JSON UNMARSHAL ERROR"}, // 前台取Redis資料轉型失敗
	"FRONTEND_VIDEO_JSON_UNMARSHAL_LIST_ERROR":     {21020, "FRONTEND VIDEO JSON UNMARSHAL LIST ERROR"},     // 前台資料轉型失敗
	"UPDATE_VIDEO_STATUS_BIND_PARAMS_FAIL":         {21021, "UPDATE VIDEO STATUS BIND PARAMS FAIL"},         // 更新影片狀態 綁定參數錯誤
	"UPDATE_VIDEO_STATUS_VALIDATE_PARAMS_FAIL":     {21022, "UPDATE VIDEO STATUS VALIDATE PARAMS FAIL"},     // 更新影片狀態 驗證參數錯誤

	/** 排行榜類錯誤 **/
	"CURL_CYPRESS_GET_TOP_WIN_API_ERROR":    {22001, "CURL CYPRESS GET TOP WIN API ERROR"},    // 取CY API 玩家派彩排行錯誤
	"GET_RANK_LIST_BY_NAME_ERROR":           {22002, "GET RANK LIST BY NAME ERROR"},           // 使用name取排行榜設定SQL錯誤
	"GET_RANK_LIST_ERROR":                   {22003, "GET RANK LIST ERROR"},                   // 取排行榜設定SQL錯誤
	"CREATE_OR_UPDATE_RANK_ERROR":           {22004, "CREATE OR UPDATE RANK ERROR"},           // 新增/編輯 排行榜設定SQL錯誤
	"CURL_CYPRESS_GET_TOP_REWARD_API_ERROR": {22005, "CURL CYPRESS GET TOP REWARD API ERROR"}, // 取CY API 遊戲大獎次數排行榜錯誤
	"CURL_CYPRESS_GET_TOP_GAMES_API_ERROR":  {22006, "CURL CYPRESS GET TOP GAMES API ERROR"},  // 取CY API 熱門遊戲排行榜錯誤
	"RATE_IS_NOT_IN_LIST":                   {22007, "RATE IS NOT IN LIST"},                   // Rate值不在清單之中

	/** 真人視訊錯誤 **/
	"GET_LIVE_LINK_ERROR":            {23001, "GET LIVE LINK ERROR"},            // 取真人視訊連結錯誤,
	"LIVE_LINK_JSON_UNMARSHAL_ERROR": {23002, "LIVE LINK JSON UNMARSHAL ERROR"}, // 真人視訊資料 unmarshal 錯誤

	/** 廣告錯誤 **/
	"AD_NAME_ALREADY_EXISTS":    {24001, "AD NAME ALREADY EXISTS"},    // 廣告已存在
	"QUERY_AD_ERROR":            {24002, "QUERY AD ERROR"},            // 搜尋廣告錯誤
	"CREATE_OR_UPDATE_AD_ERROR": {24003, "CREATE OR UPDATE AD ERROR"}, // 新增/修改廣告錯誤
	"ADID_NOT_EXISTS":           {24004, "ADID NOT EXISTS"},           // 該筆資料不存在

	/** 博彩文章錯誤 **/
	"JSON_UNMARSHAL_ARTICLE_CONTENT_LIST_ERROR":            {25001, "JSON UNMARSHAL ARTICLE CONTENT LIST ERROR"},            // 處理文章內容轉型錯誤
	"JSON_UNMARSHAL_ARTICLE_IMAGE_LIST_ERROR":              {25002, "JSON UNMARSHAL ARTICLE IMAGE LIST ERROR"},              // 處理文章圖片路徑轉型錯誤
	"FRONTEND_REDIS_GET_ARTICLE_DATA_JSON_UNMARSHAL_ERROR": {25003, "FRONTEND REDIS GET ARTICLE DATA JSON UNMARSHAL ERROR"}, // 前台處理Redis資料轉型錯誤
	"FRONTEND_JSON_UNMARSHAL_ARTICLE_CONTENT_LIST_ERROR":   {25004, "FRONTEND JSON UNMARSHAL ARTICLE CONTENT LIST ERROR"},   // 前台處理文章內容轉型錯誤
	"FRONTEND_JSON_UNMARSHAL_ARTICLE_IMAGE_LIST_ERROR":     {25005, "FRONTEND JSON UNMARSHAL ARTICLE IMAGE LIST ERROR"},     // 前台處理文章圖片路徑轉型錯誤
	"FRONTEND_WRITE_ARTICLE_DATA_REDIS_ERROR":              {25006, "FRONTEND WRITE ARTICLE DATA REDIS ERROR"},              // 前台文章寫入redis錯誤
	"CREATE_ARTICLE_BIND_PARAMS_FAIL":                      {25006, "CREATE ARTICLE BIND PARAMS FAIL"},                      // 新增文章 綁定參數錯誤
	"CREATE_ARTICLE_VALIDATE_PARAMS_FAIL":                  {25007, "CREATE ARTICLE VALIDATE PARAMS FAIL"},                  // 新增文章 驗證參數錯誤
	"CREATE_ARTICLE_PERMISSION_DENIED":                     {25008, "CREATE ARTICLE PERMISSION DENIED"},                     // 新增文章 權限不足
	"CREATE_ARTICLE_CONTENT_JSON_MARSHAL_ERROR":            {25009, "CREATE ARTICLE CONTENT JSON MARSHAL ERROR"},            // 新增文章 處理JSON錯誤
	"CREATE_ARTICLE_IMAGE_JSON_MARSHAL_ERROR":              {25010, "CREATE ARTICLE IMAGE JSON MARSHAL ERROR"},              // 新增文章圖片路徑 處理JSON錯誤
	"CREATE_OR_UPDATE_ARTICLE_ERROR":                       {25011, "CREATE OR UPDATE ARTICLE ERROR"},                       // 新增/更新 文章錯誤
	"UPDATE_ARTICLE_BIND_PARAMS_FAIL":                      {25012, "UPDATE ARTICLE BIND PARAMS FAIL"},                      // 更新文章 綁定參數錯誤
	"UPDATE_ARTICLE_VALIDATE_PARAMS_FAIL":                  {25013, "UPDATE ARTICLE VALIDATE PARAMS FAIL"},                  // 更新文章 驗證參數錯誤
	"ARTICLE_TYPE_ID_NOT_EXISTS":                           {25014, "ARTICLE TYPE ID NOT EXISTS"},                           // 更新文章分類 檢查ID錯誤
	"UPDATE_ARTICLE_PERMISSION_DENIED":                     {25015, "UPDATE ARTICLE PERMISSION DENIED"},                     // 更新文章 權限不足
	"UPDATE_ARTICLE_CONTENT_JSON_MARSHAL_ERROR":            {25016, "UPDATE ARTICLE CONTENT JSON MARSHAL ERROR"},            // 更新文章 處理JSON錯誤
	"UPDATE_ARTICLE_IMAGE_JSON_MARSHAL_ERROR":              {25017, "UPDATE ARTICLE IMAGE JSON MARSHAL ERROR"},              // 更新文章圖片路徑 處理JSON錯誤
	"DELETE_ARTICLE_BIND_PARAMS_FAIL":                      {25018, "DELETE ARTICLE BIND PARAMS FAIL"},                      // 刪除文章 綁定參數錯誤
	"DELETE_ARTICLE_VALIDATE_PARAMS_FAIL":                  {25019, "DELETE ARTICLE VALIDATE PARAMS FAIL"},                  // 刪除文章 驗證參數錯誤
	"NOT_DELETE_ENABLE_ARTICLE":                            {25020, "NOT DELETE ENABLE ARTICLE"},                            // 刪除文章 啟用中文章不能刪除
	"DELETE_ARTICLE_BY_ID_ERROR":                           {25021, "DELETE ARTICLE BY ID ERROR"},                           // 刪除文章DB錯誤
	"ARTICLE_ID_NOT_EXISTS":                                {25022, "ARTICLE ID NOT EXISTS"},                                // 更新文章 檢查ID錯誤
	"CREATE_ARTICLE_TITLE_JSON_MARSHAL_ERROR":              {25023, "CREATE ARTICLE TITLE JSON MARSHAL ERROR"},              // 新增文章標題 處理JSON錯誤
	"UPDATE_ARTICLE_TITLE_JSON_MARSHAL_ERROR":              {25024, "UPDATE ARTICLE TITLE JSON MARSHAL ERROR"},              // 更新文章標題 處理JSON錯誤
	"FRONTEND_JSON_UNMARSHAL_ARTICLE_TITLE_LIST_ERROR":     {25025, "FRONTEND JSON UNMARSHAL ARTICLE TITLE LIST ERROR"},     // 前台處理文章標題轉型錯誤
	"JSON_UNMARSHAL_ARTICLE_TITLE_LIST_ERROR":              {25026, "JSON UNMARSHAL ARTICLE TITLE LIST ERROR"},              // 處理文章標題轉型錯誤
	"JSON_UNMARSHAL_ARTICLE_DESCRIPTION_LIST_ERROR":        {25027, "JSON UNMARSHAL ARTICLE DESCRIPTION LIST ERROR"},        // 處理文章描述轉型錯誤
	"CREATE_ARTICLE_NAME_STRING_TOO_LONG_ERROR":            {25028, "CREATE ARTICLE NAME STRING TOO LONG ERROR"},            // 映射字串長度過長
	"CREATE_ARTICLE_STATUS_STRING_TOO_LONG_ERROR":          {25029, "CREATE ARTICLE STATUS STRING TOO LONG ERROR"},          // 映射字串長度過長
	"CREATE_ARTICLE_STATUS_JSON_MARSHAL_ERROR":             {25030, "CREATE ARTICLE STATUS_ SON MARSHAL ERROR"},             // 處理文章名稱轉型錯誤

	/** 博彩文章類別錯誤 **/
	"GET_ARTICLE_TYPE_LIST_ERROR":                               {26001, "GET ARTICLE TYPE LIST ERROR"},                               // 取的文章分類錯誤
	"CREATE_ARTICLE_TYPE_BIND_PARAMS_FAIL":                      {26002, "CREATE ARTICLE TYPE BIND PARAMS FAIL"},                      // 新增文章類別 綁定參數錯誤
	"CREATE_ARTICLE_TYPE_VALIDATE_PARAMS_FAIL":                  {26003, "CREATE ARTICLE TYPE VALIDATE PARAMS FAIL"},                  // 新增文章類別 驗證參數錯誤
	"CREATE_ARTICLE_TYPE_PERMISSION_DENIED":                     {26004, "CREATE ARTICLE TYPE PERMISSION DENIED"},                     // 新增文章類別 權限不足
	"ARTICLE_TYPE_NAME_ALREADY_EXISTS":                          {26005, "ARTICLE TYPE NAME ALREADY EXISTS"},                          // 新增文章類別 文章名稱重複
	"CREATE_ARTICLE_TYPE_AMOUNT_LIMIT":                          {26006, "CREATE ARTICLE TYPE AMOUNT LIMIT"},                          // 新增文章類別 數量限制
	"CREATE_OR_UPDATE_ARTICLE_TYPE_ERROR":                       {26007, "CREATE OR UPDATE ARTICLE TYPE ERROR"},                       // 新增/更新 文章類別錯誤
	"UPDATE_ARTICLE_TYPE_BIND_PARAMS_FAIL":                      {26009, "UPDATE ARTICLE TYPE BIND PARAMS FAIL"},                      // 更新文章類別 綁定參數錯誤
	"UPDATE_ARTICLE_TYPE_VALIDATE_PARAMS_FAIL":                  {26010, "UPDATE ARTICLE TYPE VALIDATE PARAMS FAIL"},                  // 更新文章類別 驗證參數錯誤
	"DELETE_ARTICLE_TYPE_BIND_PARAMS_FAIL":                      {26011, "DELETE ARTICLE TYPE BIND PARAMS FAIL"},                      // 刪除文章類別 綁定參數錯誤
	"DELETE_ARTICLE_TYPE_VALIDATE_PARAMS_FAIL":                  {26012, "DELETE ARTICLE TYPE VALIDATE PARAMS FAIL"},                  // 刪除文章類別 驗證參數錯誤
	"DELETE_ARTICLE_TYPE_BY_ID_ERROR":                           {26013, "DELETE ARTICLE TYPE BY ID ERROR"},                           // 刪除文章類別 DB錯誤
	"CREATE_ARTICLE_TYPE_NAME_JSON_MARSHAL_ERROR":               {26014, "CREATE ARTICLE TYPE NAME JSON MARSHAL ERROR"},               // 新增文章分類名稱 JSON MARSHAL錯誤
	"JSON_UNMARSHAL_ARTICLE_TYPE_NAME_LIST_ERROR":               {26015, "JSON UNMARSHAL ARTICLE TYPE NAME LIST ERROR"},               // 文章分類列表JSON MARSHAL錯誤
	"UPDATE_ARTICLE_TYPE_NAME_JSON_MARSHAL_ERROR":               {26016, "UPDATE ARTICLE TYPE NAME JSON MARSHAL ERROR"},               // 更新文章分類名稱 JSON MARSHAL錯誤
	"FRONTEND_REDIS_GET_ARTICLE_TYPE_DATA_JSON_UNMARSHAL_ERROR": {26017, "FRONTEND REDIS GET ARTICLE TYPE DATA JSON UNMARSHAL ERROR"}, // 前台取文章列表JSON UNMARSHAL錯誤
	"FRONTEND_WRITE_ARTICLE_TYPE_DATA_REDIS_ERROR":              {26018, "FRONTEND WRITE ARTICLE TYPE DATA REDIS ERROR"},              // 前台寫入Redis錯誤

	/** 連結類型 **/
	"CREATE_LINK_TYPE_BIND_PARAMS_FAIL":     {60001, "CREATE LINK TYPE BIND PARAMS FAIL"},     // bind 建立類型參數錯誤
	"CREATE_LINK_TYPE_VAILDATE_PARAMS_FAIL": {60002, "CREATE LINK TYPE VAILDATE PARAMS FAIL"}, // 驗證建立類型參數錯誤
	"UPDATE_LINK_TYPE_BIND_PARAMS_FAIL":     {60003, "UPDATE LINK TYPE BIND PARAMS FAIL"},     // bind 更新類型參數錯誤
	"UPDATE_LINK_TYPE_VAILDATE_PARAMS_FAIL": {60004, "UPDATE LINK TYPE VAILDATE PARAMS FAIL"}, // 驗證更新類型參數錯誤
	"CU_LINK_TYPE_NAME_JSON_MARSHAL_ERROR":  {60005, "CU LINK TYPE NAME JSON MARSHAL ERROR"},  // 新增/修改類型名稱資料錯誤
	"CREATE_OR_UPDATE_LINKS_TYPE_ERROR":     {60006, "CREATE OR UPDATE LINKS TYPE ERROR"},     // DB 建立/修改連結類型錯誤
	"DELETE_LINK_TYPE_BIND_PARAMS_FAIL":     {60007, "DELETE LINK TYPE BIND PARAMS FAIL"},     // bind 刪除類型參數錯誤
	"DELETE_LINK_TYPE_VAILDATE_PARAMS_FAIL": {60008, "DELETE LINK TYPE VAILDATE PARAMS FAIL"}, // 驗證刪除類型參數錯誤
	"LINK_TYPE_STILL_USING":                 {60009, "LINK TYPE STILL USING"},                 // 類型仍在使用中
	"DELETE_LINK_TYPE_ERROR":                {60010, "DELETE LINK TYPE ERROR"},                // 刪除類型資料錯誤
	"LINK_TYPE_JSON_UNMARSHAL_ERROR":        {60011, "LINK TYPE JSON UNMARSHAL ERROR"},        // Unmarshal link type 錯誤
	"GET_LINKS_TYPE_LIST_ERROR":             {60012, "GET LINKS TYPE LIST ERROR"},             // 取連結類型清單錯誤
	"LINK_TYPE_NAME_JSON_UNMARSHAL_ERROR":   {60013, "LINK TYPE NAME JSON UNMARSHAL ERROR"},   // 連結類型名稱 Unmarshal 錯誤
	"LINK_SORT_CHANGE_PARAMS_TYPE_FAIL":     {60014, "LINK SORT CHANGE PARAMS TYPE FAIL"},     // 字串轉數字錯誤

	/** 遊戲第二遊戲類型錯誤 **/
	"GAME_TYPE2_CREATE_BIND_PARAMS_FAIL":         {27001, "GAME TYPE2 CREATE BIND PARAMS FAIL"},         // 新增第二遊戲類型 綁定參數錯誤
	"GAME_TYPE2_CREATE_VALIDATE_PARAMS_FAIL":     {27002, "GAME TYPE2 CREATE VALIDATE PARAMS FAIL"},     // 新增第二遊戲類型 驗證參數錯誤
	"GET_GAME_TYPE2_LIST_ERROR":                  {27003, "GET GAME TYPE2 LIST ERROR"},                  // 取第二遊戲類型列表失敗
	"CREATE_OR_UPDATE_GAME_TYPE2_ERROR":          {27004, "CREATE OR UPDATE GAME TYPE2 ERROR"},          // 新增/更新第二遊戲類型失敗
	"DELETE_GAME_TYPE2_ERROR":                    {27005, "DELETE_GAME_TYPE2_ERROR"},                    // 刪除第二遊戲類型失敗
	"GAME_TYPE2_UPDATE_ID_VALIDATE_FAIL":         {27006, "GAME TYPE2 UPDATE ID VALIDATE FAIL"},         // 更新第二遊戲類型ID值錯誤
	"GET_GAME_TYPE2_NAME_JSON_UNMARSHAL_ERROR":   {27007, "GET GAME TYPE2 NAME JSON UNMARSHAL ERROR"},   // 第二遊戲類型NAME JSON UNMARSHAL錯誤
	"REDIS_GAME_TYPE2_LIST_JSON_UNMARSHAL_ERROR": {27008, "REDIS GAME TYPE2 LIST JSON UNMARSHAL ERROR"}, // Redis取出 第二遊戲類型列表 JSON UNMARSHAL錯誤
	"GAME_TYPE2_LIST_JSON_MARSHAL_ERROR":         {27009, "GAME TYPE2 LIST JSON MARSHAL ERROR"},         // 第二遊戲類型列表 JSON MARSHAL錯誤
	"GAME_TYPE2_ALREADY_USED":                    {27010, "GAME TYPE2 ALREADY USED"},                    // 第二遊戲類型使用中 無法刪除

	/** HTML CODE 錯誤 **/
	"UPDATE_HTML_CODE_BIND_PARAMS_FAIL":             {28001, "UPDATE HTML CODE BIND PARAMS FAIL"},             // 更新HTML參數 bind 失敗
	"UPDATE_HTML_CODE_ARTICLE_VAILDATE_PARAMS_FAIL": {28002, "UPDATE HTML CODE ARTICLE VAILDATE PARAMS FAIL"}, // 更新HTML參數驗證失敗
	"UPDATE_HTML_CODE_PERMISSION_DENIED":            {28003, "UPDATE HTML CODE PERMISSION DENIED"},            // 更新HTML 權限不足
	"HTML_CODE_JSON_MARSHAL_ERROR":                  {28004, "HTML CODE JSON MARSHAL ERROR"},                  // 更新HTML CODE JSON MARSHAL錯誤
	"HTML_STATUS_JSON_MARSHAL_ERROR":                {28005, "HTML STATUS JSON MARSHAL ERROR"},                // 更新HTML STATUS JSON MARSHAL錯誤
	"CREATE_OR_UPDATE_HTML_ERROR":                   {28006, "CREATE OR UPDATE HTML ERROR"},                   // 更新/新增 HTML DB錯誤
	"GET_HTML_DATA_ERROR":                           {28007, "GET HTML DATA ERROR"},                           // 取HTML資料DB錯誤
	"GET_HTML_DATA_BY_ID_ERROR":                     {28008, "GET HTML DATA BY ID ERROR"},                     // 取HTML資料(by ID) DB錯誤
	"HTML_CODE_BACKEND_JSON_UNMARSHAL_ERROR":        {28009, "HTML CODE BACKEND JSON UNMARSHAL ERROR"},        // 後台取HTML CODE JSON_UNMARSHAL 錯誤
	"HTML_STATUS_BACKEND_JSON_UNMARSHAL_ERROR":      {28010, "HTML STATUS BACKEND JSON UNMARSHAL ERROR"},      // 後台取HTML STATUS JSON_UNMARSHAL 錯誤
	"HTML_FRONTEND_JSON_UNMARSHAL_ERROR":            {28011, "HTML FRONTEND JSON UNMARSHAL ERROR"},            // 前台取HTML(Redis) JSON_UNMARSHAL 錯誤
	"HTML_CODE_FRONTEND_JSON_UNMARSHAL_ERROR":       {28012, "HTML CODE FRONTEND JSON UNMARSHAL ERROR"},       // 前台取HTML CODE JSON_UNMARSHAL 錯誤
	"HTML_STATUS_FRONTEND_JSON_UNMARSHAL_ERROR":     {28013, "HTML STATUS FRONTEND JSON UNMARSHAL ERROR"},     // 前台取HTML STATUS JSON_UNMARSHAL 錯誤

	/** 彩票 錯誤 **/
	"LOTTERY_GAMETOKEN_JSON_UNMARSHAL_ERROR": {29001, "LOTTERY GAMETOKEN JSON UNMARSHAL ERROR"}, // 取彩票 game token unmarshal 錯誤
	"GET_LOTTERY_GAMETOKEN_ERROR":            {29002, "GET LOTTERY GAMETOKEN ERROR"},            // 取彩票 game token 錯誤

	/** 虛擬幣大廳 錯誤 **/
	"CURL_CYPRESS_WALLET_QRCODE_API_ERROR":      {30001, "CURL CYPRESS WALLET QRCODE API ERROR"},      // 產生虛擬幣交易所綁定錢包連結錯誤
	"CURL_CYPRESS_CHECK_WALLET_API_ERROR":       {30002, "CURL CYPRESS CHECK WALLET API ERROR"},       // 檢查玩家是否有綁定交易所錢包錯誤
	"CURL_CYPRESS_WALLET_UNBIND_API_ERROR":      {30003, "CURL CYPRESS WALLET UNBIND API ERROR"},      // 玩家解除交易所錢包綁定錯誤
	"CURL_CYPRESS_CHECK_AGENT_WALLET_API_ERROR": {30004, "CURL CYPRESS CHECK AGENT WALLET API ERROR"}, // 檢查代理交易所狀態
	"CURL_CYPRESS_GET_ORDER_BY_TIME_API_ERROR":  {30005, "CURL CYPRESS GET ORDER BY TIME API ERROR"},  // 以時間區間查詢注單失敗
	"CURL_CYPRESS_GET_ORDER_BY_ID_API_ERROR":    {30006, "CURL CYPRESS GET ORDER BY ID API ERROR"},    // 以注單識別局號查詢單一注單失敗
	"CURL_CYPRESS_FREE_TICKET_API_ERROR":        {30007, "CURL CYPRESS FREE TICKET API ERROR"},        // 取得玩家免費券清單失敗
	"CQRITY_FAVORITE_GAME_JSON_MARSHAL_ERROR":   {30008, "CQRITY FAVORITE GAME JSON MARSHAL ERROR"},   // 玩家最愛遊戲列表 JSON_UNMARSHAL 錯誤

	/** 假瀏覽 錯誤 **/
	"BIND_UPDATE_SCAN_PARAMS_FAIL":     {31001, "BIND UPDATE SCAN PARAMS FAIL"},     // 取更新假瀏覽資料錯誤
	"VALIDATE_UPDATE_SCAN_PARAMS_FAIL": {31002, "VALIDATE UPDATE SCAN PARAMS FAIL"}, // 驗證更新假瀏覽資料錯誤
	"MAX_VAULE_LESS_THAN_MIN":          {31003, "MAX VAULE LESS THAN MIN"},          // 最大值小於最小值
	"CREATE_OR_UPDATE_SCAN_SQL_ERROR":  {31004, "CREATE OR UPDATE SCAN SQL ERROR"},  // 新增或修改假瀏覽規則錯誤
	"GET_SCAN_LIST_SQL_ERROR":          {31005, "GET SCAN LIST SQL ERROR"},          // 取假瀏覽清單錯誤
	"DELETE_SCAN_SQL_ERROR":            {31006, "DELETE SCAN SQL ERROR"},            // 刪除假瀏覽資料錯誤

	/** GameSite 遊戲商 錯誤**/
	"GAME_SITE_LIST_JSON_UNMARSHALL_ERROR":       {32001, "GAME SITE LIST JSON UNMARSHALL ERROR"},       // 遊戲商列表 JSON UNMARSHAL 錯誤
	"GAME_SITE_LIST_REDIS_JSON_UNMARSHALL_ERROR": {32002, "GAME SITE LIST REDIS JSON UNMARSHALL ERROR"}, // 遊戲商列表 Redis JSON UNMARSHAL 錯誤
	"GAME_SITE_LIST_JSON_MARSHALL_ERROR":         {32003, "GAME SITE LIST JSON MARSHALL ERROR"},         // 遊戲商列表 Redis JSON MARSHAL 錯誤
	"GAME_SITE_LIST_WRITE_REDIS_DATA_ERROR":      {32004, "GAME SITE LIST WRITE REDIS DATA ERROR"},      // 遊戲商列表 Redis 寫入資料 錯誤
	"GAME_SITE_NOT_EXISTS":                       {32005, "GAME SITE NOT EXISTS"},                       // 遊戲商列表 不存在
	"GAME_SITE_LIST_DELETE_REDIS_DATA_ERROR":     {32006, "GAME SITE LIST DELETE REDIS DATA ERROR"},     // 遊戲商列表 Redis 刪除資料 錯誤
	"CHECK_GAME_SITE_EXIST_DB_ERROR":             {32007, "CHECK GAME SITE EXIST DB ERROR"},             // 查詢遊戲商列表 DB 錯誤
	"GET_GAME_SITE_LIST_DB_ERROR":                {32008, "GET GAME SITE LIST DB ERROR"},                // 取得遊戲商列表 DB 錯誤
	"DELETE_GAME_SITE_DB_ERROR":                  {32009, "DELETE GAME SITE DB ERROR"},                  // 刪除遊戲商列表 DB 錯誤
	"CREATE_OR_UPDATE_GAME_SITE_ERROR":           {32010, "CREATE OR UPDATE GAME SITE ERROR"},           // 創建/修改遊戲商列表 DB 錯誤

	/** 文章類別選單錯誤 **/
	"GET_ARTICLE_MENU_LIST_ERROR":                               {33001, "GET ARTICLE MENU LIST ERROR"},                               // 取的文章分類選單錯誤
	"CREATE_ARTICLE_MENU_BIND_PARAMS_FAIL":                      {33002, "CREATE ARTICLE MENU BIND PARAMS FAIL"},                      // 新增文章類別選單 綁定參數錯誤
	"CREATE_ARTICLE_MENU_VALIDATE_PARAMS_FAIL":                  {33003, "CREATE ARTICLE MENU VALIDATE PARAMS FAIL"},                  // 新增文章類別選單 驗證參數錯誤
	"CREATE_ARTICLE_MENU_PERMISSION_DENIED":                     {33004, "CREATE ARTICLE MENU PERMISSION DENIED"},                     // 新增文章類別選單 權限不足
	"ARTICLE_MENU_NAME_ALREADY_EXISTS":                          {33005, "ARTICLE MENU NAME ALREADY EXISTS"},                          // 新增文章類別選單 文章名稱重複
	"CREATE_ARTICLE_MENU_AMOUNT_LIMIT":                          {33006, "CREATE ARTICLE MENU AMOUNT LIMIT"},                          // 新增文章類別選單 數量限制
	"CREATE_OR_UPDATE_ARTICLE_MENU_ERROR":                       {33007, "CREATE OR UPDATE ARTICLE MENU ERROR"},                       // 新增/更新 文章類別選單錯誤
	"UPDATE_ARTICLE_MENU_BIND_PARAMS_FAIL":                      {33009, "UPDATE ARTICLE MENU BIND PARAMS FAIL"},                      // 更新文章類別選單 綁定參數錯誤
	"UPDATE_ARTICLE_MENU_VALIDATE_PARAMS_FAIL":                  {33010, "UPDATE ARTICLE MENU VALIDATE PARAMS FAIL"},                  // 更新文章類別選單 驗證參數錯誤
	"DELETE_ARTICLE_MENU_BIND_PARAMS_FAIL":                      {33011, "DELETE ARTICLE MENU BIND PARAMS FAIL"},                      // 刪除文章類別選單 綁定參數錯誤
	"DELETE_ARTICLE_MENU_VALIDATE_PARAMS_FAIL":                  {33012, "DELETE ARTICLE MENU VALIDATE PARAMS FAIL"},                  // 刪除文章類別選單 驗證參數錯誤
	"DELETE_ARTICLE_MENU_BY_ID_ERROR":                           {33013, "DELETE ARTICLE MENU BY ID ERROR"},                           // 刪除文章類別選單 DB錯誤
	"CREATE_ARTICLE_MENU_NAME_JSON_MARSHAL_ERROR":               {33014, "CREATE ARTICLE MENU NAME JSON MARSHAL ERROR"},               // 新增文章分類選單名稱 JSON MARSHAL錯誤
	"JSON_UNMARSHAL_ARTICLE_MENU_NAME_LIST_ERROR":               {33015, "JSON UNMARSHAL ARTICLE MENU NAME LIST ERROR"},               // 文章分類選單列表JSON MARSHAL錯誤
	"UPDATE_ARTICLE_MENU_NAME_JSON_MARSHAL_ERROR":               {33016, "UPDATE ARTICLE MENU NAME JSON MARSHAL ERROR"},               // 更新文章分類選單名稱 JSON MARSHAL錯誤
	"FRONTEND_REDIS_GET_ARTICLE_MENU_DATA_JSON_UNMARSHAL_ERROR": {33017, "FRONTEND REDIS GET ARTICLE MENU DATA JSON UNMARSHAL ERROR"}, // 前台取文章選單列表JSON UNMARSHAL錯誤
	"FRONTEND_WRITE_ARTICLE_MENU_DATA_REDIS_ERROR":              {33018, "FRONTEND WRITE ARTICLE MENU DATA REDIS ERROR"},              // 前台寫入Redis錯誤
}
