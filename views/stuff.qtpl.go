// Code generated by qtc from "stuff.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/stuff.qtpl:1
package views

//line views/stuff.qtpl:1
import "path/filepath"

//line views/stuff.qtpl:2
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/stuff.qtpl:3
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/stuff.qtpl:4
import "github.com/bouncepaw/mycorrhiza/user"

//line views/stuff.qtpl:5
import "github.com/bouncepaw/mycorrhiza/util"

//line views/stuff.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/stuff.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/stuff.qtpl:7
func StreamBaseHTML(qw422016 *qt422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:7
	qw422016.N().S(`
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>`)
//line views/stuff.qtpl:13
	qw422016.E().S(title)
//line views/stuff.qtpl:13
	qw422016.N().S(`</title>
		<link rel="shortcut icon" href="/static/favicon.ico">
		<link rel="stylesheet" href="/static/style.css">
		<script src="/static/shortcuts.js"></script>
		`)
//line views/stuff.qtpl:17
	for _, el := range headElements {
//line views/stuff.qtpl:17
		qw422016.N().S(el)
//line views/stuff.qtpl:17
	}
//line views/stuff.qtpl:17
	qw422016.N().S(`
	</head>
	<body>
		<header>
			<nav class="header-links main-width">
				<ul class="header-links__list">
`)
//line views/stuff.qtpl:23
	for _, link := range cfg.HeaderLinks {
//line views/stuff.qtpl:23
		qw422016.N().S(`					<li class="header-links__entry"><a class="header-links__link" href="`)
//line views/stuff.qtpl:24
		qw422016.E().S(link.Href)
//line views/stuff.qtpl:24
		qw422016.N().S(`">`)
//line views/stuff.qtpl:24
		qw422016.E().S(link.Display)
//line views/stuff.qtpl:24
		qw422016.N().S(`</a></li>
`)
//line views/stuff.qtpl:25
	}
//line views/stuff.qtpl:25
	qw422016.N().S(`					`)
//line views/stuff.qtpl:26
	qw422016.N().S(UserMenuHTML(u))
//line views/stuff.qtpl:26
	qw422016.N().S(`
				</ul>
			</nav>
		</header>
		`)
//line views/stuff.qtpl:30
	qw422016.N().S(body)
//line views/stuff.qtpl:30
	qw422016.N().S(`
		<template id="dialog-template">
			<div class="dialog-backdrop"></div>
			<div class="dialog" tabindex="0">
				<div class="dialog__header">
					<h1 class="dialog__title"></h1>
					<button class="dialog__close-button" aria-label="Close this dialog"></button>
				</div>

				<div class="dialog__content"></div>
			</div>
		</template>
		`)
//line views/stuff.qtpl:42
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:42
	qw422016.N().S(`
	</body>
</html>
`)
//line views/stuff.qtpl:45
}

//line views/stuff.qtpl:45
func WriteBaseHTML(qq422016 qtio422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:45
	StreamBaseHTML(qw422016, title, body, u, headElements...)
//line views/stuff.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:45
}

//line views/stuff.qtpl:45
func BaseHTML(title, body string, u *user.User, headElements ...string) string {
//line views/stuff.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:45
	WriteBaseHTML(qb422016, title, body, u, headElements...)
//line views/stuff.qtpl:45
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:45
	return qs422016
//line views/stuff.qtpl:45
}

//line views/stuff.qtpl:47
func StreamPrimitiveSearchHTML(qw422016 *qt422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:47
	qw422016.N().S(`
<div class="layout">
<main class="main-width primitive-search">
	<h1>Search results for ‘`)
//line views/stuff.qtpl:50
	qw422016.E().S(query)
//line views/stuff.qtpl:50
	qw422016.N().S(`’</h1>
	<p>Every hypha name has been compared with the query. Hyphae that have matched the query are listed below.</p>
	<ul class="primitive-search__results">
	`)
//line views/stuff.qtpl:53
	for hyphaName := range generator(query) {
//line views/stuff.qtpl:53
		qw422016.N().S(`
		<li class="primitive-search__entry">
			<a class="primitive-search__link wikilink" href="/hypha/`)
//line views/stuff.qtpl:55
		qw422016.E().S(hyphaName)
//line views/stuff.qtpl:55
		qw422016.N().S(`">`)
//line views/stuff.qtpl:55
		qw422016.E().S(util.BeautifulName(hyphaName))
//line views/stuff.qtpl:55
		qw422016.N().S(`</a>
		</li>
	`)
//line views/stuff.qtpl:57
	}
//line views/stuff.qtpl:57
	qw422016.N().S(`
</main>
</div>
`)
//line views/stuff.qtpl:60
}

