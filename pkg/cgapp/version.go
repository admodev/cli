package cgapp

var (
	// cgapp CLI version
	version = "0.8.3"

	// Templates registry
	registry = map[string]string{
		// Backend templates
		"net/http": "create-go-app/net_http-go-template",
		"echo":     "create-go-app/echo-go-template",
		"fiber":    "create-go-app/fiber-go-template",

		// Frontend templates
		"react-js": "create-go-app/react-js-template",
		"react-ts": "create-go-app/react-ts-template",
		"preact":   "create-go-app/preact-js-template",

		// Docker containers
		"nginx":    "create-go-app/nginx-certbot-docker",
		"postgres": "create-go-app/postgres-docker",
	}
)
