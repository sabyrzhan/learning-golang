package template_method_pattern

// Social network home feed builder
// BuildHomeFeed builds home feed data by calling specific methods in specified order
// BuildHomeFeed is main method that cannot be overridden or changed, however specific implementations implement
// HomeFeedBuilder interface.

type UserType int
const (
	Free UserType = iota
	Paid
)

type HomeFeedBuilder interface {
	GetUserType() UserType
	GetHeader() string
	GetAdsSection() string
	GetFriendsSection() string
	GetPromotedPosts() string
	GetFeedsSection() string
	GetFooterSection() string
}

type WebHomeFeedBuilder struct {
	UserType UserType
}

func (h WebHomeFeedBuilder) GetUserType() UserType {
	return h.UserType
}

func (h WebHomeFeedBuilder) GetHeader() string {
	return "<h1>Home feed</h1>"
}

func (h WebHomeFeedBuilder) GetAdsSection() string {
	return "<h2>Ads</h2>"
}

func (h WebHomeFeedBuilder) GetFriendsSection() string {
	return "<h2>Friends</h2>"
}

func (h WebHomeFeedBuilder) GetPromotedPosts() string {
	return "<h2>Promoted</h2>"
}

func (h WebHomeFeedBuilder) GetFeedsSection() string {
	return "<h2>Feeds</h2>"
}

func (h WebHomeFeedBuilder) GetFooterSection() string {
	return "<h1>Footer</h1>"
}

func BuildHomeFeed(builder HomeFeedBuilder) string {
	result := builder.GetHeader() + "\n"
	if builder.GetUserType() == Free {
		result += builder.GetAdsSection() + "\n"
	}
	result += builder.GetFriendsSection() + "\n"
	if builder.GetUserType() == Free {
		result += builder.GetPromotedPosts() + "\n"
	}
	result += builder.GetFeedsSection() + "\n"
	result += builder.GetFooterSection()

	return result
}
