package seventv

import (
	"context"
	"errors"
	"html/template"
	"regexp"

	"github.com/Chatterino/api/internal/db"
	"github.com/Chatterino/api/pkg/config"
	"github.com/Chatterino/api/pkg/resolver"
	"github.com/Chatterino/api/pkg/utils"
)

const (
	thumbnailFormat = "https://cdn.7tv.app/emote/%s/4x"

	tooltipTemplate = `<div style="text-align: left;">
<b>{{.Code}}</b><br>
<b>{{.Type}} SevenTV Emote</b><br>
<b>By:</b> {{.Uploader}}` +
		`{{ if .Unlisted }}` + `
<li><b><span style="color: red;">UNLISTED</span></b></li>{{ end }}
</div>`
)

var (
	errInvalidSevenTVEmotePath = errors.New("invalid SevenTV emote path")

	domains = map[string]struct{}{
		"7tv.app": {},
	}

	emotePathRegex = regexp.MustCompile(`/emotes/([a-f0-9]+)`)

	seventvEmoteTemplate = template.Must(template.New("seventvEmoteTooltip").Parse(tooltipTemplate))
)

func Initialize(ctx context.Context, cfg config.APIConfig, pool db.Pool, resolvers *[]resolver.Resolver) {
	apiURL := utils.MustParseURL("https://api.7tv.app/v2/gql")

	*resolvers = append(*resolvers, NewEmoteResolver(ctx, cfg, pool, apiURL))
}