//line views/stuff.qtpl:60
func WritePrimitiveSearchHTML(qq422016 qtio422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:60
	StreamPrimitiveSearchHTML(qw422016, query, generator)
//line views/stuff.qtpl:60
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:60
}

//line views/stuff.qtpl:60
func PrimitiveSearchHTML(query string, generator func(string) <-chan string) string {
//line views/stuff.qtpl:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:60
	WritePrimitiveSearchHTML(qb422016, query, generator)
//line views/stuff.qtpl:60
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:60
	return qs422016
//line views/stuff.qtpl:60
}

//line views/stuff.qtpl:62
func StreamHelpHTML(qw422016 *qt422016.Writer, content string) {
//line views/stuff.qtpl:62
	qw422016.N().S(`
<div class="layout">
<main class="main-width help">
	<article>
	`)
//line views/stuff.qtpl:66
	qw422016.N().S(content)
//line views/stuff.qtpl:66
	qw422016.N().S(`
	</article>
</main>
`)
//line views/stuff.qtpl:69
	qw422016.N().S(helpTopicsHTML())
//line views/stuff.qtpl:69
	qw422016.N().S(`
</div>
`)
//line views/stuff.qtpl:71
}

//line views/stuff.qtpl:71
func WriteHelpHTML(qq422016 qtio422016.Writer, content string) {
//line views/stuff.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:71
	StreamHelpHTML(qw422016, content)
//line views/stuff.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:71
}

//line views/stuff.qtpl:71
func HelpHTML(content string) string {
//line views/stuff.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:71
	WriteHelpHTML(qb422016, content)
//line views/stuff.qtpl:71
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:71
	return qs422016
//line views/stuff.qtpl:71
}

//line views/stuff.qtpl:73
func StreamHelpEmptyErrorHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:73
	qw422016.N().S(`
<h1>This entry does not exist!</h1>
<p>Try finding a different entry that would help you.</p>
<p>If you want to write this entry by yourself, consider <a class="wikilink wikilink_external wikilink_https" href="https://github.com/bouncepaw/mycorrhiza">contributing</a> to Mycorrhiza Wiki directly.</p>
`)
//line views/stuff.qtpl:77
}

//line views/stuff.qtpl:77
func WriteHelpEmptyErrorHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:77
	StreamHelpEmptyErrorHTML(qw422016)
//line views/stuff.qtpl:77
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:77
}

//line views/stuff.qtpl:77
func HelpEmptyErrorHTML() string {
//line views/stuff.qtpl:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:77
	WriteHelpEmptyErrorHTML(qb422016)
//line views/stuff.qtpl:77
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:77
	return qs422016
//line views/stuff.qtpl:77
}

//line views/stuff.qtpl:79
func streamhelpTopicsHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:79
	qw422016.N().S(`
<aside class="help-topics layout-card">
	<h2 class="layout-card__title">Help topics</h2>
	<ul class="help-topics__list">
		<li>
			<a href="/help/en">Main</a>
		</li>
		<li>
			<a href="/help/en/hypha">Hypha</a>
			<ul>
				<li>
					<a href="/help/en/attachment">Attachment</a>
				</li>
			</ul
		</li>
	</ul>
</aside>
`)
//line views/stuff.qtpl:96
}

//line views/stuff.qtpl:96
func writehelpTopicsHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:96
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:96
	streamhelpTopicsHTML(qw422016)
//line views/stuff.qtpl:96
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:96
}

//line views/stuff.qtpl:96
func helpTopicsHTML() string {
//line views/stuff.qtpl:96
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:96
	writehelpTopicsHTML(qb422016)
//line views/stuff.qtpl:96
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:96
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:96
	return qs422016
//line views/stuff.qtpl:96
}

