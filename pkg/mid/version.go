package mid

import (
	"context"
	"net/http"
)

// APIVersionCtx set the version value in context
func APIVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), versionKey, version))
			next.ServeHTTP(w, r)
		})
	}
}

// VersionFromContext get version string value from context
func VersionFromContext(ctx context.Context) (string, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns ok=false for nil.
	version, ok := ctx.Value(versionKey).(string)
	return version, ok
}
