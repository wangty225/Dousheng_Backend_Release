namespace go relation_hertz
include "user.thrift"


/* ============== 1. define interface: /douyin/relation/action/================== */
struct DouyinRelationActionRequest {
    1:required string token (api.query="token")   // 用户鉴权token
    2:required i64 to_user_id (api.query="to_user_id") // 对方用户id
    3:required i32 action_type (api.query="action_type") // 1-关注，2-取消关注
}

struct DouyinRelationActionResponse {
    1:required i32 status_code   // 状态码，0-成功，其他值-失败
    2:optional string status_msg // 返回状态描述
}


/* ============== 2. define interface: /douyin/relation/follow/list/================== */
struct DouyinRelationFollowListRequest {
    1:required i64 user_id (api.query="user_id") // 用户id
    2:required string token (api.query="token") // 用户鉴权token
}

struct DouyinRelationFollowListResponse {
    1:required i32 status_code           // 状态码，0-成功，其他值-失败
    2:optional string status_msg         // 返回状态描述
    3:required list<user.User> user_list // 用户信息列表
}


/* ============== 3. define interface: /douyin/relation/follower/list/================== */
struct DouyinRelationFollowerListRequest {
    1:required i64 user_id (api.query="user_id") // 用户id
    2:required string token (api.query="token") // 用户鉴权token
}

struct DouyinRelationFollowerListResponse {
    1:required i32 status_code            // 状态码，0-成功，其他值-失败
    2:optional string status_msg          // 返回状态描述
    3:required list<user.User> user_list  // 用户信息列表
}


/* ============== 4. define interface: /douyin/relation/friend/list/================== */
struct DouyinRelationFriendListRequest {
    1:required i64 user_id  (api.query="user_id") // 用户id
    2:required string token (api.query="token") // 用户鉴权token
}

struct DouyinRelationFriendListResponse {
    1:required i32 status_code            // 状态码，0-成功，其他值-失败
    2:optional string status_msg          // 返回状态描述
    3:required list<FriendUser> user_list // 用户列表
}

struct FriendUser {
    1:optional string message, // 和该好友的最新聊天消息
    2:required i64 msgType,    // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
    3:required user.User baseUser,       // todo: 实现extends关键字转化，调用name属性时：friendUser.baseUser.name，如调用失败则检查此处。
}


service RelationService {
    DouyinRelationActionResponse ActionRelation(1:required DouyinRelationActionRequest req) (api.post="/douyin/relation/action/");
    DouyinRelationFollowListResponse ListFollowRelation(1:required DouyinRelationFollowListRequest req) (api.get="/douyin/relation/follow/list/");
    DouyinRelationFollowerListResponse ListFollowerRelation(1:required DouyinRelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/");
    DouyinRelationFriendListResponse ListFriendRelation(1:required DouyinRelationFriendListRequest req) (api.get="/douyin/relation/friend/list/");
}