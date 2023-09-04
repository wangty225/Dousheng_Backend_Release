namespace go user_hertz


/* ============== 1. define interface: /douyin/user/================== */
struct DouyinUserRequest {
    1:required i64 user_id (api.query="user_id")
    2:required string token (api.query="token")
}

struct DouyinUserResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required User user            // 用户信息
}

struct User {
    1:required i64 id                        // 用户id
    2:required string name                   // 用户名称
    3:required i64 follow_count              // 关注总数
    4:required i64 follower_count            // 粉丝总数
    5:required bool is_follow                // true-已关注，false-未关注
    6:optional string avator                 // 用户头像
    7:optional string background_image       // 用户个人页顶部大图
    8:optional string signature              // 个人简介
    9:optional i64 total_favorited           // 获赞数量
    10:optional i64 work_count               // 作品数量
    11:optional i64 favorite_count           // 点赞数量
}


/* ============== 2. define interface: /douyin/user/register/================= */
struct DouyinUserRegisterRequest {
    1:required string username (api.query="username")   // 注册用户名，最长32个字符
    2:required string password (api.query="password")   // 密码，最长32个字符
}

struct DouyinUserRegisterResponse {
    1:required i32 status_code    // 状态码，0-成功，其他值-失败
    2:optional string status_msg  // 返回状态描述
    3:required i64 user_id        // 用户id
    4:required string token       // 用户鉴权token
}


/* ============== 3. define interface: /douyin/user/login/================= */
struct DouyinUserLoginRequest {
    1:required string username (api.query="username")  // 登录用户名
    2:required string password (api.query="password")  // 登录密码
}

struct DouyinUserLoginResponse {
    1:required i32 status_code     // 状态码，0-成功，其他值-失败
    2:optional string status_msg   // 返回状态描述
    3:required i64 user_id         // 用户id
    4:required string token        // 用户鉴权token
}


/* ============== 4. define services:================= */
service UserService {
    DouyinUserResponse User(1:required DouyinUserRequest req) (api.get="/douyin/user/");
    DouyinUserRegisterResponse RegisterUser(1:required DouyinUserRegisterRequest req) (api.post="/douyin/user/register/");
    DouyinUserLoginResponse LoginUser(1:required DouyinUserLoginRequest req) (api.post="/douyin/user/login/")
}