//line views/stuff.qtpl:98
func StreamUserListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:98
	qw422016.N().S(`
<div class="layout">
<main class="main-width user-list">
	<h1>List of users</h1>
`)
//line views/stuff.qtpl:103
	var (
		admins     = make([]string, 0)
		moderators = make([]string, 0)
		editors    = make([]string, 0)
	)
	for u := range user.YieldUsers() {
		switch u.Group {
		case "admin":
			admins = append(admins, u.Name)
		case "moderator":
			moderators = append(moderators, u.Name)
		case "editor", "trusted":
			editors = append(editors, u.Name)
		}
	}

//line views/stuff.qtpl:118
	qw422016.N().S(`
	<section>
		<h2>Admins</h2>
		<ol>`)
//line views/stuff.qtpl:121
	for _, name := range admins {
//line views/stuff.qtpl:121
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:122
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:122
		qw422016.N().S(`/`)
//line views/stuff.qtpl:122
		qw422016.E().S(name)
//line views/stuff.qtpl:122
		qw422016.N().S(`">`)
//line views/stuff.qtpl:122
		qw422016.E().S(name)
//line views/stuff.qtpl:122
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:123
	}
//line views/stuff.qtpl:123
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Moderators</h2>
		<ol>`)
//line views/stuff.qtpl:127
	for _, name := range moderators {
//line views/stuff.qtpl:127
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:128
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:128
		qw422016.N().S(`/`)
//line views/stuff.qtpl:128
		qw422016.E().S(name)
//line views/stuff.qtpl:128
		qw422016.N().S(`">`)
//line views/stuff.qtpl:128
		qw422016.E().S(name)
//line views/stuff.qtpl:128
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:129
	}
//line views/stuff.qtpl:129
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Editors</h2>
		<ol>`)
//line views/stuff.qtpl:133
	for _, name := range editors {
//line views/stuff.qtpl:133
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:134
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:134
		qw422016.N().S(`/`)
//line views/stuff.qtpl:134
		qw422016.E().S(name)
//line views/stuff.qtpl:134
		qw422016.N().S(`">`)
//line views/stuff.qtpl:134
		qw422016.E().S(name)
//line views/stuff.qtpl:134
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:135
	}
//line views/stuff.qtpl:135
	qw422016.N().S(`</ol>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:139
}

//line views/stuff.qtpl:139
func WriteUserListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:139
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:139
	StreamUserListHTML(qw422016)
//line views/stuff.qtpl:139
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:139
}

//line views/stuff.qtpl:139
func UserListHTML() string {
//line views/stuff.qtpl:139
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:139
	WriteUserListHTML(qb422016)
//line views/stuff.qtpl:139
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:139
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:139
	return qs422016
//line views/stuff.qtpl:139
}

//line views/stuff.qtpl:141
func StreamHyphaListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:141
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>List of hyphae</h1>
	<p>This wiki has `)
//line views/stuff.qtpl:145
	qw422016.N().D(hyphae.Count())
//line views/stuff.qtpl:145
	qw422016.N().S(` hyphae.</p>
	<ul class="hypha-list">
		`)
//line views/stuff.qtpl:147
	for h := range hyphae.YieldExistingHyphae() {
//line views/stuff.qtpl:147
		qw422016.N().S(`
		<li class="hypha-list__entry">
			<a class="hypha-list__link" href="/hypha/`)
//line views/stuff.qtpl:149
		qw422016.E().S(h.Name)
//line views/stuff.qtpl:149
		qw422016.N().S(`">`)
//line views/stuff.qtpl:149
		qw422016.E().S(util.BeautifulName(h.Name))
//line views/stuff.qtpl:149
		qw422016.N().S(`</a>
			`)
//line views/stuff.qtpl:150
		if h.BinaryPath != "" {
//line views/stuff.qtpl:150
			qw422016.N().S(`
			<span class="hypha-list__amnt-type">`)
//line views/stuff.qtpl:151
			qw422016.E().S(filepath.Ext(h.BinaryPath)[1:])
//line views/stuff.qtpl:151
			qw422016.N().S(`</span>
			`)
//line views/stuff.qtpl:152
		}
//line views/stuff.qtpl:152
		qw422016.N().S(`
		</li>
		`)
//line views/stuff.qtpl:154
	}
//line views/stuff.qtpl:154
	qw422016.N().S(`
	</ul>
</main>
</div>
`)
//line views/stuff.qtpl:158
}

