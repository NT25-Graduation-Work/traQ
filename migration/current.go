package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"

	"github.com/traPtitech/traQ/model"
)

// Migrations 全てのデータベースマイグレーション
//
// 新たなマイグレーションを行う場合は、この配列の末尾に必ず追加すること
func Migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		v1(),  // インデックスidx_messages_deleted_atの削除とidx_messages_channel_id_deleted_at_created_atの追加
		v2(),  // RBAC周りのリフォーム
		v3(),  // チャンネルイベント履歴
		v4(),  // Webhook, Bot外部キー
		v5(),  // Mute, 旧Clip削除
		v6(),  // v6 ユーザーグループ拡張
		v7(),  // ファイルメタ拡張
		v8(),  // チャンネル購読拡張
		v9(),  // ユーザーテーブル拡張
		v10(), // パーミッション周りの調整
		v11(), // クリップ機能の追加
		v12(), // カスタムスタンプパレットの追加
		v13(), // パーミッション調整・インデックス付与
		v14(), // パーミッション不足修正
		v15(), // 外部ログイン機能追加
		v16(), // パーミッション修正
		v17(), // ユーザーホームチャンネル
		v18(), // インデックス追加
		v19(), // httpセッション管理テーブル変更
		v20(), // パーミッション周りの調整
		v21(), // OGPキャッシュ追加
		v22(), // BOTへのWebRTCパーミッションの付与
		v23(), // 複合インデックス追加
		v24(), // ユーザー設定追加
		v25(), // FileMetaにIsAnimatedImageを追加
		v26(), // FileMetaからThumbnail情報を分離
		v27(), // Gorm v2移行: FKの追加、FKのリネーム、一部フィールドのデータ型変更、idx_messages_channel_idの削除
		v28(), // v28 ユーザーグループにアイコンを追加
		v29(), // BotにModeを追加、WebSocket Modeを追加
		v30(), // bot_event_logsにresultを追加
		v31(), // お気に入りスタンプパーミッション削除（削除忘れ）
	}
}

// AllTables 最新のスキーマの全テーブルモデル
//
// 最新のスキーマの全テーブルのモデル構造体を記述すること
func AllTables() []interface{} {
	return []interface{}{
		&model.ChannelEvent{},
		&model.UserRole{},
		&model.RolePermission{},
		&model.DMChannelMapping{},
		&model.ChannelLatestMessage{},
		&model.BotEventLog{},
		&model.BotJoinChannel{},
		&model.Bot{},
		&model.OAuth2Client{},
		&model.OAuth2Authorize{},
		&model.OAuth2Token{},
		&model.MessageReport{},
		&model.WebhookBot{},
		&model.Stamp{},
		&model.UsersTag{},
		&model.Unread{},
		&model.Star{},
		&model.Device{},
		&model.Pin{},
		&model.FileACLEntry{},
		&model.FileThumbnail{},
		&model.FileMeta{},
		&model.UsersPrivateChannel{},
		&model.UserSubscribeChannel{},
		&model.Tag{},
		&model.ArchivedMessage{},
		&model.ClipFolderMessage{},
		&model.Message{},
		&model.StampPalette{},
		&model.UserGroup{},
		&model.UserGroupAdmin{},
		&model.UserGroupMember{},
		&model.ExternalProviderUser{},
		&model.UserProfile{},
		&model.Channel{},
		&model.ClipFolder{},
		&model.UserSettings{},
		&model.User{},
		&model.MessageStamp{},
		&model.SessionRecord{},
		&model.OgpCache{},
	}
}
