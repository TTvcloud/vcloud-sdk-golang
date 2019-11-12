package bytedance

// A Session provides a central location to create service clients, which
// is for future extension
type Session struct {
	Config   Config
}