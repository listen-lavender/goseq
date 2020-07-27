package api

const (
	Workspace_PRE    = "workspace"
	Platform_WEB     = 1
	Platform_IOS     = 2
	Platform_ANDROID = 3
	Platform_DESKTOP = 4

	PlatformType_DESKTOP_AND_WEB uint8 = 1
	PlatformType_MOBILE          uint8 = 2

	FIRST_PAGE_LIMIT int = 15
)

// 原本该状态在cloudim/api/init.go中定义 现在cloudim中不再使用了 就集中在cloudim-logic中
type DeviceStatus = uint8
type WorkspaceStatus = uint8

const (
	// 设备状态 0:物理离线 1:物理在线 2:物理在线但逻辑离线
	DeviceOffline    = DeviceStatus(0)
	DeviceOnline     = DeviceStatus(1)
	DeviceOnlineIdle = DeviceStatus(2)

	WorkspaceOffline = WorkspaceStatus(0)
	WorkspaceOnline  = WorkspaceStatus(1)
)

var (
	DefaultContent = "未知消息，请升级。"

	PlatformTypeInfo = map[int]uint8{
		Platform_WEB:     PlatformType_DESKTOP_AND_WEB,
		Platform_DESKTOP: PlatformType_DESKTOP_AND_WEB,
		Platform_IOS:     PlatformType_MOBILE,
		Platform_ANDROID: PlatformType_MOBILE,
	}
)

func ConvertPlatform(ps string) int {
	if ps == "2" {
		return Platform_IOS // Platform_IOS
	} else if ps == "3" {
		return Platform_ANDROID // Platform_ANDROID
	} else if ps == "4" {
		return Platform_DESKTOP
	} else {
		return Platform_WEB // Platform_WEB
	}
}

func PlatformType(platform int) uint8 {
	return PlatformTypeInfo[platform]
}
