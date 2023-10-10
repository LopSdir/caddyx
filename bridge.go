package izibridging

import (
	"fmt"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("hello_world", parseCaddyfile)
}

// 中间件实现了一个HTTP处理程序，将访问者的IP地址写入
// 访客的IP地址写到文件或流中。
type Middleware struct {
	Ak string
}

// CaddyModule返回Caddy模块的信息。
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.hello_world",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision实现了caddy.Provisioner。
func (m *Middleware) Provision(ctx caddy.Context) error {
	return nil
}

// Validate实现了caddy.Validator。
func (m *Middleware) Validate() error {
	return nil
}

// ServeHTTP 实现了 caddyhttp.MiddlewareHandler。
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	fmt.Printf("access key:%s\n", r.Header.Get("access_key"))
	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile实现了caddyfile.Unmarshaler。
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

// parseCaddyfile从h中解读令牌到一个新的中间件。
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
