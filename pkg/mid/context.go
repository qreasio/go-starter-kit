package mid

type contextKey string

var (
	paginationKey = contextKey("pagination")
	versionKey    = contextKey("version")
)

func (c contextKey) String() string {
	return "middleware-" + string(c)
}
