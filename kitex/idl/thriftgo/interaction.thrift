namespace go interaction
include "user.thrift"
include "video.thrift"


/* ============== 1. define interface: /douyin/favorite/action/================== */
struct DouyinFavoriteActionRequest {
    1:required string token    // 用户鉴权token
    2:required i64 video_id    // 视频id
    3:required i32 action_type // 1-点赞，2-取消点赞
}

struct DouyinFavoriteActionResponse {
    1:required i32 status_code    // 状态码，0-成功，其他值-失败
    2:optional string status_msg  // 返回状态描述
}


/* ============== 2. define interface: /douyin/favorite/list/================== */
struct DouyinFavoriteListRequest {
    1:required i64 user_id  // 用户id
    2:required string token // 用户鉴权token
}

struct DouyinFavoriteListResponse {
    1:required i32 status_code                // 状态码，0-成功，其他值-失败
    2:optional string status_msg              // 返回状态描述
    3:required list<video.Video> video_list   // 用户点赞视频列表
}


/* ============== 3. define interface: /douyin/comment/action/================== */
struct DouyinCommentActionRequest {
    1:required string token
    2:required i64 video_id
    3:required i32 action_type     // 1-发布评论，2-删除评论
    4:optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    5:optional i64 comment_id      // 要删除的评论id，在action_type=2的时候使用
}

struct DouyinCommentActionResponse {
    1:required i32 status_code   // 状态码，0-成功，其他值-失败
    2:optional string status_msg // 返回状态描述
    3:optional Comment comment   // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct Comment {
    1:required i64 id             // 视频评论id
    2:required user.User user     // 评论用户信息
    3:required string content     // 评论内容
    4:required string create_date // 评论发布日期，格式 mm-dd
}


/* ============== 4. define interface: /douyin/comment/list/================== */
struct DouyinCommentListRequest {
    1:required string token
    2:required i64 video_id
}

struct DouyinCommentListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<Comment> comment_list // 评论列表
}


/* ============== Final. define services:================== */
service InteractionService {
    DouyinFavoriteActionResponse ActionFavorite(1:required DouyinFavoriteActionRequest req)
    DouyinFavoriteListResponse ListFavorite(1:required DouyinFavoriteListRequest req)
    DouyinCommentActionResponse ActionComment(1:required DouyinCommentActionRequest req)
    DouyinCommentListResponse ListComment(1:required DouyinCommentListRequest req)
}