//line views/stuff.qtpl:158
func WriteHyphaListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:158
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:158
	StreamHyphaListHTML(qw422016)
//line views/stuff.qtpl:158
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:158
}

//line views/stuff.qtpl:158
func HyphaListHTML() string {
//line views/stuff.qtpl:158
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:158
	WriteHyphaListHTML(qb422016)
//line views/stuff.qtpl:158
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:158
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:158
	return qs422016
//line views/stuff.qtpl:158
}

//line views/stuff.qtpl:160
func StreamAboutHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:160
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section>
		<h1>About `)
//line views/stuff.qtpl:164
	qw422016.E().S(cfg.WikiName)
//line views/stuff.qtpl:164
	qw422016.N().S(`</h1>
		<ul>
			<li><b><a href="https://mycorrhiza.wiki">Mycorrhiza Wiki</a> version:</b> 1.3.0</li>
`)
//line views/stuff.qtpl:167
	if cfg.UseAuth {
//line views/stuff.qtpl:167
		qw422016.N().S(`			<li><b>User count:</b> `)
//line views/stuff.qtpl:168
		qw422016.N().DUL(user.Count())
//line views/stuff.qtpl:168
		qw422016.N().S(`</li>
			<li><b>Home page:</b> <a href="/">`)
//line views/stuff.qtpl:169
		qw422016.E().S(cfg.HomeHypha)
//line views/stuff.qtpl:169
		qw422016.N().S(`</a></li>
			<li><b>Administrators:</b>`)
//line views/stuff.qtpl:170
		for i, username := range user.ListUsersWithGroup("admin") {
//line views/stuff.qtpl:171
			if i > 0 {
//line views/stuff.qtpl:171
				qw422016.N().S(`<span aria-hidden="true">, </span>
`)
//line views/stuff.qtpl:172
			}
//line views/stuff.qtpl:172
			qw422016.N().S(`				<a href="/hypha/`)
//line views/stuff.qtpl:173
			qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:173
			qw422016.N().S(`/`)
//line views/stuff.qtpl:173
			qw422016.E().S(username)
//line views/stuff.qtpl:173
			qw422016.N().S(`">`)
//line views/stuff.qtpl:173
			qw422016.E().S(username)
//line views/stuff.qtpl:173
			qw422016.N().S(`</a>`)
//line views/stuff.qtpl:173
		}
//line views/stuff.qtpl:173
		qw422016.N().S(`</li>
`)
//line views/stuff.qtpl:174
	} else {
//line views/stuff.qtpl:174
		qw422016.N().S(`			<li>This wiki does not use authorization</li>
`)
//line views/stuff.qtpl:176
	}
//line views/stuff.qtpl:176
	qw422016.N().S(`		</ul>
		<p>See <a href="/list">/list</a> for information about hyphae on this wiki.</p>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:182
}

//line views/stuff.qtpl:182
func WriteAboutHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:182
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:182
	StreamAboutHTML(qw422016)
//line views/stuff.qtpl:182
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:182
}

//line views/stuff.qtpl:182
func AboutHTML() string {
//line views/stuff.qtpl:182
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:182
	WriteAboutHTML(qb422016)
//line views/stuff.qtpl:182
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:182
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:182
	return qs422016
//line views/stuff.qtpl:182
}

//line views/stuff.qtpl:184
func StreamCommonScripts(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:184
	qw422016.N().S(`
`)
//line views/stuff.qtpl:185
	for _, scriptPath := range cfg.CommonScripts {
//line views/stuff.qtpl:185
		qw422016.N().S(`
<script src="`)
//line views/stuff.qtpl:186
		qw422016.E().S(scriptPath)
//line views/stuff.qtpl:186
		qw422016.N().S(`"></script>
`)
//line views/stuff.qtpl:187
	}
//line views/stuff.qtpl:187
	qw422016.N().S(`
`)
//line views/stuff.qtpl:188
}

//line views/stuff.qtpl:188
func WriteCommonScripts(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:188
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:188
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:188
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:188
}

//line views/stuff.qtpl:188
func CommonScripts() string {
//line views/stuff.qtpl:188
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:188
	WriteCommonScripts(qb422016)
//line views/stuff.qtpl:188
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:188
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:188
	return qs422016
//line views/stuff.qtpl:188
}
