namespace go message_hertz


/* ============== 1. define interface: /douyin/message/chat/================== */
struct DouyinMessageChatRequest {
    1:required string token (api.query="token")   // 用户鉴权token
    2:required i64 to_user_id (api.query="to_user_id")   // 对方用户id
    3:required i64 pre_msg_time (api.query="pre_msg_time") // 上次最新消息的时间（新增字段-apk更新中）
}

struct DouyinMessageChatResponse {
    1:required i32 status_code            // 状态码，0-成功，其他值-失败
    2:optional string status_msg          // 返回状态描述
    3:required list<Message> message_list // 消息列表
}


struct Message {
    1:required i64 id             // 消息id
    2:required i64 to_user_id     // 该消息接收者的id
    3:required i64 from_user_id   // 该消息发送者的id
    4:required string content     // 消息内容
    5:optional string create_time // 消息创建时间
}


/* ============== 2. define interface: /douyin/message/action/ ================== */
struct DouyinMessageActionRequest {
    1:required string token (api.query="token")    // 用户鉴权token
    2:required i64 to_user_id (api.query="to_user_id")  // 对方用户id
    3:required i32 action_type (api.query="action_type") // 1-发送消息
    4:required string content (api.query="content")  // 消息内容
}

struct DouyinMessageActionResponse {
    1:required i32 status_code            // 状态码，0-成功，其他值-失败
    2:optional string status_msg          // 返回状态描述
}


/* ============== Final. define services ================== */
service MessageService {
    DouyinMessageChatResponse ChatMessage(1:required DouyinMessageChatRequest req) (api.get="/douyin/message/chat/");
    DouyinMessageActionResponse ActionMessage(1:required DouyinMessageActionRequest req) (api.post="/douyin/message/action/");
}