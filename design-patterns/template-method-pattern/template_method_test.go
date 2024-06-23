package template_method_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateMethodFreeUser(t *testing.T) {
	feed := WebHomeFeedBuilder{UserType: Free}
	result := BuildHomeFeed(feed)
	expected := `<h1>Home feed</h1>
<h2>Ads</h2>
<h2>Friends</h2>
<h2>Promoted</h2>
<h2>Feeds</h2>
<h1>Footer</h1>`
	assert.Equal(t, expected, result)
}

func TestTemplateMethodPaidUser(t *testing.T) {
	feed := WebHomeFeedBuilder{UserType: Paid}
	expected := `<h1>Home feed</h1>
<h2>Friends</h2>
<h2>Feeds</h2>
<h1>Footer</h1>`
	result := BuildHomeFeed(feed)
	assert.Equal(t, expected, result)
}
