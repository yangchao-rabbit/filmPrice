package all

import (
	_ "filmPrice/internal/apps/auth/http"
	_ "filmPrice/internal/apps/film/http"
	_ "filmPrice/internal/apps/system/http"
	_ "filmPrice/internal/apps/task/http"

	_ "filmPrice/internal/apps/auth/service"
	_ "filmPrice/internal/apps/film/service"
	_ "filmPrice/internal/apps/system/service"
	_ "filmPrice/internal/apps/task/service"
)
