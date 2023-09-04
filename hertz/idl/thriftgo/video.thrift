namespace go video_hertz
include "user.thrift"

/* ============== 1. define interface: /douyin/feed/================== */
struct DouyinFeedRequest {
    1:optional i64 latest_time (api.query="latest_time")   // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2:optional string token    (api.query="token")   // 可选参数，登录用户设置
}

struct DouyinFeedResponse {
    1:required i32 status_code    // 状态码，0-成功，其他值-失败
    2:optional string status_msg  // 返回状态描述
    3:required list<Video> video_list  // 视频列表
    4:optional i64 next_time      // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct Video {
    1:required i64 id             // 视频唯一标识
    2:required user.User author   // 视频作者信息
    3:required string play_url    // 视频播放地址
    4:required string cover_url   // 视频封面地址
    5:required i64 favorite_count // 视频的点赞总数
    6:required i64 comment_count  // 视频的评论总数
    7:required bool is_favorite   // true-已点赞，false-未点赞
    8:required string title       // 视频标题
}


/* ============== 2. define interface: /douyin/publish/action================== */
struct DouyinPublishActionRequest {
    1:required string token (api.form="token")   // 用户鉴权token
    2:required binary data  (api.form="data")   // 视频数据
    3:required string title (api.form="title")   // 视频标题
}

struct DouyinPublishActionResponse {
    1:required i32 status_code
    2:optional string status_msg
}


/* ============== 3. define interface: /douyin/publish/list================== */
struct DouyinPublishListRequest {
    1:required i64 user_id (api.query="user_id")
    2:required string token (api.query="token")
}

struct DouyinPublishListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<Video> video_list  // 用户发布的视频列表
}

/* ============== 4. define services: ================== */
service VideoService {
    DouyinFeedResponse Feed(1:required DouyinFeedRequest req) (api.get="/douyin/feed/");
    DouyinPublishActionResponse ActionPublish(1:required DouyinPublishActionRequest req) (api.post="/douyin/publish/action/");
    DouyinPublishListResponse ListPublish(1:required DouyinPublishListRequest req) (api.get="/douyin/publish/list/");
}